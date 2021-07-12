package Modelos
type Laptop struct {
	IdProducto int`json:"idLaptop"`
	Clase string `json:"class"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Velocidad string `json:"velocidad"`
	Generacion string `json:"generacion"`
	MarcaDisco string `json:"marcaDisco"`
	Capacidad string `json:"capacidad"`
	SerieDisco string `json:"serieDisco"`
	Bateria string `json:"bateria"`
	Eliminador string `json:"eliminador"`
	Memoria string `json:"memoria"`
	SerieOriginal string `json:"serieOriginal"`
	SerieDistribuidor string `json:"serieDistribuidor"`
	Pulgadas string `json:"pulgadas"`
}

type Desktop struct {
	IdProducto int`json:"idDesktop"`
	Clase string `json:"class"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Velocidad string `json:"velocidad"`
	Generacion string `json:"generacion"`
	MarcaDisco string `json:"marcaDisco"`
	Capacidad string `json:"capacidad"`
	SerieDisco string `json:"serieDisco"`
	Fuente_Eliminador string `json:"fuente_Eliminador"`
	Memoria string `json:"memoria"`
	SerieOriginal string `json:"serieOriginal"`
	SerieDistribuidor string `json:"serieDistribuidor"`
	Formato string `json:"formato"`
}

type Monitor struct {
	IdProducto int`json:"idMonitor"`
	Clase string `json:"class"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Pulgadas string `json:"pulgadas"`
	Tipo string `json:"tipo"`
	SerieOriginal string `json:"serieOriginal"`
	SerieDistribuidor string `json:"serieDistribuidor"`
	Salida string `json:"salida"`
}
type AllInOne struct {
	IdProducto int`json:"IdAllinOne"`
	Clase string `json:"class"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Velocidad string `json:"velocidad"`
	Generacion string `json:"generacion"`
	MarcaDisco string `json:"marcaDisco"`
	Capacidad string `json:"capacidad"`
	SerieDisco string `json:"serieDisco"`
	Fuente_Eliminador string `json:"fuente_Eliminador"`
	Memoria string `json:"memoria"`
	SerieOriginal string `json:"serieOriginal"`
	SerieDistribuidor string `json:"serieDistribuidor"`
	Pulgadas string `json:"pulgadas"`
}