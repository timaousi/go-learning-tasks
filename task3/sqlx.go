package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	username := "root"
	password := "root"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/testdb?parseTime=true", username, password)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func queryTechEmployees() {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Println("查询失败:", err)
		return
	}

	fmt.Println("技术部员工：")
	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}
}

func queryHighestSalaryEmployee() {
	var emp Employee
	err := db.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Println("查询失败:", err)
		return
	}

	fmt.Println("工资最高的员工：", emp)
}

func queryExpensiveBooks() {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		log.Println("查询失败:", err)
		return
	}

	fmt.Println("价格大于50元的书籍：")
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}

func main() {
	queryTechEmployees()
	queryHighestSalaryEmployee()
	queryExpensiveBooks()
}
