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

func Upload(w http.ResponseWriter, req *http.Request){
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
}
func GetListaExel(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)

	json.NewEncoder(w).Encode(Lista)
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
func ReadXlsx(Archivo multipart.File){
	var item Item
	var items []Item
	items=nil
	//f, err := excelize.OpenFile("PO 657.xlsx")
	f, err := excelize.OpenReader(Archivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	hojas := f.GetSheetList()

	rows, err := f.GetRows(hojas[1])

	var campos []string
	n:=1
	for i, row := range rows {

		item.Producto = make(map[string]string)
		for j, colCell := range row {
			if i==0{
				campos=row
				break
			}
			if campos[j] != ""{

				item.Producto[campos[j]] = colCell

			}
		}
		if item.Producto["Serial#"] != "" {
			item.Id=n
			items = append(items, item)
			n++
		}
	}
	fmt.Print(items)
	Lista.Data=items
}

func Iniciar() {

	router := mux.NewRouter()

	router.HandleFunc("/upload", Upload).Methods("POST")
	router.HandleFunc("/lista", GetListaExel).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS","DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

