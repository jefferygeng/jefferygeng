package response

import (
	"github.com/jefferygeng/yj/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
