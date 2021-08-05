package main

import (
	bd "ControlInvetario/BD"
	grancompuC "ControlInvetario/Controladores"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

//La función main es la primera función en ser ejecutada
func main() {
	//se ejecuta la función api

	go func() {
		tdr := time.Tick(2 * time.Minute)

		for horaActual := range tdr {
			fmt.Println("La hora es", horaActual)
		}
	}()

	//Utilidades.GetInsumosxProducto(a)
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
	fmt.Println("Servidor iniciado en el puerto 8085")
	fmt.Println("Detener Servidor con Ctrl + C")

	//INVENTARIO COMPUSI - GRANCOMPU
	gorillaRoute.HandleFunc("/lista", grancompuC.GetListaExel).Methods("GET")
	//Recibe los erchivos de exel y almacena el contenido en la BD
	gorillaRoute.HandleFunc("/Guardar", grancompuC.Guardar).Methods("POST")
	gorillaRoute.HandleFunc("/GuardarUno", grancompuC.GuardarUno).Methods("POST")
	//Rutas de consulta de contenido de la BD.
	gorillaRoute.HandleFunc("/GetMonitores", grancompuC.ObteneMonitores).Methods("GET")
	gorillaRoute.HandleFunc("/GetDesktops", grancompuC.ObteneDesktops).Methods("GET")
	gorillaRoute.HandleFunc("/GetAllinOne", grancompuC.ObteneAllione).Methods("GET")
	gorillaRoute.HandleFunc("/GetLaptops", grancompuC.ObteneLaptops).Methods("GET")


	//INVENTARIO CHOCOLATES
	//Aztualiza datos faltantes de insumo
	gorillaRoute.HandleFunc("/ActualizarInsumo", grancompuC.ActualizarInsumo).Methods("POST")
	//muestra el inventario de Insumos
	gorillaRoute.HandleFunc("/InventarioInsumos", grancompuC.InventarioInsumos).Methods("GET")
	//guarda en la base de datos nuevo registro de productos x Orden
	gorillaRoute.HandleFunc("/GuarProductoOrden", grancompuC.GuarProductoOrden).Methods("POST")
	//Actualiza un registro de productos por orden
	gorillaRoute.HandleFunc("/ActualizarProductoOrden", grancompuC.ActualizarProductoOrden).Methods("POST")
	//muestra inventario de porductos x orden
	gorillaRoute.HandleFunc("/ObtenerProductoOrden", grancompuC.ObtenerProductoOrden).Methods("GET")
	//gorillaRoute.HandleFunc("/AgregarDetalles", grancompuC.GuarProductoOrden).Methods("POST")


	//UTILIDADES
	//copia registro de productos de matrices a BD
	gorillaRoute.HandleFunc("/RegistrarIdProductos", grancompuC.CargaIdProdutos).Methods("GET")
	//Muestra los datos de los productos registrados en la bd desde matrices
	gorillaRoute.HandleFunc("/ObtenerIdProductos", grancompuC.ObtenerIdProductos).Methods("GET")
	//copia registro de Insumos de matrices a BD
	gorillaRoute.HandleFunc("/RegistrarInsumos", grancompuC.CargaDInsumos).Methods("GET")



	//ruta pra probar metodos
	gorillaRoute.HandleFunc("/Pruebas", grancompuC.Pruebas).Methods("GET")
	gorillaRoute.HandleFunc("/ObtenerProductos", grancompuC.ObtenerProductoOrden).Methods("GET")
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
	log.Fatal(http.ListenAndServe(":8085", handlerCORS))
}
