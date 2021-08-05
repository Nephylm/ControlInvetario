package Modelos
//Estrutura de inventario de Insumos
type InventarioInsumos struct {
	IdInsumo int`json:"idInsumo"`
	Nombre string `json:"nombre"`
	Existencia int `json:"existencia"`
	CantidadMinima float32`json:"cantidadMinima"`
	CantidadMaxima float64 `json:"cantidadMaxima"`
	Activo int `json:"activo"`
	UnidadDMedida string `json:"unidadDMedida"`
}
//estructura de Catlogo de Producto
type CatalogoProducto struct {
	IdProducto int`json:"idProducto"`
	Producto string `json:"producto"`
	Cantidad string `json:"cantidad"`
	Existencia string `json:"existencia"`
}
//Estructura de producto de matrices
type Producto struct {
	IdProducto int `json:"idProducto"`
	Nombre string `json:"nombre"`
	Costo float64 `json:"costo"`
	Cantidad int `json:"cantidad"`
}
//Estrutructura de identificador de productos
type IdProducto struct {
	IdProducto int `json:"idProducto"`
	NombreProducto string `json:"nombreProducto"`
	Activo int `json:"activo"`
}
//Estructura de Inventario Productos por Orden
type ProductosOrden struct {
	IdRegistro int `json:"idRegistro"`
	OrdenProduccion string `json:"ordenProduccion"`
	IdProducto int `json:"idProducto"`
	NombreProducto string `json:"nombreProducto"`
	Existencia int `json:"existencia"`
	FechaProd string `json:"fechaProd"`
	FechaCad string `json:"fechaCad"`
}
//estructura Insumo de matrices
type Insumo struct {
	IdInsumo int `json:"idInsumo"`
	Nombre string `json:"nombre"`
	UnidadMedida UnidadDMedida `json:"unidadMedida"`
	Cantidad int `json:"cantidad"`
	Costo float64 `json:"costo"`
	Activo int `json:"activo"`
}
//estructura unidad de medida en matrices
type UnidadDMedida struct {
	IdUMedida int `json:"idUMedida"`
	Nombre string `json:"nombre"`
	Abrev string `json:"abrev"`
}
//Estructura Chocolates en BD
type ProductoTerminado struct {
	Id_PTerminado int `json:"id_pterminado"`
	Producto int `json:"producto"`
	OrdenProduccion int `json:"ordenProduccion"`
}
//Estructura insumos por producto
type InsumoxProducto struct {
	IdProducto int `json:"idProducto"`
	Producto Producto `json:"producto"`
	Insumo Insumo `json:"insumo"`
	Precio float64 `json:"precio"`
	Cantidad int `json:"cantidad"`
	CantidadUsada int `json:"cantidadUsada"`
	CostoTotal float64 `json:"costoTotal"`
}
//Estructura de detalles del producto
type Detalles struct {
	IdDetalleP int `json:"idDetalleP"`
	IdProducto int `json:"idProducto"`
	Inventario int `json:"inventario"`
}
type RespuestaSencilla struct {
	CodigoRespHTTP int `json:"CodigoRespHTTP"`
	Response string `json:"Response"`
}
