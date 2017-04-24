package file

import (
	"log"
	"os"

	"github.com/moetang/moeobject/config"
)

type Volumn struct {
	VolId     uint16
	VolPath   string
	VolFolder *os.File
}

func (this *Volumn) Close() error {
	var err error
	err = this.VolFolder.Close()
	if err != nil {
		log.Println("close volume folder error.", this.VolPath, err)
	}
	return nil
}

func NewVolume(volConf config.VolumeConfig) *Volumn {
	volPath := volConf.Path
	volFolder, err := os.Open(volPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Panicln("error occurs when open volume folder.", volFolder, err)
		}
		err = os.MkdirAll(volPath, os.FileMode(0755))
		if err != nil {
			log.Panicln("create volume folder error.", volPath, err)
		}
		log.Println("volumn folder created.", volPath)
		volFolder, err = os.Open(volPath)
		if err != nil {
			log.Panicln("open created volume folder error.", volPath, err)
		}
	}

	vol := new(Volumn)
	vol.VolPath = volPath
	vol.VolFolder = volFolder
	vol.VolId = volConf.VolId
	//TODO load exist files
	//TODO create init file
	//TODO return volume
	return vol
}
