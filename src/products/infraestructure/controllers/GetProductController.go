package controllers

import (
	"Clientes_servidor_practica/src/products/aplication"
	"Clientes_servidor_practica/src/products/infraestructure"
	"encoding/json"
	"fmt"
	"net/http"
)
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	
	repo := infraestructure.NewMySQLRepository()
	useCase := aplication.NewGetProduct(repo)
	products, err := useCase.Execute()
	if err != nil {
		fmt.Printf("error al obtner los productos", err)
		return
	}

	productsJson, err := json.Marshal(products)

	fmt.Printf(string(productsJson))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)

}