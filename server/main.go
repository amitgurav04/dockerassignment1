package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World!!!")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/filedata", CallServer)
	http.HandleFunc("/checksum", GetCheckSum)
	http.ListenAndServe(":8080", nil)
}

func CallServer(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", content)
}

func GetCheckSum(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%x", sha1.Sum(content))
}
