package handlers

import (
	"net/http"
	"strconv"

	"github.com/JuHaNi654/pkg-reader/pkg/sqlite"
	"github.com/gin-gonic/gin"
)

func GetRouter(db *sqlite.SQLRepository) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", index(db))
	r.GET("/load", loadPkgs(db))

	return r
}

func index(db *sqlite.SQLRepository) gin.HandlerFunc {
	items, _ := db.GetItems(0)

	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"pkgs": items,
		})
	}
}

func loadPkgs(db *sqlite.SQLRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.DefaultQuery("page", "0")
		page, _ := strconv.Atoi(param)
		items, _ := db.GetItems(page)

		c.HTML(http.StatusOK, "items.tmpl", gin.H{
			"pkgs": items,
			"page": page + 1,
		})
	}
}
