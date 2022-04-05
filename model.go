package main

import (
	"database/sql"
	"errors"
)

// deklarasi cetakan user
type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//fungsi menampilkan 1 user
func (u *user) getUser(db *sql.DB) error {
	return errors.New("Belum ada kode apapun.")
}

//fungsi memperbaharui data user
func (u *user) updateUser(db *sql.DB) error {
	return errors.New("Belum ada kode apapun.")
}

//fungsi menghapus data user
func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("Belum ada kode apapun.")
}

//fungsi menyimpan user baru
func (u *user) createUser(db *sql.DB) error {
	return errors.New("Belum ada kode apapun.")
}

//fungsi menampilkan banyak users
func (u *user) getUsers(db *sql.DB) error {
	return errors.New("Belum ada kode apapun.")
}

func main() {

}
