package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed  bool `json:"completed"`
}

var todos =[]todo{
	{ID: "1", Item: "clean room", Completed: false},
	{ID: "2", Item: "read book", Completed: false},
	{ID: "3", Item: "record video", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}

func addTodo(context *gin.Context){
	var newTodo todo
print("HELLO");
	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

     todos = append(todos, newTodo)
      context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func updateTodo(context *gin.Context){
	id := context.Param("id")
	item := context.Param("item")
	todo, err := getTodoById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	todo.Item = item
	context.IndentedJSON(http.StatusOK, todo)
}
func getTodoById(id string) (*todo, error){
	for i, t := range todos{
		if t.ID == id{
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func main(){
	api := gin.Default()
	
	port := os.Getenv("PORT")
	if port == "" {
	port = "8080"
	}

	
	router := gin.Default()
	api.GET("/todos", getTodos)
	api.GET("/todos/:id", getTodo)
	api.PATCH("/todos/:id", toggleTodoStatus)
	api.PATCH("/update-todo/:id", updateTodo)
	api.POST("/todos", addTodo)
	api.Run(":"+port)
}
