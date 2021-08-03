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
	bd.RegistarInsumo(utilidades.GetInsumos())
	json.NewEncoder(w).Encode("ALGO")
}
func InventarioInsumos(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetInventarioInsumos())
}
func ActualizarInsumo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insumo modelos.InventarioInsumos
	json.NewDecoder(r.Body).Decode(&insumo)
	json.NewEncoder(w).Encode(bd.ActualizarInsumo(insumo))
}
//UTILIDADES

func CargaIdProdutos(w http.ResponseWriter, r *http.Request) {
	bd.RegistarIdProducto(utilidades.GetProductos())
	json.NewEncoder(w).Encode("ALGO")
}
func ObtenerIdProductos(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetIdProductos())
}
//PRODUCTO X ORDEN

func GuarProductoOrden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var productoOrden modelos.ProductosOrden
	json.NewDecoder(r.Body).Decode(&productoOrden)
	json.NewEncoder(w).Encode(bd.AgregarProductosOrden(productoOrden))
}
func ObtenerProductoOrden(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bd.GetProductosOrden())
}
func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bd.GetProductosOrden())
}
func ActualizarProductoOrden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var productoOrden modelos.ProductosOrden
	json.NewDecoder(r.Body).Decode(&productoOrden)
	json.NewEncoder(w).Encode(bd.ActualizarProductosOrden(productoOrden))
}