package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type ResourceSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Name  string `json:"name" form:"name" `
                      Type  string `json:"type" form:"type" `
                      Landlord_id  *int `json:"landlord_id" form:"landlord_id" `
    request.PageInfo
}
