package main

import (
	"client/globals"
	"client/utils"
	//"log"
)

func main() {
	utils.ConfigurarLogger()

	
	// loggear "Hola soy un log" usando la biblioteca log
	//log.Println("hola")
	
	// validar que la config este cargada correctamente
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	//Clave
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Clave)

	utils.GenerarYEnviarPaquete()


}
