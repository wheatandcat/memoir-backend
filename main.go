package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/generated"
	"github.com/wheatandcat/memoir-backend/repository"
	"github.com/wheatandcat/memoir-backend/usecase/app_trace"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"github.com/wheatandcat/memoir-backend/usecase/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func installPropagators() {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			// Putting the CloudTraceOneWayPropagator first means the TraceContext propagator
			// takes precedence if both the traceparent and the XCTC headers exist.
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		))
}

func main() {

	ctx := context.Background()

	var tr trace.Tracer = sdktrace.NewTracerProvider().Tracer("memoir-backend")

	if os.Getenv("APP_ENV") != "local" {
		projectID := os.Getenv("GCP_PROJECT_ID")
		exporter, err := texporter.New(texporter.WithProjectID(projectID))
		if err != nil {
			log.Fatalf("texporter.NewExporter: %v", err)
		}

		tp := sdktrace.NewTracerProvider(
			//sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.1)),
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
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

	a := auth.New(tr, ctx)
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
			scope.SetExtra("query", goc.RawQuery)
			scope.SetExtra("variables", goc.Variables)
			scope.SetExtra("errorCode", errorCode)

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

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
