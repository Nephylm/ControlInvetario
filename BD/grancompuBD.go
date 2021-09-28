package bd

import (
	modelos "ControlInvetario/Modelos"
	grancompu "ControlInvetario/Utilidades"
	"fmt"
	"strconv"
)

var (
	CodigoProducto 	  string
	Final             string
	Ram               string
	Fecha             string
	OC                int
	Suc               string
	Clase             string
	Existencia        int
	Modelo            string
	IdProducto        int
	Marca             string
	Procesador        string
	Velocidad         string
	Generacion        int
	Camara            string
	SerieBateria      string
	Eliminador        string
	SerieOriginal     string
	Serie             int
	SerieDesktop      string
	SerieDistribuidor string
	Tipo              string
	Pulgadas          string
	Formato           string
	Salida            string
	MemGB             string
	MemGBLaptop       int
	Licencia          string
	HddGB             string
	Familia           string
	HddSerie          string
	UnidadOp          string
	Fuente            string
	Comentarios       string
	FechaVent         []uint8
	Forma             string
	Base              string
	Monitorescol      string
	HDMI              string
	Tamaño            string
	SerieDoc          string
	DocVent           string

)

func Guardar(inventario []grancompu.Item) string {
	for _, item := range inventario {
		switch grancompu.Minusculas(item.Producto["familia"]) {
		case "desktop":
			fmt.Println("Desktop")
			Desktop(item)
		case "laptop":
			fmt.Println("laptop")
			Laptops(item)
		case "monitor":
			fmt.Println("Monitor")
			Monitores(item)
		case "monitors":
			fmt.Println("Monitor")
			Monitores(item)
		default:

		}

	}
	return "Guardado exitoso"
}

//almacena los monitores en la BD
func Monitores(item grancompu.Item) {
	var monitor modelos.Monitor
	monitor.Fecha = item.Producto["fecha"]
	monitor.OC, _ = strconv.Atoi(item.Producto["oc"])
	monitor.Suc = item.Producto["suc"]
	monitor.CodigoProducto = item.Producto["codigo producto"]
	monitor.Familia = item.Producto["familia"]
	if item.Producto["serie"]==""{
		monitor.Serie =""
	}
	monitor.SerieOriginal = item.Producto["serie original"]
	monitor.Marca = item.Producto["marca"]
	monitor.Modelo = item.Producto["modelo"]
	monitor.Forma = item.Producto["forma"]
	monitor.Tipo = item.Producto["tipo"]
	monitor.Salidas = item.Producto["salidas"]
	monitor.HDMI = item.Producto["hdmi"]
	monitor.Clase = item.Producto["clase"]
	monitor.Tamaño = item.Producto["tamaño"]
	monitor.Base = item.Producto["base"]
	stmt, es := db.Prepare("INSERT INTO Monitores (Fecha, OC, Suc, Familia, CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Forma, " +
		" Base, Tipo, Salidas, HDMI, Clase, Tamaño) " +
		" SELECT ? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? WHERE NOT EXISTS (SELECT *FROM Monitores WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(monitor.Fecha, monitor.OC, monitor.Suc, monitor.Familia, monitor.CodigoProducto, monitor.Serie, monitor.SerieOriginal,
		monitor.Marca, monitor.Modelo, monitor.Forma, monitor.Base, monitor.Tipo, monitor.Salidas, monitor.HDMI,
		monitor.Clase, monitor.Tamaño, monitor.SerieOriginal)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}

//almacena las computadoras All in One en la BD
func AllinOne(item grancompu.Item) {
	var AllOne modelos.AllInOne
	AllOne.Fecha = item.Producto["fecha"]
	AllOne.OC, _ = strconv.Atoi(item.Producto["oc"])
	AllOne.SUC = item.Producto["suc"]
	AllOne.Familia = item.Producto["familia"]
	AllOne.Serie, _ = strconv.Atoi(item.Producto["serie"])
	AllOne.SerieOriginal = item.Producto["serie original"]
	AllOne.Marca = item.Producto["marca"]
	AllOne.Modelo = item.Producto["modelo"]
	AllOne.Procesador = item.Producto["procesador"]
	AllOne.Gen, _ = strconv.Atoi(item.Producto["gen"])
	AllOne.Mem_GB, _ = item.Producto["mem/gb"]
	AllOne.Velocidad = item.Producto["vel /ghz"]
	AllOne.HDD = item.Producto["hdd/gb"]
	AllOne.HddSerie = item.Producto["hdd serie"]
	AllOne.UnidadOp = item.Producto["unidad optica"]
	AllOne.Fuente = item.Producto["fuente serie"]
	AllOne.Formato = item.Producto["formato"]
	AllOne.Pulgadas = item.Producto["pulgadas"]
	AllOne.Licencia = item.Producto["licencia"]
	AllOne.Comentarios = item.Producto["comentarios"]
	stmt, es := db.Prepare("INSERT INTO AllinOne (Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Procesador, Gen, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Pulgadas, Licencia, Comentarios)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM AllinOne WHERE SerieOriginal=? OR Serie=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(AllOne.Fecha, AllOne.OC, AllOne.SUC, AllOne.Familia, AllOne.Serie, AllOne.SerieOriginal, AllOne.Marca,
		AllOne.Modelo, AllOne.Procesador, AllOne.Gen, AllOne.Mem_GB, AllOne.Velocidad, AllOne.HDD, AllOne.HddSerie,
		AllOne.UnidadOp, AllOne.Fuente, AllOne.Formato, AllOne.Pulgadas, AllOne.Licencia, AllOne.Comentarios,
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
func Laptops(item grancompu.Item) {
	var Laptop modelos.Laptop
	Laptop.Fecha = item.Producto["fecha"]
	Laptop.OC, _ = strconv.Atoi(item.Producto["oc"])
	Laptop.Suc = item.Producto["suc"]
	Laptop.Familia = item.Producto["familia"]
	Laptop.CodigoProducto = item.Producto["codigo producto"]
	Laptop.Marca = item.Producto["marca"]
	Laptop.Modelo = item.Producto["modelo"]
	Laptop.Procesador = item.Producto["procesad"]
	Laptop.Generacion, _ = strconv.Atoi(item.Producto["gen"])
	Laptop.Velocidad = item.Producto["veloghz"]
	Laptop.MemGB, _ = strconv.Atoi(item.Producto["memgb"])
	Laptop.SerieBateria = item.Producto["serie bateria"]
	Laptop.HddGB = item.Producto["disco"]
	Laptop.HddSerie = item.Producto["serie disco"]
	Laptop.SerieOriginal = item.Producto["serie original"]
	Laptop.Pulgadas = item.Producto["pulgadas"]
	Laptop.Camara = item.Producto["camara"]
	Laptop.Eliminador = item.Producto["eliminador"]
	//Laptop.Comentarios=item.Producto["comentarios"]
	stmt, es := db.Prepare("INSERT INTO Laptop (Fecha, OC, SUC, Familia, Marca, Modelo, Procesador, Generacion,Velocidad, Mem_GB," +
		"SerieBateria , HDD, HddSerie, SerieOriginal, Pulgadas, Camara, Eliminador)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM Laptop WHERE SerieOriginal=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Laptop.Fecha, Laptop.OC, Laptop.Suc, Laptop.Familia, Laptop.Marca, Laptop.Modelo, Laptop.Procesador, Laptop.Generacion,
		Laptop.Velocidad, Laptop.MemGB, Laptop.SerieBateria, Laptop.HddGB, Laptop.HddSerie, Laptop.SerieOriginal,
		Laptop.Pulgadas, Laptop.Camara, Laptop.Eliminador, Laptop.SerieOriginal)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}

//almacena las computadoras de escritorio en la BD
func Desktop(item grancompu.Item) {
	var Escritorio modelos.Desktop
	Escritorio.Fecha = item.Producto["fecha"]
	Escritorio.OC, _ = strconv.Atoi(item.Producto["oc"])
	Escritorio.Suc = item.Producto["suc"]
	Escritorio.Familia = item.Producto["familia"]
	Escritorio.CodigoProducto=item.Producto["codigo producto"]
	Escritorio.Serie = item.Producto["serie"]
	Escritorio.SerieOriginal = item.Producto["serie original"]
	Escritorio.Marca = item.Producto["marca"]
	Escritorio.Modelo = item.Producto["modelo"]
	Escritorio.Procesador = item.Producto["procesador"]
	Escritorio.Generacion, _ = strconv.Atoi(item.Producto["gen"])
	Escritorio.MemGB = item.Producto["mem/gb"]
	Escritorio.Velocidad = item.Producto["vel /ghz"]
	Escritorio.HddGB = item.Producto["hdd/gb"]
	Escritorio.HddSerie = item.Producto["hdd serie"]
	Escritorio.UnidadOpt = item.Producto["unidad optica"]
	Escritorio.FuenteSerie = item.Producto["fuente serie"]
	Escritorio.Formato = item.Producto["formato"]
	Escritorio.Licencia = item.Producto["licencia"]
	Escritorio.Comentarios = item.Producto["comentarios"]
	if Escritorio.Serie == "" || Escritorio.Serie == "ok" {
		Final = "SerieOriginal=? AND Serie=?);"
	} else {
		Final = "SerieOriginal=? OR Serie=?);"
	}
	stmt, es := db.Prepare("INSERT INTO Desktop (Fecha, OC, SUC, Familia,CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Licencia, Comentarios)" +
		" SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT *FROM Desktop WHERE " + Final)
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Escritorio.Fecha, Escritorio.OC, Escritorio.Suc, Escritorio.Familia, Escritorio.CodigoProducto, Escritorio.Serie, Escritorio.SerieOriginal, Escritorio.Marca,
		Escritorio.Modelo, Escritorio.Procesador, Escritorio.Generacion, Escritorio.MemGB, Escritorio.Velocidad, Escritorio.HddGB, Escritorio.HddSerie,
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
	listado, _ := db.Query("SELECT IdMonitores, Fecha, OC, Suc, Familia, CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Forma," +
		" Base, Tipo, Salidas, HDMI, Clase, Tamaño FROM Monitores WHERE Activo=1;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&SerieDesktop,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Forma,
			&Base,
			&Tipo,
			&Salida,
			&HDMI,
			&Clase,
			&Tamaño,
		)
		revisarError(err)
		Data = append(Data, modelos.Monitor{
			IdProducto: IdProducto,
			Fecha:          Fecha,
			OC:             OC,
			Suc:            Suc,
			Familia:        Familia,
			CodigoProducto: CodigoProducto,
			Serie:          SerieDesktop,
			SerieOriginal:  SerieOriginal,
			Marca:          Marca,
			Modelo:         Modelo,
			Forma:          Forma,
			Base:           Base,
			Tipo:           Tipo,
			Salidas:        Salida,
			HDMI:           HDMI,
			Clase:          Clase,
			Tamaño:         Tamaño,
		})
	}
	Data = append(Data, GetMonitoresInactivo()...)
	return
}
func GetMonitoresInactivo() (Data []modelos.Monitor) {
	listado, _ := db.Query("SELECT IdMonitores, Fecha, OC, Suc, Familia, CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Forma," +
		" Base, Tipo, Salidas, HDMI, Clase, Tamaño, FechaVent, SerieDoc, DocVent  FROM Monitores WHERE Activo=0;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&SerieDesktop,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Forma,
			&Base,
			&Tipo,
			&Salida,
			&HDMI,
			&Clase,
			&Tamaño,
			&FechaVent,
			&SerieDoc,
			&DocVent,
		)
		revisarError(err)
		Data = append(Data, modelos.Monitor{
			IdProducto: IdProducto,
			Fecha:          Fecha,
			OC:             OC,
			Suc:            Suc,
			Familia:        Familia,
			CodigoProducto: CodigoProducto,
			Serie:          SerieDesktop,
			SerieOriginal:  SerieOriginal,
			Marca:          Marca,
			Modelo:         Modelo,
			Forma:          Forma,
			Base:           Base,
			Tipo:           Tipo,
			Salidas:        Salida,
			HDMI:           HDMI,
			Clase:          Clase,
			Tamaño:         Tamaño,
			FechaVent:      string(FechaVent),
			SerieDoc:       SerieDoc,
			DocVent:        DocVent,
		})
	}
	return
}

//Recupera las computadoras all in one de la base de datos
func GetAllInOne() (Data []modelos.AllInOne) {
	listado, _ := db.Query("SELECT Fecha, OC, SUC, Familia, Serie, SerieOriginal, Marca, Modelo, Procesador, Gen, Mem_GB," +
		"Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Pulgadas, Licencia, Comentarios FROM AllinOne;")
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
			&Pulgadas,
			&Licencia,
			&Comentarios,
		)
		revisarError(err)
		Data = append(Data, modelos.AllInOne{
			Fecha:         Fecha,
			OC:            OC,
			SUC:           Suc,
			Familia:       Familia,
			Serie:         Serie,
			SerieOriginal: SerieOriginal,
			Marca:         Marca,
			Modelo:        Modelo,
			Procesador:    Procesador,
			Gen:           Generacion,
			Mem_GB:        MemGB,
			Velocidad:     Velocidad,
			HDD:           HddGB,
			HddSerie:      HddSerie,
			UnidadOp:      UnidadOp,
			Fuente:        Fuente,
			Formato:       Formato,
			Pulgadas:      Pulgadas,
			Licencia:      Licencia,
			Comentarios:   Comentarios,
		})
	}
	return
}

//Recupera las laptops activas de la base de datos
func GetLaptop() (Data []modelos.Laptop) {
	listado, _ := db.Query("SELECT IdLaptop,Fecha, OC, SUC, Familia,CodigoProducto, Marca, Modelo, Procesador, Generacion,Velocidad, Mem_GB," +
		"SerieBateria , HDD, HddSerie, SerieOriginal, Pulgadas, Camara, Eliminador, Activo FROM Laptop WHERE Activo=1;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&Velocidad,
			&MemGBLaptop,
			&SerieBateria,
			&HddGB,
			&HddSerie,
			&SerieOriginal,
			&Pulgadas,
			&Camara,
			&Eliminador,
			&Activo,
		)
		revisarError(err)
		Data = append(Data, modelos.Laptop{
			IdProducto:    IdProducto,
			Fecha:         Fecha,
			OC:            OC,
			Suc:           Suc,
			Familia:       Familia,
			CodigoProducto: CodigoProducto,
			Marca:         Marca,
			Modelo:        Modelo,
			Procesador:    Procesador,
			Generacion:    Generacion,
			Velocidad:     Velocidad,
			MemGB:         MemGBLaptop,
			SerieBateria:  SerieBateria,
			HddGB:         HddGB,
			HddSerie:      HddSerie,
			SerieOriginal: SerieOriginal,
			Pulgadas:      Pulgadas,
			Camara:        Camara,
			Eliminador:    Eliminador,
			Activo:        Activo,
		})
	}

	Data = append(Data, GetLaptopInactiva()...)
	return
}

//Recupera las laptops inactivas de la base de datos
func GetLaptopInactiva() (Data []modelos.Laptop) {
	PQuery := "SELECT IdLaptop,Fecha, OC, SUC, Familia,CodigoProducto, Marca, Modelo, Procesador, Generacion,Velocidad, Mem_GB," +
		"SerieBateria , HDD, HddSerie, SerieOriginal, Pulgadas, Camara, Eliminador, Activo, FechaVent, SerieDoc, DocVent FROM Laptop WHERE Activo=0;"
	listado, _ := db.Query(PQuery)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&Velocidad,
			&MemGBLaptop,
			&SerieBateria,
			&HddGB,
			&HddSerie,
			&SerieOriginal,
			&Pulgadas,
			&Camara,
			&Eliminador,
			&Activo,
			&FechaVent,
			&SerieDoc,
			&DocVent,

		)
		revisarError(err)
		Data = append(Data, modelos.Laptop{
			IdProducto:    IdProducto,
			Fecha:         Fecha,
			OC:            OC,
			Suc:           Suc,
			Familia:       Familia,
			CodigoProducto: CodigoProducto,
			Marca:         Marca,
			Modelo:        Modelo,
			Procesador:    Procesador,
			Generacion:    Generacion,
			Velocidad:     Velocidad,
			MemGB:         MemGBLaptop,
			SerieBateria:  SerieBateria,
			HddGB:         HddGB,
			HddSerie:      HddSerie,
			SerieOriginal: SerieOriginal,
			Pulgadas:      Pulgadas,
			Camara:        Camara,
			Eliminador:    Eliminador,
			Activo:        Activo,
			FechaVent:     string(FechaVent),
			SerieDoc: SerieDoc,
			DocVent: DocVent,
		})
	}
	return
}

//Recupera las computadoras de escritorio activas de la base de datos
func GetDesktop() (Data []modelos.Desktop) {
	listado, _ := db.Query("SELECT IdDesktop, Fecha, OC, SUC, Familia, CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Licencia, Comentarios,Activo FROM Desktop WHERE Activo=1;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&SerieDesktop,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&Ram,
			&Velocidad,
			&HddGB,
			&HddSerie,
			&UnidadOp,
			&Fuente,
			&Formato,
			&Licencia,
			&Comentarios,
			&Activo,
		)
		revisarError(err)
		Data = append(Data, modelos.Desktop{
			IdProducto:    IdProducto,
			Fecha:         Fecha,
			OC:            OC,
			Suc:           Suc,
			Familia:       Familia,
			CodigoProducto: CodigoProducto,
			Serie:         SerieDesktop,
			SerieOriginal: SerieOriginal,
			Marca:         Marca,
			Modelo:        Modelo,
			Procesador:    Procesador,
			Generacion:    Generacion,
			MemGB:         Ram,
			Velocidad:     Velocidad,
			HddGB:         HddGB,
			HddSerie:      HddSerie,
			UnidadOpt:     UnidadOp,
			FuenteSerie:   Fuente,
			Formato:       Formato,
			Licencia:      Licencia,
			Comentarios:   Comentarios,
			Activo:        Activo,
		})
	}
	Data = append(Data, GetDesktopInactivo()...)
	return
}

// Recupera las computadoras de escritorio inactivos de la base de datos
func GetDesktopInactivo() (Data []modelos.Desktop) {
	listado, _ := db.Query("SELECT IdDesktop, Fecha, OC, SUC, Familia,CodigoProducto, Serie, SerieOriginal, Marca, Modelo, Procesador, Generacion, Mem_GB," +
		" Velocidad, HDD, HddSerie, UnidadOp, Fuente, Formato, Licencia, Comentarios, Activo, FechaVent, SerieDoc, DocVent FROM Desktop WHERE Activo=0;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&Fecha,
			&OC,
			&Suc,
			&Familia,
			&CodigoProducto,
			&SerieDesktop,
			&SerieOriginal,
			&Marca,
			&Modelo,
			&Procesador,
			&Generacion,
			&Ram,
			&Velocidad,
			&HddGB,
			&HddSerie,
			&UnidadOp,
			&Fuente,
			&Formato,
			&Licencia,
			&Comentarios,
			&Activo,
			&FechaVent,
			&SerieDoc,
			&DocVent,
		)
		revisarError(err)
		Data = append(Data, modelos.Desktop{
			IdProducto:    IdProducto,
			Fecha:         Fecha,
			OC:            OC,
			Suc:           Suc,
			Familia:       Familia,
			CodigoProducto: CodigoProducto,
			Serie:         SerieDesktop,
			SerieOriginal: SerieOriginal,
			Marca:         Marca,
			Modelo:        Modelo,
			Procesador:    Procesador,
			Generacion:    Generacion,
			MemGB:         Ram,
			Velocidad:     Velocidad,
			HddGB:         HddGB,
			HddSerie:      HddSerie,
			UnidadOpt:     UnidadOp,
			FuenteSerie:   Fuente,
			Formato:       Formato,
			Licencia:      Licencia,
			Comentarios:   Comentarios,
			Activo:        Activo,
			FechaVent:     string(FechaVent),
			SerieDoc: SerieDoc,
			DocVent: DocVent,
		})
	}
	return
}

// Método para reducir los productos en existencia
func Reducir(red int, modelo string) string {
	elem := grancompu.Elemento{Modelo: modelo}
	coincidencia := GetItem(elem)
	coincidencia.Existencia = coincidencia.Existencia - red
	return Actualizar(coincidencia)
}

// Recupera datos de inventario
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
			Clase:      Clase,
			Existencia: Existencia,
			Modelo:     Modelo,
			IdProducto: IdProducto,
		})
	}
	return
}

// Modifica los productos de Inventario
func Actualizar(elme grancompu.Elemento) string {
	stmt, es := db.Prepare(" UPDATE inventario SET Clase = ?, Existencia = ?,Modelo = ?, IdProducto = ? WHERE Modelo=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(elme.Clase, elme.Existencia, elme.Modelo, elme.IdProducto, elme.Modelo)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Actualizacion exitosa exitoso")
		return "exito"
	}
	return "error"
}

// Obtiene los productos del inventario en base a su modelo
func GetItem(elme grancompu.Elemento) (Data grancompu.Elemento) {
	listado, _ := db.Query("SELECT Clase, Existencia, Modelo, IdProducto FROM inventario WHERE Modelo=?;", elme.Modelo)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Clase,
			&Existencia,
			&Modelo,
			&IdProducto,
		)
		revisarError(err)
		Data = grancompu.Elemento{
			Clase:      Clase,
			Existencia: Existencia,
			Modelo:     Modelo,
			IdProducto: IdProducto,
		}
	}
	return
}

func BajaLaptop(Laptop modelos.Laptop) (resp modelos.RespuestaSencilla) {
	stmt, es := db.Prepare("UPDATE Laptop SET Activo=0, FechaVent=CURDATE() WHERE IdLaptop=?;")
	//stmt, es := db.Prepare("UPDATE Laptop SET Activo=0, FechaVent=CURDATE() WHERE CodigoProducto= AND SerieOriginal=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Laptop.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Baja exitosa"
		fmt.Println("Baja exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al dar de baja"
		fmt.Println("Error al dar de baja")
	}
	return
}
func BajaDesktop(Desktop modelos.Desktop) (resp modelos.RespuestaSencilla) {

	stmt, es := db.Prepare("UPDATE Desktop SET Activo=0, FechaVent=CURDATE() WHERE IdDesktop=?;")
	//stmt, es := db.Prepare("UPDATE Desktop SET Activo=0, FechaVent=CURDATE() WHERE CodigoProducto=? AND Serie=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Desktop.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Baja exitosa"
		fmt.Println("Baja exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al dar de baja"
		fmt.Println("Error al dar de baja")
	}
	return
}
func BajaMonitor(Monitor modelos.Monitor) (resp modelos.RespuestaSencilla) {

	stmt, es := db.Prepare("UPDATE Monitores SET Activo=0, FechaVent=CURDATE(), SerieDoc=?, DocVent=? WHERE IdMonitores=?;")
	//stmt, es := db.Prepare("UPDATE Monitores SET Activo=0, FechaVent=CURDATE(), SerieDoc=?, DocVent=? WHERE CodigoProdcuto=? Serie=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Monitor.SerieDoc, Monitor.DocVent, Monitor.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Baja exitosa"
		fmt.Println("Baja exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al dar de baja"
		fmt.Println("Error al dar de baja")
	}
	return
}

func ActualizaLaptop(Laptop modelos.Laptop) (resp modelos.RespuestaSencilla) {

	stmt, es := db.Prepare("UPDATE Laptop SET Fecha=?, OC=?, SUC=?, Familia=?,CodigoProducto=?, Marca=?, Modelo=?, Procesador=?, Generacion=?,Velocidad=?, " +
		"Mem_GB=?,SerieBateria=? , HDD=?, HddSerie=?, SerieOriginal=?, Pulgadas=?, Camara=?, Eliminador=? WHERE IdLaptop=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Laptop.Fecha, Laptop.OC, Laptop.Suc, Laptop.Familia,Laptop.CodigoProducto, Laptop.Marca, Laptop.Modelo, Laptop.Procesador, Laptop.Generacion,
		Laptop.Velocidad, Laptop.MemGB, Laptop.SerieBateria, Laptop.HddGB, Laptop.HddSerie, Laptop.SerieOriginal,
		Laptop.Pulgadas, Laptop.Camara, Laptop.Eliminador, Laptop.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Actualizacion exitosa"
		fmt.Println("Actualizacion exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al Actualizar"
		fmt.Println("Error al Actualizar")
	}
	return
}
func ActualizaDesktop(Desktop modelos.Desktop) (resp modelos.RespuestaSencilla) {

	stmt, es := db.Prepare("UPDATE Desktop SET Fecha=?, OC=?, SUC=?, Familia=?,CodigoProducto=?, Serie=?, SerieOriginal=?, Marca=?, Modelo=?, Procesador=?," +
		"Generacion=?, Mem_GB=?,Velocidad=?, HDD=?, HddSerie=?, UnidadOp=?, Fuente=?, Formato=?, Licencia=?, Comentarios=? WHERE IdDesktop=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Desktop.Fecha, Desktop.OC, Desktop.Suc, Desktop.Familia,Desktop.CodigoProducto, Desktop.Serie, Desktop.SerieOriginal, Desktop.Marca,
		Desktop.Modelo, Desktop.Procesador, Desktop.Generacion, Desktop.MemGB, Desktop.Velocidad, Desktop.HddGB, Desktop.HddSerie,
		Desktop.UnidadOpt, Desktop.FuenteSerie, Desktop.Formato, Desktop.Licencia, Desktop.Comentarios, Desktop.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Actualizacion exitosa"
		fmt.Println("Actualizacion exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al Actualizar"
		fmt.Println("Error al Actualizar")
	}
	return
}
func ActualizaMonitor(Monitor modelos.Monitor) (resp modelos.RespuestaSencilla) {

	stmt, es := db.Prepare("UPDATE Monitores SET Fecha=?, OC=?, Suc=?, Familia=?, CodigoProducto=?, Serie=?, SerieOriginal=?, Marca=?," +
		" Modelo=?, Forma=?,Base=?, Tipo=?, Salidas=?, HDMI=?, Clase=?, Tamaño=? WHERE IdMonitores=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(Monitor.Fecha, Monitor.OC, Monitor.Suc, Monitor.Familia, Monitor.CodigoProducto, Monitor.Serie, Monitor.SerieOriginal,
		Monitor.Marca, Monitor.Modelo, Monitor.Forma, Monitor.Base, Monitor.Tipo, Monitor.Salidas, Monitor.HDMI,
		Monitor.Clase, Monitor.Tamaño, Monitor.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP = 200
		resp.Response = "Actualizacion exitosa"
		fmt.Println("Actualizacion exitosa")
	} else {
		resp.CodigoRespHTTP = 400
		resp.Response = "Error al Actualizar"
		fmt.Println("Error al Actualizar")
	}
	return
}
