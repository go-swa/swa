
package service_file

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/textproto"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "demo1/service/service_file/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, deleteFileMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "DeleteFile", Remote: false}), editFileNameMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "EditFileName", Remote: false}), findFileMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "FindFile", Remote: false}), getFileRecordInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "GetFileRecordInfoList", Remote: false}), uploadMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "Upload", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, deleteFileMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "DeleteFile", Remote: true}), editFileNameMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "EditFileName", Remote: true}), findFileMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "FindFile", Remote: true}), getFileRecordInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "GetFileRecordInfoList", Remote: true}), uploadMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_file/T", Method: "Upload", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "⟦2334d5dc:wEaVeReDgE:demo1/service/service_file/T→demo1/service/service_config/T⟧\n",
	})
}

var _ weaver.InstanceOf[T] = (*impl)(nil)

var _ weaver.Unrouted = (*impl)(nil)


type t_local_stub struct {
	impl                         T
	tracer                       trace.Tracer
	deleteFileMetrics            *codegen.MethodMetrics
	editFileNameMetrics          *codegen.MethodMetrics
	findFileMetrics              *codegen.MethodMetrics
	getFileRecordInfoListMetrics *codegen.MethodMetrics
	uploadMetrics                *codegen.MethodMetrics
}

var _ T = (*t_local_stub)(nil)

func (s t_local_stub) DeleteFile(ctx context.Context, a0 ModUpload) (err error) {
	begin := s.deleteFileMetrics.Begin()
	defer func() { s.deleteFileMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_file.T.DeleteFile", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteFile(ctx, a0)
}

func (s t_local_stub) EditFileName(ctx context.Context, a0 ModUpload) (err error) {
	begin := s.editFileNameMetrics.Begin()
	defer func() { s.editFileNameMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_file.T.EditFileName", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.EditFileName(ctx, a0)
}

func (s t_local_stub) FindFile(ctx context.Context, a0 uint) (r0 ModUpload, err error) {
	begin := s.findFileMetrics.Begin()
	defer func() { s.findFileMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_file.T.FindFile", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FindFile(ctx, a0)
}

func (s t_local_stub) GetFileRecordInfoList(ctx context.Context, a0 PageInfo) (r0 []ModUpload, r1 int64, err error) {
	begin := s.getFileRecordInfoListMetrics.Begin()
	defer func() { s.getFileRecordInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_file.T.GetFileRecordInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetFileRecordInfoList(ctx, a0)
}

func (s t_local_stub) Upload(ctx context.Context, a0 ModUpload) (err error) {
	begin := s.uploadMetrics.Begin()
	defer func() { s.uploadMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_file.T.Upload", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Upload(ctx, a0)
}


type t_client_stub struct {
	stub                         codegen.Stub
	deleteFileMetrics            *codegen.MethodMetrics
	editFileNameMetrics          *codegen.MethodMetrics
	findFileMetrics              *codegen.MethodMetrics
	getFileRecordInfoListMetrics *codegen.MethodMetrics
	uploadMetrics                *codegen.MethodMetrics
}

var _ T = (*t_client_stub)(nil)

func (s t_client_stub) DeleteFile(ctx context.Context, a0 ModUpload) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteFileMetrics.Begin()
	defer func() { s.deleteFileMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_file.T.DeleteFile", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) EditFileName(ctx context.Context, a0 ModUpload) (err error) {
	var requestBytes, replyBytes int
	begin := s.editFileNameMetrics.Begin()
	defer func() { s.editFileNameMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_file.T.EditFileName", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) FindFile(ctx context.Context, a0 uint) (r0 ModUpload, err error) {
	var requestBytes, replyBytes int
	begin := s.findFileMetrics.Begin()
	defer func() { s.findFileMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_file.T.FindFile", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.Uint(a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetFileRecordInfoList(ctx context.Context, a0 PageInfo) (r0 []ModUpload, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getFileRecordInfoListMetrics.Begin()
	defer func() { s.getFileRecordInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_file.T.GetFileRecordInfoList", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	size := 0
	size += serviceweaver_size_PageInfo_8dcdfc2b(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_ModUpload_be84408f(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) Upload(ctx context.Context, a0 ModUpload) (err error) {
	var requestBytes, replyBytes int
	begin := s.uploadMetrics.Begin()
	defer func() { s.uploadMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_file.T.Upload", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 4, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)


type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

var _ codegen.Server = (*t_server_stub)(nil)

func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "DeleteFile":
		return s.deleteFile
	case "EditFileName":
		return s.editFileName
	case "FindFile":
		return s.findFile
	case "GetFileRecordInfoList":
		return s.getFileRecordInfoList
	case "Upload":
		return s.upload
	default:
		return nil
	}
}

func (s t_server_stub) deleteFile(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModUpload
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteFile(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) editFileName(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModUpload
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.EditFileName(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) findFile(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.FindFile(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getFileRecordInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 PageInfo
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetFileRecordInfoList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_ModUpload_be84408f(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) upload(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModUpload
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.Upload(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}


type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) DeleteFile(ctx context.Context, a0 ModUpload) (err error) {
	err = s.caller("DeleteFile", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) EditFileName(ctx context.Context, a0 ModUpload) (err error) {
	err = s.caller("EditFileName", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) FindFile(ctx context.Context, a0 uint) (r0 ModUpload, err error) {
	err = s.caller("FindFile", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetFileRecordInfoList(ctx context.Context, a0 PageInfo) (r0 []ModUpload, r1 int64, err error) {
	err = s.caller("GetFileRecordInfoList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) Upload(ctx context.Context, a0 ModUpload) (err error) {
	err = s.caller("Upload", ctx, []any{a0}, []any{})
	return
}


var _ codegen.AutoMarshal = (*FileHeader)(nil)

type __is_FileHeader[T ~struct {
	weaver.AutoMarshal
	Filename  string
	Header    textproto.MIMEHeader
	Size      int64
	content   []byte
	tmpfile   string
	tmpoff    int64
	tmpshared bool
}] struct{}

var _ __is_FileHeader[FileHeader]

func (x *FileHeader) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("FileHeader.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Filename)
	serviceweaver_enc_map_string_slice_string_c493bdb8(enc, (map[string][]string)(x.Header))
	enc.Int64(x.Size)
	serviceweaver_enc_slice_byte_87461245(enc, x.content)
	enc.String(x.tmpfile)
	enc.Int64(x.tmpoff)
	enc.Bool(x.tmpshared)
}

func (x *FileHeader) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("FileHeader.WeaverUnmarshal: nil receiver"))
	}
	x.Filename = dec.String()
	*(*map[string][]string)(&x.Header) = serviceweaver_dec_map_string_slice_string_c493bdb8(dec)
	x.Size = dec.Int64()
	x.content = serviceweaver_dec_slice_byte_87461245(dec)
	x.tmpfile = dec.String()
	x.tmpoff = dec.Int64()
	x.tmpshared = dec.Bool()
}

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

func serviceweaver_enc_map_string_slice_string_c493bdb8(enc *codegen.Encoder, arg map[string][]string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for k, v := range arg {
		enc.String(k)
		serviceweaver_enc_slice_string_4af10117(enc, v)
	}
}

func serviceweaver_dec_map_string_slice_string_c493bdb8(dec *codegen.Decoder) map[string][]string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make(map[string][]string, n)
	var k string
	var v []string
	for i := 0; i < n; i++ {
		k = dec.String()
		v = serviceweaver_dec_slice_string_4af10117(dec)
		res[k] = v
	}
	return res
}

func serviceweaver_enc_slice_byte_87461245(enc *codegen.Encoder, arg []byte) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Byte(arg[i])
	}
}

func serviceweaver_dec_slice_byte_87461245(dec *codegen.Decoder) []byte {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Byte()
	}
	return res
}

var _ codegen.AutoMarshal = (*ModBase)(nil)

type __is_ModBase[T ~struct {
	weaver.AutoMarshal
	ID        uint "gorm:\"primarykey\""
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time "gorm:\"index\" json:\"-\""
	CreatedBy uint      "json:\"createdBy\" gorm:\"column:createdBy;comment:创建者\""
	UpdatedBy uint      "json:\"updatedBy\" gorm:\"column:updatedBy;comment:更新者\""
	DeletedBy uint      "json:\"deletedBy\" gorm:\"column:deletedBy\""
}] struct{}

var _ __is_ModBase[ModBase]

func (x *ModBase) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModBase.WeaverMarshal: nil receiver"))
	}
	enc.Uint(x.ID)
	enc.EncodeBinaryMarshaler(&x.CreatedAt)
	enc.EncodeBinaryMarshaler(&x.UpdatedAt)
	enc.EncodeBinaryMarshaler(&x.DeletedAt)
	enc.Uint(x.CreatedBy)
	enc.Uint(x.UpdatedBy)
	enc.Uint(x.DeletedBy)
}

func (x *ModBase) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModBase.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.Uint()
	dec.DecodeBinaryUnmarshaler(&x.CreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.UpdatedAt)
	dec.DecodeBinaryUnmarshaler(&x.DeletedAt)
	x.CreatedBy = dec.Uint()
	x.UpdatedBy = dec.Uint()
	x.DeletedBy = dec.Uint()
}

var _ codegen.AutoMarshal = (*ModUpload)(nil)

type __is_ModUpload[T ~struct {
	weaver.AutoMarshal
	ModBase
	Name string "json:\"name\" gorm:\"comment:文件名\""
	Url  string "json:\"url\" gorm:\"comment:文件地址\""
	Tag  string "json:\"tag\" gorm:\"comment:文件标签\""
	Key  string "json:\"key\" gorm:\"comment:编号\""
}] struct{}

var _ __is_ModUpload[ModUpload]

func (x *ModUpload) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModUpload.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Name)
	enc.String(x.Url)
	enc.String(x.Tag)
	enc.String(x.Key)
}

func (x *ModUpload) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModUpload.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Name = dec.String()
	x.Url = dec.String()
	x.Tag = dec.String()
	x.Key = dec.String()
}

var _ codegen.AutoMarshal = (*PageInfo)(nil)

type __is_PageInfo[T ~struct {
	weaver.AutoMarshal
	Page     int    "json:\"page\" form:\"page\""
	PageSize int    "json:\"pageSize\" form:\"pageSize\""
	Keyword  string "json:\"keyword\" form:\"keyword\""
}] struct{}

var _ __is_PageInfo[PageInfo]

func (x *PageInfo) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("PageInfo.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.Page)
	enc.Int(x.PageSize)
	enc.String(x.Keyword)
}

func (x *PageInfo) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("PageInfo.WeaverUnmarshal: nil receiver"))
	}
	x.Page = dec.Int()
	x.PageSize = dec.Int()
	x.Keyword = dec.String()
}


func serviceweaver_enc_slice_ModUpload_be84408f(enc *codegen.Encoder, arg []ModUpload) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModUpload_be84408f(dec *codegen.Decoder) []ModUpload {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModUpload, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}


func serviceweaver_size_PageInfo_8dcdfc2b(x *PageInfo) int {
	size := 0
	size += 0
	size += 8
	size += 8
	size += (4 + len(x.Keyword))
	return size
}
