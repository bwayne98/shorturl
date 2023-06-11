package short

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"shorturl/m/db/store/shorturl"
	"shorturl/m/util"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	expaired_time = time.Hour * 24 * 30
)

type Controller struct {
	db *sql.DB
}

func New(db *sql.DB) *Controller {
	return &Controller{db}
}

type MakeReq struct {
	Origin string `json:"origin" binding:"required,startswith=https://"`
}

func (cont *Controller) Make(c *gin.Context) {

	var req MakeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// todo validate 套件
	if pass, err := regexp.MatchString("^https?://"+c.Request.Host, req.Origin); pass || err != nil {
		c.JSON(http.StatusBadRequest, "Origin Url not allowed.")
		return
	}

	shorturlQuery := shorturl.New(cont.db)

	var match string
	if err := getUniqueMatch(shorturlQuery, &match); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	shorturl, err := shorturlQuery.CreateShorturl(context.Background(), shorturl.CreateShorturlParams{
		Origin:    req.Origin,
		Match:     match,
		ExpiredAt: time.Now().Add(expaired_time),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// todo 解決 timezone
	c.JSON(http.StatusOK, gin.H{
		"id":        shorturl.ID,
		"shortUrl":  fmt.Sprintf("https://%s/%s", c.Request.Host, match),
		"expiredAt": shorturl.ExpiredAt.Add(time.Hour * 8).Format("2006-01-02 15:04:05"),
	})
}

func getUniqueMatch(query *shorturl.Queries, match *string) (err error) {
	exist := true
	for exist {
		*match = util.RandString(10)
		count, e := query.CountMatchShorturl(context.Background(), *match)
		if e != nil {
			err = e
			return
		}
		exist = count > 0
	}
	return
}

type MatchReq struct {
	Match string `json:"match" binding:"required,len=10"`
}

func (cont *Controller) Match(c *gin.Context) {
	var req MatchReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	shorturlQuery := shorturl.New(cont.db)
	redirectUrl, err := shorturlQuery.GetMatchShorturl(context.Background(), req.Match)

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redirectUrl": redirectUrl,
	})
}
