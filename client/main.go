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
	http.HandleFunc("/checksum", GetCheckSum)
	http.ListenAndServe(":8090", nil)
}

func GetCheckSum(w http.ResponseWriter, r *http.Request) {
	server := os.Getenv("SERVER_CONTAINER_NAME")
	resp, err := http.Get("http://"+server+":8080/filedata")
	if err != nil {
		log.Fatal(err)
	}
	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(w, "File data got from server is: %s\n", htmlData)
	fmt.Fprintf(w, "Checksum of above file data is: %x\n", sha1.Sum([]byte(htmlData)))

	// print out
	fmt.Println(os.Stdout, string(htmlData))
	resp, err = http.Get("http://"+server+":8080/checksum")
	if err != nil {
		log.Fatal(err)
	}
	checksum, err := ioutil.ReadAll(resp.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "Actual Checksum got from server is: %s", checksum)
}
