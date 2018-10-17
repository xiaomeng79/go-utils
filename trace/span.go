//传递一些上下文环境变量
package trace

import (
	"context"
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"strings"
)

//解析trace中的信息
func decodeTracer(ctx context.Context) []string {
	span := opentracing.SpanFromContext(ctx)
	s := strings.Split(fmt.Sprintf("%v", span), ":")
	if len(s) >= 3 {
		return s
	}
	return []string{"0", "0", "0"}
}

func StartSpanFromContext(ctx context.Context, operationName string) (context.Context, opentracing.Span) {
	span, _ := opentracing.StartSpanFromContext(ctx, operationName)

	ctx = opentracing.ContextWithSpan(ctx, span)
	s := decodeTracer(ctx)
	span.SetTag("trace_id", s[0])
	span.SetTag("span_id", s[1])
	span.SetTag("parent_id", s[2])
	return ctx, span
}

func StartSpanFromRoot(operationName string) (context.Context, opentracing.Span) {
	ctx := context.Background()
	span := opentracing.GlobalTracer().StartSpan(operationName)
	ctx = opentracing.ContextWithSpan(ctx, span)
	s := decodeTracer(ctx)
	span.SetTag("trace_id", s[0])
	span.SetTag("span_id", s[1])
	span.SetTag("parent_id", s[2])
	return ctx, span
}
