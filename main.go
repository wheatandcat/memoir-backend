package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/profiler"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	sentry "github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/generated"
	"github.com/wheatandcat/memoir-backend/repository"
	"github.com/wheatandcat/memoir-backend/usecase/app_trace"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"github.com/wheatandcat/memoir-backend/usecase/logger"
)

const defaultPort = "8080"

// telemetryEndpoint はGoogle Cloud Telemetry API（OTLP）のエンドポイント
const telemetryEndpoint = "telemetry.googleapis.com:443"

// newTraceExporter はGoogle Cloud Telemetry APIへ送信するOTLP trace exporterを作成する
func newTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	creds, err := oauth.NewApplicationDefault(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, err
	}

	return otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(telemetryEndpoint),
		otlptracegrpc.WithDialOption(grpc.WithPerRPCCredentials(creds)),
	)
}

// newTraceResource はトレースに付与するResourceを作成する
func newTraceResource(ctx context.Context, projectID, version string) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceName("memoir-backend"),
			semconv.ServiceVersion(version),
			attribute.String("gcp.project_id", projectID),
		),
	)
}

//lint:ignore U1000 is used in graphql generated code
func installPropagators() {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		))
}

func main() {
	v := "1.0.0"
	if os.Getenv("RELEASE_INSTANCE_VERSION") != "" {
		v = os.Getenv("RELEASE_INSTANCE_VERSION")
	}

	ctx := context.Background()

	tr := sdktrace.NewTracerProvider().Tracer("memoir-backend")

	if os.Getenv("APP_ENV") != "local" {
		projectID := os.Getenv("GCP_PROJECT_ID")

		cfg := profiler.Config{
			Service:        "memoir-backend",
			ServiceVersion: v,
			// ProjectID must be set if not running on GCP.
			ProjectID: projectID,

			// For OpenCensus users:
			// To see Profiler agent spans in APM backend,
			// set EnableOCTelemetry to true
			EnableOCTelemetry: true,
		}

		if err := profiler.Start(cfg); err != nil {
			log.Fatalf("profiler.Start: %v", err)
		}

		exporter, err := newTraceExporter(ctx)
		if err != nil {
			log.Fatalf("newTraceExporter: %v", err)
		}

		res, err := newTraceResource(ctx, projectID, v)
		if err != nil {
			log.Fatalf("newTraceResource: %v", err)
		}

		tp := sdktrace.NewTracerProvider(
			sdktrace.WithResource(res),
			sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.05)),
			//sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
		)

		defer func() {
			if err := tp.ForceFlush(ctx); err != nil {
				log.Fatal(err)
			}
		}()
		otel.SetTracerProvider(tp)
		tr = otel.GetTracerProvider().Tracer("memoir-backend")
	}

	if os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("読み込み出来ませんでした: %v", err)
		}
	}

	sco := sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}
	if os.Getenv("APP_ENV") == "production" {
		sco.Release = os.Getenv("RELEASE_INSTANCE_VERSION")
	}

	if os.Getenv("APP_ENV") != "local" {
		// ローカルの時はSentryを送信しない
		err := sentry.Init(sco)
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	f, err := repository.FirebaseApp(ctx)
	if err != nil {
		panic(err)
	}

	router.Use(logger.Middleware())

	a := auth.New()
	router.Use(a.NotLoginMiddleware())
	router.Use(a.FirebaseLoginMiddleware(f))

	fc, err := f.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	app := graph.NewApplication(tr)

	resolver := &graph.Resolver{
		FirestoreClient: fc,
		App:             app,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.Use(app_trace.NewGraphQLTracer(tr))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)

		if oc.Operation.Name != "IntrospectionQuery" {
			logger.New(ctx).Info("graphql info",
				zap.String("RawQuery", oc.RawQuery),
				zap.String("OperationName", oc.Operation.Name),
			)
		}

		return next(ctx)
	})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		goc := graphql.GetOperationContext(ctx)

		errorCode := ce.CodeDefault

		var re ce.RequestError
		if errors.As(e, &re) {
			errorCode = re.Code
		}

		err.Extensions = map[string]interface{}{
			"code": errorCode,
		}

		logger.New(ctx).Error("graphql error", zap.Error(e))

		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("kind", "GraphQL")
			scope.SetTag("operationName", goc.OperationName)
			scope.SetContext("GraphQL", sentry.Context{
				"query":     goc.RawQuery,
				"variables": goc.Variables,
				"errorCode": errorCode,
			})

			if err.Path.String() != "" {
				sentry.AddBreadcrumb(&sentry.Breadcrumb{
					Category: "GraphQL",
					Message:  "Error Path:" + err.Path.String(),
					Level:    sentry.LevelInfo,
				})
			}

			sentry.CaptureException(e)
		})

		return err
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	httpsrv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
		Addr:         ":" + port,
	}

	err = httpsrv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
