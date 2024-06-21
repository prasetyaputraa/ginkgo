package db

import "github.com/jinzhu/gorm"

type Model struct {
	db *gorm.DB
}

func (m Model) GetDB() *gorm.DB {
	return m.db
}

func (m *Model) Migrate() {
	m.db.AutoMigrate(m)
}
