package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

//Definicion de structura de JSON
type album struct{
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64 		`json:"price"`
}

var albums = []album{
	{ID:"1", Title:"Without Me", Artist:"Eminem", Price: 10},
	{ID:"2", Title:"Normal", Artist:"Feid", Price: 9.99},
	{ID:"3", Title:"Fragil", Artist:"Yahritza y Su Esenncia", Price: 15},
	{ID:"4", Title:"Por ti", Artist:"Luana", Price: 8.99},
	{ID:"5", Title:"Gata Caliente", Artist:"LOLO OG y Omar varela", Price: 6.85},
	{ID:"6", Title:"TQG", Artist:"KAROL G Y Shakira", Price: 14.99},
}

func getAlbums(c *gin.Context){
	//Envia la respuesta en json
	c.IndentedJSON(http.StatusOK, albums)
}

func main(){
	//En este caso ponemos gin.Default() para que nos permite manejar las peticiones entrantes y decir que funcion tiene que hacer ante un peticion
	//Puede manejar GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD
	router := gin.Default()
	router.GET("/albums", getAlbums)

	//Corre el servidor en el puerto 8080
	router.Run("127.0.0.1:8080")
}