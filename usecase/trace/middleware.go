package trace

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type graphqlTracer struct{}

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
	graphql.FieldInterceptor
} = &graphqlTracer{}

func NewGraphQLTracer() graphql.HandlerExtension {
	return &graphqlTracer{}
}

func (t graphqlTracer) ExtensionName() string {
	return "GraphQLTracer"
}

func (t graphqlTracer) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (t graphqlTracer) InterceptResponse(
	ctx context.Context,
	next graphql.ResponseHandler,
) *graphql.Response {
	rc := graphql.GetOperationContext(ctx)

	span, sctx := tracer.StartSpanFromContext(
		ctx,
		"graphql.request",
		tracer.ResourceName(rc.OperationName),
		tracer.Tag("graphql.operation.name", rc.OperationName),
		tracer.Tag("graphql.query", rc.RawQuery),
		tracer.Tag("graphql.variables", rc.Variables),
	)

	res := next(sctx)
	if res == nil {
		span.Finish()

		return res
	}

	if res.Errors != nil {
		err := res.Errors[0]
		span.Finish(tracer.WithError(err))
	} else {
		span.Finish()
	}

	return res
}

func (t graphqlTracer) InterceptField(
	ctx context.Context,
	next graphql.Resolver,
) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)

	span, sctx := tracer.StartSpanFromContext(
		ctx,
		"graphql.field",
		tracer.ResourceName(fc.Field.ObjectDefinition.Name+"."+fc.Field.Name),
		tracer.Tag("graphql.field.name", fc.Field.Name),
		tracer.Tag("graphql.field.type", fc.Field.ObjectDefinition.Name),
	)

	res, err := next(sctx)

	if err != nil {
		span.Finish(tracer.WithError(err))
	} else {
		span.Finish()
	}

	return res, err
}
