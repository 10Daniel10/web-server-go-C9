package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int
	Nombre          string
	Precio          float64
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func leerProductosDesdeArchivo(archivo string) ([]Producto, error) {
	var productos []Producto

	//Abre el archivo en modo lectura
	file, err := os.Open(archivo)
	if err != nil {
		return nil, err //Si hay un error al abrir el archivo, se retorna el error
	}
	defer file.Close() //Se asegura de que el archivo se cierre al final de la función

	//Crea un nuevo decodificador JSON para el archivo
	decoder := json.NewDecoder(file)
	//Decodifica el contenido del archivo en la variable 'productos'
	err = decoder.Decode(&productos)
	if err != nil {
		return nil, err //Si hay un error al decodificar el JSON, se retorna el error
	}

	return productos, nil //Retorna el slice de productos leído del archivo
}

func main() {

	r := gin.Default()

	//Leer productos desde el archivo
	productos, err := leerProductosDesdeArchivo("productos.json")
	if err != nil {
		fmt.Println("Error al leer productos:", err)
		return
	}

	//Imprimir la lista de productos por consola
	for _, producto := range productos {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d, FechaDeCreacion: %s\n", producto.Id, producto.Nombre, producto.Precio, producto.Stock, producto.FechaDeCreacion)
	}

	//Ruta para obtener la lista de productos en formato JSON
	r.GET("/productos", func(c *gin.Context) {
		c.JSON(200, productos)
	})

	r.Run(":8080")
}
