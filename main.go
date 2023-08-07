package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type person struct {
    ID     int  `json:"id"`
    Name  string  `json:"name"`
    Email string  `json:"email"`
    Phone  string `json:"phone"`
	People  string `json:"people"`
}

var persons = []person{
    {ID: 1, Name: "Yashwanth Pinneka", Email: "yashwanth.260990@gmail.com", Phone: "6312303767", People: "3"},
    {ID: 2, Name: "Likhitha Pinneka", Email: "likhitha.komm@gmail.com", Phone: "6312303767", People: "3"},
}

func getPersons(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, persons)
}

func postPersons(c *gin.Context) {
    var newPerson person

    if err := c.BindJSON(&newPerson); err != nil {
		fmt.Println("err = ", err)
        return
    }

    persons = append(persons, newPerson)
    c.IndentedJSON(http.StatusCreated, newPerson)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }
    
    router := gin.Default()
	router.Use(CORSMiddleware())
	
    router.GET("/persons", getPersons)
	router.POST("/persons", postPersons)

    // router.Run("localhost:8080")
    router.Run(":" + port)
}

// port := os.Getenv("PORT")

// if port == "" {
//     log.Fatal("$PORT must be set")
// }

// router := gin.New()
// router.Use(gin.Logger())
// router.LoadHTMLGlob("templates/*.tmpl.html")
// router.Static("/static", "static")

// router.GET("/", func(c *gin.Context) {
//     c.HTML(http.StatusOK, "index.tmpl.html", nil)
// })

// router.Run(":" + port)