package bd

import (
	modelos "ControlInvetario/Modelos"
	grancompu "ControlInvetario/Utilidades"
	"fmt"
)

var(
	Fecha string
	OC string
	Suc string
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
	Serie string
	SerieDistribuidor string
	Tipo string
	Pulgadas string
	Formato string
	Salida string
	NumNucleos string
	TipoRam string
	MemGB string
	HddTipo string
	Unidad string
	Cargador string
	Licencia string
	Extras string
	Provedor string
	HddGB string
	Familia string
	HddSerie string
	UnidadOp string
	Fuente string
	Comentarios string
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
	AllOne.Fecha=item.Producto["fecha"]
	AllOne.OC=item.Producto["oc"]
	AllOne.Suc=item.Producto["suc"]
	AllOne.Familia=item.Producto["familia"]
	AllOne.Serie=item.Producto["serie"]
	AllOne.SerieOriginal=item.Producto["serie original"]
	AllOne.Marca=item.Producto["marca"]
	AllOne.Modelo=item.Producto["modelo"]
	AllOne.Procesador=item.Producto["procesador"]
	AllOne.Generacion= item.Producto["gen"]
	AllOne.MemGB=item.Producto["mem/gb"]
	AllOne.Velocidad=item.Producto["vel /ghz"]
	AllOne.HddGB=item.Producto["hdd/gb"]
	AllOne.HddSerie=item.Producto["hdd serie"]
	AllOne.UnidadOpt=item.Producto["unidad optica"]
	AllOne.FuenteSerie=item.Producto["fuente serie"]
	AllOne.Pulgadas=item.Producto["pulgadas"]
	AllOne.Licencia=item.Producto["licencia"]
	AllOne.Comentarios=item.Producto["comentarios"]
	stmt, es := db.Prepare("INSERT INTO AllinOne (Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Prcesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Pulgadas, Licencia, Comentarios)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM AllinOne WHERE SerieOriginal=? OR Serie=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(AllOne.Fecha,AllOne.OC,AllOne.Suc,AllOne.Familia,AllOne.Serie,AllOne.SerieOriginal,AllOne.Marca,
		AllOne.Modelo,AllOne.Procesador, AllOne.Generacion, AllOne.MemGB, AllOne.Velocidad, AllOne.HddGB, AllOne.HddSerie,
		AllOne.UnidadOpt, AllOne.FuenteSerie, AllOne.Pulgadas, AllOne.Licencia, AllOne.Comentarios,
		AllOne.SerieOriginal, AllOne.Serie)
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
	Laptop.Fecha=item.Producto["fecha"]
	Laptop.OC=item.Producto["oc"]
	Laptop.Suc=item.Producto["suc"]
	Laptop.Familia=item.Producto["familia"]
	Laptop.Serie=item.Producto["serie"]
	Laptop.SerieOriginal=item.Producto["serie original"]
	Laptop.Marca=item.Producto["marca"]
	Laptop.Modelo=item.Producto["modelo"]
	Laptop.Procesador=item.Producto["procesador"]
	Laptop.Generacion= item.Producto["gen"]
	Laptop.MemGB=item.Producto["mem/gb"]
	Laptop.Velocidad=item.Producto["vel /ghz"]
	Laptop.HddGB=item.Producto["hdd/gb"]
	Laptop.HddSerie=item.Producto["hdd serie"]
	Laptop.UnidadOpt=item.Producto["unidad optica"]
	Laptop.FuenteSerie=item.Producto["fuente serie"]
	Laptop.Pulgadas=item.Producto["pulgadas"]
	Laptop.Licencia=item.Producto["licencia"]
	Laptop.Comentarios=item.Producto["comentarios"]
	stmt, es := db.Prepare("INSERT INTO Laptop (Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Prcesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Pulgadas, Licencia, Comentarios)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM Laptop WHERE SerieOriginal=? OR Serie=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Laptop.Fecha,Laptop.OC,Laptop.Suc,Laptop.Familia,Laptop.Serie,Laptop.SerieOriginal,Laptop.Marca,
		Laptop.Modelo,Laptop.Procesador, Laptop.Generacion, Laptop.MemGB, Laptop.Velocidad, Laptop.HddGB, Laptop.HddSerie,
		Laptop.UnidadOpt, Laptop.FuenteSerie, Laptop.Pulgadas, Laptop.Licencia, Laptop.Comentarios,
		Laptop.SerieOriginal, Laptop.Serie)
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
	Escritorio.Fecha=item.Producto["fecha"]
	Escritorio.OC=item.Producto["oc"]
	Escritorio.Suc=item.Producto["suc"]
	Escritorio.Familia=item.Producto["familia"]
	Escritorio.Serie=item.Producto["serie"]
	Escritorio.SerieOriginal=item.Producto["serie original"]
	Escritorio.Marca=item.Producto["marca"]
	Escritorio.Modelo=item.Producto["modelo"]
	Escritorio.Procesador=item.Producto["procesador"]
	Escritorio.Generacion= item.Producto["gen"]
	Escritorio.MemGB=item.Producto["mem/gb"]
	Escritorio.Velocidad=item.Producto["vel /ghz"]
	Escritorio.HddGB=item.Producto["hdd/gb"]
	Escritorio.HddSerie=item.Producto["hdd serie"]
	Escritorio.UnidadOpt=item.Producto["unidad optica"]
	Escritorio.FuenteSerie=item.Producto["fuente serie"]
	Escritorio.Formato=item.Producto["formato"]
	Escritorio.Licencia=item.Producto["licencia"]
	Escritorio.Comentarios=item.Producto["comentarios"]
	stmt, es := db.Prepare("INSERT INTO Desktop (Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Prcesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Licencia, Comentarios)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM Desktop WHERE SerieOriginal=? OR Serie=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Escritorio.Fecha,Escritorio.OC,Escritorio.Suc,Escritorio.Familia,Escritorio.Serie,Escritorio.SerieOriginal,Escritorio.Marca,
		Escritorio.Modelo,Escritorio.Procesador, Escritorio.Generacion, Escritorio.MemGB, Escritorio.Velocidad, Escritorio.HddGB, Escritorio.HddSerie,
		Escritorio.UnidadOpt, Escritorio.FuenteSerie, Escritorio.Formato, Escritorio.Licencia, Escritorio.Comentarios,
		Escritorio.SerieOriginal, Escritorio.Serie)
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
	listado, _ := db.Query("SELECT Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		"Velocidad, HDD, HddSerie, UnidadOp, Fuente, Pulgadas, Licencia, Comentarios FROM AllinOne;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&Serie,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&MemGB,
			&Velocidad,
			&HddGB,
			&HddSerie,
			&UnidadOp,
			&Fuente,
			&Pulgadas,
			&Licencia,
			&Comentarios,
		)
		revisarError(err)
		Data =  append(Data, modelos.AllInOne{
			Fecha: Fecha,
			OC: OC,
			Suc: Suc,
			Familia: Familia,
			Serie: Serie,
			SerieOriginal: SerieOriginal,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Generacion: Generacion,
			MemGB: MemGB,
			Velocidad: Velocidad,
			HddGB: HddGB,
			HddSerie: HddSerie,
			UnidadOpt: UnidadOp,
			FuenteSerie: Fuente,
			Pulgadas: Pulgadas,
			Licencia: Licencia,
			Comentarios: Comentarios,
		})
	}
	return
}



//Recupera las laptops de la base de datos
func GetLaptop() (Data []modelos.Laptop) {
	listado, _ := db.Query("SELECT Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		"Velocidad, HDD, HddSerie, UnidadOp, Fuente, Pulgadas, Licencia, Comentarios FROM Laptop;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&Serie,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&MemGB,
			&Velocidad,
			&HddGB,
			&HddSerie,
			&UnidadOp,
			&Fuente,
			&Pulgadas,
			&Licencia,
			&Comentarios,
		)
		revisarError(err)
		Data =  append(Data, modelos.Laptop{
			Fecha: Fecha,
			OC: OC,
			Suc: Suc,
			Familia: Familia,
			Serie: Serie,
			SerieOriginal: SerieOriginal,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Generacion: Generacion,
			MemGB: MemGB,
			Velocidad: Velocidad,
			HddGB: HddGB,
			HddSerie: HddSerie,
			UnidadOpt: UnidadOp,
			FuenteSerie: Fuente,
			Pulgadas: Pulgadas,
			Licencia: Licencia,
			Comentarios: Comentarios,
		})
	}
	return
}
//Recupera las computadoras de escritoria de la base de datos
func GetDesktop() (Data []modelos.Desktop) {
	listado, _ := db.Query("SELECT Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		"Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Licencia, Comentarios FROM Desktop;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&Serie,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&MemGB,
			&Velocidad,
			&HddGB,
			&HddSerie,
			&UnidadOp,
			&Fuente,
			&Formato,
			&Licencia,
			&Comentarios,
		)
		revisarError(err)
		Data =  append(Data, modelos.Desktop{
			Fecha: Fecha,
			OC: OC,
			Suc: Suc,
			Familia: Familia,
			Serie: Serie,
			SerieOriginal: SerieOriginal,
			Marca: Marca,
			Modelo: Modelo,
			Procesador: Procesador,
			Generacion: Generacion,
			MemGB: MemGB,
			Velocidad: Velocidad,
			HddGB: HddGB,
			HddSerie: HddSerie,
			UnidadOpt: UnidadOp,
			FuenteSerie: Fuente,
			Formato: Formato,
			Licencia: Licencia,
			Comentarios: Comentarios,
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


