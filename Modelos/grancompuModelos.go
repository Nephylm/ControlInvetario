package Modelos

//Estructura de Laptop
type Laptop struct {
	IdProducto int    `json:"idLaptop"`
	Fecha      string `json:"fecha"`
	OC         int    `json:"oc"`
	Suc        string `json:"sucursal"`
	CodigoProducto string `json:"codigoProducto"`
	Familia    string `json:"familia"`
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Procesador string `json:"procesador"`
	Generacion int    `json:"generacion"`
	Velocidad  string `json:"velocidad"`
	MemGB      string    `json:"memGB"`
	//HddTipo string `json:"hddTipo"`
	SerieBateria  string `json:"serieBateria"`
	HddGB         string `json:"hddGB"`
	HddSerie      string `json:"hddSerie"`
	SerieOriginal string `json:"serieOriginal"`
	Pulgadas      string `json:"pulgadas"`
	Camara        string `json:"camara"`
	Eliminador    string `json:"eliminador"`
	//Comentarios   string `json:"comentarios"`
	Activo    int    `json:"activo"`
	FechaVent string `json:"fechaVent"`
	SerieDoc string `json:"serieDoc"`
	DocVent string `json:"doc_vent"`
}

//Estructura de Desktop
type Desktop struct {
	IdProducto    int    `json:"idDesktop"`
	Fecha         string `json:"fecha"`
	OC            int    `json:"oc"`
	Suc           string `json:"sucursal"`
	Familia       string `json:"familia"`
	CodigoProducto string `json:"codigoProducto"`
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
	Activo      int    `json:"activo"`
	FechaVent   string `json:"fechaVent"`
	SerieDoc string `json:"serieDoc"`
	DocVent string `json:"docVent"`
	//Provedor string `json:"provedor"`
}

//Estructura de Monitor
type Monitor struct {
	IdProducto     int    `json:"idMonitores"`
	Fecha          string `json:"fecha"`
	OC             int    `json:"oc"`
	Suc            string `json:"sucursal"`
	Familia        string `json:"familia"`
	CodigoProducto string `json:"codigoProducto"`
	Serie          string `json:"serie"`
	SerieOriginal  string `json:"serieOriginal"`
	Marca          string `json:"marca"`
	Modelo         string `json:"modelo"`
	Forma          string `json:"forma"`
	Base           string `json:"base"`
	Tipo           string `json:"tipo"`
	Salidas        string `json:"salidas"`
	HDMI           string `json:"HDMI"`
	Clase          string `json:"clase"`
	Tamaño         string `json:"Tamaño"`
	Activo         int    `json:"activo"`
	FechaVent      string `json:"fechaVent"`
	SerieDoc       string `json:"serieDoc"`
	DocVent        string `json:"DocVent"`
}

//Estructura de AllInOne
type AllInOne struct {
	IdProducto    int    `json:"idAllinOne"`
	Fecha         string `json:"fecha"`
	OC            int    `json:"oc"`
	SUC           string `json:"sucursal"`
	Familia       string `json:"familia"`
	Serie         int    `json:"serie"`
	SerieOriginal string `json:"serieOriginal"`
	Marca         string `json:"marca"`
	Modelo        string `json:"modelo"`
	Procesador    string `json:"procesador"`
	Gen           int    `json:"generacion"`
	Mem_GB        string `json:"memGB"`
	Velocidad     string `json:"velocidad"`
	HDD           string `json:"hdd"`
	HddSerie      string `json:"hddSerie"`
	UnidadOp      string `json:"unidadOptica"`
	Fuente        string `json:"fuenteSerie"`
	Formato       string `json:"formato"`
	Pulgadas      string `json:"pulgadas"`
	Licencia      string `json:"licencia"`
	Comentarios   string `json:"comentarios"`
}

//Estructura General
type Mercancia struct {
	IdProducto		int		`json:"idProducto"`
	Familia			string	`json:"familia,omitempty"`
	SerieOriginal	string	`json:"serieOriginal,omitempty"`
	Fecha         string `json:"fecha,omitempty"`

	OC            	int    `json:"oc,omitempty"`
	SUC           	string `json:"sucursal,omitempty"`
	Serie         	string    `json:"serie,omitempty"`
	Marca         	string `json:"marca,omitempty"`
	Modelo        	string `json:"modelo,omitempty"`
	Procesador    	string `json:"procesador,omitempty"`
	Gen           	int    `json:"generacion,omitempty"`
	Mem_GB        	string `json:"memGB,omitempty"`
	Velocidad     	string `json:"velocidad,omitempty"`
	HDD           	string `json:"hdd,omitempty"`
	HddSerie      	string `json:"hddSerie,omitempty"`
	UnidadOp      	string `json:"unidadOptica,omitempty"`
	Fuente        	string `json:"fuenteSerie,omitempty"`
	Formato       	string `json:"formato,omitempty"`
	Pulgadas      	string `json:"pulgadas,omitempty"`
	Licencia      	string `json:"licencia,omitempty"`
	Comentarios   	string `json:"comentarios,omitempty"`
	CodigoProducto	string `json:"codigoProducto,omitempty"`
	Forma          	string `json:"forma,omitempty"`
	Base           	string `json:"base,omitempty"`
	Tipo           	string `json:"tipo,omitempty"`
	Salidas        	string `json:"salidas,omitempty"`
	HDMI           	string `json:"HDMI,omitempty"`
	Clase          	string `json:"clase,omitempty"`
	Tamanio         string `json:"tamaño,omitempty"`
	Activo         	int    `json:"activo,omitempty"`
	FechaVent      	string `json:"fechaVent,omitempty"`
	SerieDoc       	string `json:"serieDoc,omitempty"`
	HddGB       	string `json:"hddGB,omitempty"`
	Camara        	string `json:"camara,omitempty"`
	SerieBateria  string `json:"serieBateria,omitempty"`
	Eliminador    string `json:"eliminador,omitempty"`
	DocVent string `json:"doc_vent,omitempty"`
}

type CodigoProducto struct {
	IdCodigo		int 	`json:"idCodigo,omitempty"`
	Familia			string	`json:"familia,omitempty"`
	CodigoProducto	string	`json:"codigoProducto,omitempty"`
	Procesador		string	`json:"procesador,omitempty"`
	Generacion		int		`json:"generacion,omitempty"`
	Formato			string	`json:"formato,omitempty"`
}
