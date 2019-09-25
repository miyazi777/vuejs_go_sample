package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Id    int32
	Name  string
	Price int32
}

func addBook(c *gin.Context) {
	var book Book
	c.Bind(&book)

	dbmap := getDbmap()

	tx, _ := dbmap.Begin()
	tx.Insert(&book)
	tx.Commit()

	c.JSON(200, gin.H{
		"name":  book.Name,
		"price": book.Price,
	})
}

func getDbmap() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	t := dbmap.AddTableWithName(Book{}, "book").SetKeys(true, "Id")
	t.ColMap("Id").Rename("id")
	t.ColMap("Name").Rename("name")
	t.ColMap("Price").Rename("price")
	return dbmap
}

func main() {
	r := gin.Default()
	r.POST("/book", addBook)

	r.Run(":3000")
}
