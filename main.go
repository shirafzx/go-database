package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	databaseName= "mydatabase"
	username = "user"
	password = "password"
)

var db *sql.DB

type Product struct {
	ID int
	Name string
	Price int
}

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection Database Successful")

	// * create product
	// err = createProduct(&Product{Name: "Go product 2", Price: 444})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create Product Successful !")

	// * get product
	// product, err := getProduct(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get Product Successful !", product)

	// * update product
	// product, err := updateProduct(3, &Product{Name: "UUU", Price: 333})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update Product Successful !", product)

	// * delete product
	err = deleteProduct(6)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete Product Successful !")
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(
		"SELECT id, name, price FROM products WHERE id=$1;",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product

  row := db.QueryRow(
		`UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price;`,
		product.Name, 
		product.Price, 
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

  return p, err
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id = $1;",
		id,
	)

	return err
}