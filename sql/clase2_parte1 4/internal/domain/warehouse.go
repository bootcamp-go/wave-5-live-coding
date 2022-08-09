package domain

type Warehouse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WarehouseAndProducts struct {
	Warehouse
	Products []Product `json:"products"`
}
