package main

import (
	"simple-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	// router.Use(cors.Default())
	router.GET("/products", getProducts)
	router.GET("/product/:code", getProduct)
	router.POST("/products", addProduct)
	
	router.POST("/user",addUser)
	router.GET("/users",getUsers)
	router.PUT("/user/:id",updateUser)
	router.DELETE("/user/:id",deleteUser)

	router.Run("192.168.4.55:8080")
}

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(c *gin.Context) {
	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}

// Users Controller
func getUsers(c *gin.Context) {
	users := models.GetUsers()

	if users == nil || len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.IndentedJSON(http.StatusOK, users)
	}
}

func addUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddUser(user)
		c.Header("Access-Control-Allow-Origin", "*")
		c.IndentedJSON(http.StatusCreated, user)
	}
}

func deleteUser(c *gin.Context) {
	ID := c.Param("id")

	status := models.DeleteUser(ID)

	if status == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// c.IndentedJSON(http.StatusCreated, status)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"status": status})
	}
}

func updateUser(c *gin.Context) {
	var user models.User

	ID := c.Param("id")

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.EditUser(user,ID)
		c.Header("Access-Control-Allow-Origin", "*")
		c.IndentedJSON(http.StatusCreated, user)
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}