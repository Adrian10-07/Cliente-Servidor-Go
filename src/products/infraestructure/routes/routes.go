package routes

import (
	"net/http"
	"Clientes_servidor_practica/src/products/infraestructure/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/products", controllers.CreateProductHandler)
	http.HandleFunc("/view-products", controllers.GetProductHandler)
	http.HandleFunc("/delete-products/", controllers.DeleteProductHandeler)
	http.HandleFunc("/update-products/", controllers.UpdateProductHandler)
}