package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://mariainesserra.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 999)

	resp.Body.Read(bs)

	fmt.Println("status is", resp.Status)
	fmt.Println(string(bs))
}
