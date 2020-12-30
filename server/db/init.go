package db

import (
	"log"
	"path/filepath"

	"github.com/zhquiz/go-server/shared"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is the storage for current DB
type DB struct {
	Current *gorm.DB
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

	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	output = DB{
		Current: db,
	}

	output.Current.AutoMigrate(
		&User{},
		&Tag{},
		&Quiz{},
		&QuizTag{},
		&Entry{},
		&EntryTag{},
		&EntryItem{},
		&Preset{},
		&Extra{},
	)

	return output
}
