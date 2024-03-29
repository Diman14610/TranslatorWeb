package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	}
	router.POST("/load", handleImg)
	router.GET("/text/:folder/:name", handleText)

	router.StaticFS("/static", http.Dir("web/static"))
	err := router.Run(Config.Host)
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
	files := make(map[string][]string)
	nameNewFolder := ""
	e := filepath.Walk("./source/text/.", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if name == "." {
			return nil
		}
		if strings.HasSuffix(name, ".txt") {
			files[nameNewFolder] = append(files[nameNewFolder], name)
		} else {
			nameNewFolder = name
			files[nameNewFolder] = make([]string, 0)
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"links": files,
	})
}

func handleImg(c *gin.Context) {
	// single file
	file, _ := c.FormFile("image")
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

func handleText(c *gin.Context) {
	f := c.Param("folder")
	t := c.Param("name")
	content, _ := os.ReadFile("./source/text/" + f + "/" + t)
	c.HTML(http.StatusOK, "text.tmpl", gin.H{
		"content": string(content),
	})
}
