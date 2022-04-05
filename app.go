package main

import (
	"database/sql"
	"github.com/gorilla/mux"
)

//Deklarasi cetakan App
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//method initialize bertanggung jawab untuk membangun koneksi ke database
func (a *App) Initialize(user, password, dbName string) {

}

//method run digunakan untuk menjalankan script golang
func (a *App) Run(addr string) {

}
