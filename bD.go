package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	db  *sql.DB
	tabla = "membresia"
)

func main() {
	Iniciar()
}

func separador() {
	fmt.Println("")
}

func abrirConexionDB() {

	db, err = sql.Open("mysql", string("project-inventarios:Inventarios_in06;@tcp(189.236.90.166)/Inventario"))
	revisarError(err)
	err = db.Ping()
	revisarError(err)
	if err == nil{
		fmt.Println("Conexion exitosa")
	}
}



func revisarError(err error) {
	if err != nil {
		panic(err)
	}
}

func terminarConexion() {
	defer db.Close()
	fmt.Println("conexion terminada")
}

