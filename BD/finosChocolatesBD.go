package bd

import (
	modelos "ControlInvetario/Modelos"
	utilidades "ControlInvetario/Utilidades"
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
func CargarInsumos (insumo modelos.Insumo){

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
}
//modifica los datos de los insumos
func ActualizarInsumo (insumo modelos.InventarioInsumos)string{

	stmt, es := db.Prepare("UPDATE Insumo SET Existencia=?, CantidadMinima=?,CantidadMaxima=?,Activo=? WHERE Nombre=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(insumo.Existencia,insumo.CantidadMinima,insumo.CantidadMaxima,insumo.Activo,insumo.Nombre)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Actualizacion exitosa exitoso")
		return "exito"
	}
	return "error"
}
//regresa una lista con todos los insumos registrados
func GetInventarioInsumos() (Data []modelos.InventarioInsumos) {
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
	return
}
func GetInsumo(nombre string) (Data modelos.InventarioInsumos) {
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
	return
}
//metodo para registar los insumos, utilia CargarInsumos()
func RegistarInsumo(Insumos []modelos.Insumo)  {
	for _,insumo:=range Insumos{
		CargarInsumos(insumo)
	}
}

//PRODUCTO TERMINADO
//Registra un producto terminado
func AgregarPTerminado (productoT modelos.ProductoTerminado){
	stmt, es := db.Prepare("INSERT INTO Chocolates (IdProducto,IdRegistro) VALUES (?,?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoT.Producto,productoT.OrdenProduccion)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
	} else {
		fmt.Println("error al registrar")
	}
}
func EliminarPTerminado (productoT modelos.ProductoTerminado)string{
	stmt, es := db.Prepare("DELETE Chocolates WHERE IdPTerminado=?;")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoT.Id_PTerminado)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Eliminacion exitosa exitoso")
		return "exito"
	}
	return "error"
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
func AgregarProductosOrden (productoOrden modelos.ProductosOrden)string{
	var productoT modelos.ProductoTerminado

	stmt, es := db.Prepare("INSERT INTO ProductosOrden (OrdenProduccion,IdProducto,Existencia,FechaProd,FechaCad)" +
		" SELECT ?,?,?,?,?  WHERE NOT EXISTS (SELECT *FROM ProductosOrden WHERE OrdenProduccion=? AND IdProducto=?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(productoOrden.OrdenProduccion,productoOrden.IdProducto,productoOrden.Existencia,productoOrden.FechaProd,productoOrden.FechaCad,
		productoOrden.OrdenProduccion,productoOrden.IdProducto)
	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		fmt.Println("Guardado exitoso")
		productoT=modelos.ProductoTerminado{
			Producto: productoOrden.IdProducto,
			OrdenProduccion: GetIdProductosOrden(productoOrden),
		}
		i:=0
		for i < productoOrden.Existencia{
			AgregarPTerminado(productoT)
			i++
		}
		ReducirInsumo(strconv.Itoa(productoOrden.IdProducto))
		return  "guardado exitoso"
	}
		fmt.Println("error al registrar ya registrado")
	return "error al registrar ya registrado"

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
func GetProductosOrden()(Data []modelos.ProductosOrden){
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
func ActualizarProductosOrden (productoOrden modelos.ProductosOrden)string{
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
		fmt.Println("Actualizacion exitosa exitoso")
		return "exito"
	}
	return "error"
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
func RegistarIdProducto(idProductos []modelos.IdProducto)  {
	for _,idProducto:=range idProductos{
		CargarIdProducto(idProducto)
	}
}
func GetIdProductos()(Data []modelos.IdProducto)	{
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
	return
}
func ReducirInsumo(IdProducto string){
	insumosxProducto := utilidades.GetInsumosxProducto(IdProducto)
	for _, insumoXProducto :=range insumosxProducto{
		insumo:=GetInsumo(insumoXProducto.Insumo.Nombre)
		cantidad:=insumo.Existencia
		usado:=insumoXProducto.CantidadUsada
		insumo.Existencia=cantidad-usado
		ActualizarInsumo(insumo)
	}
}