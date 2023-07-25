package postgresql

import "gorm.io/gorm"

type postgresqlRepo struct {
	db *gorm.DB
}

func NewPostgresqlRepo(db *gorm.DB) *postgresqlRepo {
	return &postgresqlRepo{db: db}
}
