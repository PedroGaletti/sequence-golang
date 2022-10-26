package db

import (
	"fmt"

	"challenge/cmd/sequence"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitMyqlDb: open MySQL database connection
func InitMyqlDb(user, password, host, port, database string) (*gorm.DB, error) {
	sn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", user, password, host, port, database, "America%2FSao_Paulo")

	return gorm.Open("mysql", sn)
}

func Migrate(db *gorm.DB) {
	// tables
	db.AutoMigrate(&sequence.Sequence{})
}
