package graph

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	sentry "github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/client/authToken"
	"github.com/wheatandcat/memoir-backend/client/task"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// Graph Graph struct
type Graph struct {
	UserID          string
	FirebaseUID     string
	FirestoreClient *firestore.Client
	App             *Application
	Client          *Client
}

type Client struct {
	UUID      uuidgen.UUIDGenerator
	Time      timegen.TimeGenerator
	AuthToken authToken.AuthTokenClient
	Task      task.HTTPTaskInterface
}

// NewGraph Graphを作成
func NewGraph(ctx context.Context, app *Application, f *firestore.Client) (*Graph, error) {
	_, span := app.TraceClient.Start(ctx,
		"NewGraph",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	user := auth.ForContext(ctx)
	if user == nil {
		return nil, ce.CustomError(fmt.Errorf("access denied"))
	}

	if user.FirebaseUID == "" {
		// Firebase認証無し
		u, err := app.UserRepository.FindDatabaseDataByUID(ctx, f, user.ID)
		if err != nil {
			return nil, ce.CustomErrorWrap(err, "User Invalid")
		}
		if u.FirebaseUID != "" {
			return nil, ce.CustomError(fmt.Errorf("need to firebase auth"))
		}
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "Auth",
			Message:  "Not logged in, UserID: " + u.ID,
			Level:    sentry.LevelInfo,
		})

	} else {
		// Firebase認証有り
		u, err := app.UserRepository.FindByFirebaseUID(ctx, f, user.FirebaseUID)
		if err != nil {
			return nil, ce.CustomErrorWrap(err, "firebase auth invalid")
		}

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "Auth",
			Message:  "Logged in, UserID: " + u.ID,
			Level:    sentry.LevelInfo,
		})

		user.ID = u.ID
	}

	if user.ID == "" {
		return nil, ce.CustomError(fmt.Errorf("UserID Invalid"))
	}

	return NewGraphWithSetUserID(ctx, app, f, user.ID, user.FirebaseUID), nil
}

// NewGraphWithSetUserID Graphを作成（ログイン前）
func NewGraphWithSetUserID(ctx context.Context, app *Application, f *firestore.Client, uid, fuid string) *Graph {
	_, span := app.TraceClient.Start(ctx,
		"NewGraphWithSetUserID",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{ID: uid})
	})

	client := &Client{
		UUID:      &uuidgen.UUID{},
		Time:      &timegen.Time{},
		AuthToken: &authToken.AuthToken{},
		Task:      task.NewNotificationTask(),
	}

	return &Graph{
		UserID:          uid,
		FirebaseUID:     fuid,
		FirestoreClient: f,
		App:             app,
		Client:          client,
	}
}

func GetNestCollectFields(ctx context.Context, cfs []graphql.CollectedField, columnName string) []graphql.CollectedField {
	for _, cf := range cfs {
		if cf.Name == columnName {
			return graphql.CollectFields(graphql.GetOperationContext(ctx), cf.Selections, nil)
		}
	}

	return []graphql.CollectedField{}
}

func GetNestCollectFieldArgumentValue(ctx context.Context, cfs []graphql.CollectedField, columnName string, argumentName string) string {
	for _, cf := range cfs {
		if cf.Name == columnName {
			for _, arg := range cf.Arguments {
				if arg.Name == argumentName {
					return arg.Value.String()
				}
			}

		}
	}

	return ""
}

func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func GrpcErrorStatusCode(err error) codes.Code {
	return status.Code(errors.Cause(err))
}
