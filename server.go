package main

import (
	"fmt"
	"github.com/umahmood/haversine"
	"log"
	"net/http"
	"strings"
)

var id string
var ic string
var is string

// amd64 386
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	//fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		switch {
		case k == "id":
			fmt.Println(k, strings.Join(v, ""))
			///fmt.Println("id:", k)
			///fmt.Println("val:", strings.Join(v, ""))
		case k == "lat":
			fmt.Println("lat:", k)
			fmt.Println("val:", strings.Join(v, ""))
		case k == "len":
			fmt.Println("len:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		///fmt.Println("key:", k)
		///fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "ok") // отправляем данные на клиентскую сторону
	println("next")
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)  // установим роутер
	err := http.ListenAndServe(":3657", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func mainc() {
	/*oxford := haversine.Coord{Lat: 51.45, Lon: 1.15}  // Oxford, UK
	turin  := haversine.Coord{Lat: 45.04, Lon: 7.42}  // Turin, Italy
	mi, km := haversine.Distance(oxford, turin)
	fmt.Println("Miles:", mi, "Kilometers:", km)*/

	oxford := haversine.Coord{Lat: 47.29626385443, Lon: 39.87570985358} // Oxford, UK
	turin := haversine.Coord{Lat: 45.04, Lon: 7.42}                     // Turin, Italy
	mi, km := haversine.Distance(oxford, turin)
	fmt.Println("Miles:", mi, "Kilometers:", km)
}
