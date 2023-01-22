package fakultascontrollers

import (
	"golang-crud/entities"
	"golang-crud/models/fakultasmodels"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fakultas adalah key untuk dipanggil ke views
	fakultas := fakultasmodels.GetAll()
	data := map[string]any{
		"fakultas": fakultas,
	}

	temp, err := template.ParseFiles("views/fakultas/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	// jika request get maka tamilkan form
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/fakultas/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	// jika post

	// PERINGATAN BUAT SAYA SENDIRI : WAJIB RESTART SERVAHER SETELAH mELAKUKAN UPDATE PADA CODE
	if r.Method == "POST" {
		var inputfakultas entities.Fakultas_ent
		//panggil inpput dari forn di views
		inputfakultas.Nama_fakultas = r.FormValue("nama")
		inputfakultas.Created_at = time.Now()
		inputfakultas.Updated_at = time.Now()

		if sukses := fakultasmodels.Add(inputfakultas); !sukses {
			temp, err := template.ParseFiles("views/fakultas/add.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/Fakultas", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/fakultas/edit.html")
		if err != nil {
			panic(err)
		}
		// tangkap rparaneter id dari index
		// kenapa nennakai idString?
		// karena secara default query paraneneter yang dicapture golang itu dianggap string
		//strconv berfungsi untuk convert string ke integer
		// function atoi sendiri mengembalikan 2 parameter data yang sudah dicovnert dan error
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		// get id selesai

		fakultas := fakultasmodels.Detail(id)
		data := map[string]any{
			"fakultas": fakultas,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var fakultas entities.Fakultas_ent

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		fakultas.Nama_fakultas = r.FormValue("nama")
		fakultas.Updated_at = time.Now()

		if sukses := fakultasmodels.Edit(id, fakultas); sukses {
			http.Redirect(w, r, r.Header.Get("Refferer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/Fakultas", http.StatusSeeOther)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := fakultasmodels.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/Fakultas", http.StatusSeeOther)
}
