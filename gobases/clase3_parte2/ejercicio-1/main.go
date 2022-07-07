package main

import "fmt"

// ===========================================
// ============= Struct usuario ==============
// ===========================================

type usuario struct {
	Nombre      string
	Apellido    string
	edad        int
	correo      string
	contrasenia string
}

// ===========================================
// ================ Funciones ================
// ===========================================

func cambiarNombre(usuario *usuario, nombre string, apellido string) {
	usuario.Nombre = nombre
	usuario.Apellido = apellido
}

func cambiarEdad(usuario *usuario, edad int) {
	usuario.edad = edad
}

func cambiarCorreo(usuario *usuario, correo string) {
	usuario.correo = correo
}

func cambiarContrasenia(usuario *usuario, contrasenia string) {
	usuario.contrasenia = contrasenia
}

// ===========================================
// ================== Main ===================
// ===========================================

func main() {

	var usuario2 *usuario = &usuario{
		Nombre:      "a",
		Apellido:    "a",
		edad:        100,
		correo:      "a",
		contrasenia: "a",
	}

	fmt.Println("Usuario: ", *usuario2)
	fmt.Println("Usuario direccion de memoria: ", &usuario2)
	fmt.Println("Dirección Memoria de Nombre: ", &usuario2.Nombre)
	fmt.Println("Dirección Memoria de Apellido: ", &usuario2.Apellido)
	fmt.Println("Dirección Memoria de edad: ", &usuario2.edad)

	var nombre string = "Juan"
	var apellido string = "Perez"
	var edad int = 20
	var correo string = "juanPerez@gmail.com"
	var contrasenia string = "123456"

	cambiarNombre(usuario2, nombre, apellido)
	cambiarEdad(usuario2, edad)
	cambiarCorreo(usuario2, correo)
	cambiarContrasenia(usuario2, contrasenia)

	fmt.Println("Nuevos datos Usuario: ", *usuario2)
	fmt.Println("Usuario direcciones de memoria: ", &usuario2)
	fmt.Println("Dirección Memoria de Nombre: ", &usuario2.Nombre)
	fmt.Println("Dirección Memoria de Apellido: ", &usuario2.Apellido)
	fmt.Println("Dirección Memoria de edad: ", &usuario2.edad)
}
