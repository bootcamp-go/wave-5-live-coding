package main

import (
	"fmt"
)

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type tienda struct {
	p []producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto()
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func main() {
	producto1 := nuevoProducto(pequeño, "pegamento", 5.50)
	producto2 := nuevoProducto(mediano, "pala", 102.25)
	producto3 := nuevoProducto(grande, "motor", 2030.00)

	producto1.CalcularCosto()
	producto2.CalcularCosto()
	producto3.CalcularCosto()

	tienda := nuevaTienda()

	tienda.Agregar(producto1)
	tienda.Agregar(producto2)
	tienda.Agregar(producto3)

	fmt.Println(tienda)
	fmt.Println(tienda.Total())
}

func nuevoProducto(tipo string, nombre string, precio float64) producto {
	return producto{tipo: tipo, nombre: nombre, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (p *producto) CalcularCosto() {
	switch p.tipo {
	case "pequeño":
		p.precio = p.precio
	case "mediano":
		p.precio = p.precio * 1.03
	case "grande":
		p.precio = p.precio*1.06 + 2500
	}
}

func (t *tienda) Agregar(np producto) {
	t.p = append(t.p, np)
}

func (t *tienda) Total() float64 {
	var total float64
	for _, v := range t.p {
		total += v.precio
	}

	return total
}
