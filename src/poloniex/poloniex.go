package poloniex

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type APIQuery struct {
	command string
	currencyPair	string
}

func sendQuery(q APIQuery) (string, error) {
	if q.command == "returnTicker" || q.command == "return24hVolume" {
		resp, err := http.Get("https://poloniex.com/public?command=" + q.command)
		if err != nil {
			return "", err
		}

		ret, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		resp.Body.Close()

		return string(ret), nil
	} else if q.command == "returnOrderBook" {
		resp, err := http.Get("https://poloniex.com/public?command=" + q.command +
			"&currencyPair=" + q.currencyPair)
		if err != nil {
			return "", err
		}

		ret, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		resp.Body.Close()

		return string(ret), nil
	} else if q.command == "returnTradeHistory" {
		resp, err := http.Get("https://poloniex.com/public?command=" + "returnTradeHistory" +
			"&currencyPair=" + q.currencyPair)
		if err != nil {
			return "", err
		}

		ret, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		resp.Body.Close()

		return string(ret), nil
	}
	return "", fmt.Errorf("Code not yet implementd\n")
}

func ReturnTicker() (string, error) {
	var q APIQuery
	q.command = "returnTicker"

	return sendQuery(q)
}

func Return24hVolume() (string, error) {
	var q APIQuery
	q.command = "return24hVolume"

	return sendQuery(q)
}

func ReturnOrderBook(currencyPair string) (string, error) {
	var q APIQuery
	q.command = "returnOrderBook"
	q.currencyPair = currencyPair

	return sendQuery(q)
}

func ReturnMarketTradeHistory(currencyPair string) (string, error) {
	var q APIQuery
	q.command = "returnMarketTradeHistory"
	q.currencyPair = currencyPair

	return sendQuery(q)
}

func main() {
	print, _ := ReturnTicker()
	fmt.Print(print)
}

