package main

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GuiaBolso/darwin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type cryptoString string

func (c cryptoString) Value() (driver.Value, error) {
	return driver.Value(base64.StdEncoding.EncodeToString([]byte(c))), nil
}

func (c *cryptoString) Scan(src interface{}) error {
	var source string

	switch src.(type) {
	case string:
		source = src.(string)
	case []byte:
		source = string(src.([]byte))
	default:
		return errors.New("Incompatible type for cryptoString")
	}

	decoded, err := base64.StdEncoding.DecodeString(source)

	if err != nil {
		return err
	}

	*c = cryptoString(decoded)

	return nil
}

type person struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table people",
			Script: `CREATE TABLE people (
                        id INTEGER PRIMARY KEY,
                        name       TEXT
                     )`,
		},
		{
			Version:     2,
			Description: "Creating table hashes",
			Script: `CREATE TABLE hashes (
                        id INTEGER PRIMARY KEY,
                        hash TEXT,
                        UNIQUE(hash)
                     )`,
		},
		{
			Version:     3,
			Description: "Creating table hashes_people",
			Script: `CREATE TABLE hashes_people (
                        hash_id INTEGER,
                        person_id INTEGER
                     )`,
		},
	}
)

func addPerson(db *sqlx.DB, p person) error {
	rs, err := db.Exec("INSERT INTO people (name) VALUES (?)", cryptoString(p.Name))

	if err != nil {
		return err
	}

	personID, err := rs.LastInsertId()

	if err != nil {
		return err
	}

	tokens := strings.Split(p.Name, " ")

	for _, token := range tokens {
		crypted := fmt.Sprintf("%x", sha256.Sum256([]byte(strings.TrimSpace(token))))
		row := db.QueryRow("SELECT h.id FROM hashes h WHERE h.hash = ?", crypted)

		if err != nil {
			return err
		}

		var hashID int64

		if err = row.Scan(&hashID); err != nil {
			rs, err = db.Exec("INSERT INTO hashes (hash) VALUES (?)", crypted)
			if err != nil {
				return err
			}

			hashID, err = rs.LastInsertId()

			if err != nil {
				return err
			}
		}

		_, err = db.Exec("INSERT INTO hashes_people (hash_id, person_id) VALUES (?, ?)", hashID, personID)

		if err != nil {
			return err
		}
	}

	return nil
}

func search(db *sqlx.DB, query string) ([]person, error) {
	people := []person{}
	tokens := strings.Split(query, " ")
	cryptoTokens := []string{}

	for _, token := range tokens {
		crypted := fmt.Sprintf("%x", sha256.Sum256([]byte(strings.TrimSpace(token))))
		cryptoTokens = append(cryptoTokens, crypted)
	}

	q, args, err := sqlx.In(`SELECT DISTINCT p.id, p.name
                                 FROM people p 
                                 JOIN hashes_people hp ON hp.person_id = p.id
                                 JOIN hashes h ON h.id = hp.hash_id AND h.hash IN (?)
                             GROUP BY p.name`, cryptoTokens)

	if err != nil {
		return people, err
	}

	rows, err := db.Queryx(q, args...)

	if err != nil {
		return people, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var name cryptoString

		err := rows.Scan(&id, &name)

		if err != nil {
			return people, err
		}

		people = append(people, person{id, string(name)})
	}

	return people, err
}

func main() {
	//	database, err := sqlx.Connect("sqlite3", "file:database")
	database, err := sqlx.Connect("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("Error while opening database, %s", err)
	}

	driver := darwin.NewGenericDriver(database.DB, darwin.SqliteDialect{})
	info := make(chan darwin.MigrationInfo)

	go func() {
		for {
			select {
			case i := <-info:
				log.Printf("== %s\n", i.Migration.Description)
			}
		}
	}()

	d := darwin.New(driver, migrations, info)
	err = d.Migrate()

	if err != nil {
		log.Println(err)
	}

	if err := addPerson(database, person{Name: "Claudemiro Alves Feitosa Neto"}); err != nil {
		log.Fatal(err)
	}

	if err := addPerson(database, person{Name: "Antonio Alves Souza Feitosa"}); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		results, err := search(database, q)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	log.Println("Starting at 8080...")
	http.ListenAndServe(":8080", nil)
}
