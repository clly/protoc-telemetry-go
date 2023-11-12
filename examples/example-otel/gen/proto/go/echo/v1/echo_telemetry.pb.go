// Code generated by protoc-gen-go-telemetry version:(devel). DO NOT EDIT.

package ocechov1

import (
	context "context"
	fmt "fmt"
	attribute "go.opentelemetry.io/otel/attribute"
	trace "go.opentelemetry.io/otel/trace"
)

func (x *EchoRequest) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String("echorequest.msg", x.Msg),
		attribute.Int64("echorequest.number", int64(x.Num32)),
		attribute.Int64("echorequest.unum32", int64(x.Unum32)),
		attribute.Int64("echorequest.num64", int64(x.Num64)),
	)
	for m, v := range x.GetMeta() {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("echorequest.meta_%s", m), v),
		)
	}
}

func (x *EchoRequest) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String(pfx+".echorequest.msg", x.Msg),
		attribute.Int64(pfx+".echorequest.number", int64(x.Num32)),
		attribute.Int64(pfx+".echorequest.unum32", int64(x.Unum32)),
		attribute.Int64(pfx+".echorequest.num64", int64(x.Num64)),
	)
	for m, v := range x.GetMeta() {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("pfx.echorequest.meta_%s", m), v),
		)
	}
}

func (x *EchoRequest_Envelope) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String("echorequest.envelope.name", x.Name),
	)
}

func (x *EchoRequest_Envelope) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String(pfx+".echorequest.envelope.name", x.Name),
	)
}

func (x *MessageDetails) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String("messagedetails.details", x.Details),
	)
}

func (x *MessageDetails) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String(pfx+".messagedetails.details", x.Details),
	)
}

func (x *EchoResponse) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String("echoresponse.msg", x.Msg),
	)
}

func (x *EchoResponse) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(
		attribute.String(pfx+".echoresponse.msg", x.Msg),
	)
}