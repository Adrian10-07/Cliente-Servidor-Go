package aplication

import "Clientes_servidor_practica/src/products/domain"

type GetProduct struct {
	repo domain.IProduct
}	


func NewGetProduct(repo domain.IProduct) *GetProduct{
		return &GetProduct{repo: repo}
}

func (cp *GetProduct) Execute() ([]domain.Product, error) {
	return cp.repo.GetAll()
}