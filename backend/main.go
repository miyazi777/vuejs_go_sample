package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
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

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "msg from server",
	})
}

func test2(c *gin.Context) {
	id := c.Param("id")
	name := fmt.Sprintf("book%s", id)
	price, _ := strconv.Atoi(id)
	c.JSON(200, gin.H{
		"name":  name,
		"price": price * 600,
	})
}

func test3(c *gin.Context) {
	//var list []Book

	//list = append(list, Book{Id: 1, Name: "book1", Price: 1000})
	//list = append(list, Book{Id: 2, Name: "book2", Price: 1000})
	//list = append(list, Book{Id: 3, Name: "book3", Price: 1000})
	dbmap := getDbmap()
	list, _ := dbmap.Select(Book{}, "select * from book")
	//for _, l := range list {
	//	book := l.(*Book)
	//}

	c.JSON(200, list)
}

func test4(c *gin.Context) {
	currentPageStr := c.DefaultQuery("page", "0")
	fmt.Printf("current page = %s\n", currentPageStr)
	currentPage, _ := strconv.Atoi(currentPageStr)

	// calc pagenite
	dbmap := getDbmap()
	recordCount, _ := dbmap.SelectInt("select count(*) from book")
	page := NewPage(int(recordCount), 3)

	list, _ := dbmap.Select(Book{}, "select * from book order by id asc limit ? offset ?", page.perPage, page.calcOffset(currentPage))

	c.JSON(200, gin.H{
		"books": list,
		"page": gin.H{
			"pageCount":   page.pageCount,
			"recordCount": page.recordCount,
		},
	})
}

type Page struct {
	recordCount int
	perPage     int
	pageCount   int
}

func NewPage(recordCount int, perPage int) *Page {
	pageCount := int(recordCount) / perPage
	fmt.Printf("page %d\n", pageCount)
	flg := int(recordCount) % perPage
	fmt.Printf("flg %d\n", flg)
	if flg > 0 {
		pageCount++
	}
	fmt.Printf("page %d\n", pageCount)

	p := new(Page)
	p.recordCount = recordCount
	p.perPage = perPage
	p.pageCount = pageCount
	return p
}

func (p *Page) calcOffset(page int) int {
	if page > p.pageCount {
		return 0
	}

	offset := page * p.perPage
	return offset
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/book", addBook)
	r.GET("/test", test)
	r.GET("/test2/:id", test2)
	r.GET("/test3", test3)
	r.GET("/test4", test4)

	r.Run(":8000")
}
