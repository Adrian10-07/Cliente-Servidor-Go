package aplication

import "Clientes_servidor_practica/src/products/domain"


type DeleteProduct struct {
	rep domain.IProduct
}

func NewDeleteProduct(rep domain.IProduct) *DeleteProduct{
	return &DeleteProduct{
		rep: rep,
	}
}

func (cp *DeleteProduct) Execute(nombre string)error {
	return cp.rep.Delete(nombre)
}