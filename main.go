package main

import (
	db "MECHOPS/Db"
	routers "MECHOPS/Routers"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Connection()
	G := gin.Default()
	G.LoadHTMLGlob("Templates/*.html")
	routers.Routes(G)
	G.Run(":8080")

}
