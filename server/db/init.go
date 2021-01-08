package db

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/wangbin/jiebago"
	"github.com/zhquiz/go-zhquiz/server/zh"
	"github.com/zhquiz/go-zhquiz/shared"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var jieba jiebago.Segmenter
var zhDB zh.DB

// DB is the storage for current DB
type DB struct {
	Current *gorm.DB
}

// Connect connects to DATABASE_URL
func Connect() DB {
	jieba.LoadDictionary(filepath.Join(shared.ExecDir, "assets", "dict.txt"))
	zhDB = zh.Connect()

	output := DB{}

	db, err := gorm.Open(sqlite.Open(shared.DatabaseURL()), &gorm.Config{
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
		&Quiz{},
		&Extra{},
	)

	if r := output.Current.Raw("SELECT Name FROM sqlite_master WHERE type='table' AND name='quiz_q'").First(&struct {
		Name string
	}{}); r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			output.Current.Exec(`
			CREATE VIRTUAL TABLE quiz_q USING fts5(
				[id],
				[entry],
				[level],
				[pinyin],
				[english],
				[description],
				[tag]
			);
			`)

			var quizzes []Quiz
			output.Current.Find(&quizzes)

			output.Current.Transaction(func(tx *gorm.DB) error {
				for _, q := range quizzes {
					q.AfterCreate(tx)
				}

				return nil
			})
		} else {
			panic(r.Error)
		}
	}

	if r := output.Current.Raw("SELECT Name FROM sqlite_master WHERE type='table' AND name='extra_q'").First(&struct {
		Name string
	}{}); r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			output.Current.Exec(`
			CREATE VIRTUAL TABLE extra_q USING fts5(
				[id],
				[chinese],
				[pinyin],
				[english],
				[description],
				[tag]
			);
			`)

			var extras []Extra
			output.Current.Find(&extras)

			output.Current.Transaction(func(tx *gorm.DB) error {
				for _, ex := range extras {
					ex.AfterCreate(tx)
				}

				return nil
			})
		} else {
			panic(r.Error)
		}
	}

	t := make([]libraryType, 0)

	b, err := ioutil.ReadFile(filepath.Join(shared.ExecDir, "assets", "library.yaml"))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(b, &t); err != nil {
		panic(err)
	}

	out := readLib(t, []string{}, []libraryEntry{})

	output.Current.Exec("DROP TABLE IF EXISTS library")
	output.Current.Exec(`
	CREATE VIRTUAL TABLE library USING fts5(
		[title],
		[entries]
	);
	`)

	output.Current.Transaction(func(tx *gorm.DB) error {
		for _, a := range out {
			if r := tx.Exec("INSERT INTO library (title, entries) VALUES (?, ?)", a.Title, a.Entries); r.Error != nil {
				panic(r.Error)
			}
		}

		return nil
	})

	return output
}

func parseChinese(s string) string {
	out := make([]string, 0)
	func(ch <-chan string) {
		for word := range ch {
			out = append(out, word)
		}
	}(jieba.CutAll(s))

	return strings.Join(out, " ")
}

func parsePinyin(s string) string {
	out := make([]string, 0)
	re := regexp.MustCompile("\\d+$")

	for _, c := range strings.Split(s, " ") {
		out = append(out, re.ReplaceAllString(c, ""))
	}

	return strings.Join(out, " ")
}

type libraryType struct {
	Title    string
	Entries  []string
	Children []libraryType
}

type libraryEntry struct {
	Title   string
	Entries string
}

func readLib(t []libraryType, parent []string, current []libraryEntry) []libraryEntry {
	for _, a := range t {
		title := append(parent, a.Title)

		if len(a.Entries) != 0 {
			entries, err := yaml.Marshal(a.Entries)
			if err != nil {
				log.Fatalln(err)
			}

			current = append(current, libraryEntry{
				Title:   strings.Join(title, " / "),
				Entries: string(entries),
			})
		}

		if len(a.Children) != 0 {
			current = readLib(a.Children, title, current)
		}
	}

	return current
}
