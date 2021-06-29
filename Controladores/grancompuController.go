package Controladores

import (

	bd "ControlInvetario/BD"
	grancompu "ControlInvetario/Utilidades"
	"encoding/json"
	"net/http"
)

func Guardar(w http.ResponseWriter, r *http.Request) {
	grancompu.ReadXlsx()
	json.NewEncoder(w).Encode(bd.Guardar())
}

func ObteneInventario(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.GetInventario())
}
func Pruebas(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bd.PruebaAlmacenar())
}