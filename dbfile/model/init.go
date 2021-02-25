package model

import (
	scaffold "github.com/moetang/webapp-scaffold"
)

var db scaffold.PostgresApi

func Init(webscaff *scaffold.WebappScaffold) {
	db = webscaff
}
