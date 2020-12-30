package api

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhquiz/go-server/server/db"
	"github.com/zhquiz/go-server/server/util"
	"gopkg.in/sakura-internet/go-rison.v3"
)

func routerQuiz(apiRouter *gin.RouterGroup) {
	r := apiRouter.Group("/quiz")

	r.GET("/many", func(ctx *gin.Context) {
		userID := getUserID(ctx)
		if userID == "" {
			ctx.AbortWithStatus(401)
			return
		}

		var query struct {
			IDs     string `form:"ids"`
			Entries string `form:"entries"`
			Type    string `form:"type" binding:"oneof=hanzi vocab sentence extra"`
			Select  string `form:"select"`
		}

		if e := ctx.BindQuery(&query); e != nil {
			panic(e)
		}

		sel := []string{}
		sMap := map[string]string{
			"id":        "ID",
			"tag":       "Tag",
			"entry":     "[Entry]",
			"type":      "[Type]",
			"direction": "Direction",
			"front":     "Front",
			"back":      "Back",
			"mnemonic":  "Mnemonic",
		}

		for _, s := range strings.Split(query.Select, ",") {
			k := sMap[s]
			if k != "" {
				sel = append(sel, k)
			}
		}

		if len(sel) == 0 {
			ctx.AbortWithError(400, fmt.Errorf("not enough select"))
			return
		}

		var out []gin.H

		ids := strings.Split(query.IDs, ",")
		entries := strings.Split(query.Entries, ",")

		if len(ids) > 0 {
			if r := resource.DB.Current.Model(&db.Quiz{}).
				Select(sel).
				Where("user_id = ? AND id IN ?", userID, ids).
				Find(&out); r.Error != nil {
				panic(r.Error)
			}
		} else if len(entries) > 0 && query.Type != "" {
			if r := resource.DB.Current.Model(&db.Quiz{}).
				Select(sel).
				Where("user_id = ? AND [Type] = ? AND entry IN ?", userID, query.Type, entries).
				Find(&out); r.Error != nil {
				panic(r.Error)
			}
		} else {
			ctx.AbortWithError(400, fmt.Errorf("either IDs or Entries must be specified"))
			return
		}

		ctx.JSON(200, gin.H{
			"result": out,
		})
	})

	r.PATCH("/mark", func(ctx *gin.Context) {
		userID := getUserID(ctx)
		if userID == "" {
			ctx.AbortWithStatus(401)
			return
		}

		var query struct {
			ID   string `form:"id" binding:"required"`
			Type string `form:"type" binding:"required;oneof=right wrong repeat"`
		}

		if e := ctx.ShouldBindQuery(&query); e != nil {
			ctx.AbortWithError(400, e)
			return
		}

		var quiz db.Quiz
		if r := resource.DB.Current.
			Where("user_id = ? AND id = ?", userID, query.ID).
			First(&quiz); r.Error != nil {
			panic(r.Error)
		}

		quiz.UpdateSRSLevel(map[string]int8{
			"right":  1,
			"wrong":  -1,
			"repeat": 0,
		}[query.Type])

		if r := resource.DB.Current.Save(&quiz); r.Error != nil {
			panic(r.Error)
		}

		ctx.JSON(201, gin.H{
			"result": "updated",
		})
	})

	r.GET("/allTags", func(ctx *gin.Context) {
		userID := getUserID(ctx)
		if userID == "" {
			ctx.AbortWithStatus(401)
			return
		}

		var tagEls []struct {
			Name string
		}

		if r := resource.DB.Current.Model(&db.Quiz{}).
			Select("tag.Name").
			Joins("JOIN quiz_tag ON quiz_tag.quiz_id = quiz.id").
			Joins("JOIN tag ON quiz_tag.tag_id = tag.id").
			Group("tag.Name").
			Scan(&tagEls); r.Error != nil {
			panic(r.Error)
		}

		var result []string
		for _, t := range tagEls {
			result = append(result, t.Name)
		}

		ctx.JSON(200, gin.H{
			"result": result,
		})
	})

	r.GET("/init", func(ctx *gin.Context) {
		userID := getUserID(ctx)
		if userID == "" {
			ctx.AbortWithStatus(401)
			return
		}

		var query struct {
			RS string `form:"_"`
		}

		var rs struct {
			Type      []string `json:"type" validate:"required,min=1"`
			Stage     []string `json:"stage" validate:"required,min=1"`
			Direction []string `json:"direction" validate:"required,min=1"`
			IsDue     bool     `json:"isDue" validate:"required"`
			Tag       []string `json:"tag" validate:"required"`
		}

		if e := ctx.ShouldBindQuery(&query); e != nil {
			ctx.AbortWithError(400, e)
			return
		}

		if e := rison.Unmarshal([]byte(query.RS), &rs, rison.Rison); e != nil {
			ctx.AbortWithError(400, e)
			return
		}

		if e := validate.Struct(&rs); e != nil {
			ctx.AbortWithError(400, e)
			return
		}

		// No need to await
		go func() {
			var user db.User
			if r := resource.DB.Current.Where("id = ?", userID).First(&user); r.Error != nil {
				panic(r.Error)
			}

			user.Meta.Quiz.Direction = rs.Direction
			user.Meta.Quiz.Stage = rs.Stage
			user.Meta.Quiz.Type = rs.Type
			user.Meta.Quiz.IsDue = rs.IsDue

			if r := resource.DB.Current.Save(&user); r.Error != nil {
				panic(r.Error)
			}
		}()

		var orCond []string

		stageSet := util.MakeSet(rs.Stage)
		if stageSet["new"] {
			orCond = append(orCond, "srs_level IS NULL")
		}

		if stageSet["leech"] {
			orCond = append(orCond, "wrong_streak >= 3")
		}

		if stageSet["learning"] {
			orCond = append(orCond, "srs_level < 3")
		}

		if stageSet["graduated"] {
			orCond = append(orCond, "srs_level >= 3")
		}

		q := resource.DB.Current.
			Model(&db.Quiz{}).
			Joins("LEFT JOIN quiz_tag ON quiz_tag.quiz_id = quiz.id").
			Joins("LEFT JOIN tag ON tag.id = quiz_tag.tag_id").
			Where("user_id = ? AND [type] IN ? AND direction IN ?", userID, rs.Type, rs.Direction)

		if len(rs.Tag) > 0 {
			q = q.Where("tag.name IN ?", rs.Tag)
		}

		if len(orCond) > 0 {
			q = q.Where(strings.Join(orCond, " OR "))
		}

		var quizzes []db.Quiz

		if r := q.Group("quizzes.id").Find(&quizzes); r.Error != nil {
			panic(r.Error)
		}

		var quiz []quizInitOutput
		var upcoming []quizInitOutput

		if rs.IsDue {
			now := time.Now()

			for _, it := range quizzes {
				if it.NextReview == nil || (*it.NextReview).Before(now) {
					quiz = append(quiz, quizInitOutput{
						NextReview:  it.NextReview,
						SRSLevel:    it.SRSLevel,
						WrongStreak: it.WrongStreak,
						ID:          it.ID,
					})
				} else {
					upcoming = append(upcoming, quizInitOutput{
						NextReview: it.NextReview,
						ID:         it.ID,
					})
				}
			}
		} else {
			for _, it := range quizzes {
				quiz = append(quiz, quizInitOutput{
					NextReview:  it.NextReview,
					SRSLevel:    it.SRSLevel,
					WrongStreak: it.WrongStreak,
					ID:          it.ID,
				})
			}
		}

		rand.Shuffle(len(quiz), func(i, j int) {
			quiz[i], quiz[j] = quiz[j], quiz[i]
		})

		sort.Sort(quizInitOutputList(upcoming))

		if len(quiz) == 0 {
			quiz = make([]quizInitOutput, 0)
		}

		if len(upcoming) == 0 {
			upcoming = make([]quizInitOutput, 0)
		}

		ctx.JSON(200, gin.H{
			"quiz":     quiz,
			"upcoming": upcoming,
		})
	})
}

type quizInitOutput struct {
	NextReview  *time.Time `json:"nextReview"`
	SRSLevel    *int8      `json:"srsLevel"`
	WrongStreak *uint      `json:"wrongStreak"`
	ID          string     `json:"id"`
}

type quizInitOutputList []quizInitOutput

func (ls quizInitOutputList) Len() int {
	return len(ls)
}

func (ls quizInitOutputList) Less(i, j int) bool {
	a, b := ls[i], ls[j]
	if a.NextReview == nil {
		return true
	}
	if b.NextReview == nil {
		return false
	}

	return a.NextReview.Before(*b.NextReview)
}

func (ls quizInitOutputList) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}
