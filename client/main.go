package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	
	// loggear "Hola soy un log" usando la biblioteca log
	//log.Println("hola")
	
	// validar que la config este cargada correctamente
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")



	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Ip)
	log.Println(globals.ClientConfig.Mensaje)
	log.Println(globals.ClientConfig.Puerto)
	

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()
	utils.LeerConsolaHastaVacio()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
