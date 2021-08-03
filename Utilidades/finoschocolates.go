package Utilidades

import (
	modelos "ControlInvetario/Modelos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Peticion()  {
	response, err := http.Get("http://74.208.31.248:8086/api/productos")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	res:=string(responseData)
	fmt.Println(res)
	//var respuesta []map[string]interface{}
	var respuesta []modelos.Producto
	json.Unmarshal(responseData,&respuesta)
	fmt.Println(respuesta)
}

func GetInsumos() (LInsumos []modelos.Insumo) {
	response, err := http.Get("http://74.208.31.248:8086/insumos")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//res:=string(responseData)
	//fmt.Println(res)
	//var respuesta []map[string]interface{}
	json.Unmarshal(responseData,&LInsumos)
	fmt.Println(LInsumos)
	return LInsumos
}
func GetProductos() (LProductos []modelos.IdProducto) {
	response, err := http.Get("http://74.208.31.248:8086/api/productos")
	var LisProductos []modelos.Producto
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData,&LisProductos)
	for _,producto:=range LisProductos{
		LProductos= append(LProductos,modelos.IdProducto{
			IdProducto: producto.IdProducto,
			NombreProducto: producto.Nombre,
			Activo: 1,
		})
	}
	fmt.Println(LProductos)
	return LProductos
}

func GetInsumosxProducto(IdProducto string) (InsumosxProducto []modelos.InsumoxProducto) {
	response, err := http.Get("http://74.208.31.248:8086/api/insumosxproducto/"+IdProducto)
	var LisInsumosxProducto []modelos.InsumoxProducto
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData,&LisInsumosxProducto)

	fmt.Println(LisInsumosxProducto)
	return LisInsumosxProducto
}