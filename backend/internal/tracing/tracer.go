package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	tr "go.opentelemetry.io/otel/trace"
)

var globalTracer = trace.NewTracerProvider().Tracer("noop")

func Init(traceAPIEndpoint, serviceName string, sampleRatio float64) (cancel func(context.Context) error, err error) {
	ex, err := jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint(traceAPIEndpoint),
		),
	)

	if err != nil {
		return nil, err
	}
	prov, cancel := newTracerProvider(ex, serviceName, sampleRatio)
	otel.SetTracerProvider(prov)
	globalTracer = prov.Tracer("default tracer")
	return cancel, nil
}

func newTracerProvider(ex *jaeger.Exporter, serviceName string, sampleRatio float64) (*trace.TracerProvider, func(ctx context.Context) error) {
	batcher := trace.NewBatchSpanProcessor(ex)
	tp := trace.NewTracerProvider(
		trace.WithSpanProcessor(batcher),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
		trace.WithSampler(
			trace.ParentBased(trace.TraceIDRatioBased(sampleRatio)),
		),
	)
	return tp, batcher.Shutdown
}

func StartSpanFromContext(ctx context.Context, name string, opts ...tr.SpanStartOption) (context.Context, tr.Span) {
	return globalTracer.Start(ctx, name, opts...)
}
