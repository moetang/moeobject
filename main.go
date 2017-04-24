package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/moetang/moeobject/config"
	"github.com/moetang/moeobject/file"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "", "-config=moeobject.config")
}

func main() {
	flag.Parse()

	if configPath == "" {
		fmt.Println("need config file.")
		return
	}

	f, err := os.Open(configPath)
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
