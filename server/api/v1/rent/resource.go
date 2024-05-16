package rent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/rent"
    rentReq "github.com/flipped-aurora/gin-vue-admin/server/model/rent/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ResourceApi struct {
}

var resourceService = service.ServiceGroupApp.RentServiceGroup.ResourceService


// CreateResource 创建资源
// @Tags Resource
// @Summary 创建资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rent.Resource true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /resource/createResource [post]
func (resourceApi *ResourceApi) CreateResource(c *gin.Context) {
	var resource rent.Resource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    resource.CreatedBy = utils.GetUserID(c)

	if err := resourceService.CreateResource(&resource); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteResource 删除资源
// @Tags Resource
// @Summary 删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rent.Resource true "删除资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/deleteResource [delete]
func (resourceApi *ResourceApi) DeleteResource(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := resourceService.DeleteResource(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteResourceByIds 批量删除资源
// @Tags Resource
// @Summary 批量删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /resource/deleteResourceByIds [delete]
func (resourceApi *ResourceApi) DeleteResourceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := resourceService.DeleteResourceByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateResource 更新资源
// @Tags Resource
// @Summary 更新资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rent.Resource true "更新资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resource/updateResource [put]
func (resourceApi *ResourceApi) UpdateResource(c *gin.Context) {
	var resource rent.Resource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    resource.UpdatedBy = utils.GetUserID(c)

	if err := resourceService.UpdateResource(resource); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindResource 用id查询资源
// @Tags Resource
// @Summary 用id查询资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rent.Resource true "用id查询资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resource/findResource [get]
func (resourceApi *ResourceApi) FindResource(c *gin.Context) {
	ID := c.Query("ID")
	if reresource, err := resourceService.GetResource(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reresource": reresource}, c)
	}
}

// GetResourceList 分页获取资源列表
// @Tags Resource
// @Summary 分页获取资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rentReq.ResourceSearch true "分页获取资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/getResourceList [get]
func (resourceApi *ResourceApi) GetResourceList(c *gin.Context) {
	var pageInfo rentReq.ResourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := resourceService.GetResourceInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetResourcePublic 不需要鉴权的资源接口
// @Tags Resource
// @Summary 不需要鉴权的资源接口
// @accept application/json
// @Produce application/json
// @Param data query rentReq.ResourceSearch true "分页获取资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/getResourceList [get]
func (resourceApi *ResourceApi) GetResourcePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的资源接口信息",
    }, "获取成功", c)
}
