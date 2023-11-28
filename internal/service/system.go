// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast-token/gftoken"
)

type (
	ILogin interface {
		Login(ctx context.Context, verifyKey string, captcha string, userName string, password string) (loginUserOut *model.LoginUserOut, token string, err error)
		LoginOut(ctx context.Context) (err error)
	}
	ISysJob interface {
		JobList(ctx context.Context, input *model.GetJobListInput) (total int, out []*model.SysJobOut, err error)
		GetJobs(ctx context.Context) (jobs []*model.SysJobOut, err error)
		AddJob(ctx context.Context, input *model.SysJobAddInput) (err error)
		GetJobInfoById(ctx context.Context, id int) (job *model.SysJobOut, err error)
		EditJob(ctx context.Context, input *model.SysJobEditInput) error
		JobStart(ctx context.Context, job *model.SysJobOut) error
		JobStartMult(ctx context.Context, jobs []*model.SysJobOut) error
		JobStop(ctx context.Context, job *model.SysJobOut) error
		JobRun(ctx context.Context, job *model.SysJobOut) error
		DeleteJobByIds(ctx context.Context, ids []int) (err error)
		WithValue(ctx context.Context, value string) context.Context
		Value(ctx context.Context) uint64
	}
	ISysOperLog interface {
		GetList(ctx context.Context, input *model.SysOperLogDoInput) (total int, out []*model.SysOperLogOut, err error)
		Invoke(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, err error)
		Add(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, erro error) (err error)
		Detail(ctx context.Context, operId int) (entity *entity.SysOperLog, err error)
		Del(ctx context.Context, operIds []int) (err error)
		ClearOperationLogByDays(ctx context.Context, days int) (err error)
	}
	ISystemPluginsConfig interface {
		GetPluginsConfigList(ctx context.Context, in *model.GetPluginsConfigListInput) (total, page int, list []*model.PluginsConfigOutput, err error)
		GetPluginsConfigById(ctx context.Context, id int) (out *model.PluginsConfigOutput, err error)
		GetPluginsConfigByName(ctx context.Context, types, name string) (out *model.PluginsConfigOutput, err error)
		AddPluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error)
		EditPluginsConfig(ctx context.Context, in model.PluginsConfigEditInput) (err error)
		SavePluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error)
		DeletePluginsConfig(ctx context.Context, Ids []int) (err error)
		UpdateAllPluginsConfigCache() (err error)
		GetPluginsConfigData(pluginType, pluginName string) (res map[interface{}]interface{}, err error)
	}
	ISysUserRole interface {
		GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserRole, err error)
	}
	ICaptcha interface {
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
		VerifyString(id, answer string) bool
	}
	ISysDept interface {
		GetTree(ctx context.Context, deptName string, status int) (out []*model.DeptOut, err error)
		GetData(ctx context.Context, deptName string, status int) (data []*model.DeptOut, err error)
		Add(ctx context.Context, input *model.AddDeptInput) (err error)
		Edit(ctx context.Context, input *model.EditDeptInput) (err error)
		Detail(ctx context.Context, deptId int64) (entity *entity.SysDept, err error)
		Del(ctx context.Context, deptId int64) (err error)
		GetAll(ctx context.Context) (data []*entity.SysDept, err error)
		GetFromCache(ctx context.Context) (list []*entity.SysDept, err error)
		FindSonByParentId(deptList []*entity.SysDept, deptId int64) []*entity.SysDept
	}
	ISysNotifications interface {
		GetSysNotificationsList(ctx context.Context, input *model.GetNotificationsListInput) (total, page int, list []*model.NotificationsOut, err error)
		GetSysNotificationsById(ctx context.Context, id int) (out *model.NotificationsRes, err error)
		AddSysNotifications(ctx context.Context, in model.NotificationsAddInput) (err error)
		EditSysNotifications(ctx context.Context, in model.NotificationsEditInput) (err error)
		DeleteSysNotifications(ctx context.Context, in *system.DeleteNotificationsReq) (err error)
	}
	ISysPlugins interface {
		GetSysPluginsList(ctx context.Context, in *model.GetSysPluginsListInput) (total, page int, list []*model.SysPluginsOutput, err error)
		GetSysPluginsById(ctx context.Context, id int) (out *model.SysPluginsOutput, err error)
		AddSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error)
		EditSysPlugins(ctx context.Context, in model.SysPluginsEditInput) (err error)
		DeleteSysPlugins(ctx context.Context, Ids []int) (err error)
		SaveSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error)
		EditStatus(ctx context.Context, id int, status int) (err error)
		GetSysPluginsTypesAll(ctx context.Context, types string) (out []*model.SysPluginsInfoOut, err error)
	}
	ISysApi interface {
		GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysApi, err error)
		GetApiByMenuId(ctx context.Context, apiId int) (data []*entity.SysApi, err error)
		GetInfoById(ctx context.Context, id int) (entity *entity.SysApi, err error)
		GetApiAll(ctx context.Context) (data []*entity.SysApi, err error)
		GetApiTree(ctx context.Context, name string, address string, status int, types int) (out []*model.SysApiTreeOut, err error)
		Add(ctx context.Context, input *model.AddApiInput) (err error)
		Detail(ctx context.Context, id int) (out *model.SysApiOut, err error)
		Edit(ctx context.Context, input *model.EditApiInput) (err error)
		Del(ctx context.Context, Id int) (err error)
		EditStatus(ctx context.Context, id int, status int) (err error)
		GetInfoByAddress(ctx context.Context, address string) (entity *entity.SysApi, err error)
	}
	ISysLoginLog interface {
		Invoke(ctx context.Context, data *model.LoginLogParams)
		Add(ctx context.Context, params *model.LoginLogParams)
		GetList(ctx context.Context, req *model.SysLoginLogInput) (total, page int, list []*model.SysLoginLogOut, err error)
		Detail(ctx context.Context, infoId int) (entity *entity.SysLoginLog, err error)
		Del(ctx context.Context, infoIds []int) (err error)
	}
	ISysMenuApi interface {
		GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysMenuApi, err error)
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuApi, err error)
		GetInfoByApiId(ctx context.Context, apiId int) (data []*entity.SysMenuApi, err error)
	}
	ISysOrganization interface {
		GetTree(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error)
		GetData(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error)
		Add(ctx context.Context, input *model.AddOrganizationInput) (err error)
		Edit(ctx context.Context, input *model.EditOrganizationInput) (err error)
		Detail(ctx context.Context, id int64) (entity *entity.SysOrganization, err error)
		Del(ctx context.Context, id int64) (err error)
		GetAll(ctx context.Context) (data []*entity.SysOrganization, err error)
		Count(ctx context.Context) (count int, err error)
	}
	ISysPost interface {
		GetTree(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error)
		Add(ctx context.Context, input *model.AddPostInput) (err error)
		Edit(ctx context.Context, input *model.EditPostInput) (err error)
		Detail(ctx context.Context, postId int64) (entity *entity.SysPost, err error)
		GetData(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error)
		Del(ctx context.Context, postId int64) (err error)
		GetUsedPost(ctx context.Context) (list []*model.DetailPostRes, err error)
	}
	ISysRole interface {
		GetAll(ctx context.Context) (entity []*entity.SysRole, err error)
		GetTree(ctx context.Context, name string, status int) (out []*model.RoleTreeOut, err error)
		Add(ctx context.Context, input *model.AddRoleInput) (err error)
		Edit(ctx context.Context, input *model.EditRoleInput) (err error)
		GetInfoById(ctx context.Context, id uint) (entity *entity.SysRole, err error)
		DelInfoById(ctx context.Context, id uint) (err error)
		GetRoleList(ctx context.Context) (list []*model.RoleInfoRes, err error)
		GetInfoByIds(ctx context.Context, id []int) (entity []*entity.SysRole, err error)
		DataScope(ctx context.Context, id int, dataScope uint, deptIds []int64) (err error)
		GetAuthorizeById(ctx context.Context, id int) (menuIds []string, menuButtonIds []string, menuColumnIds []string, menuApiIds []string, err error)
	}
	ISysToken interface {
		GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
		ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	}
	ISysUser interface {
		GetUserByUsername(ctx context.Context, userName string) (data *entity.SysUser, err error)
		GetAdminUserByUsernamePassword(ctx context.Context, userName string, password string) (user *entity.SysUser, err error)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		UserList(ctx context.Context, input *model.UserListDoInput) (total int, out []*model.UserListOut, err error)
		Add(ctx context.Context, input *model.AddUserInput) (err error)
		Edit(ctx context.Context, input *model.EditUserInput) (err error)
		GetUserById(ctx context.Context, id uint) (out *model.UserInfoOut, err error)
		DelInfoById(ctx context.Context, id uint) (err error)
		ResetPassword(ctx context.Context, id uint, userPassword string) (err error)
		EditUserStatus(ctx context.Context, id uint, status uint) (err error)
		GetUserByIds(ctx context.Context, id []int) (data []*entity.SysUser, err error)
		GetAll(ctx context.Context) (data []*entity.SysUser, err error)
		CurrentUser(ctx context.Context) (userInfoOut *model.UserInfoOut, menuTreeOut []*model.UserMenuTreeOut, err error)
		EditUserAvatar(ctx context.Context, id uint, avatar string) (err error)
		EditUserInfo(ctx context.Context, input *model.EditUserInfoInput) (err error)
	}
	ISysAuthorize interface {
		AuthorizeQuery(ctx context.Context, itemsType string, menuIds []int) (out []*model.AuthorizeQueryTreeOut, err error)
		GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysAuthorize, err error)
		GetInfoByRoleIds(ctx context.Context, roleIds []int) (data []*entity.SysAuthorize, err error)
		GetInfoByRoleIdsAndItemsType(ctx context.Context, roleIds []int, itemsType string) (data []*entity.SysAuthorize, err error)
		DelByRoleId(ctx context.Context, roleId int) (err error)
		Add(ctx context.Context, authorize []*entity.SysAuthorize) (err error)
		AddAuthorize(ctx context.Context, roleId int, menuIds []string, buttonIds []string, columnIds []string, apiIds []string) (err error)
		IsAllowAuthorize(ctx context.Context, roleId int) (isAllow bool, err error)
		InitAuthorize(ctx context.Context) (err error)
		FilterDataByPermissions(ctx context.Context, model *gdb.Model) (*gdb.Model, error)
		GetDataWhere(ctx context.Context) (where g.Map, err error)
	}
	ISysUserPost interface {
		GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserPost, err error)
	}
	ISysUserOnline interface {
		Invoke(ctx context.Context, data *entity.SysUserOnline)
		Add(ctx context.Context, data *entity.SysUserOnline)
		DelByToken(ctx context.Context, token string) (err error)
		GetInfoByToken(ctx context.Context, token string) (data *entity.SysUserOnline, err error)
		DelByIds(ctx context.Context, ids []uint) (err error)
		GetAll(ctx context.Context) (data []*entity.SysUserOnline, err error)
		UserOnlineList(ctx context.Context, input *model.UserOnlineDoListInput) (total int, out []*model.UserOnlineListOut, err error)
		UserOnlineStrongBack(ctx context.Context, id int) (err error)
	}
	ISysMenuButton interface {
		GetList(ctx context.Context, status int, name string, menuId int) (data []model.UserMenuButtonRes, err error)
		GetData(ctx context.Context, status int, name string, menuId int, menuButton []model.UserMenuButtonRes) (data []model.UserMenuButtonRes, err error)
		Add(ctx context.Context, input *model.AddMenuButtonInput) (err error)
		Detail(ctx context.Context, Id int64) (entity *entity.SysMenuButton, err error)
		Edit(ctx context.Context, input *model.EditMenuButtonInput) (err error)
		Del(ctx context.Context, id int64) (err error)
		GetInfoByButtonIds(ctx context.Context, ids []int) (data []*entity.SysMenuButton, err error)
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuButton, err error)
		GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuButton, err error)
		GetAll(ctx context.Context) (data []*entity.SysMenuButton, err error)
		EditStatus(ctx context.Context, id int, menuId int, status int) (err error)
	}
	ISysMenuColumn interface {
		GetList(ctx context.Context, input *model.MenuColumnDoInput) (data []model.UserMenuColumnRes, err error)
		GetData(ctx context.Context, input *model.MenuColumnDoInput, menuColumn []model.UserMenuColumnRes) (data []model.UserMenuColumnRes, err error)
		Add(ctx context.Context, input *model.AddMenuColumnInput) (err error)
		Detail(ctx context.Context, Id int64) (entity *entity.SysMenuColumn, err error)
		Edit(ctx context.Context, input *model.EditMenuColumnInput) (err error)
		Del(ctx context.Context, Id int64) (err error)
		EditStatus(ctx context.Context, id int, menuId int, status int) (err error)
		GetInfoByColumnIds(ctx context.Context, ids []int) (data []*entity.SysMenuColumn, err error)
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuColumn, err error)
		GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuColumn, err error)
		GetAll(ctx context.Context) (data []*entity.SysMenuColumn, err error)
	}
	ISysRoleDept interface {
		GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysRoleDept, err error)
	}
	ISysMenu interface {
		GetAll(ctx context.Context) (data []*entity.SysMenu, err error)
		GetTree(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error)
		Add(ctx context.Context, input *model.AddMenuInput) (err error)
		Detail(ctx context.Context, menuId int64) (entity *entity.SysMenu, err error)
		Edit(ctx context.Context, input *model.EditMenuInput) (err error)
		Del(ctx context.Context, menuId int64) (err error)
		GetData(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error)
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenu, err error)
		GetInfoById(ctx context.Context, id int) (data *entity.SysMenu, err error)
	}
)

var (
	localSysMenu             ISysMenu
	localSysMenuButton       ISysMenuButton
	localSysMenuColumn       ISysMenuColumn
	localSysRoleDept         ISysRoleDept
	localCaptcha             ICaptcha
	localLogin               ILogin
	localSysJob              ISysJob
	localSysOperLog          ISysOperLog
	localSystemPluginsConfig ISystemPluginsConfig
	localSysUserRole         ISysUserRole
	localSysApi              ISysApi
	localSysDept             ISysDept
	localSysNotifications    ISysNotifications
	localSysPlugins          ISysPlugins
	localSysUser             ISysUser
	localSysAuthorize        ISysAuthorize
	localSysLoginLog         ISysLoginLog
	localSysMenuApi          ISysMenuApi
	localSysOrganization     ISysOrganization
	localSysPost             ISysPost
	localSysRole             ISysRole
	localSysToken            ISysToken
	localSysUserOnline       ISysUserOnline
	localSysUserPost         ISysUserPost
)

func SysApi() ISysApi {
	if localSysApi == nil {
		panic("implement not found for interface ISysApi, forgot register?")
	}
	return localSysApi
}

func RegisterSysApi(i ISysApi) {
	localSysApi = i
}

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}

func SysNotifications() ISysNotifications {
	if localSysNotifications == nil {
		panic("implement not found for interface ISysNotifications, forgot register?")
	}
	return localSysNotifications
}

func RegisterSysNotifications(i ISysNotifications) {
	localSysNotifications = i
}

func SysPlugins() ISysPlugins {
	if localSysPlugins == nil {
		panic("implement not found for interface ISysPlugins, forgot register?")
	}
	return localSysPlugins
}

func RegisterSysPlugins(i ISysPlugins) {
	localSysPlugins = i
}

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}

func SysToken() ISysToken {
	if localSysToken == nil {
		panic("implement not found for interface ISysToken, forgot register?")
	}
	return localSysToken
}

func RegisterSysToken(i ISysToken) {
	localSysToken = i
}

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

func SysAuthorize() ISysAuthorize {
	if localSysAuthorize == nil {
		panic("implement not found for interface ISysAuthorize, forgot register?")
	}
	return localSysAuthorize
}

func RegisterSysAuthorize(i ISysAuthorize) {
	localSysAuthorize = i
}

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}

func SysMenuApi() ISysMenuApi {
	if localSysMenuApi == nil {
		panic("implement not found for interface ISysMenuApi, forgot register?")
	}
	return localSysMenuApi
}

func RegisterSysMenuApi(i ISysMenuApi) {
	localSysMenuApi = i
}

func SysOrganization() ISysOrganization {
	if localSysOrganization == nil {
		panic("implement not found for interface ISysOrganization, forgot register?")
	}
	return localSysOrganization
}

func RegisterSysOrganization(i ISysOrganization) {
	localSysOrganization = i
}

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}

func SysUserPost() ISysUserPost {
	if localSysUserPost == nil {
		panic("implement not found for interface ISysUserPost, forgot register?")
	}
	return localSysUserPost
}

func RegisterSysUserPost(i ISysUserPost) {
	localSysUserPost = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysMenuButton() ISysMenuButton {
	if localSysMenuButton == nil {
		panic("implement not found for interface ISysMenuButton, forgot register?")
	}
	return localSysMenuButton
}

func RegisterSysMenuButton(i ISysMenuButton) {
	localSysMenuButton = i
}

func SysMenuColumn() ISysMenuColumn {
	if localSysMenuColumn == nil {
		panic("implement not found for interface ISysMenuColumn, forgot register?")
	}
	return localSysMenuColumn
}

func RegisterSysMenuColumn(i ISysMenuColumn) {
	localSysMenuColumn = i
}

func SysRoleDept() ISysRoleDept {
	if localSysRoleDept == nil {
		panic("implement not found for interface ISysRoleDept, forgot register?")
	}
	return localSysRoleDept
}

func RegisterSysRoleDept(i ISysRoleDept) {
	localSysRoleDept = i
}

func SystemPluginsConfig() ISystemPluginsConfig {
	if localSystemPluginsConfig == nil {
		panic("implement not found for interface ISystemPluginsConfig, forgot register?")
	}
	return localSystemPluginsConfig
}

func RegisterSystemPluginsConfig(i ISystemPluginsConfig) {
	localSystemPluginsConfig = i
}

func SysUserRole() ISysUserRole {
	if localSysUserRole == nil {
		panic("implement not found for interface ISysUserRole, forgot register?")
	}
	return localSysUserRole
}

func RegisterSysUserRole(i ISysUserRole) {
	localSysUserRole = i
}

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

func SysJob() ISysJob {
	if localSysJob == nil {
		panic("implement not found for interface ISysJob, forgot register?")
	}
	return localSysJob
}

func RegisterSysJob(i ISysJob) {
	localSysJob = i
}

func SysOperLog() ISysOperLog {
	if localSysOperLog == nil {
		panic("implement not found for interface ISysOperLog, forgot register?")
	}
	return localSysOperLog
}

func RegisterSysOperLog(i ISysOperLog) {
	localSysOperLog = i
}
