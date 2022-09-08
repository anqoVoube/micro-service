package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	s := &http.Server{
		Addr:           ":80",
		Handler:        r,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	initRoutes(r)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Couldn't run the server: %s", err)
	}

}

func initRoutes(r *gin.Engine) {
	api := r.Group("/")
	api.POST("/", ping)
	a := r.Group("/handle")
	a.POST("/", pong)

}
