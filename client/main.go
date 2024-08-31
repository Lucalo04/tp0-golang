package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("hola")
	// loggear "Hola soy un log" usando la biblioteca log

	globals.ClientConfig = utils.IniciarConfiguracion("/home/luca/Desktop/tp0-golang/client/config.json")
	// validar que la config este cargada correctamente

	log.Println(globals.ClientConfig.Ip)
	log.Println(globals.ClientConfig.Mensaje)
	log.Println(globals.ClientConfig.Puerto)

	// loggeamos el valor de la config


	utils.LeerConsolaHastaVacio()

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
