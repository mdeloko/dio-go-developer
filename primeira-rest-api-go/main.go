package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Servicos struct {
	Tipo map[string][]Produto `json:"tipo"`
}

type Produto struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}



func main() {
	ginRouter := gin.Default()

	f,err := os.Open("bd.json")
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	fileBytes,err :=io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var db Servicos
	json.Unmarshal(fileBytes,&db)

	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, db)
	})
	ginRouter.GET("/padaria",func(c *gin.Context) {
		c.JSON(http.StatusOK, db.Tipo["padaria"])
	})
	ginRouter.GET("/confeitaria",func(c *gin.Context) {
		c.JSON(http.StatusOK, db.Tipo["confeitaria"])
	})
	ginRouter.GET("/bebidas",func(c *gin.Context) {
		c.JSON(http.StatusOK, db.Tipo["bebidas"])
	})
	ginRouter.GET("/padaria/:id",func(c *gin.Context) {
		id,err:= strconv.Atoi(c.Param("id"))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"nome":db.Tipo["padaria"][id-1].Nome,
			})
		}
	})
	ginRouter.GET("/confeitaria/:id",func(c *gin.Context) {
		id,err:= strconv.Atoi(c.Param("id"))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"nome":db.Tipo["confeitaria"][id-1].Nome,
			})
		}
	})
	ginRouter.GET("/bebidas/:id",func(c *gin.Context) {
		id,err:= strconv.Atoi(c.Param("id"))
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err,
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"nome":db.Tipo["bebidas"][id-1].Nome,
			})
		}
	})
	ginRouter.Run(":3000")
}
