package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	auth := gin.BasicAuth(gin.Accounts{
		"user":  "pass",
		"user2": "pass2",
		"user3": "pass3",
	})

	router.POST("/getDataPost", getDataPost)

	//router.Run()
	//http.ListenAndServe(":9090", router)
	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}
	client := router.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)
	}

	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi I in QueryString method",
		"name": name,
		"age":  age,
	})
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hii its get cmd",
	})
}
func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"data":     "Hi I am  post method GIN Framework",
		"bodyData": string(value),
	})
}
