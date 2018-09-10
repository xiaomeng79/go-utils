package trace

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"log"
	"io"
)

func TraceingInit(address,servicename string ) io.Closer {
	//配置
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	metricsFactory := metrics.NewLocalFactory(0)
	_metrics := jaeger.NewMetrics(metricsFactory, nil)

	sender, err := jaeger.NewUDPTransport(address, 0)
	if err != nil {
		log.Println("could not initialize jaeger sender: "+err.Error())
		return nil
	}

	repoter := jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Metrics(_metrics))

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		servicename,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Reporter(repoter),
	)

	if err != nil {
		log.Println("could not initialize jaeger tracer: "+err.Error())
		return nil
	}
	return closer
}
