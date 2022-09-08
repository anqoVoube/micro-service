package main

import (
	"authentication/cmd/database"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

const webPort = "80"

//var counts int64

//type Config struct {
//	DB     *sql.DB
//	Models data.Models
//}
//
//type Env struct {
//	db *sql.DB
//}

func main() {
	conn := database.Get()
	if conn == nil {
		log.Panic("Couldn't connect to Postgres")
	}

	r := gin.Default()

	r.Use(cors.Default())
	initRoutes(r)
	r.RedirectFixedPath = true
	r.RedirectTrailingSlash = false
	fmt.Printf("RedirectTrailingSlash: %t", r.RedirectFixedPath)
	fmt.Printf("RedirectFixedPath: %t", r.RedirectFixedPath)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Couldn't run the server: %s", err)
	}

}

func initRoutes(r *gin.Engine) {
	api := r.Group("/authenticate")
	api.POST("/", ping)

}
