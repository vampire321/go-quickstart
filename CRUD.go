package main
import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`	
}

type Account struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Balance float64 `json:"balance"`
}
var users []User
var accounts []Account

func CRUD(){
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/users",getUsers)
		api.POST("/users",createUser)
		api.GET("/users/",getUserByID)
		api.PUT("/users/:id",updateUser)
		api.DELETE("/users/:id",deleteUser)
	}
	r.Run()
}

func getUsers(c *gin.Context){
	c.JSON(http.StatusOK,users)
}

func createUser(c *gin.Context){
	var u User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
}