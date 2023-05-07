package mysql

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MySQL struct {
	conn *gorm.DB
}

func New() (*MySQL, error) {
	db, err := gorm.Open(sqlite.Open("volume/storage.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to find volume: %v", err)
	}

	return &MySQL{
		conn: db,
	}, nil
}
