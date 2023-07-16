package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Created_at   string `json:"created_at"`
	Completed_at string `json:"completed_at"`
}

var TodoList []Todo

func addTodos(c *gin.Context, TodoList *[]Todo) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	} else {
		todo.ID = "6807E49C-D68F-4D38-A990-6FACCB0E51A3" + "-" + todo.ID
		*TodoList = append(*TodoList, todo)
	}

	c.JSON(http.StatusCreated, gin.H{"result": TodoList})
}

func getTodoByID(c *gin.Context, TodoList []Todo) {
	todoID := c.Param("id")

	var foundTodo Todo
	for _, item := range TodoList {
		if todoID == item.ID {
			foundTodo = item
			break
		}
	}

	if foundTodo.ID == "" {
		c.JSON(404, gin.H{"error": "Invalid request payload"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": foundTodo})
}

func getAllTodo(c *gin.Context, TodoList []Todo) {
	c.JSON(http.StatusOK, gin.H{"result": TodoList})
}

func updateTodo(c *gin.Context, TodoList *[]Todo) {
	var todo Todo
	todoID := c.Param("id")

	for index, item := range *TodoList {
		if item.ID == todoID {
			*TodoList = append((*TodoList)[:index], (*TodoList)[index+1:]...)
			todo.ID = item.ID
			*TodoList = append((*TodoList), todo)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"result": todo})
	c.JSON(http.StatusOK, gin.H{"result": *TodoList})
}

func deleteTodo(c *gin.Context, TodoList []Todo) {
	todoID := c.Param("id")

	for index, item := range TodoList {
		if todoID == item.ID {
			TodoList = append(TodoList[:index], TodoList[index+1:]...)
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{"result": TodoList})
}

func main() {
	router := gin.Default()

	TodoList = append(TodoList, Todo{ID: "6807E49C-D68F-4D38-A990-6FACCB0E51A3-1", Title: "Buy milk", Created_at: "2020-12-20T15:04:05Z", Completed_at: ""})
	TodoList = append(TodoList, Todo{ID: "6807E49C-D68F-4D38-A990-6FACCB0E51A3-2", Title: "Buy coffee", Created_at: "2020-12-21T15:04:05Z", Completed_at: ""})

	router.POST("/todos", func(c *gin.Context) {
		addTodos(c, &TodoList)
	})
	router.GET("/todos/:id", func(c *gin.Context) {
		getTodoByID(c, TodoList)
	})
	router.GET("/todos", func(c *gin.Context) {
		getAllTodo(c, TodoList)
	})
	router.PUT("/todos/:id/done", func(c *gin.Context) {
		updateTodo(c, &TodoList)
	})
	router.DELETE("/todos/:id", func(c *gin.Context) {
		deleteTodo(c, TodoList)
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}
