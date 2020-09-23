package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("*.html")

	router.GET("/", mainPage)
	router.POST("/demo", openPage)
	router.Run(":" + port)
	// router.Run(":8080")
}

func mainPage(c *gin.Context) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(c.Writer, nil)
}

func openPage(c *gin.Context) {

	if c.PostForm("Link") != "" {
		link1 := "https://xd.adobe.com/embed/"
		link2 := c.PostForm("Link")
		link := link1 + link2

		t1, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println(err)
		}

		t1.Execute(c.Writer, link)
	} else {
		text1 := "Please enter the appropriate url string!"
		t2, err := template.ParseFiles("main.html")
		if err != nil {
			log.Println(err)
		}
		t2.Execute(c.Writer, text1)
	}

}
