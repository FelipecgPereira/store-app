package models

import "store-app/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAll() []Product {
	db := db.Connect()

	raws, err := db.Query("select * from  product order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	products := []Product{}

	for raws.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = raws.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.Connect()

	insert, err := db.Prepare("insert into product (name, description, price, quantity) values($1, $2, $3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.Connect()

	insert, err := db.Prepare("update product set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity, id)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.Connect()
	delete, err := db.Prepare("delete from product where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func FindProduct(id string) Product {
	db := db.Connect()

	raw, err := db.Query("select * from  product where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for raw.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = raw.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()

	return product
}
