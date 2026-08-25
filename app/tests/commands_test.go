package config_test

import (
	"testing"
	"os"
	"fmt"
	"net/http"

)


func TestGetEnvironmentConfig(t *testing.T) {
	println("============== iniciado test de comandos ")
	
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

	os.Args = []string{
    "caminho/do/programa",
    "Walter",
    "123",
    "teste",
	}

	println(os.Args[0])
	println(os.Args[1])
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

}
func TestGetEnvironmentConfig2(t *testing.T) {
	println("============== iniciado test de comandos ")
	PORT := ":8001"
 	arguments := os.Args
    if len(arguments) != 1 {
        PORT = ":" + arguments[1]
    }
    fmt.Println("Using port number: ", PORT)

}
func TestGetEnvironmentConfig3(t *testing.T) {
	PORT := ":8001"
	err := http.ListenAndServe(PORT, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

}