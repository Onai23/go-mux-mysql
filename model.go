package main

import (
	"database/sql"
	"fmt"
)

// deklarasi cetakan user
type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//fungsi menampilkan 1 user / single row
func (u *user) getUser(db *sql.DB) error {
	//return errors.New("Belum ada kode apapun.")
	statement := fmt.Sprintf("SELECT name FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name)
}

//fungsi memperbaharui data user
func (u *user) updateUser(db *sql.DB) error {
	//return errors.New("Belum ada kode apapun.")
	statement := fmt.Sprintf("UPDATE users SET name='%s' WHERE id=%d", u.Name, u.ID)
	_, err := db.Exec(statement)
	return err
}

//fungsi menghapus data user
func (u *user) deleteUser(db *sql.DB) error {
	//return errors.New("Belum ada kode apapun.")
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	//menjalankan statement
	_, err := db.Exec(statement)
	return err
}

//fungsi menyimpan user baru
func (u *user) createUser(db *sql.DB) error {
	//return errors.New("Belum ada kode apapun.")
	statement := fmt.Sprintf("INSERT INTO users(name) VALUES('%S')", u.Name)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

//fungsi menampilkan banyak users
func (u *user) getUsers(db *sql.DB, start, count int) ([]user, error) {
	//return nil, errors.New("Belum ada kode apapun.")
	statement := fmt.Sprintf("SELECT id, name FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//set objek slice user
	users := []user{}

	for rows.Next() {
		var u user //set objek cetakan user
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		//menggabungkan ke slice users
		users = append(users, u)
	}
	return users, nil
}
