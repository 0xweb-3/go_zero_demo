package jeager

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jeagerconfig "github.com/uber/jaeger-client-go/config"
	"testing"
)

func Test_Jeager(t *testing.T) {
	cfg := jeagerconfig.Configuration{
		// 定义取样器
		Sampler: &jeagerconfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, // 取样器类型
			Param: 1,
		},
		// 采样信息发送的对象
		Reporter: &jeagerconfig.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: fmt.Sprintf("http://%s/api/traces", "127.0.0.1:14268"),
		},
	}

	// 创建jeager客户端
	jaeger, err := cfg.InitGlobalTracer("client test", jeagerconfig.Logger(jaeger.StdLogger))

	if err != nil {
		t.Log(err)
		return
	}
	defer jaeger.Close()

	// 任务的执行
	// 得到tracer
	tracer := opentracing.GlobalTracer()

	// 为任务节点定义span
	parentSpan := tracer.StartSpan("A")
	defer parentSpan.Finish()

	B(tracer, parentSpan)
}

// 子集任务
func B(tracer opentracing.Tracer, parentSpan opentracing.Span) {
	childSpan := tracer.StartSpan("B", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
}
