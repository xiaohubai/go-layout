package tracing

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	jaegerClient "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/xiaohubai/go-layout/configs/global"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func OpentracingInit() io.Closer {
	tracer, closer, err := OpentracingTracer(global.Cfg.Jaeger.Name, global.Cfg.Jaeger.Address)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	global.Tracer = tracer
	return closer
}
func OpentracingTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaegerClient.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Millisecond,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, err
}

var (
	hostname, _ = os.Hostname()
)

func OpentelemetryInit() (tp *traceSdk.TracerProvider) {
	agentConfig := strings.SplitN(global.Cfg.Jaeger.Address, ":", 2)
	endpointOption := jaeger.WithAgentEndpoint(jaeger.WithAgentHost(agentConfig[0]), jaeger.WithAgentPort(agentConfig[1]))
	exp, err := jaeger.New(endpointOption)
	if err != nil {
		panic(err)
	}
	tp = traceSdk.NewTracerProvider(
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithBatcher(exp),
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(global.Cfg.Jaeger.Name),
			semConv.ServiceVersionKey.String(global.Cfg.System.Version),
			semConv.HostNameKey.String(hostname),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp
}

// NewSpan creates a span using default tracer.
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, spanName, opts...)
}
