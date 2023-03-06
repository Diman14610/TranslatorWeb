package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if _, err := os.Stat("./source"); os.IsNotExist(err) {
		os.Mkdir("./source", os.FileMode(0522))
	}
	if _, err := os.Stat("./source/text"); os.IsNotExist(err) {
		os.Mkdir("./source/text", os.FileMode(0522))
	}
	if _, err := os.Stat("./source/img"); os.IsNotExist(err) {
		os.Mkdir("./source/img", os.FileMode(0522))
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	if _, err := os.Stat("./web"); !os.IsNotExist(err) {
		router.LoadHTMLGlob("web/templates/*")
		router.GET("/", indexPage)
		router.GET("/translate", helloPage)
	}

	router.POST("/load", handlerImg)
	router.GET("/text/:text", handlerText)

	router.StaticFS("/static", http.Dir("web/static"))
	err := router.Run(":8070")
	if err != nil {
		log.Fatalln("Start HTTP Server error", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, x-access-token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func indexPage(c *gin.Context) {
	files, err := os.ReadDir("./source/text/")
	var texts []string

	for _, v := range files {
		texts = append(texts, v.Name())
	}

	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"text":  "Главная страница",
		"texts": texts,
	})
}

func helloPage(c *gin.Context) {
	c.HTML(http.StatusOK, "translate.tmpl", gin.H{
		"version": time.Now().String(),
	})
}

func handlerImg(c *gin.Context) {
	// single file
	file, _ := c.FormFile("imageFile")

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		log.Println("Error! Input file not image/png")
		return
	}

	path := "./source/img/" + file.Filename
	c.SaveUploadedFile(file, path)

	text := Translate(ExtractText(path))
	os.WriteFile("./source/text/"+file.Filename+".txt", []byte(text), 0644)

	c.Redirect(303, "/text/"+file.Filename+".txt")
}

func handlerText(c *gin.Context) {
	t := c.Param("text")
	contents, _ := os.ReadFile("./source/text/" + t)
	c.HTML(http.StatusOK, "text.tmpl", gin.H{
		"text": string(contents),
	})
}
