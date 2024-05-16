package rent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ResourceRouter struct {
}

// InitResourceRouter 初始化 资源 路由信息
func (s *ResourceRouter) InitResourceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	resourceRouter := Router.Group("resource").Use(middleware.OperationRecord())
	resourceRouterWithoutRecord := Router.Group("resource")
	resourceRouterWithoutAuth := PublicRouter.Group("resource")

	var resourceApi = v1.ApiGroupApp.RentApiGroup.ResourceApi
	{
		resourceRouter.POST("createResource", resourceApi.CreateResource)   // 新建资源
		resourceRouter.DELETE("deleteResource", resourceApi.DeleteResource) // 删除资源
		resourceRouter.DELETE("deleteResourceByIds", resourceApi.DeleteResourceByIds) // 批量删除资源
		resourceRouter.PUT("updateResource", resourceApi.UpdateResource)    // 更新资源
	}
	{
		resourceRouterWithoutRecord.GET("findResource", resourceApi.FindResource)        // 根据ID获取资源
		resourceRouterWithoutRecord.GET("getResourceList", resourceApi.GetResourceList)  // 获取资源列表
	}
	{
	    resourceRouterWithoutAuth.GET("getResourcePublic", resourceApi.GetResourcePublic)  // 获取资源列表
	}
}
