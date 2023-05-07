package mysql

import (
	"fmt"

	"github.com/sweetie-pie/line-recommendation/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MySQL struct {
	conn *gorm.DB
}

func New() (*MySQL, error) {
	// open sqlite connection
	db, err := gorm.Open(sqlite.Open("volume/storage.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to find volume: %v", err)
	}

	// migrate models
	if err := db.AutoMigrate(
		&model.Route{},
		&model.Node{},
		&model.Search{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate models: %v", err)
	}

	return &MySQL{
		conn: db,
	}, nil
}
