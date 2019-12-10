package main

import (
		"fmt"
		"net/http"
		"io/ioutil"
		"github.com/buger/jsonparser"
		"strings"
)

func urbanDictRequest(input string) {
		cleaned := strings.Replace(input, " ", "%20", -1)
		url := "https://mashape-community-urban-dictionary.p.rapidapi.com/define?term=" + cleaned

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "mashape-community-urban-dictionary.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "39d3c9447fmshc8c96b901cd3437p1d9259jsn1f8798b3c03e")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		msg, _, _, _ := jsonparser.Get(body, "list")
		defs := string(msg)
		//fmt.Println(defs
		i := strings.Index(defs, ":\"")
		j := strings.Index(defs, "\",\"")
		single := defs[(i + 2):(j)]
		noSlash := strings.Replace(single, "\\r\\n", "", -1)
		noBracketL := strings.Replace(noSlash, "[", "", -1)
		noBracketR := strings.Replace(noBracketL, "]", "", -1)
		final := strings.Replace(noBracketR, "\\", "", -1)

		fmt.Println(final)
}
