
package frontend

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Listeners: []string{"address"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦daca1075:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→demo1/service/service_config/T⟧\n⟦d15b1802:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→demo1/service/service_swa/T⟧\n⟦cbfa2416:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→demo1/service/service_file/T⟧\n⟦0da11b29:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→address⟧\n",
	})
}

var _ weaver.InstanceOf[weaver.Main] = (*server)(nil)

var _ weaver.Unrouted = (*server)(nil)


type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

var _ weaver.Main = (*main_local_stub)(nil)


type main_client_stub struct {
	stub codegen.Stub
}

var _ weaver.Main = (*main_client_stub)(nil)

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


type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

var _ codegen.Server = (*main_server_stub)(nil)

func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}


type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

var _ weaver.Main = (*main_reflect_stub)(nil)

