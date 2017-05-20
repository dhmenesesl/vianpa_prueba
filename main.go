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

	Flightcode  string
	Origin string
	Destination string
	Price string
	Currency string

}

func main() {

	port := os.Getenv("PORT")
	router := gin.Default()

	if port == "" {
		port = "8080"
	}

	session, err := mgo.Dial("mongodb://admin:admin@ds147551.mlab.com:47551/vianca-db")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	cc := session.DB("vianca-db").C("vuelo")
	//err = cc.Insert(&flight{"001", "Medellin", "Bogota", "100000", "USD"},
	//	&flight{"002", "Medellin", "Barranquilla", "150000", "USD"})
	//if err != nil {
	//	log.Fatal(err)
	//}

	result := flight{}
	err = cc.Find(bson.M{"origin": "Medellin"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"flightcode": result.Flightcode, "origin": result.Origin, "destination": result.Destination, "price": result.Price, "currency": result.Currency })
	})

	router.Run(":" + port)
}
