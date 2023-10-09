package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hokaccha/go-prettyjson"
)

var fileName = "logFile.log"

func main() {
	r := gin.Default()
	r.GET("/event/appsflyer/:type", flyerHandler)
	r.GET("/event/operation/:operation", OperationHandler)
	r.Run(":3000")
}

func flyerHandler(ctx *gin.Context) {
	if err := logsFile(ctx, "Request Flyer"); err != nil {
		log.Panic(err)
	}
}

func OperationHandler(ctx *gin.Context) {
	if err := logsFile(ctx, "Request Operation"); err != nil {
		log.Panic(err)
	}
}

func logsFile(ctx *gin.Context, action string) error {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	wrt := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(wrt)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var req interface{}
	ctx.ShouldBindJSON(&req)
	body, err := prettyjson.Marshal(req)
	if err != nil {
		return err
	}

	log.Println(action)
	log.Println(string(body))

	return nil
}
