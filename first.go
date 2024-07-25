// First Go program
package main

import (
	"fmt"
	"log"

	// paquete para ejecución de comando en Win
	"os/exec"
	// paquete para conocer los tipos de datos
	//"reflect"

	// Paqueteria para temas de API
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestBody define la estructura del cuerpo de la solicitud POST
type RequestBody struct {
	ProgramName string `json:"program_name"`
}

func main() {
	router := gin.Default()

	router.POST("/check-program", checkProgramHandler)

	router.Run(":8080") // Inicia el servidor en el puerto 8080
}

func checkProgramHandler(c *gin.Context) {
	var requestBody RequestBody

	// Vincula el cuerpo de la solicitud JSON a la estructura RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	programName := requestBody.ProgramName
	isRunning, output := isProgramRunning(programName)

	if isRunning {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("El programa %s está en ejecución", programName), "output": output})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("El programa %s no está en ejecución", programName)})
	}
}

func isProgramRunning(programName string) (bool, string) {
	log.Println(programName)
	//programName := "brave.exe"
	cmd := exec.Command("tasklist", "/nh", "/FI", fmt.Sprintf("IMAGENAME eq %s", programName))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error ejecutando el comando:", err)
		return false, ""
	}

	// Verificamos si el programa está en la lista de tareas
	if string(output) == "INFO: No tasks are running which match the specified criteria.\r\n" {
		return false, ""
	}

	return true, "" //string(output)
}

/**
func buscarMatlab() {
	programName := "brave.exe"
	cmd := exec.Command("tasklist", "/nh", "/FI", fmt.Sprintf("IMAGENAME eq %s", programName))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("no se encontro el programa", err)
		return
	}
	fmt.Println(string(output))
	//fmt.Println(reflect.TypeOf(output))
}

// Main function
func main() {
	buscarMatlab()
}*/
