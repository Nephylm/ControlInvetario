package Utilidades

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type Elemento struct {
	Clase      string `json:"class"`
	Existencia int    `json:"existencia"`
	Modelo     string `json:"modelo"`
	IdProducto int    `json:"idProducto"`
}
type Datos struct {
	Data []Membresia `json:"data"`
}
type Productos struct {
	Data []Item `json:"data"`
}
type Mercancia struct {
	Data []Item `json:"data"`
}
type Item struct {
	Contador int               `json:"contador"`
	Producto map[string]string `json:"producto"`
}

var Lista Productos

type Membresia struct {
	Id_membresia   string `json:"id_membresia"`
	Tipo_membresia string `json:"tipo_membresia,omitempty"`
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
		ReadXlsx1(Archivo)
		//ReadCSV(Archivo)

	}
}*/
func Almacenar(w http.ResponseWriter, req *http.Request) {
	//enableCors(&w)
	//ReadXlsx1()
	//Leer()
	json.NewEncoder(w).Encode(Contador(Lista.Data, Minusculas("CLASS")))
}

// Método para el uso de Excel
//func ReadXlsx1(){
func ReadXlsx1(Archivo multipart.File) {
	var item Item
	var items []Item
	items = nil
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
	n := 1
	item.Producto = make(map[string]string)
	for i, row := range rows {
		if row == nil {
			break
		}
		for j, colCell := range row {
			if i == 0 {
				campos = row
				//fmt.Println(campos)
				break
			}

			if campos[j] != "" {

				item.Producto[Minusculas(campos[j])] = colCell

			} else {
				item.Contador = n

				if len(item.Producto) > 0 {
					items = append(items, item)
					item.Producto = clear()
					n++
				}
			}
		}
		if item.Producto[Minusculas(campos[0])] != "" {
			item.Contador = n
			items = append(items, item)
			item.Producto = make(map[string]string)
			n++
		}
	}

	Lista.Data = items
	fmt.Print(Lista)
}
func ReadXlsx(Archivo multipart.File) []Item {
	var item Item
	var items []Item
	items = nil
	f, err := excelize.OpenReader(Archivo)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	hojas := f.GetSheetList()

	rows, err := f.GetRows(hojas[0])

	var campos []string
	n := 1
	item.Producto = make(map[string]string)
	for i, row := range rows {
		if row == nil {
			break
		}
		for j, colCell := range row {
			if i == 0 {
				campos = row
				//fmt.Println(campos)
				break
			}
			if j < len(campos) {
			if campos[j] != "" {

					item.Producto[Minusculas(campos[j])] = Minusculas(colCell)

				}
			} else {
				break
			}
		}
		if item.Producto["familia"] != "" || item.Producto["clase"] != "" {
			item.Contador = n
			items = append(items, item)
			item.Producto = make(map[string]string)
			n++
		}
	}
	return items
}

// Muestra total de productos en base a su modelo
func Contador(productos []Item, Clasificacion string) []map[string]string {
	var inventario []map[string]string

	inventary := Clasificador(productos, Clasificacion)
	for _, class := range inventary {
		var item Item
		var items []Item
		for i, objeto := range class {
			item.Contador = i
			item.Producto = objeto
			items = append(items, item)
		}
		productoPorModelo := Clasificador(items, "modelo")
		for _, product := range productoPorModelo {
			producto := product[0]
			producto["existencia"] = strconv.Itoa(len(product))
			inventario = append(inventario, producto)
		}

	}

	return inventario
}
func Clasificador(productos []Item, Clasificacion string) [][]map[string]string {
	var inventary [][]map[string]string
	clasificador := Clases(productos, Clasificacion)
	for _, class := range clasificador {

		var i = Agrupacion(productos, Clasificacion, class)
		inventary = append(inventary, i)
		fmt.Println(class)
		fmt.Println(len(i))

	}
	return inventary
}

//regresa las distintas clases de productos en un archivo
func Clases(productos []Item, Clasificacion string) []string {
	estandar := Minusculas(Clasificacion)
	var clases []string
	for _, producto := range productos {
		if len(clases) > 0 {
			existe := false
			for _, class := range clases {
				if class == producto.Producto[estandar] {
					existe = true
					break
				}
			}
			if !existe {
				clases = append(clases, producto.Producto[estandar])
			}
		} else {
			clases = append(clases, producto.Producto[estandar])
		}
	}
	fmt.Println(clases)
	return clases
}

//clasifica los productos en grupos
func Agrupacion(productos []Item, Clasificacion string, Clasificador string) []map[string]string {

	var lista []map[string]string
	for _, producto := range productos {
		if producto.Producto[Clasificacion] == Clasificador {
			lista = append(lista, producto.Producto)
		}
	}
	return lista
}

//retorna un map de strings en blanco
func clear() map[string]string {
	var blanco = make(map[string]string)
	return blanco
}

//CONVIERTE TODAS LAS LETRAS DE UNA PALABRA A MINUSCULAS
func Minusculas(palabra string) string {
	return strings.ToLower(palabra)
}
