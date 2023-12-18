package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func submitData(wallet string) (*http.Response, error) {

	//Membuat permintaan
	url := fmt.Sprintf("https://airdrop-api.jup.ag/allocation/%s", wallet)
	return http.Get(url)

}

func main() {

	//Membaca File
	data, err := ioutil.ReadFile("wallet.txt")
	if err != nil {
		fmt.Println("Error", err)

	}

	//convert data to string
	dataStr := string(data)

	//Memisahkan string menjadi baris satu per satu
	lines := strings.Split(string(dataStr), "\n")

	for i, line := range lines {
		lowerCaseString := strings.ToLower(line)
		cleanLine := strings.TrimSpace(lowerCaseString)
		fmt.Printf("Baris %d: %s \n", i+1, cleanLine)

		resp, err := submitData(cleanLine)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body", err)
			return
		}

		fmt.Println("Response Body:", string(body))
	}
}
