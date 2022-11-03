package app_trace

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

type graphqlTracer struct {
	tracer trace.Tracer
}

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
	graphql.FieldInterceptor
} = &graphqlTracer{}

func NewGraphQLTracer(tracer trace.Tracer) graphql.HandlerExtension {
	return &graphqlTracer{
		tracer: tracer,
	}
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
	oc := graphql.GetOperationContext(ctx)
	if oc.Operation.Name == "IntrospectionQuery" {
		return next(ctx)
	}

	q := strings.Split(oc.RawQuery, " ")[0]
	ctx, span := t.tracer.Start(ctx, q+":"+oc.OperationName, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()
	if !span.IsRecording() {
		return next(ctx)
	}

	span.SetAttributes(
		attribute.Key("request.query").String(oc.RawQuery),
	)

	res := next(ctx)
	if res == nil {
		return res
	}

	if len(res.Errors) > 0 {
		span.SetStatus(codes.Error, res.Errors.Error())
		span.RecordError(fmt.Errorf(res.Errors.Error()))
		err := res.Errors[0]
		span.SetAttributes(attribute.Key("error.message").String(err.Message))

	}
	return res
}

func (t graphqlTracer) InterceptField(
	ctx context.Context,
	next graphql.Resolver,
) (interface{}, error) {
	oc := graphql.GetOperationContext(ctx)
	if oc.Operation.Name == "IntrospectionQuery" {
		return next(ctx)
	}

	fc := graphql.GetFieldContext(ctx)

	ctx, span := t.tracer.Start(ctx,
		fc.Field.ObjectDefinition.Name+"/"+fc.Field.Name,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	if !span.IsRecording() {
		return next(ctx)
	}

	span.SetAttributes(
		attribute.Key("resolver.path").String(fc.Path().String()),
		attribute.Key("resolver.object").String(fc.Field.ObjectDefinition.Name),
		attribute.Key("resolver.field").String(fc.Field.Name),
		attribute.Key("resolver.alias").String(fc.Field.Alias),
	)

	argKV := []attribute.KeyValue{}
	for _, arg := range fc.Field.Arguments {
		if arg.Value != nil {
			argKV = append(argKV, attribute.Key(fmt.Sprintf("resolver.args.%s", arg.Name)).String(arg.Value.String()))
		}
	}

	if len(argKV) > 0 {
		span.SetAttributes(argKV...)
	}

	res, err := next(ctx)

	if err != nil {
		span.SetAttributes(attribute.Key("error.message").String(err.Error()))
	}

	errList := graphql.GetFieldErrors(ctx, fc)
	if len(errList) != 0 {
		span.SetStatus(codes.Error, errList.Error())
		span.RecordError(fmt.Errorf(errList.Error()))
		err := errList[0]
		span.SetAttributes(attribute.Key("error.message").String(err.Message))
	}

	return res, ce.CustomError(err)
}
