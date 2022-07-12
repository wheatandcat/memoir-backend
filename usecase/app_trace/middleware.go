package app_trace

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"go.opencensus.io/trace"
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
	sctx, span := trace.StartSpan(ctx, rc.OperationName)
	defer span.End()

	if !span.IsRecordingEvents() {
		return next(sctx)
	}

	span.AddAttributes(
		trace.StringAttribute("request.query", rc.RawQuery),
	)

	res := next(sctx)
	if res == nil {
		return res
	}

	if res.Errors != nil {
		err := res.Errors[0]
		span.AddAttributes(trace.StringAttribute("error.message", err.Message))
	}
	return res
}

func (t graphqlTracer) InterceptField(
	ctx context.Context,
	next graphql.Resolver,
) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)
	sctx, span := trace.StartSpan(ctx, fc.Field.ObjectDefinition.Name+"/"+fc.Field.Name)
	defer span.End()
	if !span.IsRecordingEvents() {
		return next(sctx)
	}
	span.AddAttributes(
		trace.StringAttribute("resolver.path", fc.Path().String()),
		trace.StringAttribute("resolver.object", fc.Field.ObjectDefinition.Name),
		trace.StringAttribute("resolver.field", fc.Field.Name),
		trace.StringAttribute("resolver.alias", fc.Field.Alias),
	)
	for _, arg := range fc.Field.Arguments {
		if arg.Value != nil {
			span.AddAttributes(
				trace.StringAttribute(fmt.Sprintf("resolver.args.%s", arg.Name), arg.Value.String()),
			)
		}
	}

	res, err := next(sctx)

	if err != nil {
		span.AddAttributes(trace.StringAttribute("error.message", err.Error()))

	}

	errList := graphql.GetFieldErrors(sctx, fc)
	if len(errList) != 0 {
		span.SetStatus(trace.Status{
			Code:    2,
			Message: errList.Error(),
		})
		span.AddAttributes(
			trace.BoolAttribute("resolver.hasError", true),
			trace.Int64Attribute("resolver.errorCount", int64(len(errList))),
		)
		for idx, err := range errList {
			span.AddAttributes(
				trace.StringAttribute(fmt.Sprintf("resolver.error.%d.message", idx), err.Error()),
				trace.StringAttribute(fmt.Sprintf("resolver.error.%d.kind", idx), fmt.Sprintf("%T", err)),
			)
		}
	}

	return res, err
}
