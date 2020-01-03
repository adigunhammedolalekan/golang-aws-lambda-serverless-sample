package database

import (
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/jinzhu/gorm"
	"os"
)

type ConnectionProvider struct {
	db *gorm.DB
}

func New() (ConnectionProvider, error) {
	uri := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", uri)
	if err != nil {
		return ConnectionProvider{}, err
	}
	if err := db.DB().Ping(); err != nil {
		return ConnectionProvider{}, err
	}
	d := ConnectionProvider{db: db}
	// d.runMigration()
	return d, nil
}

func (d ConnectionProvider) Close() error {
	return d.db.Close()
}

func (d ConnectionProvider) DB() *gorm.DB {
	return d.db
}

func (d ConnectionProvider) runMigration() {
	d.db.AutoMigrate(&types.Post{})
}

func (d ConnectionProvider) Destroy() {

}