package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getFundIdFromName(name string) string {
	values := map[string]string{"name": name}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://www.avanza.se/_api/fund-guide/list", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res := AvanzaFundSearch{}
	json.NewDecoder(resp.Body).Decode(&res)

	if len(res.FundListViews) == 1 {
		return strconv.Itoa(res.FundListViews[0].OrderbookID)
	} else {
		fmt.Println("Warning: fund id not found for", name, "in", res)
	}
	return ""
}

func getFundFromId(id string) (fund AvanzaFund) {
	resp, err := http.Get("https://www.avanza.se/_api/fund-guide/guide/" + id)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&fund)

	return
}
