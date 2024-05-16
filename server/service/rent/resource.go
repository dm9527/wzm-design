package rent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rent"
    rentReq "github.com/flipped-aurora/gin-vue-admin/server/model/rent/request"
    "gorm.io/gorm"
)

type ResourceService struct {
}

// CreateResource 创建资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService) CreateResource(resource *rent.Resource) (err error) {
	err = global.GVA_DB.Create(resource).Error
	return err
}

// DeleteResource 删除资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)DeleteResource(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&rent.Resource{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&rent.Resource{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteResourceByIds 批量删除资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)DeleteResourceByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&rent.Resource{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&rent.Resource{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateResource 更新资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)UpdateResource(resource rent.Resource) (err error) {
	err = global.GVA_DB.Model(&rent.Resource{}).Where("id = ?",resource.ID).Updates(&resource).Error
	return err
}

// GetResource 根据ID获取资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)GetResource(ID string) (resource rent.Resource, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&resource).Error
	return
}

// GetResourceInfoList 分页获取资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)GetResourceInfoList(info rentReq.ResourceSearch) (list []rent.Resource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&rent.Resource{})
    var resources []rent.Resource
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Type != "" {
        db = db.Where("type = ?",info.Type)
    }
    if info.Landlord_id != nil {
        db = db.Where("landlord_id = ?",info.Landlord_id)
    }
    if info.Location != "" {
        db = db.Where("location LIKE ?","%"+ info.Location+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&resources).Error
	return  resources, total, err
}