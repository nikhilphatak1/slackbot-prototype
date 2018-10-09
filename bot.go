package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	//"math/rand"
)



func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me a number: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSuffix(num, "\n")
	url := fmt.Sprintf("http://numbersapi.com/%s/math",num)

	resp, e := http.Get(url)
	if e != nil {
		log.Fatal("GET request: ", e)
		return
	} else {
		msg, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println("ERR:", e)
			os.Exit(1)
		}
		fmt.Printf("%s\n", msg)
		resp.Body.Close()
	}

}
