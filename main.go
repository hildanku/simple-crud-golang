package main

import (
	"golang-crud/config"
	"golang-crud/controllers/fakultascontrollers"
	"golang-crud/controllers/homecontrollers"
	"log"
	"net/http"
)

func main() {
	config.ConnDB()
	// homepage
	http.HandleFunc("/", homecontrollers.Welcome)

	// fakultas controller
	http.HandleFunc("/Fakultas/", fakultascontrollers.Index)
	http.HandleFunc("/Fakultas/add", fakultascontrollers.Add)
	http.HandleFunc("/Fakultas/edit", fakultascontrollers.Edit)
	http.HandleFunc("/Fakultas/delete", fakultascontrollers.Delete)

	log.Println("serverruninngport8080")
	http.ListenAndServe(":8080", nil)
}
