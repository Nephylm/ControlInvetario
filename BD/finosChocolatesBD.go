package bd

import (
	modelos "ControlInvetario/Modelos"
	utilidades "ControlInvetario/Utilidades"
	"database/sql"
	"fmt"
	"strconv"
)

var(
	IdRegistro int
	IdInsumo int
	Nombre string
	NombreProducto string
	CantidadMinima float32
	CantidadMaxima float64
	Activo int
	UnidadMedida string
	OrdenProduccion string
	FechaProd string
	FechaCad string
	Id_PTerminado int
)
//INSUMOS
//Recupera los insumos de las matrices y los registra
func CargarInsumos (insumo modelos.Insumo)(resp modelos.RespuestaSencilla){

	stmt, es := db.Prepare("INSERT INTO Insumo (IdInsumo, Nombre,Existencia,CantidadMinima,CantidadMaxima,Activo,UnidadMedida)" +
		" SELECT ?, ?,?,?,?,?,? WHERE NOT EXISTS (SELECT *FROM Insumo WHERE Nombre=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(insumo.IdInsumo,insumo.Nombre,0,0,0,1,insumo.UnidadMedida.Abrev, insumo.Nombre)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
	switch err {
	case nil:
		resp.CodigoRespHTTP=200
		resp.Response="Guardado exitoso"
	case sql.ErrNoRows:
		resp.CodigoRespHTTP=404
		resp.Response=err.Error()
	default:
		resp.Response="error al registrar"
	}
	return
}
//modifica los datos de los insumos
func ActualizarInsumo (insumo modelos.InventarioInsumos)(resp modelos.RespuestaSencilla){

	stmt, es := db.Prepare("UPDATE Insumo SET Existencia=?, CantidadMinima=?,CantidadMaxima=?,Activo=? WHERE Nombre=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(insumo.Existencia,insumo.CantidadMinima,insumo.CantidadMaxima,insumo.Activo,insumo.Nombre)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
			resp.CodigoRespHTTP=200
			resp.Response="Actualizacion exitosa"
			return
	}
	resp.CodigoRespHTTP=400
	resp.Response="Error al Actualizar"
	return
}
//regresa una lista con todos los insumos registrados
func GetInventarioInsumos() (Data []modelos.InventarioInsumos, resp modelos.RespuestaSencilla) {
	listado, _ := db.Query("SELECT IdInsumo,Nombre,Existencia,CantidadMinima,CantidadMaxima,Activo,UnidadMedida FROM Insumo;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdInsumo,
			&Nombre,
			&Existencia,
			&CantidadMinima,
			&CantidadMaxima,
			&Activo,
			&UnidadMedida,
		)
		revisarError(err)

		Data = append(Data,
			modelos.InventarioInsumos{
			IdInsumo: IdInsumo,
			Nombre: Nombre,
			Existencia: Existencia,
			CantidadMinima: CantidadMinima,
			CantidadMaxima: CantidadMaxima,
			Activo: Activo,
			UnidadDMedida: UnidadMedida,
		})
	}
	switch err {
	case nil:
		resp.CodigoRespHTTP=200
	case sql.ErrNoRows:
		resp.CodigoRespHTTP=404
		resp.Response=err.Error()
	default:
	}
	return
}
func GetInsumo(nombre string) (Data modelos.InventarioInsumos,resp modelos.RespuestaSencilla) {
	listado, _ := db.Query("SELECT IdInsumo,Nombre,Existencia,CantidadMinima,CantidadMaxima,Activo FROM Insumo WHERE Nombre=?;",nombre)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdInsumo,
			&Nombre,
			&Existencia,
			&CantidadMinima,
			&CantidadMaxima,
			&Activo,
		)
		revisarError(err)

		Data = modelos.InventarioInsumos{
				IdInsumo: IdInsumo,
				Nombre: Nombre,
				Existencia: Existencia,
				CantidadMinima: CantidadMinima,
				CantidadMaxima: CantidadMaxima,
				Activo: Activo,
			}

	}
	switch err {
	case nil:
		resp.CodigoRespHTTP=200
	case sql.ErrNoRows:
		resp.CodigoRespHTTP=404
		resp.Response=err.Error()
	default:
		fmt.Println("Error en la consulta")
	}
	return
}
//metodo para registar los insumos, utilia CargarInsumos()
func RegistarInsumo(Insumos []modelos.Insumo)(resp modelos.RespuestaSencilla)  {
	for _,insumo:=range Insumos{
		CargarInsumos(insumo)
	}
	resp.CodigoRespHTTP=200
	resp.Response="Guardado exitoso"
	return
}

//PRODUCTO TERMINADO
//Registra un producto terminado
func AgregarPTerminado (productoT modelos.ProductoTerminado)(resp modelos.RespuestaSencilla){
	stmt, es := db.Prepare("INSERT INTO Chocolates (IdProducto,IdRegistro) VALUES (?,?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoT.Producto,productoT.OrdenProduccion)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP=200
		resp.Response="Guardado exitoso"
		fmt.Println("Guardado exitoso")

	} else {
		resp.CodigoRespHTTP=400
		resp.Response="Error al registrar"
		fmt.Println("error al registrar")
	}
	return
}
func EliminarPTerminado (productoT modelos.ProductoTerminado)(resp modelos.RespuestaSencilla){
	stmt, es := db.Prepare("DELETE Chocolates WHERE IdPTerminado=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoT.Id_PTerminado)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP=200
		resp.Response="Eliminado con exito"
		return
	}
	resp.CodigoRespHTTP=400
	resp.Response="Error al eliminar"
	return
}
func GetPTerminado (idRegistro int)(Data []modelos.ProductoTerminado){
	listado, _ := db.Query("SELECT Id_PTerminado, IdRegistro,IdProduto FROM ProductosOrden WHERE OrdenProduccion=? AND IdRegistro=?;",idRegistro)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&Id_PTerminado,
			&IdRegistro,
			&IdProducto,
		)
		revisarError(err)
		Data = append(Data,modelos.ProductoTerminado{
			Id_PTerminado: Id_PTerminado,
			Producto: IdProducto,
			OrdenProduccion: idRegistro,
		})
	}
	return
}
//PRODUCTOS X ORDEN
func AgregarProductosOrden (productoOrden modelos.ProductosOrden)(resp modelos.RespuestaSencilla){
	var productoT modelos.ProductoTerminado

	stmt, es := db.Prepare("INSERT INTO ProductosOrden (OrdenProduccion,IdProducto,Existencia,FechaProd,FechaCad)" +
		" SELECT ?,?,?,?,?  WHERE NOT EXISTS (SELECT *FROM ProductosOrden WHERE OrdenProduccion=? AND IdProducto=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoOrden.OrdenProduccion,productoOrden.IdProducto,productoOrden.Existencia,productoOrden.FechaProd,productoOrden.FechaCad,
		productoOrden.OrdenProduccion,productoOrden.IdProducto)
	Id:=productoOrden.IdProducto
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
		switch err {
		case nil:
			resp.CodigoRespHTTP=200
			resp.Response="Guardado existoso"
		case sql.ErrNoRows:
			resp.CodigoRespHTTP=404
			resp.Response=err.Error()
		default:

		}
		productoT=modelos.ProductoTerminado{
			Producto: Id,
			OrdenProduccion: GetIdProductosOrden(productoOrden),
		}
		i:=0
		for i < productoOrden.Existencia{
			AgregarPTerminado(productoT)
			i++
		}
		ReducirInsumo(strconv.Itoa(Id))
		detalleP:=modelos.Detalles{
			IdProducto: productoOrden.IdProducto,
			Inventario: ContadorProductos(Id),
		}
		AgregraDetalleP(detalleP)
		return
	}
		resp.CodigoRespHTTP=400
		resp.Response="Error al registrar, ya estaba registrado"
		fmt.Println("error al registrar ya registrado")
	return

}
func GetIdProductosOrden(productoOrden modelos.ProductosOrden)(Data int){
	listado, _ := db.Query("SELECT IdRegistro FROM ProductosOrden WHERE OrdenProduccion=? AND IdProducto=?;",productoOrden.OrdenProduccion,productoOrden.IdProducto)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdRegistro,
		)
		revisarError(err)
		Data = IdRegistro
	}
	return
}
func GetProductosOrden()(Data []modelos.ProductosOrden,resp modelos.RespuestaSencilla){
	listado, _ := db.Query("SELECT PO.IdRegistro,PO.OrdenProduccion,PO.IdProducto,Producto.NombreProducto,PO.Existencia, PO.FechaProd," +
		" PO.FechaCad FROM ProductosOrden AS PO INNER JOIN Producto ON PO.IdProducto=Producto.IdProducto;")
	revisarError(err)

	for listado.Next() {
		err = listado.Scan(
			&IdRegistro,
			&OrdenProduccion,
			&IdProducto,
			&NombreProducto,
			&Existencia,
			&FechaProd,
			&FechaCad,
		)
		revisarError(err)
		Data = append(Data,modelos.ProductosOrden{
			IdRegistro: IdRegistro,
			OrdenProduccion: OrdenProduccion,
			IdProducto: IdProducto,
			NombreProducto: NombreProducto,
			Existencia: Existencia,
			FechaProd: FechaProd,
			FechaCad: FechaCad,
		})
	}
	switch err {
	case nil:
		resp.CodigoRespHTTP=200
	case sql.ErrNoRows:
		resp.CodigoRespHTTP=404
		resp.Response="Error al cargar datos"
	default:
		resp.CodigoRespHTTP=400
		resp.Response="Error en la consulta"
		fmt.Println("error en la consulta")
	}
	return
}
func ActualizarProductosOrden (productoOrden modelos.ProductosOrden)(resp modelos.RespuestaSencilla){
	anterior :=GetProductoOrden(productoOrden.IdRegistro)
	stmt, es := db.Prepare("UPDATE ProductosOrden SET Existencia=?, FechaProd=?,FechaCad=? WHERE IdRegistro=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoOrden.Existencia,productoOrden.FechaCad,productoOrden.FechaCad,productoOrden.IdRegistro)
	if anterior.Existencia < productoOrden.Existencia{
		productoT:=modelos.ProductoTerminado{
			Producto: productoOrden.IdProducto,
			OrdenProduccion: GetIdProductosOrden(productoOrden),
		}
		i:=anterior.Existencia
		for i<productoOrden.Existencia{
			AgregarPTerminado(productoT)
			ReducirInsumo(strconv.Itoa(productoOrden.IdProducto))
			i++
		}
	}else if anterior.Existencia > productoOrden.Existencia{
		i:=productoOrden.Existencia
		list:=GetPTerminado(productoOrden.IdRegistro)
		for _,eliminar:=range list{
			EliminarPTerminado(eliminar)
			i++
			if anterior.Existencia==productoOrden.Existencia{
				break
			}
		}
	}
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP=200
		resp.Response="Actualizacion exitosa"
		fmt.Println("Actualizacion exitosa exitoso")
		return
	}
	resp.CodigoRespHTTP=400
	resp.Response="Error al actualizar, revise sus datos"
	return
}
func GetProductoOrden(id int)(Data modelos.ProductosOrden){
	listado, _ := db.Query("SELECT IdRegistro,OrdenProduccion,IdProducto,Existencia, FechaProd," +
		" FechaCad FROM ProductosOrden WHERE IdRegistro=?;",id)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdRegistro,
			&OrdenProduccion,
			&IdProducto,
			&Existencia,
			&FechaProd,
			&FechaCad,
		)
		revisarError(err)
		Data = modelos.ProductosOrden{
			IdRegistro: IdRegistro,
			OrdenProduccion: OrdenProduccion,
			IdProducto: IdProducto,
			Existencia: Existencia,
			FechaProd: FechaProd,
			FechaCad: FechaCad,
		}
	}
	return
}
func GetProductos()(Data []modelos.ProductosOrden){
	listado, _ := db.Query("SELECT PO.IdRegistro,PO.OrdenProduccion,PO.IdProducto,Producto.NombreProducto,PO.Existencia, PO.FechaProd," +
		" PO.FechaCad FROM ProductosOrden AS PO INNER JOIN Producto ON PO.IdProducto=Producto.IdProducto;")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdRegistro,
			&OrdenProduccion,
			&IdProducto,
			&NombreProducto,
			&Existencia,
			&FechaProd,
			&FechaCad,
		)
		revisarError(err)
		Data = append(Data,modelos.ProductosOrden{
			IdRegistro: IdRegistro,
			OrdenProduccion: OrdenProduccion,
			IdProducto: IdProducto,
			NombreProducto: NombreProducto,
			Existencia: Existencia,
			FechaProd: FechaProd,
			FechaCad: FechaCad,
		})
	}
	return
}
func GetProductosOrdenProduccion(productoOrden modelos.ProductosOrden)(Data []modelos.ProductosOrden){
	listado, _ := db.Query("SELECT IdRegistro,OrdenProduccion,IdProducto,Existencia, FechaProd," +
		" FechaCad FROM ProductosOrden WHERE OrdenProduccion=?;",productoOrden.OrdenProduccion)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdRegistro,
			&OrdenProduccion,
			&IdProducto,
			&Existencia,
			&FechaProd,
			&FechaCad,
		)
		revisarError(err)
		Data = append(Data,modelos.ProductosOrden{
			IdRegistro: IdRegistro,
			OrdenProduccion: OrdenProduccion,
			IdProducto: IdProducto,
			Existencia: Existencia,
			FechaProd: FechaProd,
			FechaCad: FechaCad,
		})
	}
	return
}
func ContadorProductos(idProducto int)(existencia int){
	listado, _ := db.Query("SELECT SUM(Existencia)  FROM ProductosOrden WHERE IdProducto=?;",idProducto)
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&existencia,
		)
		revisarError(err)
	}
	return
}

//Utilidades
//Agrega IdProducto
func CargarIdProducto (idProducto modelos.IdProducto){
	stmt, es := db.Prepare("INSERT INTO Producto (IdProducto,NombreProducto,Activo) SELECT ?,?,? " +
		"WHERE NOT EXISTS (SELECT *FROM Producto WHERE IdProducto=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(idProducto.IdProducto, idProducto.NombreProducto,idProducto.Activo,idProducto.IdProducto)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("producto ya registrado")
	}
}
func RegistarIdProducto(idProductos []modelos.IdProducto)( resp modelos.RespuestaSencilla)  {
	for _,idProducto:=range idProductos{
		CargarIdProducto(idProducto)
	}
	resp.CodigoRespHTTP=200
	resp.Response="Guardado exitoso"
	return
}
func GetIdProductos()(Data []modelos.IdProducto,resp modelos.RespuestaSencilla)	{
	listado, _ := db.Query("SELECT IdProducto,NombreProducto,Activo FROM Producto")
	revisarError(err)
	for listado.Next() {
		err = listado.Scan(
			&IdProducto,
			&NombreProducto,
			&Activo,
		)
		revisarError(err)
		Data = append(Data,
			modelos.IdProducto{
				IdProducto: IdProducto,
				NombreProducto: NombreProducto,
				Activo: Activo,
			})

	}
	switch err {
	case nil:
		resp.CodigoRespHTTP=200
	case sql.ErrNoRows:
		resp.CodigoRespHTTP=404
		resp.Response="Error al consultar datos"
		return
	default:
		resp.CodigoRespHTTP=400
		resp.Response="Error en la consulta"
		fmt.Println("error en la consulta")
	}
	return
}
func ReducirInsumo(IdProducto string){
	insumosxProducto := utilidades.GetInsumosxProducto(IdProducto)
	for _, insumoXProducto :=range insumosxProducto{
		insumo,_:=GetInsumo(insumoXProducto.Insumo.Nombre)
		cantidad:=insumo.Existencia
		usado:=insumoXProducto.CantidadUsada
		insumo.Existencia=cantidad-usado
		ActualizarInsumo(insumo)
	}
}
func AgregraDetalleP (DetalleP modelos.Detalles)(resp modelos.RespuestaSencilla){

	stmt, es := db.Prepare("INSERT INTO DetalleProducto (IdProducto,Inventario)" +
		" SELECT ?,?  WHERE NOT EXISTS (SELECT *FROM DetalleProducto WHERE IdProducto=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(DetalleP.IdProducto,DetalleP.Inventario,DetalleP.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		resp.CodigoRespHTTP=200
		resp.Response="Guardado exitoso"
		fmt.Println("Guardado exitoso")

		return
	}else {
		stmt, es := db.Prepare("UPDATE DetalleProducto SET Inventario=? WHERE IdProducto=?;")
		if es != nil {
			panic(es.Error())
		}
		_, err := stmt.Exec(DetalleP.Inventario,DetalleP.IdProducto)
		revisarError(err)
		if affected > 0 {
			resp.CodigoRespHTTP=200
			resp.Response="Actualizacion exitosa"
			fmt.Println("Actualizacion exitoso")
			return
		}
	}
	resp.CodigoRespHTTP=400
	resp.Response="Error al registrar"
	return
}
