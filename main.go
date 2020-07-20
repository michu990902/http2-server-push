package main

import (
	"log"
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(tpl)
	
	r.GET("/", Home)

	r.RunTLS(":8080", "tls/cert_test_example.pem", "tls/key_test_example.pem")
}

func Home(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		if err := pusher.Push("/assets/script.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
		if err := pusher.Push("/assets/style.css", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"status": "success",
	})
}