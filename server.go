package main

import (
	"encoding/csv"
	"fmt"
	"github.com/umahmood/haversine"
	"log"
	"net/http"
	"os"
	"strings"
)


var mcc string
var mnc string
var lac string
var cellid string

// amd64 386
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	//fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])

	if len(r.Form) != 4 { // 4 параметра :з
		fmt.Fprintf(w, "error 91237")
		return
	}	else {
		fmt.Fprintf(w, "ok")
	}

	// key - value
	for k, v := range r.Form {
		switch {
		case k == "mcc":
			mcc = strings.Join(v, "")
			//fmt.Println(k, strings.Join(v, ""))
		case k == "mnc":
			mnc = strings.Join(v, "")
			//fmt.Println("lat:", k)
			//fmt.Println("val:", strings.Join(v, ""))
		case k == "lac":
			lac = strings.Join(v, "")
			//fmt.Println("len:", k)
			//fmt.Println("val:", strings.Join(v, ""))
		case k == "cellid":
			cellid = strings.Join(v, "")
			//fmt.Println("len:", k)
			//fmt.Println("val:", strings.Join(v, ""))
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

	empData := [][]string{
		{"Name", "City", "Skills"},
	}
	println(empData)
	print("asdasd")
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

func test(id string, mcc string, mnc string, lac string, cellid string,) {
	csvFile, err := os.Create(id + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	empData := [][]string{
		{"Name", "City", "Skills"},
	}

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}