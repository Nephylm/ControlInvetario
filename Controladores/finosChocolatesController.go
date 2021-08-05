package Controladores

import (
	bd "ControlInvetario/BD"
	modelos "ControlInvetario/Modelos"
	utilidades "ControlInvetario/Utilidades"
	"encoding/json"
	"net/http"
)
//INUSMO
func CargaDInsumos(w http.ResponseWriter, r *http.Request) {
	resp:=bd.RegistarInsumo(utilidades.GetInsumos())
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func InventarioInsumos(w http.ResponseWriter, r *http.Request) {
	a,b:=bd.GetInventarioInsumos()
	w.WriteHeader(b.CodigoRespHTTP)
	json.NewEncoder(w).Encode(a)
}
func ActualizarInsumo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insumo modelos.InventarioInsumos
	json.NewDecoder(r.Body).Decode(&insumo)
	resp:=bd.ActualizarInsumo(insumo)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func AgregarDetalles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Detalles modelos.Detalles
	json.NewDecoder(r.Body).Decode(&Detalles)
	resp :=bd.AgregraDetalleP(Detalles)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
//UTILIDADES
func CargaIdProdutos(w http.ResponseWriter, r *http.Request) {
	resp:=bd.RegistarIdProducto(utilidades.GetProductos())
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func ObtenerIdProductos(w http.ResponseWriter, r *http.Request) {
	IdProductos,resp:=bd.GetIdProductos()
	w.WriteHeader(resp.CodigoRespHTTP)
	if resp.CodigoRespHTTP==200{
		json.NewEncoder(w).Encode(IdProductos)
	}else {
		json.NewEncoder(w).Encode(resp.Response)
	}

}
//PRODUCTO X ORDEN

func GuarProductoOrden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var productoOrden modelos.ProductosOrden
	json.NewDecoder(r.Body).Decode(&productoOrden)
	resp :=bd.AgregarProductosOrden(productoOrden)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func ObtenerProductoOrden(w http.ResponseWriter, r *http.Request) {
	ProductosOrden,resp:=bd.GetProductosOrden()
	w.WriteHeader(resp.CodigoRespHTTP)
	if resp.CodigoRespHTTP==200{
		json.NewEncoder(w).Encode(ProductosOrden)
	}else {
		json.NewEncoder(w).Encode(resp.Response)
	}
}
func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(bd.GetProductosOrden())
}
func ActualizarProductoOrden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var productoOrden modelos.ProductosOrden
	json.NewDecoder(r.Body).Decode(&productoOrden)
	resp:=bd.ActualizarProductosOrden(productoOrden)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}