package main

import "github.com/gin-gonic/gin"



func main(){
	G:=gin.Default()
	G.Run(":8080")
}