package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*Constantes de la app
============================================== */
const (
	basePath   = "http://gateway.marvel.com/v1/public/characters?"
	publicKey  = "2f7180348dab8d57d616d1472feb94d3"
	privateKey = "b7702d0082e2eb64a44d7c0964a0cc0581ca31fa"
)

/*============================================== */

/*Variables de la app
============================================== */
var myLink string
var optionUser int
var responseObject Response
var text string

/*============================================== */

/*Estructura de los heroes
============================================== */
// A Response struct to map the Entire Response
type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

// Data Struct
type Data struct {
	Limit int     `json:"limit"`
	Total int     `json:"total"`
	Count int     `json:"count"`
	Heroe []Heroe `json:"results"`
}

// Heroe Struct
type Heroe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Modified    string   `json:"modified"`
	Comics      []Comic  `json:"comics"`
	Stories     []Storie `json:"stories"`
	Events      []Event  `json:"Events"`
	Series      []Serie  `json:"series"`
}

// Comic Struct
type Comic struct {
	Available int    `json:"available"`
	Returned  int    `json:"returned"`
	Items     []Item `json:"items"`
}

// Storie Struct
type Storie struct {
	Available int    `json:"available"`
	Returned  int    `json:"returned"`
	Items     []Item `json:"items"`
}

// Event Struct
type Event struct {
	Available int    `json:"available"`
	Returned  int    `json:"returned"`
	Items     []Item `json:"items"`
}

// Serie Struct
type Serie struct {
	Available int    `json:"available"`
	Returned  int    `json:"returned"`
	Items     []Item `json:"items"`
}

// Item Struct
type Item struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

/*============================================== */

/*FUNCIONES

función para establecer un timestamp del día en que se ingrese, hacer el md5 con las llaves pública y privada junto con el timestamp para obtener el enlace de la api
============================================== */
func setLinkHash(heroName string) string {
	ts := time.Now().Format("2006-01-02T15:04:05")

	hash := md5.New()
	io.WriteString(hash, ts)
	io.WriteString(hash, privateKey)
	io.WriteString(hash, publicKey)

	key := fmt.Sprintf("%x", hash.Sum(nil))

	name := heroName

	if len(name) != 0 {
		myLink = basePath + "name=" + name + "&ts=" + ts + "&apikey=" + publicKey + "&hash=" + key
	} else {
		myLink = basePath + "ts=" + ts + "&apikey=" + publicKey + "&hash=" + key
	}

	return myLink
}

/*---------------------------------------------------------*/

/*
función para dibujar pantalla
============================================== */
func drawInterface() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Bienvenido a la App 'MarvelGo'")
	fmt.Println(" ")
	fmt.Println("*************************************")
	fmt.Println("Podrás buscar un héroe por su nombre o listar una serie de 20 héroes desde la Marvel Comics API ")
	fmt.Println(" ")
	fmt.Println("Opción 1 - Buscar por nombre ")
	fmt.Println("Opción 2 - Listar ")
}

func drawInterfaceSearch() {
	fmt.Println(" ")
	fmt.Println("  --------------------------------------------  ")
	fmt.Println("Escribe el nombre del héroe que deseas buscar: ")
}

/*============================================== */

/*Main
============================================== */
func main() {

	text = ""

	res, err := http.Get(setLinkHash(text))
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	json.Unmarshal(responseData, &responseObject)

	drawInterface()

	fmt.Scanln(&optionUser)
	/*  Use of Switch Case in Golang */
	switch optionUser {
	case 1:
		drawInterfaceSearch()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text = scanner.Text()
		if len(text) != 0 {
			fmt.Println("Buscando: " + text)
			res, err := http.Get(setLinkHash(text))
			if err != nil {
				fmt.Print(err.Error())
			}

			responseData, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			json.Unmarshal(responseData, &responseObject)

			if len(responseObject.Data.Heroe) > 0 {
				fmt.Println(" ")
				fmt.Println("Coincidencias: ")
				fmt.Println("  --------------------------------------------  ")
				for i := 0; i < len(responseObject.Data.Heroe); i++ {
					fmt.Println(responseObject.Data.Heroe[i].Name)
					fmt.Println(responseObject.Data.Heroe[i].Description)
					fmt.Println(" ")
					if len(responseObject.Data.Heroe[i].Comics) > 0 {
						fmt.Println("Comics: ")
						for j := 0; j < len(responseObject.Data.Heroe[i].Comics); j++ {
							fmt.Println(responseObject.Data.Heroe[i].Comics[j])
						}
						fmt.Println(" ")
					}
					if len(responseObject.Data.Heroe[i].Stories) > 0 {
						fmt.Println("Historias: ")
						for j := 0; j < len(responseObject.Data.Heroe[i].Stories); j++ {
							fmt.Println(responseObject.Data.Heroe[i].Stories[j])
						}
						fmt.Println(" ")
					}
					if len(responseObject.Data.Heroe[i].Events) > 0 {
						fmt.Println("Eventos: ")
						for j := 0; j < len(responseObject.Data.Heroe[i].Events); j++ {
							fmt.Println(responseObject.Data.Heroe[i].Events[j])
						}
						fmt.Println(" ")
					}
					if len(responseObject.Data.Heroe[i].Series) > 0 {
						fmt.Println("Series: ")
						for j := 0; j < len(responseObject.Data.Heroe[i].Series); j++ {
							fmt.Println(responseObject.Data.Heroe[i].Series[j])
						}
						fmt.Println(" ")
					}
					fmt.Println("Modificado en: " + responseObject.Data.Heroe[i].Modified)
					fmt.Println(" ")
					fmt.Println("  --------------------------------------------  ")
				}
			} else {
				fmt.Println("No hay resultados...")
			}

		} else {
			fmt.Println("Tienes que escribir un nombre ")
			drawInterfaceSearch()
		}
	case 2:
		//Listando los heroes
		fmt.Println(" ")
		fmt.Println("Lista de héroes: ")
		fmt.Println("----------------------------------------------------")
		for i := 0; i < len(responseObject.Data.Heroe); i++ {
			fmt.Println(responseObject.Data.Heroe[i].Name)
			fmt.Println(responseObject.Data.Heroe[i].Description)
			fmt.Println(" ")
			if len(responseObject.Data.Heroe[i].Comics) > 0 {
				fmt.Println("Comics: ")
				for j := 0; j < len(responseObject.Data.Heroe[i].Comics); j++ {
					fmt.Println(responseObject.Data.Heroe[i].Comics[j])
				}
				fmt.Println(" ")
			}
			if len(responseObject.Data.Heroe[i].Stories) > 0 {
				fmt.Println("Historias: ")
				for j := 0; j < len(responseObject.Data.Heroe[i].Stories); j++ {
					fmt.Println(responseObject.Data.Heroe[i].Stories[j])
				}
				fmt.Println(" ")
			}
			if len(responseObject.Data.Heroe[i].Events) > 0 {
				fmt.Println("Eventos: ")
				for j := 0; j < len(responseObject.Data.Heroe[i].Events); j++ {
					fmt.Println(responseObject.Data.Heroe[i].Events[j])
				}
				fmt.Println(" ")
			}
			if len(responseObject.Data.Heroe[i].Series) > 0 {
				fmt.Println("Series: ")
				for j := 0; j < len(responseObject.Data.Heroe[i].Series); j++ {
					fmt.Println(responseObject.Data.Heroe[i].Series[j])
				}
				fmt.Println(" ")
			}
			fmt.Println("Modificado en: " + responseObject.Data.Heroe[i].Modified)
			fmt.Println(" ")
			fmt.Println("  --------------------------------------------  ")
		}
	}

}

/*============================================== */
