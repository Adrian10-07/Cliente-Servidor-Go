package infraestructure

import (
	"Clientes_servidor_practica/src/products/domain"
	"Clientes_servidor_practica/src/core"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Product) error {
	query := "INSERT INTO Product (Nombre, Precio, Cantidad, CodigoDeBarras) VALUES (?, ?, ?, ?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Precio, p.Cantidad, p.CodigoDeBarras)
	return err
}

func (r *MySQLRepository) Delete(p string)error{
	nombre :=p
	query := "DELETE FROM Product WHERE Nombre = ?"
	_,err :=r.conn.DB.Exec(query,nombre)
	return err
}

func (r *MySQLRepository) Update(id int,p *domain.Product)error{
	query := "UPDATE Product SET Nombre = ?, Precio = ?, Cantidad = ?, CodigoDeBarras = ?, WHERE id = ?"
    _, err := r.conn.DB.Exec(query, p.Nombre, p.Precio,p.Cantidad,p.CodigoDeBarras,id)
    if err != nil {
        return err
    }
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Product, error) {
	query := "SELECT Nombre, Precio, Cantidad, CodigoDeBarras FROM Product"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Nombre, &product.Precio, product.Cantidad, product.CodigoDeBarras); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}