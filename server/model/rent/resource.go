// 自动生成模板Resource
package rent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 资源 结构体  Resource
type Resource struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:资源名称;" binding:"required"`  //名称 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:资源类型;" binding:"required"`  //资源类型 
      Landlord_id  *int `json:"landlord_id" form:"landlord_id" gorm:"column:landlord_id;comment:房东id;" binding:"required"`  //房东id 
      Location  string `json:"location" form:"location" gorm:"column:location;comment:资源位置;" binding:"required"`  //资源位置 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 资源 Resource自定义表名 Resource
func (Resource) TableName() string {
  return "Resource"
}

