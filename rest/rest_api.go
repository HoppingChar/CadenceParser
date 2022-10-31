package rest

import (
	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	on_get_ping := func(c *gin.Context) {
		c.String(http.StatusOK, "Cadence parser rest service.")
	}

	on_post_parse := func(c *gin.Context) {
		// var body struct {
		// 	code string `json:"code" binding:"required"`
		// }

		// if c.Bind(&body) == nil {
		// 	ast_json := parser.Parse2Json([]byte(body.code))
		// 	c.JSON(http.StatusOK, ast_json)
		// }

		code, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		ast_json := Parse2Json(code)
		c.String(http.StatusOK, string(ast_json))
	}

	on_get_parse := func(c *gin.Context) {
		c.String(http.StatusOK, "Use post instead.")
	}

	r.GET("/ping", on_get_ping)
	r.GET("/parse", on_get_parse)
	r.POST("/parse", on_post_parse)

	return r
}
