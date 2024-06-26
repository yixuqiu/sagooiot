// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta         `orm:"table:sys_dict_type, do:true"`
	DictId         interface{} // 字典主键
	ParentId       interface{} // 父主键ID
	DictName       interface{} // 字典名称
	DictType       interface{} // 字典类型
	ModuleClassify interface{} // 模块分类
	Remark         interface{} // 备注
	Status         interface{} // 状态（0正常 1停用）
	IsDeleted      interface{} // 是否删除 0未删除 1已删除
	CreatedBy      interface{} // 创建者
	CreatedAt      *gtime.Time // 创建日期
	UpdatedBy      interface{} // 更新者
	UpdatedAt      *gtime.Time // 修改日期
	DeletedBy      interface{} // 删除人
	DeletedAt      *gtime.Time // 删除时间
}
