package Controladores

import (
	bd "ControlInvetario/BD"
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
