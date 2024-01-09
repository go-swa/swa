
package service_config

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "demo1/service/service_config/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, getCfgGormDbMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgGormDb", Remote: false}), getCfgJwtMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgJwt", Remote: false}), getCfgSwaConfigMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgSwaConfig", Remote: false}), setCfgGormDbMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "SetCfgGormDb", Remote: false}), setSystemConfigMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "SetSystemConfig", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, getCfgGormDbMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgGormDb", Remote: true}), getCfgJwtMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgJwt", Remote: true}), getCfgSwaConfigMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "GetCfgSwaConfig", Remote: true}), setCfgGormDbMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "SetCfgGormDb", Remote: true}), setSystemConfigMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_config/T", Method: "SetSystemConfig", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

var _ weaver.InstanceOf[T] = (*impl)(nil)

var _ weaver.Unrouted = (*impl)(nil)


type t_local_stub struct {
	impl                   T
	tracer                 trace.Tracer
	getCfgGormDbMetrics    *codegen.MethodMetrics
	getCfgJwtMetrics       *codegen.MethodMetrics
	getCfgSwaConfigMetrics *codegen.MethodMetrics
	setCfgGormDbMetrics    *codegen.MethodMetrics
	setSystemConfigMetrics *codegen.MethodMetrics
}

var _ T = (*t_local_stub)(nil)

func (s t_local_stub) GetCfgGormDb(ctx context.Context) (r0 ModGormDb, err error) {
	begin := s.getCfgGormDbMetrics.Begin()
	defer func() { s.getCfgGormDbMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_config.T.GetCfgGormDb", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetCfgGormDb(ctx)
}

func (s t_local_stub) GetCfgJwt(ctx context.Context) (r0 ModJwt, err error) {
	begin := s.getCfgJwtMetrics.Begin()
	defer func() { s.getCfgJwtMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_config.T.GetCfgJwt", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetCfgJwt(ctx)
}

func (s t_local_stub) GetCfgSwaConfig(ctx context.Context) (r0 ModSwaConfig, err error) {
	begin := s.getCfgSwaConfigMetrics.Begin()
	defer func() { s.getCfgSwaConfigMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_config.T.GetCfgSwaConfig", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetCfgSwaConfig(ctx)
}

func (s t_local_stub) SetCfgGormDb(ctx context.Context, a0 ModGormDb) (err error) {
	begin := s.setCfgGormDbMetrics.Begin()
	defer func() { s.setCfgGormDbMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_config.T.SetCfgGormDb", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetCfgGormDb(ctx, a0)
}

func (s t_local_stub) SetSystemConfig(ctx context.Context, a0 ModSwaConfig) (err error) {
	begin := s.setSystemConfigMetrics.Begin()
	defer func() { s.setSystemConfigMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_config.T.SetSystemConfig", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetSystemConfig(ctx, a0)
}


type t_client_stub struct {
	stub                   codegen.Stub
	getCfgGormDbMetrics    *codegen.MethodMetrics
	getCfgJwtMetrics       *codegen.MethodMetrics
	getCfgSwaConfigMetrics *codegen.MethodMetrics
	setCfgGormDbMetrics    *codegen.MethodMetrics
	setSystemConfigMetrics *codegen.MethodMetrics
}

var _ T = (*t_client_stub)(nil)

func (s t_client_stub) GetCfgGormDb(ctx context.Context) (r0 ModGormDb, err error) {
	var requestBytes, replyBytes int
	begin := s.getCfgGormDbMetrics.Begin()
	defer func() { s.getCfgGormDbMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_config.T.GetCfgGormDb", trace.WithSpanKind(trace.SpanKindClient))
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

	var shardKey uint64

	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
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

func (s t_client_stub) GetCfgJwt(ctx context.Context) (r0 ModJwt, err error) {
	var requestBytes, replyBytes int
	begin := s.getCfgJwtMetrics.Begin()
	defer func() { s.getCfgJwtMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_config.T.GetCfgJwt", trace.WithSpanKind(trace.SpanKindClient))
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

	var shardKey uint64

	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
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

func (s t_client_stub) GetCfgSwaConfig(ctx context.Context) (r0 ModSwaConfig, err error) {
	var requestBytes, replyBytes int
	begin := s.getCfgSwaConfigMetrics.Begin()
	defer func() { s.getCfgSwaConfigMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_config.T.GetCfgSwaConfig", trace.WithSpanKind(trace.SpanKindClient))
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

	var shardKey uint64

	var results []byte
	results, err = s.stub.Run(ctx, 2, nil, shardKey)
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

func (s t_client_stub) SetCfgGormDb(ctx context.Context, a0 ModGormDb) (err error) {
	var requestBytes, replyBytes int
	begin := s.setCfgGormDbMetrics.Begin()
	defer func() { s.setCfgGormDbMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_config.T.SetCfgGormDb", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_ModGormDb_1df36f7c(&a0)
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
	err = dec.Error()
	return
}

func (s t_client_stub) SetSystemConfig(ctx context.Context, a0 ModSwaConfig) (err error) {
	var requestBytes, replyBytes int
	begin := s.setSystemConfigMetrics.Begin()
	defer func() { s.setSystemConfigMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_config.T.SetSystemConfig", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_ModSwaConfig_54667f3d(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

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
	case "GetCfgGormDb":
		return s.getCfgGormDb
	case "GetCfgJwt":
		return s.getCfgJwt
	case "GetCfgSwaConfig":
		return s.getCfgSwaConfig
	case "SetCfgGormDb":
		return s.setCfgGormDb
	case "SetSystemConfig":
		return s.setSystemConfig
	default:
		return nil
	}
}

func (s t_server_stub) getCfgGormDb(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetCfgGormDb(ctx)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getCfgJwt(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetCfgJwt(ctx)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getCfgSwaConfig(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetCfgSwaConfig(ctx)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setCfgGormDb(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModGormDb
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetCfgGormDb(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setSystemConfig(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModSwaConfig
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetSystemConfig(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}


type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) GetCfgGormDb(ctx context.Context) (r0 ModGormDb, err error) {
	err = s.caller("GetCfgGormDb", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetCfgJwt(ctx context.Context) (r0 ModJwt, err error) {
	err = s.caller("GetCfgJwt", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetCfgSwaConfig(ctx context.Context) (r0 ModSwaConfig, err error) {
	err = s.caller("GetCfgSwaConfig", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) SetCfgGormDb(ctx context.Context, a0 ModGormDb) (err error) {
	err = s.caller("SetCfgGormDb", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) SetSystemConfig(ctx context.Context, a0 ModSwaConfig) (err error) {
	err = s.caller("SetSystemConfig", ctx, []any{a0}, []any{})
	return
}


var _ codegen.AutoMarshal = (*Local)(nil)

type __is_Local[T ~struct {
	weaver.AutoMarshal
	Path      string "mapstructure:\"path\" json:\"path\" yaml:\"path\""
	StorePath string "mapstructure:\"store-path\" json:\"store-path\" yaml:\"store-path\""
}] struct{}

var _ __is_Local[Local]

func (x *Local) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Local.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Path)
	enc.String(x.StorePath)
}

func (x *Local) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Local.WeaverUnmarshal: nil receiver"))
	}
	x.Path = dec.String()
	x.StorePath = dec.String()
}

var _ codegen.AutoMarshal = (*ModCaptcha)(nil)

type __is_ModCaptcha[T ~struct {
	weaver.AutoMarshal
	KeyLong            int "mapstructure:\"key-long\" json:\"key-long\" yaml:\"key-long\""
	ImgWidth           int "mapstructure:\"img-width\" json:\"img-width\" yaml:\"img-width\""
	ImgHeight          int "mapstructure:\"img-height\" json:\"img-height\" yaml:\"img-height\""
	OpenCaptcha        int "mapstructure:\"open-captcha\" json:\"open-captcha\" yaml:\"open-captcha\""
	OpenCaptchaTimeOut int "mapstructure:\"open-captcha-timeout\" json:\"open-captcha-timeout\" yaml:\"open-captcha-timeout\""
}] struct{}

var _ __is_ModCaptcha[ModCaptcha]

func (x *ModCaptcha) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModCaptcha.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.KeyLong)
	enc.Int(x.ImgWidth)
	enc.Int(x.ImgHeight)
	enc.Int(x.OpenCaptcha)
	enc.Int(x.OpenCaptchaTimeOut)
}

func (x *ModCaptcha) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModCaptcha.WeaverUnmarshal: nil receiver"))
	}
	x.KeyLong = dec.Int()
	x.ImgWidth = dec.Int()
	x.ImgHeight = dec.Int()
	x.OpenCaptcha = dec.Int()
	x.OpenCaptchaTimeOut = dec.Int()
}

var _ codegen.AutoMarshal = (*ModGormDb)(nil)

type __is_ModGormDb[T ~struct {
	weaver.AutoMarshal
	DbType       string "mapstructure:\"db-type\" json:\"db-type\" yaml:\"db-type\""
	Ip           string "mapstructure:\"ip\" json:\"ip\" yaml:\"ip\""
	Port         string "mapstructure:\"port\" json:\"port\" yaml:\"port\""
	Username     string "mapstructure:\"username\" json:\"username\" yaml:\"username\""
	Password     string "mapstructure:\"password\" json:\"password\" yaml:\"password\""
	Dbname       string "mapstructure:\"db-name\" json:\"db-name\" yaml:\"db-name\""
	Config       string "mapstructure:\"config\" json:\"config\" yaml:\"config\""
	Prefix       string "mapstructure:\"prefix\" json:\"prefix\" yaml:\"prefix\""
	Singular     bool   "mapstructure:\"singular\" json:\"singular\" yaml:\"singular\""
	Engine       string "mapstructure:\"engine\" json:\"engine\" yaml:\"engine\" default:\"InnoDB\""
	MaxIdleConns int    "mapstructure:\"max-idle-conns\" json:\"max-idle-conns\" yaml:\"max-idle-conns\""
	MaxOpenConns int    "mapstructure:\"max-open-conns\" json:\"max-open-conns\" yaml:\"max-open-conns\""
	LogMode      string "mapstructure:\"log-mode\" json:\"log-mode\" yaml:\"log-mode\""
	LogZap       bool   "mapstructure:\"log-zap\" json:\"log-zap\" yaml:\"log-zap\""
	Dsn          string
}] struct{}

var _ __is_ModGormDb[ModGormDb]

func (x *ModGormDb) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModGormDb.WeaverMarshal: nil receiver"))
	}
	enc.String(x.DbType)
	enc.String(x.Ip)
	enc.String(x.Port)
	enc.String(x.Username)
	enc.String(x.Password)
	enc.String(x.Dbname)
	enc.String(x.Config)
	enc.String(x.Prefix)
	enc.Bool(x.Singular)
	enc.String(x.Engine)
	enc.Int(x.MaxIdleConns)
	enc.Int(x.MaxOpenConns)
	enc.String(x.LogMode)
	enc.Bool(x.LogZap)
	enc.String(x.Dsn)
}

func (x *ModGormDb) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModGormDb.WeaverUnmarshal: nil receiver"))
	}
	x.DbType = dec.String()
	x.Ip = dec.String()
	x.Port = dec.String()
	x.Username = dec.String()
	x.Password = dec.String()
	x.Dbname = dec.String()
	x.Config = dec.String()
	x.Prefix = dec.String()
	x.Singular = dec.Bool()
	x.Engine = dec.String()
	x.MaxIdleConns = dec.Int()
	x.MaxOpenConns = dec.Int()
	x.LogMode = dec.String()
	x.LogZap = dec.Bool()
	x.Dsn = dec.String()
}

var _ codegen.AutoMarshal = (*ModJwt)(nil)

type __is_ModJwt[T ~struct {
	weaver.AutoMarshal
	SigningKey  string "mapstructure:\"signing-key\" json:\"signing-key\" yaml:\"signing-key\""
	ExpiresTime string "mapstructure:\"expires-time\" json:\"expires-time\" yaml:\"expires-time\""
	BufferTime  string "mapstructure:\"buffer-time\" json:\"buffer-time\" yaml:\"buffer-time\""
	Issuer      string "mapstructure:\"issuer\" json:\"issuer\" yaml:\"issuer\""
}] struct{}

var _ __is_ModJwt[ModJwt]

func (x *ModJwt) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModJwt.WeaverMarshal: nil receiver"))
	}
	enc.String(x.SigningKey)
	enc.String(x.ExpiresTime)
	enc.String(x.BufferTime)
	enc.String(x.Issuer)
}

func (x *ModJwt) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModJwt.WeaverUnmarshal: nil receiver"))
	}
	x.SigningKey = dec.String()
	x.ExpiresTime = dec.String()
	x.BufferTime = dec.String()
	x.Issuer = dec.String()
}

var _ codegen.AutoMarshal = (*ModRedis)(nil)

type __is_ModRedis[T ~struct {
	weaver.AutoMarshal
	DB       int    "mapstructure:\"db\" json:\"db\" yaml:\"db\""
	Addr     string "mapstructure:\"addr\" json:\"addr\" yaml:\"addr\""
	Password string "mapstructure:\"password\" json:\"password\" yaml:\"password\""
}] struct{}

var _ __is_ModRedis[ModRedis]

func (x *ModRedis) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModRedis.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.DB)
	enc.String(x.Addr)
	enc.String(x.Password)
}

func (x *ModRedis) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModRedis.WeaverUnmarshal: nil receiver"))
	}
	x.DB = dec.Int()
	x.Addr = dec.String()
	x.Password = dec.String()
}

var _ codegen.AutoMarshal = (*ModSwaConfig)(nil)

type __is_ModSwaConfig[T ~struct {
	weaver.AutoMarshal
	Jwt     ModJwt     "mapstructure:\"jwt\" json:\"jwt\" toml:\"jwt\""
	Redis   ModRedis   "mapstructure:\"redis\" json:\"redis\" toml:\"redis\""
	Captcha ModCaptcha "mapstructure:\"captcha\" json:\"captcha\" toml:\"captcha\""
	GormDB  ModGormDb  "mapstructure:\"gormDB\" json:\"gormDB\" toml:\"gormDB\""
	System  ModSystem  "mapstructure:\"system\" json:\"system\" toml:\"system\""
	Local   Local      "mapstructure:\"local\" json:\"local\" toml:\"local\""
}] struct{}

var _ __is_ModSwaConfig[ModSwaConfig]

func (x *ModSwaConfig) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModSwaConfig.WeaverMarshal: nil receiver"))
	}
	(x.Jwt).WeaverMarshal(enc)
	(x.Redis).WeaverMarshal(enc)
	(x.Captcha).WeaverMarshal(enc)
	(x.GormDB).WeaverMarshal(enc)
	(x.System).WeaverMarshal(enc)
	(x.Local).WeaverMarshal(enc)
}

func (x *ModSwaConfig) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModSwaConfig.WeaverUnmarshal: nil receiver"))
	}
	(&x.Jwt).WeaverUnmarshal(dec)
	(&x.Redis).WeaverUnmarshal(dec)
	(&x.Captcha).WeaverUnmarshal(dec)
	(&x.GormDB).WeaverUnmarshal(dec)
	(&x.System).WeaverUnmarshal(dec)
	(&x.Local).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*ModSystem)(nil)

type __is_ModSystem[T ~struct {
	weaver.AutoMarshal
	Env           string "mapstructure:\"env\" json:\"env\" yaml:\"env\""
	Addr          int    "mapstructure:\"addr\" json:\"addr\" yaml:\"addr\""
	DbType        string "mapstructure:\"db-type\" json:\"db-type\" yaml:\"db-type\""
	OssType       string "mapstructure:\"oss-type\" json:\"oss-type\" yaml:\"oss-type\""
	UseMultipoint bool   "mapstructure:\"use-multipoint\" json:\"use-multipoint\" yaml:\"use-multipoint\""
	UseRedis      bool   "mapstructure:\"use-redis\" json:\"use-redis\" yaml:\"use-redis\""
	LimitCountIP  int    "mapstructure:\"ip-limit-count\" json:\"ip-limit-count\" yaml:\"ip-limit-count\""
	LimitTimeIP   int    "mapstructure:\"ip-limit-time\" json:\"ip-limit-time\" yaml:\"ip-limit-time\""
	RouterPrefix  string "mapstructure:\"router-prefix\" json:\"router-prefix\" yaml:\"router-prefix\""
}] struct{}

var _ __is_ModSystem[ModSystem]

func (x *ModSystem) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModSystem.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Env)
	enc.Int(x.Addr)
	enc.String(x.DbType)
	enc.String(x.OssType)
	enc.Bool(x.UseMultipoint)
	enc.Bool(x.UseRedis)
	enc.Int(x.LimitCountIP)
	enc.Int(x.LimitTimeIP)
	enc.String(x.RouterPrefix)
}

func (x *ModSystem) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModSystem.WeaverUnmarshal: nil receiver"))
	}
	x.Env = dec.String()
	x.Addr = dec.Int()
	x.DbType = dec.String()
	x.OssType = dec.String()
	x.UseMultipoint = dec.Bool()
	x.UseRedis = dec.Bool()
	x.LimitCountIP = dec.Int()
	x.LimitTimeIP = dec.Int()
	x.RouterPrefix = dec.String()
}


func serviceweaver_size_Local_39802d8d(x *Local) int {
	size := 0
	size += 0
	size += (4 + len(x.Path))
	size += (4 + len(x.StorePath))
	return size
}

func serviceweaver_size_ModCaptcha_8669fd4f(x *ModCaptcha) int {
	size := 0
	size += 0
	size += 8
	size += 8
	size += 8
	size += 8
	size += 8
	return size
}

func serviceweaver_size_ModGormDb_1df36f7c(x *ModGormDb) int {
	size := 0
	size += 0
	size += (4 + len(x.DbType))
	size += (4 + len(x.Ip))
	size += (4 + len(x.Port))
	size += (4 + len(x.Username))
	size += (4 + len(x.Password))
	size += (4 + len(x.Dbname))
	size += (4 + len(x.Config))
	size += (4 + len(x.Prefix))
	size += 1
	size += (4 + len(x.Engine))
	size += 8
	size += 8
	size += (4 + len(x.LogMode))
	size += 1
	size += (4 + len(x.Dsn))
	return size
}

func serviceweaver_size_ModJwt_71a8c3b1(x *ModJwt) int {
	size := 0
	size += 0
	size += (4 + len(x.SigningKey))
	size += (4 + len(x.ExpiresTime))
	size += (4 + len(x.BufferTime))
	size += (4 + len(x.Issuer))
	return size
}

func serviceweaver_size_ModRedis_79c3af32(x *ModRedis) int {
	size := 0
	size += 0
	size += 8
	size += (4 + len(x.Addr))
	size += (4 + len(x.Password))
	return size
}

func serviceweaver_size_ModSwaConfig_54667f3d(x *ModSwaConfig) int {
	size := 0
	size += 0
	size += serviceweaver_size_ModJwt_71a8c3b1(&x.Jwt)
	size += serviceweaver_size_ModRedis_79c3af32(&x.Redis)
	size += serviceweaver_size_ModCaptcha_8669fd4f(&x.Captcha)
	size += serviceweaver_size_ModGormDb_1df36f7c(&x.GormDB)
	size += serviceweaver_size_ModSystem_8e8da7a8(&x.System)
	size += serviceweaver_size_Local_39802d8d(&x.Local)
	return size
}

func serviceweaver_size_ModSystem_8e8da7a8(x *ModSystem) int {
	size := 0
	size += 0
	size += (4 + len(x.Env))
	size += 8
	size += (4 + len(x.DbType))
	size += (4 + len(x.OssType))
	size += 1
	size += 1
	size += 8
	size += 8
	size += (4 + len(x.RouterPrefix))
	return size
}
