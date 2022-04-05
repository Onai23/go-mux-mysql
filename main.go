package main

func main() {
	//set objek cetakan App
	a := App{}

	//konfigurasi database
	a.Initialize("root", "root", "db_go_mux_mysql")
	a.Run(":8080")
}
