package controller

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/moetang/moeobject/dbfile/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	RespStatusSucc        = 0
	RespStatusErrParam    = 1
	RespStatusErrOversize = 2
	RespStatusErrInternal = 3
)

const (
	MaxFileSize = 1 * 1024 * 1024
)

func FileUpload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("[ERROR] parse multipart form error.", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrParam,
			"errormessage": "param error",
		})
		return
	}

	lf := form.File["file"]
	if len(lf) != 1 {
		log.Println("[ERROR] need exact one file.")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrParam,
			"errormessage": "param error",
		})
		return
	}
	fh := lf[0]
	if fh.Size > MaxFileSize {
		log.Println("[ERROR] file oversize.")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrOversize,
			"errormessage": "param error",
		})
		return
	}

	f, err := fh.Open()
	if err != nil {
		log.Println("[ERROR] open uploaded file error.", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrInternal,
			"errormessage": "internal error",
		})
		return
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Println("[ERROR] read file error.", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrInternal,
			"errormessage": "internal error",
		})
		return
	}

	df := new(model.Dbfile)
	df.FileId = uuid.NewV4().String()
	df.Filecontent = b
	err = model.SaveDbFile(df)
	if err != nil {
		log.Println("[ERROR] save db file error.", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":       RespStatusErrInternal,
			"errormessage": "internal error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  RespStatusSucc,
		"file_id": df.FileId,
	})
}

func GetFile(ctx *gin.Context) {
	fileId := ctx.Param("file_id")

	f, err := model.LoadDbFileByFileId(fileId)
	if err != nil {
		log.Println("[ERROR] load file error.", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-type", "application/octet-stream")
	_, err = io.Copy(ctx.Writer, bytes.NewReader(f.Filecontent))
	if err != nil {
		log.Println("[ERROR] send error.", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
}
