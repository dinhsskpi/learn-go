package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	db, ctx := ConnectDB()
	cursor, err := db.Collection("albums").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var albums []bson.M
	for cursor.Next(ctx) {
		var album bson.M
		if err = cursor.Decode(&album); err != nil {
			log.Fatal(err)
		}
		albums = append(albums, album)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	db, ctx := ConnectDB()
	var album bson.M
	filter := bson.M{"_id": objID}
	collection := db.Collection("albums")
	if err := collection.FindOne(ctx, filter).Decode(&album); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	if album != nil {
		c.IndentedJSON(http.StatusOK, album)
	}
}

type CreateAblumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func createAlbum(c *gin.Context) {
	var input CreateAblumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db, ctx := ConnectDB()
	albumCollection := db.Collection("albums")

	newAlbum := bson.D{{"title", input.Title}, {"artist", input.Artist}, {"price", input.Price}}
	res, err := albumCollection.InsertOne(ctx, newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	c.IndentedJSON(http.StatusCreated, id)
}

func ConnectDB() (*mongo.Database, context.Context) {
	var cred options.Credential
	cred.Username = "dinhpv"
	cred.Password = "12345678"

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("demo")
	return db, ctx
}

func main() {
	server := gin.Default()
	server.GET("albums", getAlbums)
	server.GET("albums/:id", getAlbumsById)
	server.POST("albums", createAlbum)

	server.Run("localhost:8080")
}
