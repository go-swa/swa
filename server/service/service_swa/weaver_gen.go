
package service_swa

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/gofrs/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "demo1/service/service_swa/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, addBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "AddBaseMenu", Remote: false}), addMenuRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "AddMenuRole", Remote: false}), changePasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "ChangePassword", Remote: false}), createApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateApi", Remote: false}), createBlacklistMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateBlacklist", Remote: false}), createRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateRole", Remote: false}), createSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaDict", Remote: false}), createSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaDictail", Remote: false}), createSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaRecord", Remote: false}), deleteApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteApi", Remote: false}), deleteApisByIdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteApisByIds", Remote: false}), deleteBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteBaseMenu", Remote: false}), deleteRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteRole", Remote: false}), deleteSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaDict", Remote: false}), deleteSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaDictail", Remote: false}), deleteSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaRecord", Remote: false}), deleteSwaRecordByIdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaRecordByIds", Remote: false}), deleteUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteUser", Remote: false}), getAPIInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAPIInfoList", Remote: false}), getAllApisMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAllApis", Remote: false}), getAllBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAllBaseMenu", Remote: false}), getApiByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetApiById", Remote: false}), getBaseMenuByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetBaseMenuById", Remote: false}), getInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetInfoList", Remote: false}), getMenuRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetMenuRole", Remote: false}), getPolicyPathByRoleIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetPolicyPathByRoleId", Remote: false}), getRoleInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleInfo", Remote: false}), getRoleListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleList", Remote: false}), getRoleMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleMenu", Remote: false}), getServerInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetServerInfo", Remote: false}), getSubRolesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSubRoles", Remote: false}), getSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDict", Remote: false}), getSwaDictInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictInfoList", Remote: false}), getSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictail", Remote: false}), getSwaDictailInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictailInfoList", Remote: false}), getSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaRecord", Remote: false}), getSwaRecordInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaRecordInfoList", Remote: false}), getUserDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserDict", Remote: false}), getUserInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserInfo", Remote: false}), getUserInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserInfoList", Remote: false}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Login", Remote: false}), registerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Register", Remote: false}), resetPasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "ResetPassword", Remote: false}), scasbinMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Scasbin", Remote: false}), setRoleMenusMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetRoleMenus", Remote: false}), setSelfInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetSelfInfo", Remote: false}), setSubRolesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetSubRoles", Remote: false}), setUserAuthoritiesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserAuthorities", Remote: false}), setUserInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserInfo", Remote: false}), setUserRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserRole", Remote: false}), updateApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateApi", Remote: false}), updateBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateBaseMenu", Remote: false}), updateCasbinMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateCasbin", Remote: false}), updateRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateRole", Remote: false}), updateSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateSwaDict", Remote: false}), updateSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateSwaDictail", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, addBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "AddBaseMenu", Remote: true}), addMenuRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "AddMenuRole", Remote: true}), changePasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "ChangePassword", Remote: true}), createApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateApi", Remote: true}), createBlacklistMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateBlacklist", Remote: true}), createRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateRole", Remote: true}), createSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaDict", Remote: true}), createSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaDictail", Remote: true}), createSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "CreateSwaRecord", Remote: true}), deleteApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteApi", Remote: true}), deleteApisByIdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteApisByIds", Remote: true}), deleteBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteBaseMenu", Remote: true}), deleteRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteRole", Remote: true}), deleteSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaDict", Remote: true}), deleteSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaDictail", Remote: true}), deleteSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaRecord", Remote: true}), deleteSwaRecordByIdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteSwaRecordByIds", Remote: true}), deleteUserMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "DeleteUser", Remote: true}), getAPIInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAPIInfoList", Remote: true}), getAllApisMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAllApis", Remote: true}), getAllBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetAllBaseMenu", Remote: true}), getApiByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetApiById", Remote: true}), getBaseMenuByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetBaseMenuById", Remote: true}), getInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetInfoList", Remote: true}), getMenuRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetMenuRole", Remote: true}), getPolicyPathByRoleIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetPolicyPathByRoleId", Remote: true}), getRoleInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleInfo", Remote: true}), getRoleListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleList", Remote: true}), getRoleMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetRoleMenu", Remote: true}), getServerInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetServerInfo", Remote: true}), getSubRolesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSubRoles", Remote: true}), getSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDict", Remote: true}), getSwaDictInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictInfoList", Remote: true}), getSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictail", Remote: true}), getSwaDictailInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaDictailInfoList", Remote: true}), getSwaRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaRecord", Remote: true}), getSwaRecordInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetSwaRecordInfoList", Remote: true}), getUserDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserDict", Remote: true}), getUserInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserInfo", Remote: true}), getUserInfoListMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "GetUserInfoList", Remote: true}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Login", Remote: true}), registerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Register", Remote: true}), resetPasswordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "ResetPassword", Remote: true}), scasbinMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "Scasbin", Remote: true}), setRoleMenusMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetRoleMenus", Remote: true}), setSelfInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetSelfInfo", Remote: true}), setSubRolesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetSubRoles", Remote: true}), setUserAuthoritiesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserAuthorities", Remote: true}), setUserInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserInfo", Remote: true}), setUserRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "SetUserRole", Remote: true}), updateApiMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateApi", Remote: true}), updateBaseMenuMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateBaseMenu", Remote: true}), updateCasbinMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateCasbin", Remote: true}), updateRoleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateRole", Remote: true}), updateSwaDictMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateSwaDict", Remote: true}), updateSwaDictailMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "demo1/service/service_swa/T", Method: "UpdateSwaDictail", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "⟦f74170a9:wEaVeReDgE:demo1/service/service_swa/T→demo1/service/service_config/T⟧\n",
	})
}

var _ weaver.InstanceOf[T] = (*impl)(nil)

var _ weaver.Unrouted = (*impl)(nil)


type t_local_stub struct {
	impl                         T
	tracer                       trace.Tracer
	addBaseMenuMetrics           *codegen.MethodMetrics
	addMenuRoleMetrics           *codegen.MethodMetrics
	changePasswordMetrics        *codegen.MethodMetrics
	createApiMetrics             *codegen.MethodMetrics
	createBlacklistMetrics       *codegen.MethodMetrics
	createRoleMetrics            *codegen.MethodMetrics
	createSwaDictMetrics         *codegen.MethodMetrics
	createSwaDictailMetrics      *codegen.MethodMetrics
	createSwaRecordMetrics       *codegen.MethodMetrics
	deleteApiMetrics             *codegen.MethodMetrics
	deleteApisByIdsMetrics       *codegen.MethodMetrics
	deleteBaseMenuMetrics        *codegen.MethodMetrics
	deleteRoleMetrics            *codegen.MethodMetrics
	deleteSwaDictMetrics         *codegen.MethodMetrics
	deleteSwaDictailMetrics      *codegen.MethodMetrics
	deleteSwaRecordMetrics       *codegen.MethodMetrics
	deleteSwaRecordByIdsMetrics  *codegen.MethodMetrics
	deleteUserMetrics            *codegen.MethodMetrics
	getAPIInfoListMetrics        *codegen.MethodMetrics
	getAllApisMetrics            *codegen.MethodMetrics
	getAllBaseMenuMetrics        *codegen.MethodMetrics
	getApiByIdMetrics            *codegen.MethodMetrics
	getBaseMenuByIdMetrics       *codegen.MethodMetrics
	getInfoListMetrics           *codegen.MethodMetrics
	getMenuRoleMetrics           *codegen.MethodMetrics
	getPolicyPathByRoleIdMetrics *codegen.MethodMetrics
	getRoleInfoMetrics           *codegen.MethodMetrics
	getRoleListMetrics           *codegen.MethodMetrics
	getRoleMenuMetrics           *codegen.MethodMetrics
	getServerInfoMetrics         *codegen.MethodMetrics
	getSubRolesMetrics           *codegen.MethodMetrics
	getSwaDictMetrics            *codegen.MethodMetrics
	getSwaDictInfoListMetrics    *codegen.MethodMetrics
	getSwaDictailMetrics         *codegen.MethodMetrics
	getSwaDictailInfoListMetrics *codegen.MethodMetrics
	getSwaRecordMetrics          *codegen.MethodMetrics
	getSwaRecordInfoListMetrics  *codegen.MethodMetrics
	getUserDictMetrics           *codegen.MethodMetrics
	getUserInfoMetrics           *codegen.MethodMetrics
	getUserInfoListMetrics       *codegen.MethodMetrics
	loginMetrics                 *codegen.MethodMetrics
	registerMetrics              *codegen.MethodMetrics
	resetPasswordMetrics         *codegen.MethodMetrics
	scasbinMetrics               *codegen.MethodMetrics
	setRoleMenusMetrics          *codegen.MethodMetrics
	setSelfInfoMetrics           *codegen.MethodMetrics
	setSubRolesMetrics           *codegen.MethodMetrics
	setUserAuthoritiesMetrics    *codegen.MethodMetrics
	setUserInfoMetrics           *codegen.MethodMetrics
	setUserRoleMetrics           *codegen.MethodMetrics
	updateApiMetrics             *codegen.MethodMetrics
	updateBaseMenuMetrics        *codegen.MethodMetrics
	updateCasbinMetrics          *codegen.MethodMetrics
	updateRoleMetrics            *codegen.MethodMetrics
	updateSwaDictMetrics         *codegen.MethodMetrics
	updateSwaDictailMetrics      *codegen.MethodMetrics
}

var _ T = (*t_local_stub)(nil)

func (s t_local_stub) AddBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	begin := s.addBaseMenuMetrics.Begin()
	defer func() { s.addBaseMenuMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.AddBaseMenu", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AddBaseMenu(ctx, a0)
}

func (s t_local_stub) AddMenuRole(ctx context.Context, a0 []SysBaseMenu, a1 uint) (err error) {
	begin := s.addMenuRoleMetrics.Begin()
	defer func() { s.addMenuRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.AddMenuRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AddMenuRole(ctx, a0, a1)
}

func (s t_local_stub) ChangePassword(ctx context.Context, a0 uint, a1 string, a2 string) (r0 *SysUser, err error) {
	begin := s.changePasswordMetrics.Begin()
	defer func() { s.changePasswordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.ChangePassword", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ChangePassword(ctx, a0, a1, a2)
}

func (s t_local_stub) CreateApi(ctx context.Context, a0 ModApi) (err error) {
	begin := s.createApiMetrics.Begin()
	defer func() { s.createApiMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateApi", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateApi(ctx, a0)
}

func (s t_local_stub) CreateBlacklist(ctx context.Context, a0 JwtBlacklist) (err error) {
	begin := s.createBlacklistMetrics.Begin()
	defer func() { s.createBlacklistMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateBlacklist", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateBlacklist(ctx, a0)
}

func (s t_local_stub) CreateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	begin := s.createRoleMetrics.Begin()
	defer func() { s.createRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateRole(ctx, a0)
}

func (s t_local_stub) CreateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	begin := s.createSwaDictMetrics.Begin()
	defer func() { s.createSwaDictMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateSwaDict", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateSwaDict(ctx, a0)
}

func (s t_local_stub) CreateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	begin := s.createSwaDictailMetrics.Begin()
	defer func() { s.createSwaDictailMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateSwaDictail", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateSwaDictail(ctx, a0)
}

func (s t_local_stub) CreateSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	begin := s.createSwaRecordMetrics.Begin()
	defer func() { s.createSwaRecordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.CreateSwaRecord", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.CreateSwaRecord(ctx, a0)
}

func (s t_local_stub) DeleteApi(ctx context.Context, a0 ModApi) (err error) {
	begin := s.deleteApiMetrics.Begin()
	defer func() { s.deleteApiMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteApi", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteApi(ctx, a0)
}

func (s t_local_stub) DeleteApisByIds(ctx context.Context, a0 []uint) (err error) {
	begin := s.deleteApisByIdsMetrics.Begin()
	defer func() { s.deleteApisByIdsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteApisByIds", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteApisByIds(ctx, a0)
}

func (s t_local_stub) DeleteBaseMenu(ctx context.Context, a0 int) (err error) {
	begin := s.deleteBaseMenuMetrics.Begin()
	defer func() { s.deleteBaseMenuMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteBaseMenu", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteBaseMenu(ctx, a0)
}

func (s t_local_stub) DeleteRole(ctx context.Context, a0 RoleUserMenu) (err error) {
	begin := s.deleteRoleMetrics.Begin()
	defer func() { s.deleteRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteRole(ctx, a0)
}

func (s t_local_stub) DeleteSwaDict(ctx context.Context, a0 ModDict) (err error) {
	begin := s.deleteSwaDictMetrics.Begin()
	defer func() { s.deleteSwaDictMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteSwaDict", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteSwaDict(ctx, a0)
}

func (s t_local_stub) DeleteSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	begin := s.deleteSwaDictailMetrics.Begin()
	defer func() { s.deleteSwaDictailMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteSwaDictail", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteSwaDictail(ctx, a0)
}

func (s t_local_stub) DeleteSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	begin := s.deleteSwaRecordMetrics.Begin()
	defer func() { s.deleteSwaRecordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteSwaRecord", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteSwaRecord(ctx, a0)
}

func (s t_local_stub) DeleteSwaRecordByIds(ctx context.Context, a0 []uint) (err error) {
	begin := s.deleteSwaRecordByIdsMetrics.Begin()
	defer func() { s.deleteSwaRecordByIdsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteSwaRecordByIds", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteSwaRecordByIds(ctx, a0)
}

func (s t_local_stub) DeleteUser(ctx context.Context, a0 int) (err error) {
	begin := s.deleteUserMetrics.Begin()
	defer func() { s.deleteUserMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.DeleteUser", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.DeleteUser(ctx, a0)
}

func (s t_local_stub) GetAPIInfoList(ctx context.Context, a0 ModApi, a1 PageInfo, a2 string, a3 bool) (r0 []ModApi, r1 int64, err error) {
	begin := s.getAPIInfoListMetrics.Begin()
	defer func() { s.getAPIInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetAPIInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAPIInfoList(ctx, a0, a1, a2, a3)
}

func (s t_local_stub) GetAllApis(ctx context.Context) (r0 []ModApi, err error) {
	begin := s.getAllApisMetrics.Begin()
	defer func() { s.getAllApisMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetAllApis", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAllApis(ctx)
}

func (s t_local_stub) GetAllBaseMenu(ctx context.Context) (r0 []SysBaseMenu, err error) {
	begin := s.getAllBaseMenuMetrics.Begin()
	defer func() { s.getAllBaseMenuMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetAllBaseMenu", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAllBaseMenu(ctx)
}

func (s t_local_stub) GetApiById(ctx context.Context, a0 int) (r0 ModApi, err error) {
	begin := s.getApiByIdMetrics.Begin()
	defer func() { s.getApiByIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetApiById", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetApiById(ctx, a0)
}

func (s t_local_stub) GetBaseMenuById(ctx context.Context, a0 int) (r0 SysBaseMenu, err error) {
	begin := s.getBaseMenuByIdMetrics.Begin()
	defer func() { s.getBaseMenuByIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetBaseMenuById", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetBaseMenuById(ctx, a0)
}

func (s t_local_stub) GetInfoList(ctx context.Context) (r0 []SysBaseMenu, r1 int64, err error) {
	begin := s.getInfoListMetrics.Begin()
	defer func() { s.getInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetInfoList(ctx)
}

func (s t_local_stub) GetMenuRole(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	begin := s.getMenuRoleMetrics.Begin()
	defer func() { s.getMenuRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetMenuRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetMenuRole(ctx, a0)
}

func (s t_local_stub) GetPolicyPathByRoleId(ctx context.Context, a0 uint) (r0 []CasbinInfo, err error) {
	begin := s.getPolicyPathByRoleIdMetrics.Begin()
	defer func() { s.getPolicyPathByRoleIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetPolicyPathByRoleId", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetPolicyPathByRoleId(ctx, a0)
}

func (s t_local_stub) GetRoleInfo(ctx context.Context, a0 uint) (r0 SwaRole, err error) {
	begin := s.getRoleInfoMetrics.Begin()
	defer func() { s.getRoleInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetRoleInfo", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetRoleInfo(ctx, a0)
}

func (s t_local_stub) GetRoleList(ctx context.Context, a0 PageInfo) (r0 []SwaRole, r1 int64, err error) {
	begin := s.getRoleListMetrics.Begin()
	defer func() { s.getRoleListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetRoleList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetRoleList(ctx, a0)
}

func (s t_local_stub) GetRoleMenu(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	begin := s.getRoleMenuMetrics.Begin()
	defer func() { s.getRoleMenuMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetRoleMenu", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetRoleMenu(ctx, a0)
}

func (s t_local_stub) GetServerInfo(ctx context.Context) (r0 Server, err error) {
	begin := s.getServerInfoMetrics.Begin()
	defer func() { s.getServerInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetServerInfo", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetServerInfo(ctx)
}

func (s t_local_stub) GetSubRoles(ctx context.Context, a0 uint) (r0 []SwaRole, err error) {
	begin := s.getSubRolesMetrics.Begin()
	defer func() { s.getSubRolesMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSubRoles", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSubRoles(ctx, a0)
}

func (s t_local_stub) GetSwaDict(ctx context.Context, a0 string, a1 uint, a2 *bool) (r0 ModDict, err error) {
	begin := s.getSwaDictMetrics.Begin()
	defer func() { s.getSwaDictMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaDict", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaDict(ctx, a0, a1, a2)
}

func (s t_local_stub) GetSwaDictInfoList(ctx context.Context, a0 SwaDictSearch) (r0 []ModDict, r1 int64, err error) {
	begin := s.getSwaDictInfoListMetrics.Begin()
	defer func() { s.getSwaDictInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaDictInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaDictInfoList(ctx, a0)
}

func (s t_local_stub) GetSwaDictail(ctx context.Context, a0 uint) (r0 ModDictail, err error) {
	begin := s.getSwaDictailMetrics.Begin()
	defer func() { s.getSwaDictailMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaDictail", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaDictail(ctx, a0)
}

func (s t_local_stub) GetSwaDictailInfoList(ctx context.Context, a0 SwaDictailSearch) (r0 []ModDictail, r1 int64, err error) {
	begin := s.getSwaDictailInfoListMetrics.Begin()
	defer func() { s.getSwaDictailInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaDictailInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaDictailInfoList(ctx, a0)
}

func (s t_local_stub) GetSwaRecord(ctx context.Context, a0 uint) (r0 SwaRecord, err error) {
	begin := s.getSwaRecordMetrics.Begin()
	defer func() { s.getSwaRecordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaRecord", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaRecord(ctx, a0)
}

func (s t_local_stub) GetSwaRecordInfoList(ctx context.Context, a0 SwaRecordSearch) (r0 []SwaRecord, r1 int64, err error) {
	begin := s.getSwaRecordInfoListMetrics.Begin()
	defer func() { s.getSwaRecordInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetSwaRecordInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSwaRecordInfoList(ctx, a0)
}

func (s t_local_stub) GetUserDict(ctx context.Context) (r0 []SysUser, err error) {
	begin := s.getUserDictMetrics.Begin()
	defer func() { s.getUserDictMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetUserDict", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetUserDict(ctx)
}

func (s t_local_stub) GetUserInfo(ctx context.Context, a0 uuid.UUID) (r0 SysUser, err error) {
	begin := s.getUserInfoMetrics.Begin()
	defer func() { s.getUserInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetUserInfo", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetUserInfo(ctx, a0)
}

func (s t_local_stub) GetUserInfoList(ctx context.Context, a0 PageInfo) (r0 []SysUser, r1 int64, err error) {
	begin := s.getUserInfoListMetrics.Begin()
	defer func() { s.getUserInfoListMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.GetUserInfoList", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetUserInfoList(ctx, a0)
}

func (s t_local_stub) Login(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.Login", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Login(ctx, a0)
}

func (s t_local_stub) Register(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	begin := s.registerMetrics.Begin()
	defer func() { s.registerMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.Register", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Register(ctx, a0)
}

func (s t_local_stub) ResetPassword(ctx context.Context, a0 uint) (err error) {
	begin := s.resetPasswordMetrics.Begin()
	defer func() { s.resetPasswordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.ResetPassword", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ResetPassword(ctx, a0)
}

func (s t_local_stub) Scasbin(ctx context.Context, a0 string, a1 string, a2 string) (r0 bool, err error) {
	begin := s.scasbinMetrics.Begin()
	defer func() { s.scasbinMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.Scasbin", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Scasbin(ctx, a0, a1, a2)
}

func (s t_local_stub) SetRoleMenus(ctx context.Context, a0 SwaRole) (err error) {
	begin := s.setRoleMenusMetrics.Begin()
	defer func() { s.setRoleMenusMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetRoleMenus", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetRoleMenus(ctx, a0)
}

func (s t_local_stub) SetSelfInfo(ctx context.Context, a0 SysUser) (err error) {
	begin := s.setSelfInfoMetrics.Begin()
	defer func() { s.setSelfInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetSelfInfo", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetSelfInfo(ctx, a0)
}

func (s t_local_stub) SetSubRoles(ctx context.Context, a0 SwaRole) (err error) {
	begin := s.setSubRolesMetrics.Begin()
	defer func() { s.setSubRolesMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetSubRoles", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetSubRoles(ctx, a0)
}

func (s t_local_stub) SetUserAuthorities(ctx context.Context, a0 uint, a1 []uint) (err error) {
	begin := s.setUserAuthoritiesMetrics.Begin()
	defer func() { s.setUserAuthoritiesMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetUserAuthorities", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetUserAuthorities(ctx, a0, a1)
}

func (s t_local_stub) SetUserInfo(ctx context.Context, a0 SysUser) (err error) {
	begin := s.setUserInfoMetrics.Begin()
	defer func() { s.setUserInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetUserInfo", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetUserInfo(ctx, a0)
}

func (s t_local_stub) SetUserRole(ctx context.Context, a0 uint, a1 uint) (err error) {
	begin := s.setUserRoleMetrics.Begin()
	defer func() { s.setUserRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.SetUserRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SetUserRole(ctx, a0, a1)
}

func (s t_local_stub) UpdateApi(ctx context.Context, a0 ModApi) (err error) {
	begin := s.updateApiMetrics.Begin()
	defer func() { s.updateApiMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateApi", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateApi(ctx, a0)
}

func (s t_local_stub) UpdateBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	begin := s.updateBaseMenuMetrics.Begin()
	defer func() { s.updateBaseMenuMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateBaseMenu", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateBaseMenu(ctx, a0)
}

func (s t_local_stub) UpdateCasbin(ctx context.Context, a0 uint, a1 []CasbinInfo) (err error) {
	begin := s.updateCasbinMetrics.Begin()
	defer func() { s.updateCasbinMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateCasbin", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateCasbin(ctx, a0, a1)
}

func (s t_local_stub) UpdateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	begin := s.updateRoleMetrics.Begin()
	defer func() { s.updateRoleMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateRole", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateRole(ctx, a0)
}

func (s t_local_stub) UpdateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	begin := s.updateSwaDictMetrics.Begin()
	defer func() { s.updateSwaDictMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateSwaDict", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateSwaDict(ctx, a0)
}

func (s t_local_stub) UpdateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	begin := s.updateSwaDictailMetrics.Begin()
	defer func() { s.updateSwaDictailMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "service_swa.T.UpdateSwaDictail", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UpdateSwaDictail(ctx, a0)
}


type t_client_stub struct {
	stub                         codegen.Stub
	addBaseMenuMetrics           *codegen.MethodMetrics
	addMenuRoleMetrics           *codegen.MethodMetrics
	changePasswordMetrics        *codegen.MethodMetrics
	createApiMetrics             *codegen.MethodMetrics
	createBlacklistMetrics       *codegen.MethodMetrics
	createRoleMetrics            *codegen.MethodMetrics
	createSwaDictMetrics         *codegen.MethodMetrics
	createSwaDictailMetrics      *codegen.MethodMetrics
	createSwaRecordMetrics       *codegen.MethodMetrics
	deleteApiMetrics             *codegen.MethodMetrics
	deleteApisByIdsMetrics       *codegen.MethodMetrics
	deleteBaseMenuMetrics        *codegen.MethodMetrics
	deleteRoleMetrics            *codegen.MethodMetrics
	deleteSwaDictMetrics         *codegen.MethodMetrics
	deleteSwaDictailMetrics      *codegen.MethodMetrics
	deleteSwaRecordMetrics       *codegen.MethodMetrics
	deleteSwaRecordByIdsMetrics  *codegen.MethodMetrics
	deleteUserMetrics            *codegen.MethodMetrics
	getAPIInfoListMetrics        *codegen.MethodMetrics
	getAllApisMetrics            *codegen.MethodMetrics
	getAllBaseMenuMetrics        *codegen.MethodMetrics
	getApiByIdMetrics            *codegen.MethodMetrics
	getBaseMenuByIdMetrics       *codegen.MethodMetrics
	getInfoListMetrics           *codegen.MethodMetrics
	getMenuRoleMetrics           *codegen.MethodMetrics
	getPolicyPathByRoleIdMetrics *codegen.MethodMetrics
	getRoleInfoMetrics           *codegen.MethodMetrics
	getRoleListMetrics           *codegen.MethodMetrics
	getRoleMenuMetrics           *codegen.MethodMetrics
	getServerInfoMetrics         *codegen.MethodMetrics
	getSubRolesMetrics           *codegen.MethodMetrics
	getSwaDictMetrics            *codegen.MethodMetrics
	getSwaDictInfoListMetrics    *codegen.MethodMetrics
	getSwaDictailMetrics         *codegen.MethodMetrics
	getSwaDictailInfoListMetrics *codegen.MethodMetrics
	getSwaRecordMetrics          *codegen.MethodMetrics
	getSwaRecordInfoListMetrics  *codegen.MethodMetrics
	getUserDictMetrics           *codegen.MethodMetrics
	getUserInfoMetrics           *codegen.MethodMetrics
	getUserInfoListMetrics       *codegen.MethodMetrics
	loginMetrics                 *codegen.MethodMetrics
	registerMetrics              *codegen.MethodMetrics
	resetPasswordMetrics         *codegen.MethodMetrics
	scasbinMetrics               *codegen.MethodMetrics
	setRoleMenusMetrics          *codegen.MethodMetrics
	setSelfInfoMetrics           *codegen.MethodMetrics
	setSubRolesMetrics           *codegen.MethodMetrics
	setUserAuthoritiesMetrics    *codegen.MethodMetrics
	setUserInfoMetrics           *codegen.MethodMetrics
	setUserRoleMetrics           *codegen.MethodMetrics
	updateApiMetrics             *codegen.MethodMetrics
	updateBaseMenuMetrics        *codegen.MethodMetrics
	updateCasbinMetrics          *codegen.MethodMetrics
	updateRoleMetrics            *codegen.MethodMetrics
	updateSwaDictMetrics         *codegen.MethodMetrics
	updateSwaDictailMetrics      *codegen.MethodMetrics
}

var _ T = (*t_client_stub)(nil)

func (s t_client_stub) AddBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	var requestBytes, replyBytes int
	begin := s.addBaseMenuMetrics.Begin()
	defer func() { s.addBaseMenuMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.AddBaseMenu", trace.WithSpanKind(trace.SpanKindClient))
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

func (s t_client_stub) AddMenuRole(ctx context.Context, a0 []SysBaseMenu, a1 uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.addMenuRoleMetrics.Begin()
	defer func() { s.addMenuRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.AddMenuRole", trace.WithSpanKind(trace.SpanKindClient))
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
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, a0)
	enc.Uint(a1)
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

func (s t_client_stub) ChangePassword(ctx context.Context, a0 uint, a1 string, a2 string) (r0 *SysUser, err error) {
	var requestBytes, replyBytes int
	begin := s.changePasswordMetrics.Begin()
	defer func() { s.changePasswordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.ChangePassword", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + len(a1))
	size += (4 + len(a2))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.Uint(a0)
	enc.String(a1)
	enc.String(a2)
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
	r0 = serviceweaver_dec_ptr_SysUser_342d2fab(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) CreateApi(ctx context.Context, a0 ModApi) (err error) {
	var requestBytes, replyBytes int
	begin := s.createApiMetrics.Begin()
	defer func() { s.createApiMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateApi", trace.WithSpanKind(trace.SpanKindClient))
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

func (s t_client_stub) CreateBlacklist(ctx context.Context, a0 JwtBlacklist) (err error) {
	var requestBytes, replyBytes int
	begin := s.createBlacklistMetrics.Begin()
	defer func() { s.createBlacklistMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateBlacklist", trace.WithSpanKind(trace.SpanKindClient))
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

func (s t_client_stub) CreateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	var requestBytes, replyBytes int
	begin := s.createRoleMetrics.Begin()
	defer func() { s.createRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateRole", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 5, enc.Data(), shardKey)
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

func (s t_client_stub) CreateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	var requestBytes, replyBytes int
	begin := s.createSwaDictMetrics.Begin()
	defer func() { s.createSwaDictMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateSwaDict", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 6, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) CreateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	var requestBytes, replyBytes int
	begin := s.createSwaDictailMetrics.Begin()
	defer func() { s.createSwaDictailMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateSwaDictail", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 7, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) CreateSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	var requestBytes, replyBytes int
	begin := s.createSwaRecordMetrics.Begin()
	defer func() { s.createSwaRecordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.CreateSwaRecord", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 8, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteApi(ctx context.Context, a0 ModApi) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteApiMetrics.Begin()
	defer func() { s.deleteApiMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteApi", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 9, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteApisByIds(ctx context.Context, a0 []uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteApisByIdsMetrics.Begin()
	defer func() { s.deleteApisByIdsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteApisByIds", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + (len(a0) * 8))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	serviceweaver_enc_slice_uint_e3009941(enc, a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 10, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteBaseMenu(ctx context.Context, a0 int) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteBaseMenuMetrics.Begin()
	defer func() { s.deleteBaseMenuMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteBaseMenu", trace.WithSpanKind(trace.SpanKindClient))
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

	enc.Int(a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 11, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteRole(ctx context.Context, a0 RoleUserMenu) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteRoleMetrics.Begin()
	defer func() { s.deleteRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteRole", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 12, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteSwaDict(ctx context.Context, a0 ModDict) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteSwaDictMetrics.Begin()
	defer func() { s.deleteSwaDictMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteSwaDict", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 13, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteSwaDictailMetrics.Begin()
	defer func() { s.deleteSwaDictailMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteSwaDictail", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 14, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteSwaRecordMetrics.Begin()
	defer func() { s.deleteSwaRecordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteSwaRecord", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 15, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteSwaRecordByIds(ctx context.Context, a0 []uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteSwaRecordByIdsMetrics.Begin()
	defer func() { s.deleteSwaRecordByIdsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteSwaRecordByIds", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + (len(a0) * 8))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	serviceweaver_enc_slice_uint_e3009941(enc, a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 16, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) DeleteUser(ctx context.Context, a0 int) (err error) {
	var requestBytes, replyBytes int
	begin := s.deleteUserMetrics.Begin()
	defer func() { s.deleteUserMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.DeleteUser", trace.WithSpanKind(trace.SpanKindClient))
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

	enc.Int(a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 17, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) GetAPIInfoList(ctx context.Context, a0 ModApi, a1 PageInfo, a2 string, a3 bool) (r0 []ModApi, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getAPIInfoListMetrics.Begin()
	defer func() { s.getAPIInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetAPIInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	(a1).WeaverMarshal(enc)
	enc.String(a2)
	enc.Bool(a3)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 18, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_ModApi_d17cefdc(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetAllApis(ctx context.Context) (r0 []ModApi, err error) {
	var requestBytes, replyBytes int
	begin := s.getAllApisMetrics.Begin()
	defer func() { s.getAllApisMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetAllApis", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 19, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_ModApi_d17cefdc(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetAllBaseMenu(ctx context.Context) (r0 []SysBaseMenu, err error) {
	var requestBytes, replyBytes int
	begin := s.getAllBaseMenuMetrics.Begin()
	defer func() { s.getAllBaseMenuMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetAllBaseMenu", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 20, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetApiById(ctx context.Context, a0 int) (r0 ModApi, err error) {
	var requestBytes, replyBytes int
	begin := s.getApiByIdMetrics.Begin()
	defer func() { s.getApiByIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetApiById", trace.WithSpanKind(trace.SpanKindClient))
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

	enc.Int(a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 21, enc.Data(), shardKey)
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

func (s t_client_stub) GetBaseMenuById(ctx context.Context, a0 int) (r0 SysBaseMenu, err error) {
	var requestBytes, replyBytes int
	begin := s.getBaseMenuByIdMetrics.Begin()
	defer func() { s.getBaseMenuByIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetBaseMenuById", trace.WithSpanKind(trace.SpanKindClient))
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

	enc.Int(a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 22, enc.Data(), shardKey)
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

func (s t_client_stub) GetInfoList(ctx context.Context) (r0 []SysBaseMenu, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getInfoListMetrics.Begin()
	defer func() { s.getInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 23, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetMenuRole(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	var requestBytes, replyBytes int
	begin := s.getMenuRoleMetrics.Begin()
	defer func() { s.getMenuRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetMenuRole", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 24, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetPolicyPathByRoleId(ctx context.Context, a0 uint) (r0 []CasbinInfo, err error) {
	var requestBytes, replyBytes int
	begin := s.getPolicyPathByRoleIdMetrics.Begin()
	defer func() { s.getPolicyPathByRoleIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetPolicyPathByRoleId", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 25, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_CasbinInfo_0ebf7d6e(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetRoleInfo(ctx context.Context, a0 uint) (r0 SwaRole, err error) {
	var requestBytes, replyBytes int
	begin := s.getRoleInfoMetrics.Begin()
	defer func() { s.getRoleInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetRoleInfo", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 26, enc.Data(), shardKey)
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

func (s t_client_stub) GetRoleList(ctx context.Context, a0 PageInfo) (r0 []SwaRole, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getRoleListMetrics.Begin()
	defer func() { s.getRoleListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetRoleList", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_PageInfo_c29f03f8(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 27, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SwaRole_8aab416c(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetRoleMenu(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	var requestBytes, replyBytes int
	begin := s.getRoleMenuMetrics.Begin()
	defer func() { s.getRoleMenuMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetRoleMenu", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 28, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetServerInfo(ctx context.Context) (r0 Server, err error) {
	var requestBytes, replyBytes int
	begin := s.getServerInfoMetrics.Begin()
	defer func() { s.getServerInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetServerInfo", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 29, nil, shardKey)
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

func (s t_client_stub) GetSubRoles(ctx context.Context, a0 uint) (r0 []SwaRole, err error) {
	var requestBytes, replyBytes int
	begin := s.getSubRolesMetrics.Begin()
	defer func() { s.getSubRolesMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSubRoles", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 30, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SwaRole_8aab416c(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetSwaDict(ctx context.Context, a0 string, a1 uint, a2 *bool) (r0 ModDict, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaDictMetrics.Begin()
	defer func() { s.getSwaDictMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaDict", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + len(a0))
	size += 8
	size += serviceweaver_size_ptr_bool_31f02903(a2)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.String(a0)
	enc.Uint(a1)
	serviceweaver_enc_ptr_bool_31f02903(enc, a2)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 31, enc.Data(), shardKey)
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

func (s t_client_stub) GetSwaDictInfoList(ctx context.Context, a0 SwaDictSearch) (r0 []ModDict, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaDictInfoListMetrics.Begin()
	defer func() { s.getSwaDictInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaDictInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 32, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_ModDict_3a395239(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetSwaDictail(ctx context.Context, a0 uint) (r0 ModDictail, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaDictailMetrics.Begin()
	defer func() { s.getSwaDictailMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaDictail", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 33, enc.Data(), shardKey)
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

func (s t_client_stub) GetSwaDictailInfoList(ctx context.Context, a0 SwaDictailSearch) (r0 []ModDictail, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaDictailInfoListMetrics.Begin()
	defer func() { s.getSwaDictailInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaDictailInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 34, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_ModDictail_42b2abbf(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetSwaRecord(ctx context.Context, a0 uint) (r0 SwaRecord, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaRecordMetrics.Begin()
	defer func() { s.getSwaRecordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaRecord", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 35, enc.Data(), shardKey)
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

func (s t_client_stub) GetSwaRecordInfoList(ctx context.Context, a0 SwaRecordSearch) (r0 []SwaRecord, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getSwaRecordInfoListMetrics.Begin()
	defer func() { s.getSwaRecordInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetSwaRecordInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 36, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SwaRecord_bba98daa(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) GetUserDict(ctx context.Context) (r0 []SysUser, err error) {
	var requestBytes, replyBytes int
	begin := s.getUserDictMetrics.Begin()
	defer func() { s.getUserDictMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetUserDict", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 37, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysUser_00eb59d9(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) GetUserInfo(ctx context.Context, a0 uuid.UUID) (r0 SysUser, err error) {
	var requestBytes, replyBytes int
	begin := s.getUserInfoMetrics.Begin()
	defer func() { s.getUserInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetUserInfo", trace.WithSpanKind(trace.SpanKindClient))
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
	enc.EncodeBinaryMarshaler(&a0)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 38, enc.Data(), shardKey)
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

func (s t_client_stub) GetUserInfoList(ctx context.Context, a0 PageInfo) (r0 []SysUser, r1 int64, err error) {
	var requestBytes, replyBytes int
	begin := s.getUserInfoListMetrics.Begin()
	defer func() { s.getUserInfoListMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.GetUserInfoList", trace.WithSpanKind(trace.SpanKindClient))
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
	size += serviceweaver_size_PageInfo_c29f03f8(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	(a0).WeaverMarshal(enc)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 39, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_SysUser_00eb59d9(dec)
	r1 = dec.Int64()
	err = dec.Error()
	return
}

func (s t_client_stub) Login(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	var requestBytes, replyBytes int
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.Login", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 40, enc.Data(), shardKey)
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

func (s t_client_stub) Register(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	var requestBytes, replyBytes int
	begin := s.registerMetrics.Begin()
	defer func() { s.registerMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.Register", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 41, enc.Data(), shardKey)
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

func (s t_client_stub) ResetPassword(ctx context.Context, a0 uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.resetPasswordMetrics.Begin()
	defer func() { s.resetPasswordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.ResetPassword", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 42, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) Scasbin(ctx context.Context, a0 string, a1 string, a2 string) (r0 bool, err error) {
	var requestBytes, replyBytes int
	begin := s.scasbinMetrics.Begin()
	defer func() { s.scasbinMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.Scasbin", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + len(a0))
	size += (4 + len(a1))
	size += (4 + len(a2))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.String(a0)
	enc.String(a1)
	enc.String(a2)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 43, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

func (s t_client_stub) SetRoleMenus(ctx context.Context, a0 SwaRole) (err error) {
	var requestBytes, replyBytes int
	begin := s.setRoleMenusMetrics.Begin()
	defer func() { s.setRoleMenusMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetRoleMenus", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 44, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) SetSelfInfo(ctx context.Context, a0 SysUser) (err error) {
	var requestBytes, replyBytes int
	begin := s.setSelfInfoMetrics.Begin()
	defer func() { s.setSelfInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetSelfInfo", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 45, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) SetSubRoles(ctx context.Context, a0 SwaRole) (err error) {
	var requestBytes, replyBytes int
	begin := s.setSubRolesMetrics.Begin()
	defer func() { s.setSubRolesMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetSubRoles", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 46, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) SetUserAuthorities(ctx context.Context, a0 uint, a1 []uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.setUserAuthoritiesMetrics.Begin()
	defer func() { s.setUserAuthoritiesMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetUserAuthorities", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + (len(a1) * 8))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.Uint(a0)
	serviceweaver_enc_slice_uint_e3009941(enc, a1)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 47, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) SetUserInfo(ctx context.Context, a0 SysUser) (err error) {
	var requestBytes, replyBytes int
	begin := s.setUserInfoMetrics.Begin()
	defer func() { s.setUserInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetUserInfo", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 48, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) SetUserRole(ctx context.Context, a0 uint, a1 uint) (err error) {
	var requestBytes, replyBytes int
	begin := s.setUserRoleMetrics.Begin()
	defer func() { s.setUserRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.SetUserRole", trace.WithSpanKind(trace.SpanKindClient))
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
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	enc.Uint(a0)
	enc.Uint(a1)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 49, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) UpdateApi(ctx context.Context, a0 ModApi) (err error) {
	var requestBytes, replyBytes int
	begin := s.updateApiMetrics.Begin()
	defer func() { s.updateApiMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateApi", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 50, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) UpdateBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	var requestBytes, replyBytes int
	begin := s.updateBaseMenuMetrics.Begin()
	defer func() { s.updateBaseMenuMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateBaseMenu", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 51, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) UpdateCasbin(ctx context.Context, a0 uint, a1 []CasbinInfo) (err error) {
	var requestBytes, replyBytes int
	begin := s.updateCasbinMetrics.Begin()
	defer func() { s.updateCasbinMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateCasbin", trace.WithSpanKind(trace.SpanKindClient))
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
	enc.Uint(a0)
	serviceweaver_enc_slice_CasbinInfo_0ebf7d6e(enc, a1)
	var shardKey uint64

	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 52, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) UpdateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	var requestBytes, replyBytes int
	begin := s.updateRoleMetrics.Begin()
	defer func() { s.updateRoleMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateRole", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 53, enc.Data(), shardKey)
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

func (s t_client_stub) UpdateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	var requestBytes, replyBytes int
	begin := s.updateSwaDictMetrics.Begin()
	defer func() { s.updateSwaDictMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateSwaDict", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 54, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) UpdateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	var requestBytes, replyBytes int
	begin := s.updateSwaDictailMetrics.Begin()
	defer func() { s.updateSwaDictailMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.stub.Tracer().Start(ctx, "service_swa.T.UpdateSwaDictail", trace.WithSpanKind(trace.SpanKindClient))
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
	results, err = s.stub.Run(ctx, 55, enc.Data(), shardKey)
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
	case "AddBaseMenu":
		return s.addBaseMenu
	case "AddMenuRole":
		return s.addMenuRole
	case "ChangePassword":
		return s.changePassword
	case "CreateApi":
		return s.createApi
	case "CreateBlacklist":
		return s.createBlacklist
	case "CreateRole":
		return s.createRole
	case "CreateSwaDict":
		return s.createSwaDict
	case "CreateSwaDictail":
		return s.createSwaDictail
	case "CreateSwaRecord":
		return s.createSwaRecord
	case "DeleteApi":
		return s.deleteApi
	case "DeleteApisByIds":
		return s.deleteApisByIds
	case "DeleteBaseMenu":
		return s.deleteBaseMenu
	case "DeleteRole":
		return s.deleteRole
	case "DeleteSwaDict":
		return s.deleteSwaDict
	case "DeleteSwaDictail":
		return s.deleteSwaDictail
	case "DeleteSwaRecord":
		return s.deleteSwaRecord
	case "DeleteSwaRecordByIds":
		return s.deleteSwaRecordByIds
	case "DeleteUser":
		return s.deleteUser
	case "GetAPIInfoList":
		return s.getAPIInfoList
	case "GetAllApis":
		return s.getAllApis
	case "GetAllBaseMenu":
		return s.getAllBaseMenu
	case "GetApiById":
		return s.getApiById
	case "GetBaseMenuById":
		return s.getBaseMenuById
	case "GetInfoList":
		return s.getInfoList
	case "GetMenuRole":
		return s.getMenuRole
	case "GetPolicyPathByRoleId":
		return s.getPolicyPathByRoleId
	case "GetRoleInfo":
		return s.getRoleInfo
	case "GetRoleList":
		return s.getRoleList
	case "GetRoleMenu":
		return s.getRoleMenu
	case "GetServerInfo":
		return s.getServerInfo
	case "GetSubRoles":
		return s.getSubRoles
	case "GetSwaDict":
		return s.getSwaDict
	case "GetSwaDictInfoList":
		return s.getSwaDictInfoList
	case "GetSwaDictail":
		return s.getSwaDictail
	case "GetSwaDictailInfoList":
		return s.getSwaDictailInfoList
	case "GetSwaRecord":
		return s.getSwaRecord
	case "GetSwaRecordInfoList":
		return s.getSwaRecordInfoList
	case "GetUserDict":
		return s.getUserDict
	case "GetUserInfo":
		return s.getUserInfo
	case "GetUserInfoList":
		return s.getUserInfoList
	case "Login":
		return s.login
	case "Register":
		return s.register
	case "ResetPassword":
		return s.resetPassword
	case "Scasbin":
		return s.scasbin
	case "SetRoleMenus":
		return s.setRoleMenus
	case "SetSelfInfo":
		return s.setSelfInfo
	case "SetSubRoles":
		return s.setSubRoles
	case "SetUserAuthorities":
		return s.setUserAuthorities
	case "SetUserInfo":
		return s.setUserInfo
	case "SetUserRole":
		return s.setUserRole
	case "UpdateApi":
		return s.updateApi
	case "UpdateBaseMenu":
		return s.updateBaseMenu
	case "UpdateCasbin":
		return s.updateCasbin
	case "UpdateRole":
		return s.updateRole
	case "UpdateSwaDict":
		return s.updateSwaDict
	case "UpdateSwaDictail":
		return s.updateSwaDictail
	default:
		return nil
	}
}

func (s t_server_stub) addBaseMenu(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysBaseMenu
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.AddBaseMenu(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) addMenuRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 []SysBaseMenu
	a0 = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
	var a1 uint
	a1 = dec.Uint()

	appErr := s.impl.AddMenuRole(ctx, a0, a1)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) changePassword(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()
	var a1 string
	a1 = dec.String()
	var a2 string
	a2 = dec.String()

	r0, appErr := s.impl.ChangePassword(ctx, a0, a1, a2)

	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_SysUser_342d2fab(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createApi(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModApi
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.CreateApi(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createBlacklist(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 JwtBlacklist
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.CreateBlacklist(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRole
	(&a0).WeaverUnmarshal(dec)

	r0, appErr := s.impl.CreateRole(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createSwaDict(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDict
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.CreateSwaDict(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createSwaDictail(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDictail
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.CreateSwaDictail(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) createSwaRecord(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRecord
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.CreateSwaRecord(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteApi(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModApi
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteApi(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteApisByIds(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 []uint
	a0 = serviceweaver_dec_slice_uint_e3009941(dec)

	appErr := s.impl.DeleteApisByIds(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteBaseMenu(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	appErr := s.impl.DeleteBaseMenu(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 RoleUserMenu
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteRole(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteSwaDict(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDict
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteSwaDict(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteSwaDictail(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDictail
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteSwaDictail(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteSwaRecord(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRecord
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.DeleteSwaRecord(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteSwaRecordByIds(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 []uint
	a0 = serviceweaver_dec_slice_uint_e3009941(dec)

	appErr := s.impl.DeleteSwaRecordByIds(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) deleteUser(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	appErr := s.impl.DeleteUser(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getAPIInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModApi
	(&a0).WeaverUnmarshal(dec)
	var a1 PageInfo
	(&a1).WeaverUnmarshal(dec)
	var a2 string
	a2 = dec.String()
	var a3 bool
	a3 = dec.Bool()

	r0, r1, appErr := s.impl.GetAPIInfoList(ctx, a0, a1, a2, a3)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_ModApi_d17cefdc(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getAllApis(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetAllApis(ctx)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_ModApi_d17cefdc(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getAllBaseMenu(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetAllBaseMenu(ctx)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getApiById(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	r0, appErr := s.impl.GetApiById(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getBaseMenuById(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	r0, appErr := s.impl.GetBaseMenuById(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, r1, appErr := s.impl.GetInfoList(ctx)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getMenuRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetMenuRole(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getPolicyPathByRoleId(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetPolicyPathByRoleId(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_CasbinInfo_0ebf7d6e(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getRoleInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetRoleInfo(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getRoleList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 PageInfo
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetRoleList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SwaRole_8aab416c(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getRoleMenu(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetRoleMenu(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getServerInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetServerInfo(ctx)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSubRoles(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetSubRoles(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SwaRole_8aab416c(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaDict(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 uint
	a1 = dec.Uint()
	var a2 *bool
	a2 = serviceweaver_dec_ptr_bool_31f02903(dec)

	r0, appErr := s.impl.GetSwaDict(ctx, a0, a1, a2)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaDictInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaDictSearch
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetSwaDictInfoList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_ModDict_3a395239(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaDictail(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetSwaDictail(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaDictailInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaDictailSearch
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetSwaDictailInfoList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_ModDictail_42b2abbf(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaRecord(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	r0, appErr := s.impl.GetSwaRecord(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getSwaRecordInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRecordSearch
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetSwaRecordInfoList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SwaRecord_bba98daa(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getUserDict(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	r0, appErr := s.impl.GetUserDict(ctx)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysUser_00eb59d9(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getUserInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uuid.UUID
	dec.DecodeBinaryUnmarshaler(&a0)

	r0, appErr := s.impl.GetUserInfo(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getUserInfoList(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 PageInfo
	(&a0).WeaverUnmarshal(dec)

	r0, r1, appErr := s.impl.GetUserInfoList(ctx, a0)

	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_SysUser_00eb59d9(enc, r0)
	enc.Int64(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) login(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysUser
	(&a0).WeaverUnmarshal(dec)

	r0, appErr := s.impl.Login(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) register(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysUser
	(&a0).WeaverUnmarshal(dec)

	r0, appErr := s.impl.Register(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) resetPassword(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()

	appErr := s.impl.ResetPassword(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) scasbin(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()
	var a2 string
	a2 = dec.String()

	r0, appErr := s.impl.Scasbin(ctx, a0, a1, a2)

	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setRoleMenus(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRole
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetRoleMenus(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setSelfInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysUser
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetSelfInfo(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setSubRoles(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRole
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetSubRoles(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setUserAuthorities(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()
	var a1 []uint
	a1 = serviceweaver_dec_slice_uint_e3009941(dec)

	appErr := s.impl.SetUserAuthorities(ctx, a0, a1)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setUserInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysUser
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.SetUserInfo(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) setUserRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()
	var a1 uint
	a1 = dec.Uint()

	appErr := s.impl.SetUserRole(ctx, a0, a1)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateApi(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModApi
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.UpdateApi(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateBaseMenu(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SysBaseMenu
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.UpdateBaseMenu(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateCasbin(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 uint
	a0 = dec.Uint()
	var a1 []CasbinInfo
	a1 = serviceweaver_dec_slice_CasbinInfo_0ebf7d6e(dec)

	appErr := s.impl.UpdateCasbin(ctx, a0, a1)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateRole(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 SwaRole
	(&a0).WeaverUnmarshal(dec)

	r0, appErr := s.impl.UpdateRole(ctx, a0)

	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateSwaDict(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDict
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.UpdateSwaDict(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) updateSwaDictail(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	dec := codegen.NewDecoder(args)
	var a0 ModDictail
	(&a0).WeaverUnmarshal(dec)

	appErr := s.impl.UpdateSwaDictail(ctx, a0)

	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}


type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) AddBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	err = s.caller("AddBaseMenu", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) AddMenuRole(ctx context.Context, a0 []SysBaseMenu, a1 uint) (err error) {
	err = s.caller("AddMenuRole", ctx, []any{a0, a1}, []any{})
	return
}

func (s t_reflect_stub) ChangePassword(ctx context.Context, a0 uint, a1 string, a2 string) (r0 *SysUser, err error) {
	err = s.caller("ChangePassword", ctx, []any{a0, a1, a2}, []any{&r0})
	return
}

func (s t_reflect_stub) CreateApi(ctx context.Context, a0 ModApi) (err error) {
	err = s.caller("CreateApi", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) CreateBlacklist(ctx context.Context, a0 JwtBlacklist) (err error) {
	err = s.caller("CreateBlacklist", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) CreateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	err = s.caller("CreateRole", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) CreateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	err = s.caller("CreateSwaDict", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) CreateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	err = s.caller("CreateSwaDictail", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) CreateSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	err = s.caller("CreateSwaRecord", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteApi(ctx context.Context, a0 ModApi) (err error) {
	err = s.caller("DeleteApi", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteApisByIds(ctx context.Context, a0 []uint) (err error) {
	err = s.caller("DeleteApisByIds", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteBaseMenu(ctx context.Context, a0 int) (err error) {
	err = s.caller("DeleteBaseMenu", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteRole(ctx context.Context, a0 RoleUserMenu) (err error) {
	err = s.caller("DeleteRole", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteSwaDict(ctx context.Context, a0 ModDict) (err error) {
	err = s.caller("DeleteSwaDict", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	err = s.caller("DeleteSwaDictail", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteSwaRecord(ctx context.Context, a0 SwaRecord) (err error) {
	err = s.caller("DeleteSwaRecord", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteSwaRecordByIds(ctx context.Context, a0 []uint) (err error) {
	err = s.caller("DeleteSwaRecordByIds", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) DeleteUser(ctx context.Context, a0 int) (err error) {
	err = s.caller("DeleteUser", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) GetAPIInfoList(ctx context.Context, a0 ModApi, a1 PageInfo, a2 string, a3 bool) (r0 []ModApi, r1 int64, err error) {
	err = s.caller("GetAPIInfoList", ctx, []any{a0, a1, a2, a3}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetAllApis(ctx context.Context) (r0 []ModApi, err error) {
	err = s.caller("GetAllApis", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetAllBaseMenu(ctx context.Context) (r0 []SysBaseMenu, err error) {
	err = s.caller("GetAllBaseMenu", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetApiById(ctx context.Context, a0 int) (r0 ModApi, err error) {
	err = s.caller("GetApiById", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetBaseMenuById(ctx context.Context, a0 int) (r0 SysBaseMenu, err error) {
	err = s.caller("GetBaseMenuById", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetInfoList(ctx context.Context) (r0 []SysBaseMenu, r1 int64, err error) {
	err = s.caller("GetInfoList", ctx, []any{}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetMenuRole(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	err = s.caller("GetMenuRole", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetPolicyPathByRoleId(ctx context.Context, a0 uint) (r0 []CasbinInfo, err error) {
	err = s.caller("GetPolicyPathByRoleId", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetRoleInfo(ctx context.Context, a0 uint) (r0 SwaRole, err error) {
	err = s.caller("GetRoleInfo", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetRoleList(ctx context.Context, a0 PageInfo) (r0 []SwaRole, r1 int64, err error) {
	err = s.caller("GetRoleList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetRoleMenu(ctx context.Context, a0 uint) (r0 []SysBaseMenu, err error) {
	err = s.caller("GetRoleMenu", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetServerInfo(ctx context.Context) (r0 Server, err error) {
	err = s.caller("GetServerInfo", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetSubRoles(ctx context.Context, a0 uint) (r0 []SwaRole, err error) {
	err = s.caller("GetSubRoles", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetSwaDict(ctx context.Context, a0 string, a1 uint, a2 *bool) (r0 ModDict, err error) {
	err = s.caller("GetSwaDict", ctx, []any{a0, a1, a2}, []any{&r0})
	return
}

func (s t_reflect_stub) GetSwaDictInfoList(ctx context.Context, a0 SwaDictSearch) (r0 []ModDict, r1 int64, err error) {
	err = s.caller("GetSwaDictInfoList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetSwaDictail(ctx context.Context, a0 uint) (r0 ModDictail, err error) {
	err = s.caller("GetSwaDictail", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetSwaDictailInfoList(ctx context.Context, a0 SwaDictailSearch) (r0 []ModDictail, r1 int64, err error) {
	err = s.caller("GetSwaDictailInfoList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetSwaRecord(ctx context.Context, a0 uint) (r0 SwaRecord, err error) {
	err = s.caller("GetSwaRecord", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetSwaRecordInfoList(ctx context.Context, a0 SwaRecordSearch) (r0 []SwaRecord, r1 int64, err error) {
	err = s.caller("GetSwaRecordInfoList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) GetUserDict(ctx context.Context) (r0 []SysUser, err error) {
	err = s.caller("GetUserDict", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) GetUserInfo(ctx context.Context, a0 uuid.UUID) (r0 SysUser, err error) {
	err = s.caller("GetUserInfo", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) GetUserInfoList(ctx context.Context, a0 PageInfo) (r0 []SysUser, r1 int64, err error) {
	err = s.caller("GetUserInfoList", ctx, []any{a0}, []any{&r0, &r1})
	return
}

func (s t_reflect_stub) Login(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	err = s.caller("Login", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) Register(ctx context.Context, a0 SysUser) (r0 SysUser, err error) {
	err = s.caller("Register", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) ResetPassword(ctx context.Context, a0 uint) (err error) {
	err = s.caller("ResetPassword", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) Scasbin(ctx context.Context, a0 string, a1 string, a2 string) (r0 bool, err error) {
	err = s.caller("Scasbin", ctx, []any{a0, a1, a2}, []any{&r0})
	return
}

func (s t_reflect_stub) SetRoleMenus(ctx context.Context, a0 SwaRole) (err error) {
	err = s.caller("SetRoleMenus", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) SetSelfInfo(ctx context.Context, a0 SysUser) (err error) {
	err = s.caller("SetSelfInfo", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) SetSubRoles(ctx context.Context, a0 SwaRole) (err error) {
	err = s.caller("SetSubRoles", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) SetUserAuthorities(ctx context.Context, a0 uint, a1 []uint) (err error) {
	err = s.caller("SetUserAuthorities", ctx, []any{a0, a1}, []any{})
	return
}

func (s t_reflect_stub) SetUserInfo(ctx context.Context, a0 SysUser) (err error) {
	err = s.caller("SetUserInfo", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) SetUserRole(ctx context.Context, a0 uint, a1 uint) (err error) {
	err = s.caller("SetUserRole", ctx, []any{a0, a1}, []any{})
	return
}

func (s t_reflect_stub) UpdateApi(ctx context.Context, a0 ModApi) (err error) {
	err = s.caller("UpdateApi", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) UpdateBaseMenu(ctx context.Context, a0 SysBaseMenu) (err error) {
	err = s.caller("UpdateBaseMenu", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) UpdateCasbin(ctx context.Context, a0 uint, a1 []CasbinInfo) (err error) {
	err = s.caller("UpdateCasbin", ctx, []any{a0, a1}, []any{})
	return
}

func (s t_reflect_stub) UpdateRole(ctx context.Context, a0 SwaRole) (r0 SwaRole, err error) {
	err = s.caller("UpdateRole", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) UpdateSwaDict(ctx context.Context, a0 ModDict) (err error) {
	err = s.caller("UpdateSwaDict", ctx, []any{a0}, []any{})
	return
}

func (s t_reflect_stub) UpdateSwaDictail(ctx context.Context, a0 ModDictail) (err error) {
	err = s.caller("UpdateSwaDictail", ctx, []any{a0}, []any{})
	return
}


var _ codegen.AutoMarshal = (*CasbinInfo)(nil)

type __is_CasbinInfo[T ~struct {
	weaver.AutoMarshal
	Path   string "json:\"path\""
	Method string "json:\"method\""
}] struct{}

var _ __is_CasbinInfo[CasbinInfo]

func (x *CasbinInfo) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("CasbinInfo.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Path)
	enc.String(x.Method)
}

func (x *CasbinInfo) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("CasbinInfo.WeaverUnmarshal: nil receiver"))
	}
	x.Path = dec.String()
	x.Method = dec.String()
}

var _ codegen.AutoMarshal = (*Cpu)(nil)

type __is_Cpu[T ~struct {
	weaver.AutoMarshal
	Cpus  []float64 "json:\"cpus\""
	Cores int       "json:\"cores\""
}] struct{}

var _ __is_Cpu[Cpu]

func (x *Cpu) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Cpu.WeaverMarshal: nil receiver"))
	}
	serviceweaver_enc_slice_float64_946dd0da(enc, x.Cpus)
	enc.Int(x.Cores)
}

func (x *Cpu) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Cpu.WeaverUnmarshal: nil receiver"))
	}
	x.Cpus = serviceweaver_dec_slice_float64_946dd0da(dec)
	x.Cores = dec.Int()
}

func serviceweaver_enc_slice_float64_946dd0da(enc *codegen.Encoder, arg []float64) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Float64(arg[i])
	}
}

func serviceweaver_dec_slice_float64_946dd0da(dec *codegen.Decoder) []float64 {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Float64()
	}
	return res
}

var _ codegen.AutoMarshal = (*Disk)(nil)

type __is_Disk[T ~struct {
	weaver.AutoMarshal
	UsedGB      int "json:\"usedGb\""
	TotalGB     int "json:\"totalGb\""
	UsedPercent int "json:\"usedPercent\""
}] struct{}

var _ __is_Disk[Disk]

func (x *Disk) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Disk.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.UsedGB)
	enc.Int(x.TotalGB)
	enc.Int(x.UsedPercent)
}

func (x *Disk) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Disk.WeaverUnmarshal: nil receiver"))
	}
	x.UsedGB = dec.Int()
	x.TotalGB = dec.Int()
	x.UsedPercent = dec.Int()
}

var _ codegen.AutoMarshal = (*JwtBlacklist)(nil)

type __is_JwtBlacklist[T ~struct {
	weaver.AutoMarshal
	ModBase
	Jwt string "gorm:\"type:text;comment:jwt\""
}] struct{}

var _ __is_JwtBlacklist[JwtBlacklist]

func (x *JwtBlacklist) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("JwtBlacklist.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Jwt)
}

func (x *JwtBlacklist) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("JwtBlacklist.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Jwt = dec.String()
}

var _ codegen.AutoMarshal = (*Meta)(nil)

type __is_Meta[T ~struct {
	weaver.AutoMarshal
	ActiveName  string "json:\"activeName\" gorm:\"comment:高亮菜单\""
	KeepAlive   bool   "json:\"keepAlive\" gorm:\"comment:是否缓存\""
	DefaultMenu bool   "json:\"defaultMenu\" gorm:\"comment:是否是基础路由（开发中）\""
	Title       string "json:\"title\" gorm:\"comment:菜单名\""
	Icon        string "json:\"icon\" gorm:\"comment:菜单图标\""
	CloseTab    bool   "json:\"closeTab\" gorm:\"comment:自动关闭tab\""
}] struct{}

var _ __is_Meta[Meta]

func (x *Meta) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Meta.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ActiveName)
	enc.Bool(x.KeepAlive)
	enc.Bool(x.DefaultMenu)
	enc.String(x.Title)
	enc.String(x.Icon)
	enc.Bool(x.CloseTab)
}

func (x *Meta) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Meta.WeaverUnmarshal: nil receiver"))
	}
	x.ActiveName = dec.String()
	x.KeepAlive = dec.Bool()
	x.DefaultMenu = dec.Bool()
	x.Title = dec.String()
	x.Icon = dec.String()
	x.CloseTab = dec.Bool()
}

var _ codegen.AutoMarshal = (*ModApi)(nil)

type __is_ModApi[T ~struct {
	weaver.AutoMarshal
	ModBase
	AppID       uint   "json:\"appID\" gorm:\"comment:appID\""
	TableID     uint   "json:\"tableID\" gorm:\"comment:tableID\""
	Path        string "json:\"path\" gorm:\"comment:api路径\""
	Description string "json:\"description\" gorm:\"comment:api中文描述\""
	ApiGroup    string "json:\"apiGroup\" gorm:\"comment:api组\""
	Method      string "json:\"method\" gorm:\"default:POST;comment:方法\""
}] struct{}

var _ __is_ModApi[ModApi]

func (x *ModApi) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModApi.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.Uint(x.AppID)
	enc.Uint(x.TableID)
	enc.String(x.Path)
	enc.String(x.Description)
	enc.String(x.ApiGroup)
	enc.String(x.Method)
}

func (x *ModApi) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModApi.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.AppID = dec.Uint()
	x.TableID = dec.Uint()
	x.Path = dec.String()
	x.Description = dec.String()
	x.ApiGroup = dec.String()
	x.Method = dec.String()
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

var _ codegen.AutoMarshal = (*ModDict)(nil)

type __is_ModDict[T ~struct {
	weaver.AutoMarshal
	ModBase
	Name        string       "json:\"name\" form:\"name\" gorm:\"column:name;comment:字典名（中）\""
	Type        string       "json:\"type\" form:\"type\" gorm:\"column:type;comment:字典名（英）\""
	Status      *bool        "json:\"status\" form:\"status\" gorm:\"column:status;comment:状态\""
	Detail      string       "json:\"detail\" form:\"detail\" gorm:\"column:detail;comment:描述\""
	ModDictails []ModDictail "json:\"swaDictail\" gorm:\"-\""
}] struct{}

var _ __is_ModDict[ModDict]

func (x *ModDict) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModDict.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Name)
	enc.String(x.Type)
	serviceweaver_enc_ptr_bool_31f02903(enc, x.Status)
	enc.String(x.Detail)
	serviceweaver_enc_slice_ModDictail_42b2abbf(enc, x.ModDictails)
}

func (x *ModDict) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModDict.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Name = dec.String()
	x.Type = dec.String()
	x.Status = serviceweaver_dec_ptr_bool_31f02903(dec)
	x.Detail = dec.String()
	x.ModDictails = serviceweaver_dec_slice_ModDictail_42b2abbf(dec)
}

func serviceweaver_enc_ptr_bool_31f02903(enc *codegen.Encoder, arg *bool) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.Bool(*arg)
	}
}

func serviceweaver_dec_ptr_bool_31f02903(dec *codegen.Decoder) *bool {
	if !dec.Bool() {
		return nil
	}
	var res bool
	res = dec.Bool()
	return &res
}

func serviceweaver_enc_slice_ModDictail_42b2abbf(enc *codegen.Encoder, arg []ModDictail) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModDictail_42b2abbf(dec *codegen.Decoder) []ModDictail {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModDictail, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*ModDictail)(nil)

type __is_ModDictail[T ~struct {
	weaver.AutoMarshal
	ModBase
	Label     string "json:\"label\" form:\"label\" gorm:\"column:label;comment:展示值\""
	Value     int    "json:\"value\" form:\"value\" gorm:\"column:value;comment:字典值\""
	Status    *bool  "json:\"status\" form:\"status\" gorm:\"column:status;comment:启用状态\""
	Sort      int    "json:\"sort\" form:\"sort\" gorm:\"column:sort;comment:排序标记\""
	SwaDictID int    "json:\"swaDictID\" form:\"swaDictID\" gorm:\"column:swaDictID;comment:关联标记\""
}] struct{}

var _ __is_ModDictail[ModDictail]

func (x *ModDictail) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModDictail.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Label)
	enc.Int(x.Value)
	serviceweaver_enc_ptr_bool_31f02903(enc, x.Status)
	enc.Int(x.Sort)
	enc.Int(x.SwaDictID)
}

func (x *ModDictail) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModDictail.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Label = dec.String()
	x.Value = dec.Int()
	x.Status = serviceweaver_dec_ptr_bool_31f02903(dec)
	x.Sort = dec.Int()
	x.SwaDictID = dec.Int()
}

var _ codegen.AutoMarshal = (*ModMenu)(nil)

type __is_ModMenu[T ~struct {
	weaver.AutoMarshal
	ModBase
	MenuLevel uint   "json:\"-\""
	ParentId  string "json:\"parentId\" gorm:\"comment:父菜单ID\""
	Path      string "json:\"path\" gorm:\"comment:路由path\""
	Name      string "json:\"name\" gorm:\"comment:路由name\""
	Hidden    bool   "json:\"hidden\" gorm:\"comment:是否在列表隐藏\""
	Component string "json:\"component\" gorm:\"comment:对应前端文件路径\""
	Sort      int    "json:\"sort\" gorm:\"comment:排序标记\""
	Meta      "json:\"meta\" gorm:\"embedded;comment:附加属性\""
}] struct{}

var _ __is_ModMenu[ModMenu]

func (x *ModMenu) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModMenu.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.Uint(x.MenuLevel)
	enc.String(x.ParentId)
	enc.String(x.Path)
	enc.String(x.Name)
	enc.Bool(x.Hidden)
	enc.String(x.Component)
	enc.Int(x.Sort)
	(x.Meta).WeaverMarshal(enc)
}

func (x *ModMenu) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModMenu.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.MenuLevel = dec.Uint()
	x.ParentId = dec.String()
	x.Path = dec.String()
	x.Name = dec.String()
	x.Hidden = dec.Bool()
	x.Component = dec.String()
	x.Sort = dec.Int()
	(&x.Meta).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*ModRole)(nil)

type __is_ModRole[T ~struct {
	weaver.AutoMarshal
	ModBase
	RoleId        uint   "json:\"roleId\" gorm:\"column:roleId;not null;comment:角色ID;size:90\""
	RoleName      string "json:\"roleName\" gorm:\"column:roleName;comment:角色名\""
	ParentId      uint   "json:\"parentId\" gorm:\"column:parentId;comment:父角色ID\""
	DefaultRouter string "json:\"defaultRouter\" gorm:\"column:defaultRouter;comment:默认菜单;default:dashboard\""
}] struct{}

var _ __is_ModRole[ModRole]

func (x *ModRole) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("ModRole.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.Uint(x.RoleId)
	enc.String(x.RoleName)
	enc.Uint(x.ParentId)
	enc.String(x.DefaultRouter)
}

func (x *ModRole) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("ModRole.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.RoleId = dec.Uint()
	x.RoleName = dec.String()
	x.ParentId = dec.Uint()
	x.DefaultRouter = dec.String()
}

var _ codegen.AutoMarshal = (*Os)(nil)

type __is_Os[T ~struct {
	weaver.AutoMarshal
	GOOS         string "json:\"goos\""
	NumCPU       int    "json:\"numCpu\""
	Compiler     string "json:\"compiler\""
	GoVersion    string "json:\"goVersion\""
	NumGoroutine int    "json:\"numGoroutine\""
}] struct{}

var _ __is_Os[Os]

func (x *Os) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Os.WeaverMarshal: nil receiver"))
	}
	enc.String(x.GOOS)
	enc.Int(x.NumCPU)
	enc.String(x.Compiler)
	enc.String(x.GoVersion)
	enc.Int(x.NumGoroutine)
}

func (x *Os) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Os.WeaverUnmarshal: nil receiver"))
	}
	x.GOOS = dec.String()
	x.NumCPU = dec.Int()
	x.Compiler = dec.String()
	x.GoVersion = dec.String()
	x.NumGoroutine = dec.Int()
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

var _ codegen.AutoMarshal = (*Ram)(nil)

type __is_Ram[T ~struct {
	weaver.AutoMarshal
	UsedMB      int "json:\"usedMb\""
	TotalMB     int "json:\"totalMb\""
	UsedPercent int "json:\"usedPercent\""
}] struct{}

var _ __is_Ram[Ram]

func (x *Ram) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Ram.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.UsedMB)
	enc.Int(x.TotalMB)
	enc.Int(x.UsedPercent)
}

func (x *Ram) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Ram.WeaverUnmarshal: nil receiver"))
	}
	x.UsedMB = dec.Int()
	x.TotalMB = dec.Int()
	x.UsedPercent = dec.Int()
}

var _ codegen.AutoMarshal = (*RoleUserMenu)(nil)

type __is_RoleUserMenu[T ~struct {
	weaver.AutoMarshal
	ModRole
	Users        []SysUser     "json:\"users\" gorm:\"-\""
	SysBaseMenus []SysBaseMenu "json:\"menus\" gorm:\"-\""
}] struct{}

var _ __is_RoleUserMenu[RoleUserMenu]

func (x *RoleUserMenu) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("RoleUserMenu.WeaverMarshal: nil receiver"))
	}
	(x.ModRole).WeaverMarshal(enc)
	serviceweaver_enc_slice_SysUser_00eb59d9(enc, x.Users)
	serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc, x.SysBaseMenus)
}

func (x *RoleUserMenu) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("RoleUserMenu.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModRole).WeaverUnmarshal(dec)
	x.Users = serviceweaver_dec_slice_SysUser_00eb59d9(dec)
	x.SysBaseMenus = serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec)
}

func serviceweaver_enc_slice_SysUser_00eb59d9(enc *codegen.Encoder, arg []SysUser) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SysUser_00eb59d9(dec *codegen.Decoder) []SysUser {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SysUser, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_SysBaseMenu_61e4d06d(enc *codegen.Encoder, arg []SysBaseMenu) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SysBaseMenu_61e4d06d(dec *codegen.Decoder) []SysBaseMenu {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SysBaseMenu, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*SearchInfoApi)(nil)

type __is_SearchInfoApi[T ~struct {
	weaver.AutoMarshal
	ModApi
	PageInfo
	StartCreatedAt time.Time "json:\"startCreatedAt\" form:\"startCreatedAt\""
	EndCreatedAt   time.Time "json:\"endCreatedAt\" form:\"endCreatedAt\""
	OrderKey       string    "json:\"orderKey\""
	Desc           bool      "json:\"desc\""
}] struct{}

var _ __is_SearchInfoApi[SearchInfoApi]

func (x *SearchInfoApi) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoApi.WeaverMarshal: nil receiver"))
	}
	(x.ModApi).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.StartCreatedAt)
	enc.EncodeBinaryMarshaler(&x.EndCreatedAt)
	enc.String(x.OrderKey)
	enc.Bool(x.Desc)
}

func (x *SearchInfoApi) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoApi.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModApi).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.StartCreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.EndCreatedAt)
	x.OrderKey = dec.String()
	x.Desc = dec.Bool()
}

var _ codegen.AutoMarshal = (*SearchInfoBlacklist)(nil)

type __is_SearchInfoBlacklist[T ~struct {
	weaver.AutoMarshal
	JwtBlacklist
	PageInfo
	StartCreatedAt time.Time "json:\"startCreatedAt\" form:\"startCreatedAt\""
	EndCreatedAt   time.Time "json:\"endCreatedAt\" form:\"endCreatedAt\""
	OrderKey       string    "json:\"orderKey\""
	Desc           bool      "json:\"desc\""
}] struct{}

var _ __is_SearchInfoBlacklist[SearchInfoBlacklist]

func (x *SearchInfoBlacklist) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoBlacklist.WeaverMarshal: nil receiver"))
	}
	(x.JwtBlacklist).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.StartCreatedAt)
	enc.EncodeBinaryMarshaler(&x.EndCreatedAt)
	enc.String(x.OrderKey)
	enc.Bool(x.Desc)
}

func (x *SearchInfoBlacklist) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoBlacklist.WeaverUnmarshal: nil receiver"))
	}
	(&x.JwtBlacklist).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.StartCreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.EndCreatedAt)
	x.OrderKey = dec.String()
	x.Desc = dec.Bool()
}

var _ codegen.AutoMarshal = (*SearchInfoDict)(nil)

type __is_SearchInfoDict[T ~struct {
	weaver.AutoMarshal
	ModDict
	PageInfo
	StartCreatedAt time.Time "json:\"startCreatedAt\" form:\"startCreatedAt\""
	EndCreatedAt   time.Time "json:\"endCreatedAt\" form:\"endCreatedAt\""
	OrderKey       string    "json:\"orderKey\""
	Desc           bool      "json:\"desc\""
}] struct{}

var _ __is_SearchInfoDict[SearchInfoDict]

func (x *SearchInfoDict) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoDict.WeaverMarshal: nil receiver"))
	}
	(x.ModDict).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.StartCreatedAt)
	enc.EncodeBinaryMarshaler(&x.EndCreatedAt)
	enc.String(x.OrderKey)
	enc.Bool(x.Desc)
}

func (x *SearchInfoDict) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoDict.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModDict).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.StartCreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.EndCreatedAt)
	x.OrderKey = dec.String()
	x.Desc = dec.Bool()
}

var _ codegen.AutoMarshal = (*SearchInfoDictail)(nil)

type __is_SearchInfoDictail[T ~struct {
	weaver.AutoMarshal
	ModDictail
	PageInfo
	StartCreatedAt time.Time "json:\"startCreatedAt\" form:\"startCreatedAt\""
	EndCreatedAt   time.Time "json:\"endCreatedAt\" form:\"endCreatedAt\""
	OrderKey       string    "json:\"orderKey\""
	Desc           bool      "json:\"desc\""
}] struct{}

var _ __is_SearchInfoDictail[SearchInfoDictail]

func (x *SearchInfoDictail) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoDictail.WeaverMarshal: nil receiver"))
	}
	(x.ModDictail).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.StartCreatedAt)
	enc.EncodeBinaryMarshaler(&x.EndCreatedAt)
	enc.String(x.OrderKey)
	enc.Bool(x.Desc)
}

func (x *SearchInfoDictail) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoDictail.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModDictail).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.StartCreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.EndCreatedAt)
	x.OrderKey = dec.String()
	x.Desc = dec.Bool()
}

var _ codegen.AutoMarshal = (*SearchInfoRecord)(nil)

type __is_SearchInfoRecord[T ~struct {
	weaver.AutoMarshal
	SwaRecord
	PageInfo
	StartCreatedAt time.Time "json:\"startCreatedAt\" form:\"startCreatedAt\""
	EndCreatedAt   time.Time "json:\"endCreatedAt\" form:\"endCreatedAt\""
	OrderKey       string    "json:\"orderKey\""
	Desc           bool      "json:\"desc\""
}] struct{}

var _ __is_SearchInfoRecord[SearchInfoRecord]

func (x *SearchInfoRecord) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoRecord.WeaverMarshal: nil receiver"))
	}
	(x.SwaRecord).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.StartCreatedAt)
	enc.EncodeBinaryMarshaler(&x.EndCreatedAt)
	enc.String(x.OrderKey)
	enc.Bool(x.Desc)
}

func (x *SearchInfoRecord) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SearchInfoRecord.WeaverUnmarshal: nil receiver"))
	}
	(&x.SwaRecord).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.StartCreatedAt)
	dec.DecodeBinaryUnmarshaler(&x.EndCreatedAt)
	x.OrderKey = dec.String()
	x.Desc = dec.Bool()
}

var _ codegen.AutoMarshal = (*Server)(nil)

type __is_Server[T ~struct {
	weaver.AutoMarshal
	Os   Os   "json:\"os\""
	Cpu  Cpu  "json:\"cpu\""
	Ram  Ram  "json:\"ram\""
	Disk Disk "json:\"disk\""
}] struct{}

var _ __is_Server[Server]

func (x *Server) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Server.WeaverMarshal: nil receiver"))
	}
	(x.Os).WeaverMarshal(enc)
	(x.Cpu).WeaverMarshal(enc)
	(x.Ram).WeaverMarshal(enc)
	(x.Disk).WeaverMarshal(enc)
}

func (x *Server) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Server.WeaverUnmarshal: nil receiver"))
	}
	(&x.Os).WeaverUnmarshal(dec)
	(&x.Cpu).WeaverUnmarshal(dec)
	(&x.Ram).WeaverUnmarshal(dec)
	(&x.Disk).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*SwaDictSearch)(nil)

type __is_SwaDictSearch[T ~struct {
	weaver.AutoMarshal
	ModDict
	PageInfo
}] struct{}

var _ __is_SwaDictSearch[SwaDictSearch]

func (x *SwaDictSearch) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SwaDictSearch.WeaverMarshal: nil receiver"))
	}
	(x.ModDict).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
}

func (x *SwaDictSearch) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SwaDictSearch.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModDict).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*SwaDictailSearch)(nil)

type __is_SwaDictailSearch[T ~struct {
	weaver.AutoMarshal
	ModDictail
	PageInfo
}] struct{}

var _ __is_SwaDictailSearch[SwaDictailSearch]

func (x *SwaDictailSearch) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SwaDictailSearch.WeaverMarshal: nil receiver"))
	}
	(x.ModDictail).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
}

func (x *SwaDictailSearch) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SwaDictailSearch.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModDictail).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*SwaRecord)(nil)

type __is_SwaRecord[T ~struct {
	weaver.AutoMarshal
	ModBase
	Ip           string        "json:\"ip\" form:\"ip\" gorm:\"column:ip;comment:请求ip\""
	Method       string        "json:\"method\" form:\"method\" gorm:\"column:method;comment:请求方法\""
	Path         string        "json:\"path\" form:\"path\" gorm:\"column:path;comment:请求路径\""
	Status       int           "json:\"status\" form:\"status\" gorm:\"column:status;comment:请求状态\""
	Latency      time.Duration "json:\"latency\" form:\"latency\" gorm:\"column:latency;comment:延迟\" swaggertype:\"string\""
	Agent        string        "json:\"agent\" form:\"agent\" gorm:\"type:text;column:agent;comment:代理\""
	ErrorMessage string        "json:\"error_message\" form:\"error_message\" gorm:\"type:text;column:error_message;comment:错误信息\""
	Body         string        "json:\"body\" form:\"body\" gorm:\"type:text;column:body;comment:请求Body\""
	Resp         string        "json:\"resp\" form:\"resp\" gorm:\"type:text;column:resp;comment:响应Body\""
	UserID       int           "json:\"user_id\" form:\"user_id\" gorm:\"column:user_id;comment:用户id\""
	User         SysUser       "json:\"user\""
}] struct{}

var _ __is_SwaRecord[SwaRecord]

func (x *SwaRecord) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRecord.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Ip)
	enc.String(x.Method)
	enc.String(x.Path)
	enc.Int(x.Status)
	enc.Int64((int64)(x.Latency))
	enc.String(x.Agent)
	enc.String(x.ErrorMessage)
	enc.String(x.Body)
	enc.String(x.Resp)
	enc.Int(x.UserID)
	(x.User).WeaverMarshal(enc)
}

func (x *SwaRecord) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRecord.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Ip = dec.String()
	x.Method = dec.String()
	x.Path = dec.String()
	x.Status = dec.Int()
	*(*int64)(&x.Latency) = dec.Int64()
	x.Agent = dec.String()
	x.ErrorMessage = dec.String()
	x.Body = dec.String()
	x.Resp = dec.String()
	x.UserID = dec.Int()
	(&x.User).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*SwaRecordSearch)(nil)

type __is_SwaRecordSearch[T ~struct {
	weaver.AutoMarshal
	SwaRecord
	PageInfo
}] struct{}

var _ __is_SwaRecordSearch[SwaRecordSearch]

func (x *SwaRecordSearch) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRecordSearch.WeaverMarshal: nil receiver"))
	}
	(x.SwaRecord).WeaverMarshal(enc)
	(x.PageInfo).WeaverMarshal(enc)
}

func (x *SwaRecordSearch) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRecordSearch.WeaverUnmarshal: nil receiver"))
	}
	(&x.SwaRecord).WeaverUnmarshal(dec)
	(&x.PageInfo).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*SwaRole)(nil)

type __is_SwaRole[T ~struct {
	weaver.AutoMarshal
	ModRole
	SubRoles  []ModRole "json:\"subRoles\" gorm:\"-\""
	RoleMenus []ModMenu "json:\"menus\" gorm:\"-\""
}] struct{}

var _ __is_SwaRole[SwaRole]

func (x *SwaRole) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRole.WeaverMarshal: nil receiver"))
	}
	(x.ModRole).WeaverMarshal(enc)
	serviceweaver_enc_slice_ModRole_d56700dc(enc, x.SubRoles)
	serviceweaver_enc_slice_ModMenu_c14b750e(enc, x.RoleMenus)
}

func (x *SwaRole) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SwaRole.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModRole).WeaverUnmarshal(dec)
	x.SubRoles = serviceweaver_dec_slice_ModRole_d56700dc(dec)
	x.RoleMenus = serviceweaver_dec_slice_ModMenu_c14b750e(dec)
}

func serviceweaver_enc_slice_ModRole_d56700dc(enc *codegen.Encoder, arg []ModRole) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModRole_d56700dc(dec *codegen.Decoder) []ModRole {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModRole, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_ModMenu_c14b750e(enc *codegen.Encoder, arg []ModMenu) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModMenu_c14b750e(dec *codegen.Decoder) []ModMenu {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModMenu, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*SysBaseMenu)(nil)

type __is_SysBaseMenu[T ~struct {
	weaver.AutoMarshal
	ModMenu
	SwaRoles   []SwaRole              "json:\"roles\" gorm:\"many2many:swa_menu_role;\""
	Parameters []SysBaseMenuParameter "json:\"parameters\" gorm:\"-\""
	MenuBtn    []SysBaseMenuBtn       "json:\"menuBtn\" gorm:\"-\""
	Meta       "json:\"meta\" gorm:\"embedded;comment:附加属性\""
}] struct{}

var _ __is_SysBaseMenu[SysBaseMenu]

func (x *SysBaseMenu) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenu.WeaverMarshal: nil receiver"))
	}
	(x.ModMenu).WeaverMarshal(enc)
	serviceweaver_enc_slice_SwaRole_8aab416c(enc, x.SwaRoles)
	serviceweaver_enc_slice_SysBaseMenuParameter_d5b3b57f(enc, x.Parameters)
	serviceweaver_enc_slice_SysBaseMenuBtn_5d92de74(enc, x.MenuBtn)
	(x.Meta).WeaverMarshal(enc)
}

func (x *SysBaseMenu) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenu.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModMenu).WeaverUnmarshal(dec)
	x.SwaRoles = serviceweaver_dec_slice_SwaRole_8aab416c(dec)
	x.Parameters = serviceweaver_dec_slice_SysBaseMenuParameter_d5b3b57f(dec)
	x.MenuBtn = serviceweaver_dec_slice_SysBaseMenuBtn_5d92de74(dec)
	(&x.Meta).WeaverUnmarshal(dec)
}

func serviceweaver_enc_slice_SwaRole_8aab416c(enc *codegen.Encoder, arg []SwaRole) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SwaRole_8aab416c(dec *codegen.Decoder) []SwaRole {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SwaRole, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_SysBaseMenuParameter_d5b3b57f(enc *codegen.Encoder, arg []SysBaseMenuParameter) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SysBaseMenuParameter_d5b3b57f(dec *codegen.Decoder) []SysBaseMenuParameter {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SysBaseMenuParameter, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_SysBaseMenuBtn_5d92de74(enc *codegen.Encoder, arg []SysBaseMenuBtn) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SysBaseMenuBtn_5d92de74(dec *codegen.Decoder) []SysBaseMenuBtn {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SysBaseMenuBtn, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*SysBaseMenuBtn)(nil)

type __is_SysBaseMenuBtn[T ~struct {
	weaver.AutoMarshal
	ModBase
	Name          string "json:\"name\" gorm:\"comment:按钮关键key\""
	Desc          string "json:\"desc\" gorm:\"按钮备注\""
	SysBaseMenuID uint   "json:\"sysBaseMenuID\" gorm:\"comment:菜单ID\""
}] struct{}

var _ __is_SysBaseMenuBtn[SysBaseMenuBtn]

func (x *SysBaseMenuBtn) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenuBtn.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.String(x.Name)
	enc.String(x.Desc)
	enc.Uint(x.SysBaseMenuID)
}

func (x *SysBaseMenuBtn) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenuBtn.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.Name = dec.String()
	x.Desc = dec.String()
	x.SysBaseMenuID = dec.Uint()
}

var _ codegen.AutoMarshal = (*SysBaseMenuParameter)(nil)

type __is_SysBaseMenuParameter[T ~struct {
	weaver.AutoMarshal
	ModBase
	SysBaseMenuID uint   "json:\"menuId\""
	Type          string "json:\"type\" gorm:\"comment:地址栏携带参数为params还是query\""
	Key           string "json:\"key\" gorm:\"comment:地址栏携带参数的key\""
	Value         string "json:\"value\" gorm:\"comment:地址栏携带参数的值\""
}] struct{}

var _ __is_SysBaseMenuParameter[SysBaseMenuParameter]

func (x *SysBaseMenuParameter) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenuParameter.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.Uint(x.SysBaseMenuID)
	enc.String(x.Type)
	enc.String(x.Key)
	enc.String(x.Value)
}

func (x *SysBaseMenuParameter) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SysBaseMenuParameter.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	x.SysBaseMenuID = dec.Uint()
	x.Type = dec.String()
	x.Key = dec.String()
	x.Value = dec.String()
}

var _ codegen.AutoMarshal = (*SysUser)(nil)

type __is_SysUser[T ~struct {
	weaver.AutoMarshal
	ModBase
	UUID        uuid.UUID "json:\"uuid\" gorm:\"index;comment:用户UUID\""
	Username    string    "json:\"userName\" gorm:\"index;comment:用户登录名\""
	Password    string    "json:\"-\"  gorm:\"comment:用户登录密码\""
	NickName    string    "json:\"nickName\" gorm:\"default:系统用户;comment:用户昵称\""
	SideMode    string    "json:\"sideMode\" gorm:\"default:dark;comment:用户侧边主题\""
	HeaderImg   string    "json:\"headerImg\" gorm:\"comment:用户头像\""
	BaseColor   string    "json:\"baseColor\" gorm:\"default:#fff;comment:基础颜色\""
	ActiveColor string    "json:\"activeColor\" gorm:\"default:#1890ff;comment:活跃颜色\""
	Phone       string    "json:\"phone\"  gorm:\"comment:用户手机号\""
	Email       string    "json:\"email\"  gorm:\"comment:用户邮箱\""
	Team        string    "json:\"team\"  gorm:\"comment:项目组\""
	Dept        string    "json:\"dept\"  gorm:\"comment:部门\""
	Company     string    "json:\"company\"  gorm:\"comment:公司/组织\""
	Enable      int       "json:\"enable\" gorm:\"default:1;comment:用户是否被冻结 1正常 2冻结\""
	RoleId      uint      "json:\"roleId\" gorm:\"default:888;comment:用户角色ID\""
	RoleIds     []uint    "json:\"roleIds\" gorm:\"-\""
	SwaRole     SwaRole   "json:\"swaRole\" gorm:\"-\""
	Authorities []SwaRole "json:\"authorities\" gorm:\"-\""
}] struct{}

var _ __is_SysUser[SysUser]

func (x *SysUser) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SysUser.WeaverMarshal: nil receiver"))
	}
	(x.ModBase).WeaverMarshal(enc)
	enc.EncodeBinaryMarshaler(&x.UUID)
	enc.String(x.Username)
	enc.String(x.Password)
	enc.String(x.NickName)
	enc.String(x.SideMode)
	enc.String(x.HeaderImg)
	enc.String(x.BaseColor)
	enc.String(x.ActiveColor)
	enc.String(x.Phone)
	enc.String(x.Email)
	enc.String(x.Team)
	enc.String(x.Dept)
	enc.String(x.Company)
	enc.Int(x.Enable)
	enc.Uint(x.RoleId)
	serviceweaver_enc_slice_uint_e3009941(enc, x.RoleIds)
	(x.SwaRole).WeaverMarshal(enc)
	serviceweaver_enc_slice_SwaRole_8aab416c(enc, x.Authorities)
}

func (x *SysUser) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SysUser.WeaverUnmarshal: nil receiver"))
	}
	(&x.ModBase).WeaverUnmarshal(dec)
	dec.DecodeBinaryUnmarshaler(&x.UUID)
	x.Username = dec.String()
	x.Password = dec.String()
	x.NickName = dec.String()
	x.SideMode = dec.String()
	x.HeaderImg = dec.String()
	x.BaseColor = dec.String()
	x.ActiveColor = dec.String()
	x.Phone = dec.String()
	x.Email = dec.String()
	x.Team = dec.String()
	x.Dept = dec.String()
	x.Company = dec.String()
	x.Enable = dec.Int()
	x.RoleId = dec.Uint()
	x.RoleIds = serviceweaver_dec_slice_uint_e3009941(dec)
	(&x.SwaRole).WeaverUnmarshal(dec)
	x.Authorities = serviceweaver_dec_slice_SwaRole_8aab416c(dec)
}

func serviceweaver_enc_slice_uint_e3009941(enc *codegen.Encoder, arg []uint) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Uint(arg[i])
	}
}

func serviceweaver_dec_slice_uint_e3009941(dec *codegen.Decoder) []uint {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]uint, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Uint()
	}
	return res
}


func serviceweaver_enc_ptr_SysUser_342d2fab(enc *codegen.Encoder, arg *SysUser) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_SysUser_342d2fab(dec *codegen.Decoder) *SysUser {
	if !dec.Bool() {
		return nil
	}
	var res SysUser
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_slice_ModApi_d17cefdc(enc *codegen.Encoder, arg []ModApi) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModApi_d17cefdc(dec *codegen.Decoder) []ModApi {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModApi, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_CasbinInfo_0ebf7d6e(enc *codegen.Encoder, arg []CasbinInfo) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_CasbinInfo_0ebf7d6e(dec *codegen.Decoder) []CasbinInfo {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]CasbinInfo, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_ModDict_3a395239(enc *codegen.Encoder, arg []ModDict) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_ModDict_3a395239(dec *codegen.Decoder) []ModDict {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]ModDict, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_SwaRecord_bba98daa(enc *codegen.Encoder, arg []SwaRecord) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_SwaRecord_bba98daa(dec *codegen.Decoder) []SwaRecord {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]SwaRecord, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}


func serviceweaver_size_ptr_bool_31f02903(x *bool) int {
	if x == nil {
		return 1
	} else {
		return 1 + 1
	}
}

func serviceweaver_size_PageInfo_c29f03f8(x *PageInfo) int {
	size := 0
	size += 0
	size += 8
	size += 8
	size += (4 + len(x.Keyword))
	return size
}
