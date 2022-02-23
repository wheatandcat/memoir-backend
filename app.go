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
	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/generated"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

const defaultPort = "8080"

func main() {
	if os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("読み込み出来ませんでした: %v", err)
		}
	}

	sco := sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}
	if os.Getenv("APP_ENV") != "local" {
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

	ctx := context.Background()
	f, err := repository.FirebaseApp(ctx)
	if err != nil {
		panic(err)
	}

	router.Use(auth.NotLoginMiddleware())
	router.Use(auth.FirebaseLoginMiddleware(f))

	fc, err := f.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	app := graph.NewApplication()

	resolver := &graph.Resolver{
		FirestoreClient: fc,
		App:             app,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
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

		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("kind", "GraphQL")
			scope.SetTag("operationName", goc.OperationName)
			scope.SetExtra("query", goc.RawQuery)
			scope.SetExtra("variables", goc.Variables)
			scope.SetExtra("Error Code", errorCode)

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
