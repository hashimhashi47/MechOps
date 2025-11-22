package main

import (
	db "MECHOPS/Db"
	routers "MECHOPS/Routers"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)



func main(){
	 defer func() {
        if r := recover(); r != nil {
            log.Println("🔥 PANIC:", r)
            debug.PrintStack() // add this
        }
    }()


	db.Connection()
	G:=gin.Default()
	routers.Routes(G)	
	G.Run(":8080")

}