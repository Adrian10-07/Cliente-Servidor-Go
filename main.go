package main

import(
	"log"
	"net/http"
	"Clientes_servidor_practica/src/products/infraestructure/routes"
)

func main()  {
	routes.SetupRoutes()
	log.Print("iniciando server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}