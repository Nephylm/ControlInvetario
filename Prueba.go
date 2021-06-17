package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

type Datos struct {
	Data []Membresia `json:"data"`
}
type Productos struct {

	Data []Item `json:"data"`
}

type Item struct{
	Id int `json:"id"`
	Producto map[string]string `json:"producto"`
}
var Lista Productos

type Membresia struct {
	Id_membresia string `json:"id_membresia"`
	Tipo_membresia string `json:"tipo_membresia,omitempty"`
}
type Monitores struct {
	Data []Monitor `json:"data"`
}
type Monitor struct {
	Warehouse string `json:"warehouse"`
	Class string `json:"class"`
	AssetSSE int`json:"assetSSE"`
	Serial string`json:"serial"`
	Manuf string`json:"manuf"`
	Model string`json:"model"`
	ModelN string`json:"modelN"`
	Cond string`json:"cond"`
	ScreenSize int`json:"screenSize"`
	NotesComments string`json:"notesComments"`
	NoSKU string`json:"noSKU"`
	Price int`json:"price"`
}

// EndPoints

/*func Upload(w http.ResponseWriter, req *http.Request){
	req.ParseMultipartForm(32 << 20)
	Archivo, _, _ := req.FormFile("file")
	//fmt.Println(nombre)
	//fmt.Println(r)
	if (*req).Method == "OPTIONS" {
		return
	}
	if (*req).Method=="POST"{
		//var monitor Monitor
		ReadXlsx(Archivo)
		//ReadCSV(Archivo)

	}
}*/
func GetListaExel(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)
	ReadXlsx()
	//Leer()
	json.NewEncoder(w).Encode(contador(Lista.Data,minusculas("CLASS")))
}
func Prueba(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)

	json.NewEncoder(w).Encode("mensaje de prueba")
}
func ReadCSV(Archivo multipart.File){
	var item Item
	var items []Item
	r:=csv.NewReader(Archivo)
	r.Comma=','
	r.FieldsPerRecord = -1
	if err != nil{
		log.Println(err)
	}
	campos,_:=r.Read()

	fmt.Println(campos)
	for {
		mon, err := r.Read()
		if err == io.EOF {
			break
		}
		item.Producto = make(map[string]string)
		for i, c := range mon {
			item.Producto[campos[i]] = c
		}
		items = append(items, item)

	}
	fmt.Println(items)
}
func ReadXlsx(){
//func ReadXlsx(Archivo multipart.File){
	var item Item
	var items []Item
	items=nil
	f, err := excelize.OpenFile("viaje 1305 26MAYO2021.xlsx")
	//f, err := excelize.OpenFile("VIAJE 1306   9JUN2021.xlsx")
	//f, err := excelize.OpenReader(Archivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	hojas := f.GetSheetList()

	rows, err := f.GetRows(hojas[0])

	var campos []string
	n:=1
	item.Producto = make(map[string]string)
	for i, row := range rows {
		if row==nil{
			break
		}
		for j, colCell := range row {
			if i==0{
				campos=row
				//fmt.Println(campos)
				break
			}

			if campos[j] != ""{

				item.Producto[minusculas(campos[j])] = colCell

			}else{
				item.Id=n

				if len(item.Producto)>0 {
					items = append(items, item)
					item.Producto = clear()
					n++
				}
			}
		}
		if item.Producto[minusculas(campos[0])] != "" {
			item.Id=n
			items = append(items, item)
			item.Producto = make(map[string]string)
			n++
		}
	}

	contador(items,minusculas("CLASS"))
	Lista.Data=items
	fmt.Print(Lista)
}
func contador(productos [] Item, Clasificacion string) [][]map[string]string{
	var inventary [][]map[string]string
	clasificador :=Clases(productos,Clasificacion)
	for _,class:=range clasificador {

		var i = Agrupacion(productos, Clasificacion, class)
		inventary=append(inventary,i)
		fmt.Println(class)
		fmt.Println(len(i))

	}
	return inventary
}
//regresa las distintas clases de productos en un archivo
func Clases (productos [] Item, Clasificacion string) []string{
	estandar:=minusculas(Clasificacion)
	var clases []string
	for _, producto := range productos{
		if len(clases)>0{
			existe:=false
			for _, class := range clases{
				if class==producto.Producto[estandar] {
					existe=true
					break
				}
			}
			if !existe{
				clases=append(clases,producto.Producto[estandar])
			}
		}else{
			clases=append(clases,producto.Producto[estandar])
		}
	}
	fmt.Println(clases)
	return clases
}

//clasifica los productos en grupos
func Agrupacion (productos [] Item, Clasificacion string,Clasificador string) []map[string]string{

	var lista []map[string]string
	for _, producto := range productos{
		if producto.Producto[Clasificacion]==Clasificador{
			lista = append(lista,producto.Producto)
		}
	}
	return lista
}

//retorna un map de strings en blanco
func clear() map[string]string{
	var blanco=make(map[string]string)
	return blanco
}
//CONVIERTE TODAS LAS LETRAS DE UNA PALABRA A MINUSCULAS
func minusculas(palabra string) string{
	return strings.ToLower(palabra)
}
func Iniciar() {

	router := mux.NewRouter()

	//router.HandleFunc("/upload", Upload).Methods("POST")
	router.HandleFunc("/lista", GetListaExel).Methods("GET")
	router.HandleFunc("/prueba", Prueba).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS","DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
