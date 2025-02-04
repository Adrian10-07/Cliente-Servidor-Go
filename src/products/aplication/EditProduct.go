package aplication

import "Clientes_servidor_practica/src/products/domain"

type EditProduct struct {
	repo domain.IProduct
}

func NewUpdateProduct(repo domain.IProduct) *EditProduct{
	return &EditProduct{repo: repo}
}

func (cp *EditProduct) Execute(id int,product *domain.Product)error{
	return cp.repo.Update(id,product)
}