package api

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhquiz/go-server/server/db"
	"gopkg.in/sakura-internet/go-rison.v3"
)

type tExtraRouter struct {
	Router *gin.RouterGroup
}

func (r tExtraRouter) init() {
	r.getQ()
	r.getMatch()
	r.doCreate()
	r.doUpdate()
	r.doDelete()
}

func (r tExtraRouter) getQ() {
	r.Router.GET("/q", cache.CachePage(store, time.Hour, func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("userID").(string)
		if userID == "" {
			ctx.AbortWithStatus(401)
		}

		var query struct {
			RS string `form:"_"`
		}

		var rs struct {
			Select  []string
			Sort    *string
			Page    *int
			PerPage *int `json:"perPage"`
		}

		if e := ctx.ShouldBindQuery(&query); e != nil {
			panic(e)
		}

		if e := rison.Unmarshal([]byte(query.RS), &rs, rison.Rison); e != nil {
			panic(e)
		}

		// if e := validate.Struct(&rs); e != nil {
		// 	panic(e)
		// }

		if rs.Sort == nil {
			*rs.Sort = "-updatedAt"
		}

		if string((*rs.Sort)[0]) == "-" {
			*rs.Sort = string((*rs.Sort)[1:]) + " desc"
		}

		if rs.Page == nil {
			*rs.Page = 1
		}

		if rs.PerPage == nil {
			*rs.PerPage = 10
		}

		sel := []string{}
		sMap := map[string]string{
			"chinese": "Chinese",
			"pinyin":  "Pinyin",
			"english": "English",
		}

		for _, s := range rs.Select {
			k := sMap[s]
			if k != "" {
				sel = append(sel, k)
			}
		}

		if len(sel) == 0 {
			sel = []string{"Chinese", "Pinyin", "English"}
		}

		var out struct {
			result []gin.H
			count  int
		}

		if r := resource.DB.Current.
			Model(&db.Extra{}).
			Select("COUNT(ID) AS [Count]").
			Where("userID = ?", userID).
			Find(&out); r.Error != nil {
			panic(r.Error)
		}

		if r := resource.DB.Current.
			Model(&db.Extra{}).
			Select(sel).
			Order(*rs.Sort).
			Limit(*rs.PerPage).
			Offset((*rs.Page-1)**rs.PerPage).
			Where("userID = ?", userID).
			Find(&out.result); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(200, out)
	}))
}

func (r tExtraRouter) getMatch() {
	r.Router.GET("/", cache.CachePage(store, time.Hour, func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("userID").(string)
		if userID == "" {
			ctx.AbortWithStatus(401)
		}

		var query struct {
			Entry string
			RS    string `form:"_"`
		}

		var rs struct {
			Select  []string
			Sort    *string
			Page    *int
			PerPage *int `json:"perPage"`
		}

		if e := ctx.ShouldBindQuery(&query); e != nil {
			panic(e)
		}

		if e := rison.Unmarshal([]byte(query.RS), &rs, rison.Rison); e != nil {
			panic(e)
		}

		sel := []string{}
		sMap := map[string]string{
			"chinese": "Chinese",
			"pinyin":  "Pinyin",
			"english": "English",
		}

		for _, s := range rs.Select {
			k := sMap[s]
			if k != "" {
				sel = append(sel, k)
			}
		}

		if len(sel) == 0 {
			sel = []string{"Chinese", "Pinyin", "English"}
		}

		var out gin.H

		if r := resource.DB.Current.
			Model(&db.Extra{}).
			Select(sel).
			Where("userID = ?", userID).
			First(&out); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(200, out)
	}))
}

func (r tExtraRouter) doCreate() {
	r.Router.PUT("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("userID").(string)
		if userID == "" {
			ctx.AbortWithStatus(401)
		}

		var body struct {
			Chinese string `binding:"required"`
			Pinyin  string `binding:"required"`
			English string `binding:"required"`
		}

		if e := ctx.BindJSON(&body); e != nil {
			panic(e)
		}

		checkVocab := func() bool {
			stmt, e := resource.Zh.Current.Prepare(`
			SELECT simplified FROM cedict
			WHERE simplified = ? OR traditional = ?
			`)

			if e != nil {
				panic(e)
			}

			r := stmt.QueryRow(body.Chinese, body.Chinese)

			if e := r.Err(); e != nil {
				panic(e)
			}

			var entry string
			if e := r.Scan(entry); e == sql.ErrNoRows {
				return false
			}

			ctx.JSON(200, gin.H{
				"existing": gin.H{
					"type":  "vocab",
					"entry": entry,
				},
			})

			return true
		}

		checkHanzi := func() bool {
			stmt, e := resource.Zh.Current.Prepare(`
			SELECT entry FROM token
			WHERE entry = ? AND english IS NOT NULL
			`)

			if e != nil {
				panic(e)
			}

			r := stmt.QueryRow(body.Chinese)

			if e := r.Err(); e != nil {
				panic(e)
			}

			var entry string
			if e := r.Scan(entry); e == sql.ErrNoRows {
				return false
			}

			ctx.JSON(200, gin.H{
				"existing": gin.H{
					"type":  "hanzi",
					"entry": entry,
				},
			})

			return true
		}

		checkSentence := func() bool {
			stmt, e := resource.Zh.Current.Prepare(`
			SELECT chinese FROM sentence
			WHERE chinese = ?
			`)

			if e != nil {
				panic(e)
			}

			r := stmt.QueryRow(body.Chinese)

			if e := r.Err(); e != nil {
				panic(e)
			}

			var entry string
			if e := r.Scan(entry); e == sql.ErrNoRows {
				return false
			}

			ctx.JSON(200, gin.H{
				"existing": gin.H{
					"type":  "sentence",
					"entry": entry,
				},
			})

			return true
		}

		if checkVocab() {
			return
		}

		if len([]rune(body.Chinese)) == 1 {
			if checkHanzi() {
				return
			}
		} else {
			if checkSentence() {
				return
			}
		}

		it := db.Extra{
			Chinese: body.Chinese,
			Pinyin:  body.Pinyin,
			English: body.English,
			UserID:  userID,
		}

		if r := resource.DB.Current.Create(&it); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(201, gin.H{
			"id": it.ID,
		})
	})
}

func (r tExtraRouter) doUpdate() {
	r.Router.PUT("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("userID").(string)
		if userID == "" {
			ctx.AbortWithStatus(401)
		}

		id := ctx.Query("id")
		if id == "" {
			ctx.AbortWithError(400, fmt.Errorf("id to update not specified"))
		}

		var body struct {
			Chinese string `binding:"required"`
			Pinyin  string `binding:"required"`
			English string `binding:"required"`
		}

		if e := ctx.BindJSON(&body); e != nil {
			panic(e)
		}

		var it db.Extra

		if r := resource.DB.Current.Where("UserID = ? AND ID = ?", userID, id).First(&it); r.Error != nil {
			panic(r.Error)
		}

		it.Chinese = body.Chinese
		it.Pinyin = body.Pinyin
		it.English = body.English

		if r := resource.DB.Current.Save(&it); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(201, gin.H{
			"result": "updated",
		})
	})
}

func (r tExtraRouter) doDelete() {
	r.Router.PUT("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userID := session.Get("userID").(string)
		if userID == "" {
			ctx.AbortWithStatus(401)
		}

		id := ctx.Query("id")
		if id == "" {
			ctx.AbortWithError(400, fmt.Errorf("id to update not specified"))
		}

		if r := resource.DB.Current.Where("UserID = ? AND ID = ?", userID, id).Delete(&db.Extra{}); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(201, gin.H{
			"result": "deleted",
		})
	})
}