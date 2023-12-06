package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type postgresqlProductRepo struct {
	db *sql.DB
}

func NewPostgressqlProductRepo(db *sql.DB) domain.ProductRepo {
	return &postgresqlProductRepo{db}
}

func (p *postgresqlProductRepo) GetProductByID(ctx context.Context, id int64) (*swagger.ProductInfoWithLabelAndDiscount, error) {
	sqlStatement := `

	`
	row := p.db.QueryRow(sqlStatement, id)

	product := &swagger.ProductInfoWithLabelAndDiscount{}
	if err := row.Scan(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return product, nil
}

// func (p *postgresqlProductRepo) GetByProductID(ctx context.Context, id int64) (*swagger.ProductInfo, error) {
// row := p.db.QueryRow("SELECT product_id, store_id, name, description, picture, price, stock, status FROM products WHERE product_id = $1", id)

// product := &swagger.ProductInfo{}
// if err := row.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Description, &product.Picture, &product.Price, &product.Stock, &product.Status); err != nil {
// 	logrus.Error(err)
// 	return nil, err
// }
// return product, nil
// }

// func (p *postgresqlProductRepo) AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error {
// sqlStatement := `
// INSERT INTO products (store_id, name, description, picture, price, stock, status) VALUES
// ($1, $2, $3, $4, $5, $6, $7)
// `

// _, err := p.db.Exec(sqlStatement, id, product.Name, product.Description, product.Picture, product.Price, product.Stock, product.Status)
// if err != nil {
// 	logrus.Error(err)
// 	return err
// }
// return nil
// }
