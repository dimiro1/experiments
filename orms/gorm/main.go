package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Recipe represents single recipe
type Recipe struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Recipe{})

	if result := db.Create(&Recipe{Name: "This is the name of a recipe"}); result.Error != nil {
		panic(err)
	}

	var recipes []Recipe
	if result := db.Find(&recipes); result.Error != nil {
		panic(err)
	}

	fmt.Printf("%#+v\n", recipes)
}
