package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/shurcooL/github_flavored_markdown"
)

func sendIndex(c *gin.Context) {
    file, _ := os.ReadFile("/home/div/notes/todo.md");
    content := string(github_flavored_markdown.Markdown(file))
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"content": template.HTML(content), 
	})
}

func sendSpecificFile(c *gin.Context) {
    path := c.Param("path")
    dirString := os.Args[1] 
    if path == "/" {
        file, _ := os.ReadFile(dirString + "index.md");
        content := string(github_flavored_markdown.Markdown(file))
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "content": template.HTML(content), 
            "title": "index.md",
        })
    } else {
        pathItems := strings.Split(path[1:], "/")
        for _, item := range pathItems {
            if(strings.Contains(item,".")){
                dirString += item
                file, _ := os.ReadFile(dirString);
                content := string(github_flavored_markdown.Markdown(file))
                c.HTML(http.StatusOK, "index.tmpl", gin.H{
                    "content": template.HTML(content), 
                    "title": item,
                })
            } else {
                dirString += item + "/"
            }
        }
    }
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("public/index.tmpl")
    //router.Static("/static", "public")
    router.GET("/*path", sendSpecificFile)
	router.Run(":8080")
}
