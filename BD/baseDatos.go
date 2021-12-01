package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Importa libreria compatible para database/sql
)

var (
	db  *sql.DB
	err error
)

//NuevaConexionBD configura la conexión a la base de datos
func NuevaConexionBD() {
	//Se configura la conexion a la base de datos
	db, err = sql.Open("mysql", "intel:Intel06!@tcp(74.208.31.248:3306)/Inventario")
	//db, err = sql.Open("mysql", "project-inventarios:Inventarios_in06;@tcp(189.236.207.53:3306)/Inventario")
	revisarError(err)
	//Se comprueba que la conexion siga activa
	err = db.Ping()
	revisarError(err)
	fmt.Println("Conectado a la BD!!")
}

//TerminarConexionBD termina la conexión de manera segura a la base de datos
func TerminarConexionBD() {
	db.Close()
}

//función para imprimir un error de consulta o conexión a la base de datos en caso de existir
func revisarError(err error) {

	if err != nil {
		panic(err)
	}
}
