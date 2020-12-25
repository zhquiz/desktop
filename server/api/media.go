package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhquiz/go-server/shared"
)

type tMediaRouter struct {
	Router *gin.RouterGroup
}

func (r tMediaRouter) init() {
	r.doUpload()
}

func (r tMediaRouter) doUpload() {
	r.Router.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.AbortWithError(400, err)
		}

		ctx.SaveUploadedFile(file, shared.Paths().MediaPath())

		ctx.JSON(201, gin.H{
			"url": "/media/" + file.Filename,
		})
	})
}
