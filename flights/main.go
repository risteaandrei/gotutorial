package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := `https://skyscanner-skyscanner-flight-search-v1.p.rapidapi.com/apiservices/browsequotes/v1.0/RO/RON/en-GB/BCN-sky/CRA-sky/2020-07`
	//url := "https://skyscanner-skyscanner-flight-search-v1.p.rapidapi.com/apiservices/autosuggest/v1.0/RO/RON/en-GB/?query=craiova"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Add("x-rapidapi-host", "skyscanner-skyscanner-flight-search-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "7a3ed73447msha1004e86b191c4ap11d360jsnffac3cc0acb5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
