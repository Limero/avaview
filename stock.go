package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getStockIdFromName(name string) (id string) {
	resp, err := http.Get("https://www.avanza.se/frontend/template.html/marketing/advanced-filter/advanced-filter-template?1608150968335=&widgets.marketCapitalInSek.filter.lower=&widgets.marketCapitalInSek.active=true&widgets.textSearch.filter.value=" + url.QueryEscape(name) + "&widgets.textSearch.active=true&parameters.startIndex=0&parameters.maxResults=20")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	splits := strings.Split(string(body), "/aktier/om-aktien.html/")[1]
	id = strings.Split(splits, "/")[0]

	return
}

func getStockFromId(id string) (stock AvanzaStock) {
	values := map[string]string{
		"orderbookId": id,
		"chartType":   "AREA",
		"timePeriod":  "today",
	}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://www.avanza.se/ab/component/highstockchart/getchart/orderbook", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&stock)

	return
}
