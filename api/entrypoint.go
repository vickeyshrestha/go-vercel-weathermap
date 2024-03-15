package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	app *gin.Engine
)

// create endpoint
func myRoute(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Vercel")
	})
	r.GET("/album", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	})
	r.GET("/info", func(c *gin.Context) {
		resp := make(map[string]string)
		resp["message"] = "Hello World from Go! 👋"
		resp["language"] = "go"
		resp["cloud"] = "Hosted on Vercel! ▲"
		resp["github"] = "https://github.com/vickeyshrestha/go-vercel-serverless"
		c.IndentedJSON(http.StatusOK, resp)
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

// Handler is the entry point for serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
