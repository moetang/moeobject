package main

import (
	"encoding/json"
	"github.com/moetang/moeobject/dbfile"
	"io/ioutil"
	"os"

	"github.com/moetang/moeobject/config"
	"github.com/moetang/moeobject/file"

	scaffold "github.com/moetang/webapp-scaffold"
)

func main() {
	webscaff, err := scaffold.NewFromConfigFile("moeobject.toml")
	if err != nil {
		panic(err)
	}
	dbfile.SyncStart(webscaff)
}

func todo() {
	f, err := os.Open("")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	c := new(config.Config)
	json.Unmarshal(data, c)

	vol := file.NewVolume(c.Volumes[0])
	vol.Close()
}
