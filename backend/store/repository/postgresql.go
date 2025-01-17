package repository

import (
	"context"
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/sirupsen/logrus"
)

type postgresqlStoreRepo struct {
	db *sql.DB
}

func NewPostgressqlStoreRepo(db *sql.DB) domain.StoreRepo {
	return &postgresqlStoreRepo{db}
}

func (p *postgresqlStoreRepo) GetByID(ctx context.Context, id string) (*swagger.StoreInfo, error) {
	row := p.db.QueryRow("SELECT store_id, name, rate, rate_count, address, picture, description, shipping_fee, status FROM stores WHERE store_id = $1", id)
	s := &swagger.StoreInfo{}
	if err := row.Scan(&s.StoreId, &s.Name, &s.Rate, &s.RateCount, &s.Address, &s.Picture, &s.Description, &s.ShippingFee, &s.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}
