package Controladores

import (
	bd "ControlInvetario/BD"
	modelos "ControlInvetario/Modelos"
	"encoding/json"
	"fmt"
	"net/http"
)

func GuardarMercancias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//crea item y array de items
	var mercancias []modelos.Mercancia
	//des serializa el json a un map
	json.NewDecoder(r.Body).Decode(&mercancias)
	fmt.Println(mercancias)
	//guarda el producto
	json.NewEncoder(w).Encode(bd.Registro(mercancias))
}

func ActualizaMercancia(w http.ResponseWriter, r *http.Request) {
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
func ConsultarMercancias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bd.ConsultaMercancias())
}