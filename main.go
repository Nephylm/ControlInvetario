package main

import (
	bd "ControlInvetario/BD"
	"fmt"
	"log"
	"net/http"

	grancompuC "ControlInvetario/Controladores"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//La función main es la primera función en ser ejecutada
func main() {
	//se ejecuta la función api
	api()
}

//api realiza la configuración del servidor con el paquete gorilla mux
func api() {
	//Se hace llamar a la funcion para crear una nueva conexion a la base de datos
	bd.NuevaConexionBD()
	//Se termina la conexion con la base de datos
	//'defer' terminará la conexión a la base de datos hasta que todos los demás
	// procesos de la función 'api' terminen
	defer bd.TerminarConexionBD()

	//Se crea el nuevo router para crear la api
	//El metodo 'StrictSlash', cuando está en 'true' no lee la última diagonal de la URL
	//Es decir, que si se solicita '/agregaUsuario/' es lo mismo que solicitar '/agregaUsuario'
	gorillaRoute := mux.NewRouter().StrictSlash(true)
	fmt.Println("Servidor iniciado en el puerto 8081")
	fmt.Println("Detener Servidor con Ctrl + C")

	gorillaRoute.HandleFunc("/lista", GetListaExel).Methods("GET")
	gorillaRoute.HandleFunc("/Guardar", grancompuC.Guardar).Methods("GET")
	gorillaRoute.HandleFunc("/Inventario", grancompuC.ObteneInventario).Methods("GET")
	gorillaRoute.HandleFunc("/Pruebas", grancompuC.Pruebas).Methods("GET")
	//'PathPrefix' estamos indicando apartir de que direccion se va a escuchar para publicar archivos
	//'http.FileServer' indica que va a estar imprimiendo archivos en el navegador
	//'http.Dir' indica que todo lo que esté dentro de ese directorio se podrá ejecutar con
	// 'http.FileServer' y dentro se indica el nombre de la carpeta
	gorillaRoute.PathPrefix("/").Handler(http.FileServer(http.Dir("./vistas/")))

	//Toda la configuracion y evento registra en la URL será manejado por la variable 'gorillaRoute'
	http.Handle("/", gorillaRoute)
	//Implmentación de la política CORS
	handlerCORS := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		//AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders: []string{"*"},
		//ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		/*MaxAge: 7200,
		OptionsPassthrough: true,
		Debug: true,*/
	}).Handler(gorillaRoute)

	//Se lanza el servidor en el puerto 8081
	//en caso de existir error, se mostrará con 'log.Fatal'
	//log.Fatal(http.ListenAndServe(":8081", gorillaRoute))
	log.Fatal(http.ListenAndServe(":8081", handlerCORS))
}
