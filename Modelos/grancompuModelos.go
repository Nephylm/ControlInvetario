package Modelos

//Estructura de Laptop
type Laptop struct {
	IdProducto int    `json:"idLaptop"`
	Fecha      string `json:"fecha"`
	OC         int    `json:"oc"`
	Suc        string `json:"sucursal"`
	Familia    string `json:"familia"`
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Procesador string `json:"procesador"`
	Generacion int    `json:"generacion"`
	Velocidad  string `json:"velocidad"`
	MemGB      int    `json:"memGB"`
	//HddTipo string `json:"hddTipo"`
	SerieBateria  string `json:"serieBateria"`
	HddGB         string `json:"hddGB"`
	HddSerie      string `json:"hddSerie"`
	SerieOriginal string `json:"serieOriginal"`
	Pulgadas      string `json:"pulgadas"`
	Camara        string `json:"camara"`
	Eliminador    string `json:"eliminador"`
	Comentarios   string `json:"comentarios"`
}

//Estructura de Desktop
type Desktop struct {
	IdProducto    int    `json:"idDesktop"`
	Fecha         string `json:"fecha"`
	OC            int    `json:"oc"`
	Suc           string `json:"sucursal"`
	Familia       string `json:"familia"`
	Serie         string `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca         string `json:"marca"`
	Modelo        string `json:"modelo"`
	Procesador    string `json:"procesador"`
	Generacion    int    `json:"generacion"`
	MemGB         string `json:"memGB"`
	Velocidad     string `json:"velocidad"`
	//HddTipo string `json:"hddTipo"`
	HddGB       string `json:"hddGB"`
	HddSerie    string `json:"hddSerie"`
	UnidadOpt   string `json:"unidadOptica"`
	FuenteSerie string `json:"fuenteSerie"`
	Formato     string `json:"formato"`
	Licencia    string `json:"licencia"`
	Comentarios string `json:"comentarios"`
	//Provedor string `json:"provedor"`
}

//Estructura de Monitor
type Monitor struct {
	IdProducto        int    `json:"idMonitor"`
	Clase             string `json:"class"`
	Marca             string `json:"marca"`
	Modelo            string `json:"modelo"`
	Pulgadas          string `json:"pulgadas"`
	Tipo              string `json:"tipo"`
	SerieOriginal     string `json:"serieOriginal"`
	SerieDistribuidor string `json:"serieDistribuidor"`
	Salida            string `json:"salida"`
}

//Estructura de AllInOne
type AllInOne struct {
	IdProducto    int    `json:"idAllinOne"`
	Fecha         string `json:"fecha"`
	OC            int    `json:"oc"`
	Suc           string `json:"sucursal"`
	Familia       string `json:"familia"`
	Serie         string `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca         string `json:"marca"`
	Modelo        string `json:"modelo"`
	Procesador    string `json:"procesador"`
	Generacion    int    `json:"generacion"`
	MemGB         int    `json:"memGB"`
	Velocidad     string `json:"velocidad"`
	//HddTipo string `json:"hddTipo"`
	HddGB       string `json:"hddGB"`
	HddSerie    string `json:"hddSerie"`
	UnidadOpt   string `json:"unidadOptica"`
	FuenteSerie string `json:"fuenteSerie"`
	Pulgadas    string `json:"pulgadas"`
	Licencia    string `json:"licencia"`
	Comentarios string `json:"comentarios"`
}
