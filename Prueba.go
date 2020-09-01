package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Datos struct {
	Data []Membresia `json:"data"`
}
type Productos struct {
	Data []Item `json:"data"`
}

type Item struct{
	Producto map[string]string `json:"producto"`
}
var Lista Productos
var item Item
var items []Item
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

var membresias []Membresia
var membresia Membresia
var data Datos
var pantallas Monitores
var monitores [] Monitor
// EndPoints
func GetMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
    for _,item:= range membresias{
    	if item.Id_membresia==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//RecuperarXId(tabla,1)
	json.NewEncoder(w).Encode(membresia)
}

func GetMembershipsEndpoint(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)
	data.Data=membresias
	json.NewEncoder(w).Encode(data)
}
func CreateMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)
	if (*req).Method == "OPTIONS" {
		return
	}
	if (*req).Method=="POST"{
		var memberish Membresia
		_ = json.NewDecoder(req.Body).Decode(&memberish)
		abrirConexionDB()
		agregarDatosBD(memberish)
		resultadosQuery(tabla)
		terminarConexion()
		data.Data=membresias
		json.NewEncoder(w).Encode(data)
	}
}
func UpdateMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)
		var memberish Membresia
		params := mux.Vars(req)
		_ = json.NewDecoder(req.Body).Decode(&memberish)
		abrirConexionDB()
		actualizarDatosBD(memberish,params["id"])
		terminarConexion()
		data.Data=membresias
		json.NewEncoder(w).Encode(data)
}

func DeleteMembershipEndpoint(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	params := mux.Vars(req)
	abrirConexionDB()
	eliminarDatosBD(strconv.ParseInt(params["id"],10,64) )
	terminarConexion()
	data.Data=membresias
	json.NewEncoder(w).Encode(data)
}
func Upload(w http.ResponseWriter, req *http.Request){
	if (*req).Method == "OPTIONS" {
		return
	}
	if (*req).Method=="POST"{
		//var monitor Monitor
		r:=csv.NewReader(req.Body)
		r.Comma=','
		r.FieldsPerRecord = -1
		if err != nil{
			log.Println(err)
		}
		campos,_:=r.Read()

		fmt.Println(campos)
		for{
			i:=0
			mon, err:=r.Read()
			if err == io.EOF{
				break
			}
			item.Producto=make(map[string]string)
			for _,c :=range mon{

				item.Producto[campos[i]]=c
				i++

			}
			items=append(items,item)

			/*
			var m Monitor
			m.Warehouse=mon[0]
			m.Class=mon[1]
			m.AssetSSE, _ =strconv.Atoi(mon[2])
			m.Serial =mon[3]
			m.Manuf=mon[4]
			m.Model =mon[5]
			m.ModelN =mon[6]
			m.Cond =mon[7]
			m.ScreenSize,_ =strconv.Atoi(mon[8])
			m.NotesComments =mon[9]
			m.NoSKU =mon[10]
			m.Price,_=strconv.Atoi(mon[11])
			monitores = append(monitores,m)*/

		}
		//pantallas.Data=monitores
		Lista.Data=items
		json.NewEncoder(w).Encode(Lista)
		fmt.Println(Lista)
		//fmt.Println(monitores)
	}
}
func PruebaMap(w http.ResponseWriter, req *http.Request){
	type User struct {
		Name string `json:"name"`
		Age  int16  `json:"age"`
	}
	u := User{Name: "Alexys", Age: 37}
	amount := 1500
	json.NewEncoder(w).Encode(struct{
		User `json:"user"`
		Amount int `json:"amount"`
	}{
		u,
		amount,
	})

}
func ReadExel(){
//exel := excelize.NewFile ()
	/*f, err := excelize.OpenFile("Horarios.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell, err := f.GetCellValue("Hoja1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(cell)
	rows, err := f.GetRows("Hoja1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t\t")
		}
		fmt.Println()
	}*/
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func Iniciar() {

	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/memberships", GetMembershipsEndpoint).Methods("GET")
	router.HandleFunc("/memberships/{id}", GetMembershipEndpoint).Methods("GET")
	router.HandleFunc("/memberships/agregar", CreateMembershipEndpoint).Methods("POST")
	router.HandleFunc("/upload", Upload).Methods("POST")
	router.HandleFunc("/memberships/actualizar/{id}",UpdateMembershipEndpoint).Methods("POST")
	router.HandleFunc("/memberships/eliminar/{id}", DeleteMembershipEndpoint).Methods("DELETE")
	router.HandleFunc("/encode", PruebaMap).Methods("GET")

	//
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS","DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

