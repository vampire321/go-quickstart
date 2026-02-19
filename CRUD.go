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
		api.GET("/users/",getUserById)
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
	u.ID = len(users)+1
	users = append(users,u)
	c.JSON(http.StatusCreated,u)
}
func getUserById(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	for _,u := range users{
		if u.ID == id {
			c.JSON(http.StatusOK,u)
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"error":"User not found"})
}

func updateUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var updated User
	if err := c.BindJSON(&updated);
	err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	for i,u := range users{
		if u.ID == id {
			users[i].Name=updated.Name
			c.JSON(http.StatusOK,users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
func deleteUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	for i,u := range users{
		if u.ID == id {
			users = append(users[:i],users[i+1:]...)
			c.JSON(http.StatusOK,gin.H{"message":"User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{
		"error":"User not found",
	})

}

// --- Account Handlers --- 

func getAccounts(c *gin.Context) { 
	c.JSON(http.StatusOK, accounts) 
	} 

func createAccount(c *gin.Context) { 
	var a Account 
	if err := c.BindJSON(&a); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		return 
		} 
		a.ID = len(accounts) + 1 
		accounts = append(accounts, a) 
		c.JSON(http.StatusCreated, a) 
}
func getAccountByID(c *gin.Context) { 
	id, _ := strconv.Atoi(c.Param("id")) 
	for _, a := range accounts { 
		if a.ID == id { 
			c.JSON(http.StatusOK, a) 
			return 
			} 
			} 
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"}) 
		}

func updateAccount(c *gin.Context) { 
	id, _ := strconv.Atoi(c.Param("id")) 
	var updated Account 
	if err := c.BindJSON(&updated); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		 return 
		 } 
		 for i, a := range accounts { 
			if a.ID == id { 
				accounts[i].Balance = updated.Balance 
				c.JSON(http.StatusOK, accounts[i]) 
				return 
				} 
				} 
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"}) 
			}

func deleteAccount(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, a := range accounts {
        if a.ID == id {
            accounts = append(accounts[:i], accounts[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
}