// Code generated by protoc-gen-go-telemetry version:(devel). DO NOT EDIT.

package testv1

import (
	context "context"
	trace "go.opencensus.io/trace"
)

func (x *StringMessage) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("stringmessage.msg", x.Msg),
	)
}

func (x *StringMessage) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".stringmessage.msg", x.Msg),
	)
}

func (x *Int32Message) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute("int32message.num32", int64(x.Num32)),
	)
}

func (x *Int32Message) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute(pfx+".int32message.num32", int64(x.Num32)),
	)
}

func (x *Uint32Message) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute("uint32message.unum32", int64(x.Unum32)),
	)
}

func (x *Uint32Message) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute(pfx+".uint32message.unum32", int64(x.Unum32)),
	)
}

func (x *Int64Message) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute("int64message.num64", int64(x.Num64)),
	)
}

func (x *Int64Message) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.Int64Attribute(pfx+".int64message.num64", int64(x.Num64)),
	)
}

func (x *SubMessage) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes()
}

func (x *SubMessage) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes()
}

func (x *SubMessage_Envelope) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("submessage.envelope.name", x.Name),
	)
}

func (x *SubMessage_Envelope) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".submessage.envelope.name", x.Name),
	)
}

func (x *MapMessage) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes()
}

func (x *MapMessage) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes()
}

func (x *MessageDetails) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("messagedetails.details", x.Details),
	)
}

func (x *MessageDetails) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".messagedetails.details", x.Details),
	)
}

func (x *ExcludeField) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("excludefield.non_masked", x.NonMasked),
	)
}

func (x *ExcludeField) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".excludefield.non_masked", x.NonMasked),
	)
}

func (x *ExcludeMessage) TraceAttributes(ctx context.Context) {
}

func (x *ExcludeMessage) TraceNamedAttributes(ctx context.Context, pfx string) {
}

func (x *RenameMessagePrefix) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("pfx.msg", x.Msg),
	)
}

func (x *RenameMessagePrefix) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".pfx.msg", x.Msg),
	)
}

func (x *NameField) TraceAttributes(ctx context.Context) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute("namefield.message", x.Msg),
	)
}

func (x *NameField) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.FromContext(ctx)
	if !span.IsRecordingEvents() {
		return
	}

	span.AddAttributes(
		trace.StringAttribute(pfx+".namefield.message", x.Msg),
	)
}
