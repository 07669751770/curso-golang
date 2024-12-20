package main

import (
	"fmt"
	"time"
)

type ConexaoBD struct {
	BDName string
}
type ConexaoAPI struct {
	ApiName string
}

func getDataBaseConnection(c chan ConexaoBD) {
	time.Sleep(2 * time.Second)
	c <- ConexaoBD{"MySQL"}
	fmt.Println("Database connected")
}

func getAPIConnection(c chan ConexaoAPI) {
	time.Sleep(10 * time.Second)
	c <- ConexaoAPI{"REST"}
	fmt.Println("API connected")
}

func main() {

	cbd := make(chan ConexaoBD, 1)
	capi := make(chan ConexaoAPI, 1)

	go getDataBaseConnection(cbd)
	go getAPIConnection(capi)

	fmt.Println("Waiting for connections.....")
	bdc, apic := <-cbd, <-capi
	fmt.Println("Connections established")
	fmt.Printf("..... Database Name: %s\n", bdc.BDName)
	fmt.Printf("..... API Name: %s\n", apic.ApiName)
}
