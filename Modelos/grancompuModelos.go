package Modelos
type Laptop struct {
	IdProducto int`json:"idLaptop"`
	Fecha string `json:"fecha"`
	OC string `json:"oc"`
	Suc string `json:"sucursal"`
	Familia string `json:"familia"`
	Serie string `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Generacion string `json:"generacion"`
	MemGB string`json:"memGB"`
	Velocidad string `json:"velocidad"`
	//HddTipo string `json:"hddTipo"`
	HddGB string `json:"hddGB"`
	HddSerie string `json:"hddSerie"`
	UnidadOpt string`json:"unidadOptica"`
	FuenteSerie string `json:"fuenteSerie"`
	Pulgadas string `json:"pulgadas"`
	Licencia string`json:"licencia"`
	Comentarios string `json:"comentarios"`
}

type Desktop struct {
	IdProducto int`json:"idDesktop"`
	Fecha string `json:"fecha"`
	OC string `json:"oc"`
	Suc string `json:"sucursal"`
	Familia string `json:"familia"`
	Serie string `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Generacion string `json:"generacion"`
	MemGB string`json:"memGB"`
	Velocidad string `json:"velocidad"`
	//HddTipo string `json:"hddTipo"`
	HddGB string `json:"hddGB"`
	HddSerie string `json:"hddSerie"`
	UnidadOpt string`json:"unidadOptica"`
	FuenteSerie string `json:"fuenteSerie"`
	Formato string `json:"formato"`
	Licencia string`json:"licencia"`
	Comentarios string `json:"comentarios"`
	//Provedor string `json:"provedor"`
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
	IdProducto int`json:"idAllinOne"`
	Fecha string `json:"fecha"`
	OC string `json:"oc"`
	Suc string `json:"sucursal"`
	Familia string `json:"familia"`
	Serie string `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca string `json:"marca"`
	Modelo string`json:"modelo"`
	Procesador string `json:"procesador"`
	Generacion string `json:"generacion"`
	MemGB string`json:"memGB"`
	Velocidad string `json:"velocidad"`
	//HddTipo string `json:"hddTipo"`
	HddGB string `json:"hddGB"`
	HddSerie string `json:"hddSerie"`
	UnidadOpt string`json:"unidadOptica"`
	FuenteSerie string `json:"fuenteSerie"`
	Pulgadas string `json:"pulgadas"`
	Licencia string`json:"licencia"`
	Comentarios string `json:"comentarios"`
}