package main

import (
	"fmt"
	//"bufio"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	//"strings"
	"math/rand"
)



func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Give me a number: ")
	//num, _ := reader.ReadString('\n')
	//num = strings.TrimSuffix(num, "\n")
	num := rand.Intn(100)
	fmt.Printf("Number selected: %d\n",num)
	url := fmt.Sprintf("http://numbersapi.com/%d/math",num)

	resp, e := http.Get(url)
	if e != nil {
		log.Fatal("Response: ", e)
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
