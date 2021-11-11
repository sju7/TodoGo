package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var todos = []Todo{
	{ID: 1, Title: "Dishes", Description: "Do the dishes"},
	{ID: 2, Title: "trash", Description: "Take out the trash"},
}

func getTodos() []Todo {
	return todos
}
func addTodo(newTodo Todo) {
	todos = append(todos, newTodo)
}
func deleteTodo(id int) {
	var res []Todo
	for _, x := range todos {

		if x.ID != id {
			res = append(res, x)
		}
	}
	todos = res
}

func nextId() (max int) {
	if len(todos) == 0 {
		return 0
	}
	return todos[len(todos)-1].ID + 1
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", home)
	r.POST("/", postTodo)
	r.POST("/finish", finishTodo)

	r.Run()
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
		"todos": getTodos(),
	})
}

func postTodo(c *gin.Context) {
	newtodo := Todo{
		ID:          nextId(),
		Title:       c.PostForm("Title"),
		Description: c.PostForm("Description"),
	}
	addTodo(newtodo)
	c.Redirect(http.StatusFound, "/")
}

func finishTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("ID"))
	deleteTodo(id)
	c.Redirect(http.StatusFound, "/")
}
