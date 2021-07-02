package bd

import (
	grancompu "ControlInvetario/Utilidades"
	"fmt"
	"strconv"
)

var(
Clase string
Existencia int
Modelo string
IdProducto int)


func Guardar() string {

	inventario := grancompu.Contador(grancompu.Lista.Data, grancompu.Minusculas("CLASS"))
	var elme grancompu.Elemento
	for _,clases := range inventario{



			elme.Clase=clases["class"]
			elme.Existencia, _ =strconv.Atoi(clases["existencia"])
			elme.Modelo=clases["modelo"]
			if clases["class"]=="Desktops"{
				elme.IdProducto=4
			}else if clases["class"]=="Monitors"{
				elme.IdProducto=1
			}else if clases["class"]=="bateria"{
				elme.IdProducto=2
			}else if clases["class"]=="Adaptador"{
				elme.IdProducto=3
			}else if clases["class"]=="disco duro"{
				elme.IdProducto=5
			}
			stmt, es := db.Prepare("INSERT INTO inventario (Clase,Existencia,Modelo,IdProducto)" +
				" SELECT ?, ?, ?,? WHERE NOT EXISTS (SELECT *FROM inventario WHERE Modelo=?);")
			if es != nil {
				panic(es.Error())
			}
			a, err := stmt.Exec(elme.Clase, elme.Existencia,elme.Modelo,elme.IdProducto,elme.Modelo)

			revisarError(err)
			affected, _ := a.RowsAffected()
			if affected > 0 {
				fmt.Println("Guardado exitoso")
			}else{
				conicidencia :=GetItem(elme)
				elme.Existencia=conicidencia.Existencia+elme.Existencia
				return Actualizar(elme)
			}


	}
	return "Guardado exitoso"
}
func Reducir(red int, modelo string) string{
	elem:=grancompu.Elemento{Modelo: modelo}
	coincidencia:=GetItem(elem)
	coincidencia.Existencia=coincidencia.Existencia-red;
	return Actualizar(coincidencia)
}
func GetInventario() (Data []grancompu.Elemento) {
	listado, _ := db.Query("SELECT Clase, Existencia, Modelo, IdProducto FROM inventario ")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Existencia,
			&Modelo,
			&IdProducto,
		)
		revisarError(err)
		Data = append(Data, grancompu.Elemento{
			Clase:   Clase,
			Existencia: Existencia,
			Modelo: Modelo,
			IdProducto: IdProducto,
		})
	}
	return
}
func PruebaAlmacenar() string{
	inventario := grancompu.Contador(grancompu.Lista.Data, grancompu.Minusculas("CLASS"))
	var elme grancompu.Elemento
	elme.IdProducto=1
	elme.Clase=inventario[0]["class"]
	elme.Existencia,_=strconv.Atoi(inventario[0]["existencia"])
	elme.Modelo=inventario[0]["modelo"]
	//stmt, es := db.Prepare("INSERT INTO inventario (Clase,Existencia,Modelo,IdProducto) VALUES (?, ?, ?,?);")
	stmt, es := db.Prepare("INSERT INTO inventario (Clase,Existencia,Modelo,IdProducto)" +
		" SELECT ?, ?, ?,? WHERE NOT EXISTS (SELECT *FROM inventario WHERE Modelo=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(elme.Clase, elme.Existencia,elme.Modelo,elme.IdProducto,elme.Modelo)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		return"exito"
		fmt.Println("Guardado exitoso")
	} else{
		conicidencia :=GetItem(elme)
		elme.Existencia=conicidencia.Existencia+elme.Existencia
		return Actualizar(elme)
	}
	return "error"
}
func Actualizar(elme grancompu.Elemento)string{
	stmt, es := db.Prepare(" UPDATE inventario SET Clase = ?, Existencia = ?,Modelo = ?, IdProducto = ? WHERE Modelo=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(elme.Clase, elme.Existencia,elme.Modelo,elme.IdProducto,elme.Modelo)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Actualizacion exitosa exitoso")
		return "exito"
	}
	return "error"
}
func GetItem(elme grancompu.Elemento) (Data grancompu.Elemento) {
	listado, _ := db.Query("SELECT Clase, Existencia, Modelo, IdProducto FROM inventario WHERE Modelo=?;",elme.Modelo)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Existencia,
			&Modelo,
			&IdProducto,
		)
		revisarError(err)
		Data =  grancompu.Elemento{
			Clase:   Clase,
			Existencia: Existencia,
			Modelo: Modelo,
			IdProducto: IdProducto,
		}
	}
	return
}


