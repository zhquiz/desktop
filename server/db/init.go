package db

import (
	"log"
	"path/filepath"

	"github.com/zhquiz/go-server/shared"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the storage for current DB
type DB struct {
	Current *gorm.DB
	Type    string
}

// Connect connects to DATABASE_URL
func Connect() DB {
	databaseURL := shared.GetenvOrDefaultFn("DATABASE_URL", func() string {
		paths := []string{"data.db"}
		if root := shared.Paths().Root; root != "" {
			paths = append([]string{root}, paths...)
		}

		return filepath.Join(paths...)
	})

	output := DB{}

	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	output = DB{
		Current: db,
		Type:    "sqlite",
	}

	output.Current.AutoMigrate(&User{}, &Tag{}, &Quiz{}, &Entry{}, &EntryItem{}, &Preset{}, &Extra{})

	return output
}