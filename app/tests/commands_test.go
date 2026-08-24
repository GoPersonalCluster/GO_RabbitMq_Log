package config_test

import (
	"testing"
	"os"
	"fmt"
)


func TestGetEnvironmentConfig(t *testing.T) {
	println("============== iniciado test de comandos ")
	
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