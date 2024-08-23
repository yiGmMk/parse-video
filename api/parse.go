package api

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wujunwei928/parse-video/handler"
)

//go:embed parse/templates/*
var files embed.FS
var router *gin.Engine

func init() {
	router = gin.Default()
	sub, err := fs.Sub(files, "templates")
	if err != nil {
		panic(err)
	}
	handler.RegisterHandler(router, sub)
}

// vercel Function
func Parse(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
