package controllers

import (
	"log"
	"net/http"
	"store-app/models"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price")
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity")
		}

		models.CreateProduct(name, description, priceConvert, quantityConvert)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProduct(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting quantity")
		}

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price")
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity")
		}

		models.UpdateProduct(idConvert, name, description, priceConvert, quantityConvert)
	}

	http.Redirect(w, r, "/", 301)
}
