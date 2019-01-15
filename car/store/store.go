package store

import (
	"github.com/jinzhu/gorm"
)

//Store Хранилище машинок
type Store struct {
	Db *gorm.DB
}

//Migrate мигрируем таблицы
func (s *Store) Migrate() {
	s.Db.AutoMigrate(&Car{})
}
