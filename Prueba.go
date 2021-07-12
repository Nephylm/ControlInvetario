package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
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
type Elemento struct {
	Clase string `json:"class"`
	Existencia string `json:"existencia"`
	Modelo string`json:"modelo"`
	IdProducto int`json:"idProducto"`
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
	ReadXlsx(Archivo)
}
func Almacenar(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)
	//ReadXlsx()
	json.NewEncoder(w).Encode(Contador(Lista.Data, Minusculas("CLASS")))
}
func pruebaAlmacenar() string{
	inventario := Contador(Lista.Data, Minusculas("CLASS"))
	var elme Elemento
	elme.IdProducto=1
	elme.Clase=inventario[0]["class"]
	elme.Existencia=inventario[0]["existencia"]
	elme.Modelo=inventario[0]["modelo"]
	stmt, es := db.Prepare("INSERT INTO ListaPrecio (Clase,Existencia,Modelo,IdProducto) VALUES (?, ?, ?,?);")
	if es != nil {
		panic(es.Error())
	}
	a, err := stmt.Exec(elme.Clase, elme.Existencia,elme.Modelo,elme.IdProducto)

	revisarError(err)
	affected, _ := a.RowsAffected()
	if affected > 0 {
		return"exito"
		fmt.Println("Guardado exitoso")
	}
	return "error"
}
func Guardar()  {
	inventario := Contador(Lista.Data, Minusculas("CLASS"))
	var c int

		for _,producto := range inventario{

				if producto["class"]=="Desktops"{
					c=4
				}else if producto["class"]=="Monitors"{
					c=1
				}else if producto["class"]=="bateria"{
					c=2
				}else if producto["class"]=="Adaptador"{
					c=3
				}else if producto["class"]=="disco duro"{
					c=5
				}

				a, err := db.Exec("INSERT INTO ListaPrecio (Clase,Existencia,Modelo,IdProducto) VALUES (?, ?, ?,?)"+
					" WHERE NOT EXISTS(SELECT * FROM ListaPrecio WHERE Codigo = ?);",
					producto["class"], producto["existencia"],producto["modelo"],c,producto["modelo"])

				revisarError(err)
				affected, _ := a.RowsAffected()
				if affected > 0 {
					fmt.Println("Guardado exitoso")
				}

		}

}
func GetListaExel(w http.ResponseWriter, req *http.Request){

	//enableCors(&w)
	//ReadXlsx()
	//Guardar()
	//pruebaAlmacenar()
	//Leer()
	json.NewEncoder(w).Encode(Contador(Lista.Data, Minusculas("CLASS")))
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
//func ReadXlsx(){
func ReadXlsx(Archivo multipart.File){
	var item Item
	var items []Item
	items=nil
	//f, err := excelize.OpenFile("viaje 1305 26MAYO2021.xlsx")
	//f, err := excelize.OpenFile("VIAJE 1306   9JUN2021.xlsx")
	f, err := excelize.OpenReader(Archivo)
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

				item.Producto[Minusculas(campos[j])] = colCell

			}else{
				item.Id=n

				if len(item.Producto)>0 {
					items = append(items, item)
					item.Producto = clear()
					n++
				}
			}
		}
		if item.Producto[Minusculas(campos[0])] != "" {
			item.Id=n
			items = append(items, item)
			item.Producto = make(map[string]string)
			n++
		}
	}

	//Contador(items, Minusculas("CLASS"))
	Lista.Data=items
	fmt.Print(Lista)
}
func Contador(productos [] Item, Clasificacion string) []map[string]string{
	var inventario []map[string]string

	inventary := Clasificador(productos,Clasificacion)
	for _,class:=range inventary{
		var item Item
		var items []Item
		for i,objeto:= range class{
			item.Id=i
			item.Producto=objeto
			items=append(items,item)
		}
		productoPorModelo :=Clasificador(items,"modelo")
		for _,product:= range productoPorModelo{
			producto:=product[0]
			producto["existencia"]=strconv.Itoa(len(product))
			inventario=append(inventario,producto)
		}

	}
	print(inventario)
	return inventario
}
func 	Clasificador(productos [] Item, Clasificacion string) [][]map[string]string{
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
	estandar:= Minusculas(Clasificacion)
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
func Minusculas(palabra string) string{
	return strings.ToLower(palabra)
}


