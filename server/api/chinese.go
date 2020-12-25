package api

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/tebeka/atexit"
	"github.com/yanyiwu/gojieba"
)

type tChineseRouter struct {
	Router *gin.RouterGroup
	jieba  *gojieba.Jieba
}

func (r tChineseRouter) init() {
	r.jieba = gojieba.NewJieba()
	atexit.Register(r.jieba.Free)

	r.getJieba()
}

func (r tChineseRouter) getJieba() {
	r.Router.GET("/jieba", cache.CachePage(store, time.Hour, func(ctx *gin.Context) {
		var query struct {
			Q string
		}

		if e := ctx.ShouldBindQuery(&query); e != nil {
			ctx.AbortWithError(400, e)
		}

		ctx.JSON(200, gin.H{
			"result": r.jieba.CutAll(query.Q),
		})
	}))
}
