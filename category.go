package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type categorie struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

func relasi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var products product
	db.Model(&product{}).Association("Categorie")

	result := db.Preload("categorie").First(&products)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	hasil, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(hasil))
}

func createCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var categorie categorie
	err := json.NewDecoder(r.Body).Decode(&categorie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Select("Nama").Create(&categorie)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func getAllCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var categories []categorie
	result := db.Find(&categories)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	hasil, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func getCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var categories []categorie
	result := db.First(&categories, id)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	hasil, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func updateCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var categories categorie
	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	db, errGorm := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Model(&categorie{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nama": categories.Nama,
	})
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func deleteCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Delete(&categorie{}, id)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}
