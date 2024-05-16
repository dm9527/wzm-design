import service from '@/utils/request'

// @Tags Resource
// @Summary 创建资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resource true "创建资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /resource/createResource [post]
export const createResource = (data) => {
  return service({
    url: '/resource/createResource',
    method: 'post',
    data
  })
}

// @Tags Resource
// @Summary 删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resource true "删除资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/deleteResource [delete]
export const deleteResource = (params) => {
  return service({
    url: '/resource/deleteResource',
    method: 'delete',
    params
  })
}

// @Tags Resource
// @Summary 批量删除资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/deleteResource [delete]
export const deleteResourceByIds = (params) => {
  return service({
    url: '/resource/deleteResourceByIds',
    method: 'delete',
    params
  })
}

// @Tags Resource
// @Summary 更新资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resource true "更新资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resource/updateResource [put]
export const updateResource = (data) => {
  return service({
    url: '/resource/updateResource',
    method: 'put',
    data
  })
}

// @Tags Resource
// @Summary 用id查询资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Resource true "用id查询资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resource/findResource [get]
export const findResource = (params) => {
  return service({
    url: '/resource/findResource',
    method: 'get',
    params
  })
}

// @Tags Resource
// @Summary 分页获取资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/getResourceList [get]
export const getResourceList = (params) => {
  return service({
    url: '/resource/getResourceList',
    method: 'get',
    params
  })
}
