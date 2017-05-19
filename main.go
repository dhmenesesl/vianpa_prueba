package main

import (

	"log"
	"net/http"
	"os"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type flight struct {
	Origen  string
	Destino string
}

func main() {

	port := os.Getenv("PORT")
	router := gin.Default()

	if port == "" {
		port = "8080"
	}

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	cc := session.DB("db").C("vuelo")
	err = cc.Insert(&flight{"Medellin", "Bogota"},
					&flight{"Medellin2", "Barranquilla"})
	if err != nil {
		log.Fatal(err)
	}

	result := flight{}
	err = cc.Find(bson.M{"origen": "Medellin2"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hola mundo" , "salida": result.Origen, "destino": result.Destino})
	})

	router.Run(":" + port)
}
