package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Articles []Article

func main() {
	Articles = []Article{
		Article{Titulo: "Detergente", Descrip: "Descripcion de articulo", Contenido: "Descripcion de articulo"},
		Article{Titulo: "Crema", Descrip: "Descripcion de articulo", Contenido: "Descripcion de articulo"},
	}
	handleRequest()

}

//homePage Enviará las solciitudes a la URL root
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Bienvenido a la pagina Home")
	fmt.Println("Home")
}

//handleRequest Unirá la URL con main.
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/Articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Article struct {
	Titulo    string `json:"title"`
	Descrip   string `json:"descrip"`
	Contenido string `json:"contenido"`
}

//returnAllArticles devuelve todos los articulos en formato JSON
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
