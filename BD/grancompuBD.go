package bd

import (
	modelos "ControlInvetario/Modelos"
	grancompu "ControlInvetario/Utilidades"
	"fmt"
)

var(
Clase string
Existencia int
Modelo string
IdProducto int
Marca string
Procesador string
Velocidad string
Generacion string
MarcaDisco string
Capacidad string
SerieDisco string
Bateria string
Eliminador string
Memoria string
SerieOriginal string
SerieDistribuidor string
Tipo string
Pulgadas string
Formato string
Salida string
)


func Guardar(inventario []grancompu.Item) string {

	for _, item := range inventario{
		switch item.Producto["clase"] {
		case "monitor":
			Monitores(item)
		case "desktop":
			Desktop(item)
		case "allinone":
			AllinOne(item)
		case "laptop":
			Laptops(item)
		case "disco duro":

		default:
			return "error"
		}
	}
	return "Guardado exitoso"
}
//almacena los monitores en la BD
func Monitores (item grancompu.Item){
	var monitor modelos.Monitor
	monitor.Clase=item.Producto["clase"]
	monitor.Modelo=item.Producto["modelo"]
	monitor.Marca=item.Producto["marca"]
	monitor.Pulgadas=item.Producto["pulgadas"]
	monitor.SerieDistribuidor=item.Producto["serie distribuidor"]
	monitor.SerieOriginal=item.Producto["serie original"]
	monitor.Tipo=item.Producto["tipo"]
	monitor.Salida=item.Producto["salidas"]
	stmt, es := db.Prepare("INSERT INTO Monitores (Clase,Modelo,Marca,Pulgadas, SerieDistri, SerieOriginal, Tipo, Salidas)" +
		" SELECT ?, ?, ?, ?, ?,?,?,? WHERE NOT EXISTS (SELECT *FROM Monitores WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(monitor.Clase,monitor.Modelo,monitor.Marca,monitor.Pulgadas,
		monitor.SerieDistribuidor,monitor.SerieOriginal,monitor.Tipo,monitor.Salida,monitor.SerieOriginal)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
//almacena las computadoras All in One en la BD
func AllinOne (item grancompu.Item){
	var AllOne modelos.AllInOne
	AllOne.Clase= item.Producto["clase"]
	AllOne.Modelo=item.Producto["modelo"]
	AllOne.Marca=item.Producto["marca"]
	AllOne.Procesador=item.Producto["procesador"]
	AllOne.Velocidad=item.Producto["velocidad"]
	AllOne.Generacion= item.Producto["generacion"]
	AllOne.MarcaDisco=item.Producto["marca hdd"]
	AllOne.Capacidad=item.Producto["capacidad"]
	AllOne.SerieDisco=item.Producto["serie hdd"]
	AllOne.Fuente_Eliminador=item.Producto["fuente/eliminador"]
	AllOne.Memoria=item.Producto["memoria"]
	AllOne.SerieOriginal=item.Producto["serie original"]
	AllOne.SerieDistribuidor=item.Producto["serie distribuidor"]
	AllOne.Pulgadas= item.Producto["pulgadas"]
	stmt, es := db.Prepare("INSERT INTO AllinOne (Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHHD" +
		", Capacidad, SerieHHD, Eliminador, Memoria, SerieOriginal, SerieDistri, Pulgadas)" +
		" SELECT ?, ?, ?, ?, ?,?,?, ?, ?, ?, ?, ?,?,? WHERE NOT EXISTS (SELECT *FROM AllinOne WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(AllOne.Clase,AllOne.Marca,AllOne.Modelo,AllOne.Procesador,AllOne.Velocidad,
		AllOne.Generacion,AllOne.MarcaDisco,AllOne.Capacidad,AllOne.SerieDisco,AllOne.Fuente_Eliminador,
		AllOne.Memoria,AllOne.SerieOriginal,AllOne.SerieDistribuidor,AllOne.Pulgadas,AllOne.SerieOriginal)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
//almacena las Laptops All in One en la BD
func Laptops (item grancompu.Item){
	var Laptop modelos.Laptop
	Laptop.Clase= item.Producto["clase"]
	Laptop.Modelo=item.Producto["modelo"]
	Laptop.Marca=item.Producto["marca"]
	Laptop.Procesador=item.Producto["procesador"]
	Laptop.Velocidad=item.Producto["velocidad"]
	Laptop.Generacion= item.Producto["generacion"]
	Laptop.MarcaDisco=item.Producto["marca hdd"]
	Laptop.Capacidad=item.Producto["capacidad"]
	Laptop.SerieDisco=item.Producto["serie hdd"]
	Laptop.Bateria=item.Producto["bateria"]
	Laptop.Eliminador=item.Producto["eliminador"]
	Laptop.Memoria=item.Producto["memoria"]
	Laptop.SerieOriginal=item.Producto["serie original"]
	Laptop.SerieDistribuidor=item.Producto["serie distribuidor"]
	Laptop.Pulgadas= item.Producto["pulgadas"]
	stmt, es := db.Prepare("INSERT INTO Laptops (Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHDD" +
		", Capacidad, SerieHDD,Bateria, Eliminador, Memoria, SerieOriginal, SerieDistri, Pulgadas)" +
		" SELECT ?, ?, ?, ?, ?,?,?, ?, ?, ?, ?, ?,?,?,? WHERE NOT EXISTS (SELECT *FROM Laptops WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Laptop.Clase, Laptop.Marca, Laptop.Modelo, Laptop.Procesador, Laptop.Velocidad,
		Laptop.Generacion, Laptop.MarcaDisco, Laptop.Capacidad, Laptop.SerieDisco, Laptop.Bateria, Laptop.Eliminador,
		Laptop.Memoria, Laptop.SerieOriginal, Laptop.SerieDistribuidor, Laptop.Pulgadas, Laptop.SerieOriginal)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
//almacena las computadoras de escritorio en la BD
func Desktop (item grancompu.Item){
	var Escritorio modelos.Desktop
	Escritorio.Clase= item.Producto["clase"]
	Escritorio.Modelo=item.Producto["modelo"]
	Escritorio.Marca=item.Producto["marca"]
	Escritorio.Procesador=item.Producto["procesador"]
	Escritorio.Velocidad=item.Producto["velocidad"]
	Escritorio.Generacion= item.Producto["generacion"]
	Escritorio.MarcaDisco=item.Producto["marca hdd"]
	Escritorio.Capacidad=item.Producto["capacidad"]
	Escritorio.SerieDisco=item.Producto["serie hdd"]
	Escritorio.Fuente_Eliminador=item.Producto["fuente/eliminador"]
	Escritorio.Memoria=item.Producto["memoria"]
	Escritorio.SerieOriginal=item.Producto["serie original"]
	Escritorio.SerieDistribuidor=item.Producto["serie distribuidor"]
	Escritorio.Formato=item.Producto["formato"]
	stmt, es := db.Prepare("INSERT INTO Desktop (Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHDD" +
		", Capacidad, SerieHDD, Eliminador, Memoria, SerieOriginal, SerieDistri, Formato)" +
		" SELECT ?, ?, ?, ?, ?,?,?, ?, ?, ?, ?, ?,?,? WHERE NOT EXISTS (SELECT *FROM Desktop WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Escritorio.Clase, Escritorio.Marca, Escritorio.Modelo, Escritorio.Procesador, Escritorio.Velocidad,
		Escritorio.Generacion, Escritorio.MarcaDisco, Escritorio.Capacidad, Escritorio.SerieDisco, Escritorio.Fuente_Eliminador,
		Escritorio.Memoria, Escritorio.SerieOriginal, Escritorio.SerieDistribuidor,Escritorio.Formato, Escritorio.SerieOriginal)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
//Recupera los monitores de la base de datos
func GetMonitores() (Data []modelos.Monitor) {
	listado, _ := db.Query("SELECT Clase,Marca,Modelo,Pulgadas, Tipo, SerieOriginal,SerieDistri, Salidas FROM Monitores;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Marca,
			&Modelo,
			&Pulgadas,
			&Tipo,
			&SerieOriginal,
			&SerieDistribuidor,
			&Salida,
		)
		revisarError(err)
		Data =  append(Data,modelos.Monitor{
			Clase:   Clase,
			Marca: Marca,
			Modelo: Modelo,
			Pulgadas: Pulgadas,
			Tipo: Tipo,
			SerieOriginal: SerieOriginal,
			SerieDistribuidor: SerieDistribuidor,
			Salida: Salida,
		})
	}
	return
}
//Recupera las computadoras all in one de la base de datos
func GetAllInOne() (Data []modelos.AllInOne) {
	listado, _ := db.Query("SELECT Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHHD," +
		" Capacidad, SerieHHD, Eliminador, Memoria, SerieOriginal, SerieDistri, Pulgadas FROM AllinOne;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Marca,
			&Modelo,
			&Procesador,
			&Velocidad,
			&Generacion,
			&MarcaDisco,
			&Capacidad,
			&SerieDisco,
			&Eliminador,
			&Memoria,
			&SerieOriginal,
			&SerieDistribuidor,
			&Pulgadas,
		)
		revisarError(err)
		Data =  append(Data,modelos.AllInOne{
			Clase:   Clase,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Velocidad: Velocidad,
			Generacion: Generacion,
			MarcaDisco: MarcaDisco,
			Capacidad: Capacidad,
			SerieDisco: SerieDisco,
			Fuente_Eliminador: Eliminador,
			Memoria: Memoria,
			SerieOriginal: SerieOriginal,
			SerieDistribuidor: SerieDistribuidor,
			Pulgadas: Pulgadas,
		})
	}
	return
}
//Recupera las laptops de la base de datos
func GetLaptop() (Data []modelos.Laptop) {
	listado, _ := db.Query("SELECT Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHDD," +
		" Capacidad, SerieHDD,Bateria, Eliminador, Memoria, SerieOriginal, SerieDistri, Pulgadas FROM Laptops;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Marca,
			&Modelo,
			&Procesador,
			&Velocidad,
			&Generacion,
			&MarcaDisco,
			&Capacidad,
			&SerieDisco,
			&Bateria,
			&Eliminador,
			&Memoria,
			&SerieOriginal,
			&SerieDistribuidor,
			&Pulgadas,
		)
		revisarError(err)
		Data =  append(Data,modelos.Laptop{
			Clase:   Clase,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Velocidad: Velocidad,
			Generacion: Generacion,
			MarcaDisco: MarcaDisco,
			Capacidad: Capacidad,
			SerieDisco: SerieDisco,
			Bateria: Bateria,
			Eliminador: Eliminador,
			Memoria: Memoria,
			SerieOriginal: SerieOriginal,
			SerieDistribuidor: SerieDistribuidor,
			Pulgadas: Pulgadas,
		})
	}
	return
}
//Recupera las computadoras de escritoria de la base de datos
func GetDesktop() (Data []modelos.Desktop) {
	listado, _ := db.Query("SELECT Clase,Marca,Modelo,Procesador, Velocidad, Generacion, MarcaHDD, Capacidad, SerieHDD, Eliminador, Memoria, SerieOriginal, SerieDistri, Formato FROM Desktop;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Marca,
			&Modelo,
			&Procesador,
			&Velocidad,
			&Generacion,
			&MarcaDisco,
			&Capacidad,
			&SerieDisco,
			&Eliminador,
			&Memoria,
			&SerieOriginal,
			&SerieDistribuidor,
			&Formato,
		)
		revisarError(err)
		Data =  append(Data, modelos.Desktop{
			Clase:   Clase,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Velocidad: Velocidad,
			Generacion: Generacion,
			MarcaDisco: MarcaDisco,
			Capacidad: Capacidad,
			SerieDisco: SerieDisco,
			Fuente_Eliminador: Eliminador,
			Memoria: Memoria,
			SerieOriginal: SerieOriginal,
			SerieDistribuidor: SerieDistribuidor,
			Formato: Formato,
		})
	}
	return
}
func Reducir(red int, modelo string) string{
	elem:=grancompu.Elemento{Modelo: modelo}
	coincidencia:=GetItem(elem)
	coincidencia.Existencia=coincidencia.Existencia-red
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


