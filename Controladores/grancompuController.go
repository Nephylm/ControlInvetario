package Controladores

import (
	bd "ControlInvetario/BD"
	modelos "ControlInvetario/Modelos"
	utilidades "ControlInvetario/Utilidades"
	"encoding/json"
	"fmt"
	"net/http"
)

func Guardar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	Archivo, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	//items:=utilidades.ReadXlsx(Archivo)
	//bd.Guardar(items)
	json.NewEncoder(w).Encode(bd.Guardar(utilidades.ReadXlsx(Archivo)))
}

func ObteneInventario(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetInventario())
}
func ObteneMonitores(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetMonitores())
}
func ObteneDesktops(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetDesktop())
}
func ObteneLaptops(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bd.GetLaptop())
}
func ObteneAllione(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetAllInOne())
}

func GetListaExel(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	Archivo, _, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	utilidades.ReadXlsx1(Archivo)
	//Leer()
	json.NewEncoder(w).Encode(utilidades.Contador(utilidades.Lista.Data, utilidades.Minusculas("CLASS")))
}
func GuardarUno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var items []utilidades.Item
	var item utilidades.Item
	//des serializa el json a un map
	json.NewDecoder(r.Body).Decode(&item.Producto)
	fmt.Println(item)
	items = append(items, item)
	//guarda el producto
	json.NewEncoder(w).Encode(bd.Guardar(items))
}
func Pruebas(w http.ResponseWriter, r *http.Request) {
	utilidades.Peticion()
	utilidades.GetInsumos()
	json.NewEncoder(w).Encode("ALGO")
}
func Upload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	Archivo, _, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(nombre)
	//fmt.Println(r)
	utilidades.ReadXlsx1(Archivo)
	json.NewEncoder(w).Encode("archivo recibido")
}
func BajaDesktop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Desktop modelos.Desktop
	//des serializa el json a una estructura de Desktop
	json.NewDecoder(r.Body).Decode(&Desktop)
	fmt.Println(Desktop)
	resp := bd.BajaDesktop(Desktop)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func BajaLaptop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Laptop modelos.Laptop
	//des serializa el json a una estructura de Laptop
	json.NewDecoder(r.Body).Decode(&Laptop)
	fmt.Println(Laptop)
	resp := bd.BajaLaptop(Laptop)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func BajaMonitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Monitor modelos.Monitor
	//des serializa el json a una estructura de Monitor
	json.NewDecoder(r.Body).Decode(&Monitor)
	fmt.Println(Monitor)
	resp := bd.BajaMonitor(Monitor)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func ActualizaDesktop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Desktop modelos.Desktop
	//des serializa el json a una estructura de Desktop
	json.NewDecoder(r.Body).Decode(&Desktop)
	fmt.Println(Desktop)
	resp := bd.ActualizaDesktop(Desktop)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func ActualizaLaptop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Laptop modelos.Laptop
	//des serializa el json a una estructura de Laptop
	json.NewDecoder(r.Body).Decode(&Laptop)
	fmt.Println(Laptop)
	resp := bd.ActualizaLaptop(Laptop)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
func ActualizaMonitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var Monitor modelos.Monitor
	//des serializa el json a una estructura de Monitor
	json.NewDecoder(r.Body).Decode(&Monitor)
	fmt.Println(Monitor)
	resp := bd.ActualizaMonitor(Monitor)
	w.WriteHeader(resp.CodigoRespHTTP)
	json.NewEncoder(w).Encode(resp.Response)
}
