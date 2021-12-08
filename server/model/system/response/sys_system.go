package response

import "github.com/jefferygeng/yj/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
