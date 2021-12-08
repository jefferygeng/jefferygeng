package example

import "github.com/jefferygeng/yj/server/service"

type ApiGroup struct {
	CustomerApi
	ExcelApi
	FileUploadAndDownloadApi
}

var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
var customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
var excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
