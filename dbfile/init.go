package dbfile

import (
	"github.com/moetang/moeobject/dbfile/controller"
	"github.com/moetang/moeobject/dbfile/model"

	scaffold "github.com/moetang/webapp-scaffold"

	"github.com/gin-gonic/gin"
)

func SyncStart(webscaf *scaffold.WebappScaffold) {
	webscaf.GetGin().Use(gin.Logger())
	webscaf.GetGin().Use(gin.Recovery())

	model.Init(webscaf)
	controller.Init(webscaf)

	err := webscaf.SyncStart()
	if err != nil {
		panic(err)
	}
}
