package Controladores

import (

	bd "ControlInvetario/BD"
	grancompu "ControlInvetario/Utilidades"
	"encoding/json"
	"fmt"
	"net/http"
)

func Guardar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	Archivo, _, err := r.FormFile("file")
	if err != nil{
		fmt.Println(err)
	}
	//items:=grancompu.ReadXlsx(Archivo)
	//bd.Guardar(items)
	json.NewEncoder(w).Encode(bd.Guardar(grancompu.ReadXlsx(Archivo)))
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

func GetListaExel(w http.ResponseWriter, req *http.Request){
	req.ParseMultipartForm(32 << 20)
	Archivo, _, err := req.FormFile("file")
	if err != nil{
		fmt.Println(err)
	}
	grancompu.ReadXlsx1(Archivo)
	//Leer()
	json.NewEncoder(w).Encode(grancompu.Contador(grancompu.Lista.Data, grancompu.Minusculas("CLASS")))
}
func Pruebas(w http.ResponseWriter, r *http.Request) {
	ip:=r.Body
	json.NewEncoder(w).Encode(ip)
}
func Upload(w http.ResponseWriter, req *http.Request){
	req.ParseMultipartForm(32 << 20)
	Archivo, _, err := req.FormFile("file")
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(nombre)
	//fmt.Println(r)
	grancompu.ReadXlsx1(Archivo)
	json.NewEncoder(w).Encode("archivo recibido")
}