package utils

var (
	IdVerify             = Rules{"ID": []string{NotEmpty()}}
	ApiVerify            = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify           = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify       = Rules{"Title": {NotEmpty()}}
	LoginVerify          = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "RoleId": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify       = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify       = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify    = Rules{"PackageName": {NotEmpty()}}
	RoleVerify           = Rules{"RoleId": {NotEmpty()}, "RoleName": {NotEmpty()}}
	RoleIdVerify         = Rules{"RoleId": {NotEmpty()}}
	OldRoleVerify        = Rules{"OldRoleId": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserRoleVerify    = Rules{"RoleId": {NotEmpty()}}
)
