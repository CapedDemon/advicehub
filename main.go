package main

import (
	"fmt"
	"getAdvice/models"
	"getAdvice/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)
	server := gin.Default()

	server.Static("/css", "./templates/css")
	server.Static("/js", "./templates/js")
	server.Use(static.Serve("/img", static.LocalFile("./templates/css/images", true)))
	server.LoadHTMLGlob("templates/index.html")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Advice Hub",
		})
	})

	//getting the keu
	content, err := os.ReadFile("./classified/key.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	webpaths := server.Group("/")

	{
		webpaths.GET("advices", routes.Getadvices)
		webpaths.GET("advices/:id", routes.GetAdviceid)
		webpaths.POST(string(content)+"advice", routes.Addadvice)
		webpaths.PUT(string(content)+"advice/:id", routes.Updateadvice)
		webpaths.GET("suradvice", routes.GetsurpriseAdvice)
	}

	server.Run()

}
