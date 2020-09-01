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
	//abrirConexionDB()
	//resultadosQuery(tabla)
	//terminarConexion()
	Iniciar()

	//separador()

	//RecuperarXId(tabla,1)
	//agregarDatosBD(nombre)


}

func separador() {
	fmt.Println("")
}

func abrirConexionDB() {

	//db, err = sql.Open("mysql", string("root:Root@tcp(127.0.0.1:3306)/bienhechor"))
	db, err = sql.Open("mysql", string("bienhechor:Bienhechor_1234;@tcp(74.208.31.248:3306)/bienhechor"))
	revisarError(err)
	err = db.Ping()
	revisarError(err)
	if err == nil{
		fmt.Println("Conexion exitosa")
	}
}

func RecuperarXId(tabla string, where int64) {
	var tipo_membresia, id_membresia string

	unicoDato := db.QueryRow("SELECT tipo_membresia, id_membresia  FROM "+tabla+" WHERE id_membresia = ? ", where)

	err = unicoDato.Scan(&tipo_membresia, &id_membresia)
	revisarError(err)
	membresia.Id_membresia=id_membresia
	membresia.Tipo_membresia=tipo_membresia
	fmt.Println(tipo_membresia,"ID:", id_membresia)
}
func VerificarNombre(tabla string, where string) {
	var tipo_membresia, id_membresia string

	unicoDato := db.QueryRow("SELECT tipo_membresia, id_membresia  FROM "+tabla+" WHERE tipo_membresia = ? ", where)

	err = unicoDato.Scan(&tipo_membresia, &id_membresia)
	revisarError(err)
	fmt.Println(tipo_membresia,"ID:", id_membresia)
}

func resultadosQuery(tabla string) {
	membresias=nil
	query, _ := db.Query("SELECT * FROM " + tabla)
	separador()
	fmt.Println("id|membresia")
	for query.Next() {
		var idmembresia, tipomembresia string

		err = query.Scan(&idmembresia, &tipomembresia)
		revisarError(err)
		membresias= append(membresias,Membresia{Id_membresia: idmembresia, Tipo_membresia:tipomembresia})

		fmt.Println(idmembresia, tipomembresia)
	}
}

func agregarDatosBD(memberish Membresia) {
	var agregar sql.Result=nil
	var err error=nil
	if memberish.Id_membresia==""{
		queryTexto := "Insert into membresia (tipo_membresia)"
		agregar, err = db.Exec(queryTexto+" values (?)", memberish.Tipo_membresia)
	}else {
		queryTexto := "Insert into membresia (id_membresia,tipo_membresia)"
		agregar, err = db.Exec(queryTexto+" values (?,?)", memberish.Id_membresia,memberish.Tipo_membresia)
	}

	revisarError(err)
	status, err := agregar.LastInsertId()
	revisarError(err)
	separador()
	RecuperarXId(tabla,status)
	separador()
	//eliminarDatosBD(status)
}
func actualizarDatosBD(memberish Membresia, ID string) {
	if memberish.Id_membresia==""{
		db.Exec("Update membresia SET tipo_membresia= ? WHERE id_membresia= ?", memberish.Tipo_membresia,ID)

	}else {
		db.Exec("Update membresia SET id_membresia= ?, tipo_membresia= ?  WHERE id_membresia= ?",memberish.Id_membresia, memberish.Tipo_membresia,ID)

	}
	resultadosQuery(tabla)
}
func eliminarDatosBD(ID int64, err2 error) {
	eliminar, err := db.Exec("delete from membresia where id_membresia = ?", ID)
	revisarError(err)
	status, err := eliminar.RowsAffected()
	revisarError(err)
	fmt.Println(status)
	separador()
	resultadosQuery(tabla)
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

