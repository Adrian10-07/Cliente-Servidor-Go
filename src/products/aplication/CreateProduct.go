package aplication

import "Clientes_servidor_practica/src/products/domain"

type CreateProduct struct {
	repo domain.IProduct
}


func NewCreateProduct(repo domain.IProduct) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Execute(p domain.Product)error{
	return cp.repo.Save(&p)
}