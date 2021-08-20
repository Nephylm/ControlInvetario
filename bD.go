package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err   error
	db    *sql.DB
	tabla = "membresia"
)

func separador() {
	fmt.Println("")
}

// Configura la conexión a la base de datos
func abrirConexionDB() {
	//db, err = sql.Open("mysql", string("project-inventarios:Inventarios_in06;@tcp(189.236.90.166)/Inventario"))
	db, err = sql.Open("mysql", string("root:Root;@tcp(localhost:3306)/Inventario"))
	revisarError(err)
	err = db.Ping()
	revisarError(err)
	if err == nil {
		fmt.Println("Conexion exitosa")
	}
}

//función para imprimir un error de consulta o conexión a la base de datos en caso de existir
func revisarError(err error) {
	if err != nil {
		panic(err)
	}
}

//terminarConexion termina la conexión de manera segura a la base de datos
func terminarConexion() {
	defer db.Close()
	fmt.Println("conexion terminada")
}
