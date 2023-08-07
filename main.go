package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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

func deletePerson(c *gin.Context) {
    id := c.Param("id")
    num, err := strconv.Atoi(id)
    if err != nil {
      fmt.Println(err)
      return
    }

    for i, a := range persons {
        if a.ID == num + 1 {
            persons = append(persons[:i], persons[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "person successfully deleted"})
            return
        }
    }
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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
    router.DELETE("/persons/:id", deletePerson)

    router.Run("localhost:8080")
    // router.Run(":" + port)
}