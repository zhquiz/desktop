package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/zhquiz/go-zhquiz/server/rand"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// User holds user data
type User struct {
	ID        string `gorm:"primaryKey;check:length(id) > 20"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Email  string `gorm:"index:,unique;not null;check:length(email) > 4"`
	APIKey string `gorm:"index,not null;check:length(api_key) > 20"`

	Meta UserMeta

	// Relations
	Decks   []Deck   `gorm:"constraint:OnDelete:CASCADE"`
	Entries []Entry  `gorm:"constraint:OnDelete:CASCADE"`
	Presets []Preset `gorm:"constraint:OnDelete:CASCADE"`
	Quizzes []Quiz   `gorm:"constraint:OnDelete:CASCADE"`
}

// UserMeta holds User's settings
type UserMeta struct {
	Forvo    *string
	Level    *uint
	LevelMin *uint
	Settings struct {
		Level struct {
			WhatToShow []string
		}
		Quiz struct {
			Type      []string
			Stage     []string
			Direction []string
			IsDue     bool
		}
		Sentence struct {
			Min *uint
			Max *uint
		}
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *UserMeta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := UserMeta{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (j UserMeta) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// GormDBDataType represents UserMeta's data type
func (UserMeta) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return "TEXT"
}

// New creates new User record
func (u *User) New(id string, email string) {
	u.ID = id
	u.Email = email

	u.NewAPIKey()
}

// NewAPIKey generates a new API key to the User, and returns it
func (u *User) NewAPIKey() string {
	apiKey, err := rand.GenerateRandomString(64)
	if err != nil {
		log.Fatalln(err)
	}

	u.APIKey = apiKey

	return apiKey
}
