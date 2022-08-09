package domain

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Count        int     `json:"count"`
	Price        float64 `json:"price"`
	Id_warehouse int     `json:"id_warehouse"`
}

type ProductAndWarehouse struct {
	Product
	Warehouse Warehouse
}
