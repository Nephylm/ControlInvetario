package bd

import (
	modelos "ControlInvetario/Modelos"
	"fmt"
	"time"
)
func Registro(mercacancia []modelos.Mercancia) interface{} {
	for _, item := range mercacancia{
		RegistroMercancia(item)
	}

	return nil
}

func RegistroMercancia (mercancia modelos.Mercancia) {
	today := time.Now()
	stmt, es := db.Prepare("INSERT INTO Registro (Fecha,Familia,SerieOriginal) SELECT "+today.Format("2006/1/2")+",?,? " +
		"WHERE NOT EXISTS (SELECT *FROM Registro WHERE Familia=? AND SerieOriginal=?")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(mercancia.Familia,mercancia.SerieOriginal,mercancia.Familia,mercancia.SerieOriginal)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		stmt, er:=db.Prepare("INSERT INTO Inventario (IdRegistro) SELECT (SELECT IdRegistro FROM Registro WHERE SerieOriginal=? AND Familia=?);")
		if er != nil {
			panic(es.Error())
		}
		a,err=stmt.Exec(mercancia.SerieOriginal,mercancia.Familia)
		revisarError(err)
		affected, _ = a.RowsAffected()
		if affected > 0 {
			fmt.Println("Guardado exitoso")
		}else {
			fmt.Println("Error al registra tabla Inventario")
		}
	} else {
		fmt.Println("producto ya registrado")
	}
}

func ActualizacionMercancia (mercancia modelos.Mercancia) {
	stmt, es := db.Prepare("UPDATE Registro SET IdRegistro=?,Fecha=?,OC=?,SUC=?,Familia=?,CodigoProducto=?,Serie=?,SerieOriginal=?,Marca=?," +
		"Modelo=?,Procesador=?,Generacion=?,Mem_GB=?,Velocidad=?,HDD=?,HddSerie=?,UnidadOp=?,Fuente=?,Formato=?,Licencia=?,Comentarios=?,Pulgadas=?,Camara=?," +
		"Eliminador=?,SerieBateria=?,Forma=?,Base=?,TipoPantalla=?,Salidas=?,HDMI=?,Clase=?,Tamaño=? WHERE IdRegistro=?")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(mercancia.IdProducto,mercancia.Fecha,mercancia.OC,mercancia.SUC,mercancia.Familia,mercancia.CodigoProducto,mercancia.SerieOriginal,
		mercancia.Marca,mercancia.Modelo,mercancia.Procesador,mercancia.Gen,mercancia.Mem_GB,mercancia.Velocidad,mercancia.HDD,mercancia.HddSerie,
		mercancia.UnidadOp,mercancia.Fuente,mercancia.Formato,mercancia.Licencia,mercancia.Comentarios,mercancia.Pulgadas,mercancia.Camara,
		mercancia.Eliminador,mercancia.SerieBateria,mercancia.Forma,mercancia.Base,mercancia.Tipo,mercancia.Salidas,mercancia.HDMI,mercancia.Clase,
		mercancia.Tamanio,mercancia.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
			fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
func ConsultaMercancias() (Data []modelos.Mercancia) {
	listado, _ := db.Query("SELECT R.IdRegistro,R.Fecha,R.OC,R.SUC,R.Familia,R.CodigoProducto,R.Serie,R.SerieOriginal,R.Marca,R." +
		"Modelo,R.Procesador,R.Generacion,R.Mem_GB,R.Velocidad,R.HDD,R.HddSerie,R.UnidadOp,R.Fuente,R.Formato,R.Licencia,R.Comentarios,R.Pulgadas,R.Camara,R." +
		"Eliminador,R.SerieBateria,R.Forma,R.Base,R.TipoPantalla,R.Salidas,R.HDMI,R.Clase,R.Tamaño, I.Estado FROM Registro AS R INNER JOIN Inventario AS I ON R.IdRegistro = I.IdRegistro " +
		"WHERE I.Estado=1")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto, &Fecha, &OC, &Suc,
			&Familia, &CodigoProducto, &SerieDesktop, &SerieOriginal,
			&Marca, &Modelo, &Procesador, &Generacion,
			&MemGB, &Velocidad, &HddGB, &HddSerie,
			&UnidadOp, &Fuente, &Formato, &Licencia,
			&Comentarios, &Pulgadas, &Camara, &Eliminador,
			&SerieBateria, &Forma, &Base, &Tipo,
			&Salida, &HDMI, &Clase, &Tamaño,
			&Activo,
		)
		revisarError(err)
		Data = append(Data, modelos.Mercancia{
			IdProducto: IdProducto,
			Fecha:          Fecha,
			OC:             OC,
			SUC:            Suc,
			Familia:        Familia,
			CodigoProducto: CodigoProducto,
			Serie:          SerieDesktop,
			SerieOriginal:  SerieOriginal,
			Marca:          Marca,
			Modelo:         Modelo,
			Procesador: Procesador,
			Gen: Generacion,
			Mem_GB: MemGB,
			Velocidad: Velocidad,
			HDD: HddGB,
			HddSerie: HddSerie,
			UnidadOp: UnidadOp,
			Fuente: Fuente,
			Formato: Formato,
			Licencia: Licencia,
			Comentarios: Comentarios,
			Pulgadas: Pulgadas,
			Camara: Camara,
			Eliminador: Eliminador,
			SerieBateria: SerieBateria,
			Forma:          Forma,
			Base:           Base,
			Tipo:           Tipo,
			Salidas:        Salida,
			HDMI:           HDMI,
			Clase:          Clase,
			Tamanio:         Tamaño,
			Activo: Activo,
		})
	}
	return
}