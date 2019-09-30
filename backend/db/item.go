package db

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

type ItemRepository struct{}

func getDbmap() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	t := dbmap.AddTableWithName(Item{}, "item").SetKeys(true, "Id")
	t.ColMap("Id").Rename("id")
	t.ColMap("Name").Rename("name")
	t.ColMap("Price").Rename("price")
	return dbmap
}

func (i *ItemRepository) GetList() *[]*Item {
	dbmap := getDbmap()
	list, _ := dbmap.Select(Item{}, "select * from item")

	var items []*Item
	for _, entity := range list {
		items = append(items, entity.(*Item))
	}

	return &items
}

func (i *ItemRepository) Get(id int) *Item {
	dbmap := getDbmap()
	item := Item{}
	dbmap.SelectOne(&item, "select * from item where id = ?", id)

	return &item
}

func (i *ItemRepository) Add(item *Item) {
	dbmap := getDbmap()

	tx, _ := dbmap.Begin()
	tx.Insert(item)
	tx.Commit()
}

func (i *ItemRepository) Update(item *Item) {
	dbmap := getDbmap()

	tx, _ := dbmap.Begin()
	tx.Update(item)
	tx.Commit()
}
