package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

var Db *sql.DB 

func sendIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

//func sendMarkdown(c *gin.Context) {
	//file, _ := os.ReadFile("./public/todo.md")
	//c.JSON(200, gin.H{"file": string(file)})
//}

//func writeMarkdown(c *gin.Context) {
	//type WriteRequest struct {
        //Text string `json:"text"`
	//}
    //var data WriteRequest
    //c.ShouldBind(&data)
    //os.WriteFile("./public/todo.md", []byte(data.Text), 0644)

//}

type Todo struct {
	Id       int `json:"id"`
	Content  string `json:"content"`
	Checked  bool `json:"checked"`
}

func deleteTodo(c *gin.Context) {
	type TodoDeleteRequest struct {
		Id  int `json:"id"`
	}
	var todo TodoDeleteRequest 
    err := c.ShouldBind(&todo)
	if err != nil {
		println(err)
	}
    if(todo.Id != 1){
        statement, _ := Db.Prepare("DELETE FROM todo WHERE id=?")
        statement.Exec(todo.Id)
    }

	c.JSON(200, gin.H{
		"message": "Sucessful delete",
	})
}



func insertTodo(c *gin.Context) {
	type TodoInsertRequest struct {
		Content  string `json:"content"`
	}
	var todo TodoInsertRequest 
    err := c.ShouldBind(&todo)
	if err != nil {
		println(err)
	}
    statement, _ := Db.Prepare("INSERT INTO todo (content) VALUES (?)")
	statement.Exec(todo.Content)

	c.JSON(200, gin.H{
		"message": "Sucessful insert",
	})
}

func updateTodo(c *gin.Context) {
	type TodoUpdateRequest struct {
        Id       int `json:"id"`
        Content  string `json:"content"`
	}
	var todo TodoUpdateRequest 
    err := c.ShouldBind(&todo)
	if err != nil {
		println(err)
	}
    statement, _ := Db.Prepare("UPDATE todo SET content=? WHERE id=?")
	statement.Exec(todo.Content,todo.Id)

	c.JSON(200, gin.H{
		"message": "Sucessful update",
	})
}

func changeTodoChecked(c *gin.Context) {
	type TodoUpdateRequest struct {
        Id       int `json:"id"`
        Checked  bool `json:"checked"`
	}
	var todo TodoUpdateRequest 
    err := c.ShouldBind(&todo)
	if err != nil {
		println(err)
	}
    statement, _ := Db.Prepare("UPDATE todo SET checked=? WHERE id=?")
	statement.Exec(todo.Checked,todo.Id)

	c.JSON(200, gin.H{
		"message": "Sucessful change state",
	})
}


func sendTodos(c *gin.Context) {
	rows, _ := Db.Query("SELECT id, content, checked FROM todo")
    var todos []Todo
	for rows.Next() {
        var tempTodo Todo
		rows.Scan(&tempTodo.Id, &tempTodo.Content, &tempTodo.Checked)
        todos = append(todos, tempTodo)
	}
	c.JSON(200, gin.H{
		"todos": todos ,
	})
}










func main() {
	router := gin.Default()

    var err error
	Db, err = sql.Open("sqlite", "todo.db")
	if err != nil {
		log.Println(err)
	}


	router.LoadHTMLFiles("public/index.html")
	//router.GET("/api/gettext", sendMarkdown)

    router.GET("/api/todo", sendTodos)
    router.POST("/api/todo", insertTodo)

    router.POST("/api/todo/update", updateTodo)
    router.POST("/api/todo/check", changeTodoChecked)
    router.POST("/api/todo/delete", deleteTodo)

	router.Static("/static", "public")
	router.GET("/", sendIndex)
	router.Run(":8080")
}


