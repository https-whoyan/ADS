package main

import (
	funcs "github.com/https-whoyan/ADS/Assigment1/pkg/functions"
	"log"
)

func main() {
	// Сюда нужно вписать номер задания от 1 до 10:
	// К примеру, task4 или task9
	err := funcs.Task10()
	if err != nil {
		log.Fatal(err)
	}
}
