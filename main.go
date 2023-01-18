package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func sendIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func sendMarkdown(c *gin.Context) {
	file, _ := os.ReadFile("./public/todo.md")
	c.JSON(200, gin.H{"file": string(file)})
}

func writeMarkdown(c *gin.Context) {
	type WriteRequest struct {
        Text string `json:"text"`
	}
    var data WriteRequest
    c.ShouldBind(&data)
    os.WriteFile("./public/todo.md", []byte(data.Text), 0644)

}
func main() {
	router := gin.Default()
	router.LoadHTMLFiles("public/index.html")
	router.GET("/api/gettext", sendMarkdown)
	router.POST("/api/writetext", writeMarkdown)
	router.Static("/static", "public")
	router.GET("/", sendIndex)
	router.Run(":8080")
}
