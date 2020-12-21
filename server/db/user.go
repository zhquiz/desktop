package db

import (
	"log"

	"github.com/zhquiz/go-server/server/rand"
	"gorm.io/gorm"
)

// User holds user data
type User struct {
	gorm.Model

	Email  string  `gorm:"index:,unique;not null;check:email <> ''"`
	Name   string  `gorm:"not null;check:name <> ''"`
	Image  string  `gorm:"not null;check:image <> ''"`
	APIKey string  `gorm:"index,not null;check:api_key <> ''"`
	API    UserAPI `gorm:"embedded"`

	// Relations
	Decks   []Deck   `gorm:"constraint:OnDelete:CASCADE"`
	Entries []Entry  `gorm:"constraint:OnDelete:CASCADE"`
	Presets []Preset `gorm:"constraint:OnDelete:CASCADE"`
	Quizzes []Quiz   `gorm:"constraint:OnDelete:CASCADE"`
}

// UserAPI holds User's API keys
type UserAPI struct {
	Forvo *string
}

// New creates new User record
func (u *User) New(email string, name string, image string) {
	u.Email = email
	u.Name = name
	u.Image = image

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
