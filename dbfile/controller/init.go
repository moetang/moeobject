package controller

import (
	scaffold "github.com/moetang/webapp-scaffold"

	limits "github.com/gin-contrib/size"
)

func Init(webscaff *scaffold.WebappScaffold) {
	api := webscaff.GetGin().Group("/api/v1")

	api.Use(limits.RequestSizeLimiter(1*1024*1024 + 512*1024)) // 1.5MB

	// post file
	api.POST("/file", FileUpload)
	// get file by id
	api.GET("/file/:file_id", GetFile)
}
