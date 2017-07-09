package poloniex

import (
	"net/http"
	"io/ioutil"
)

type APIQuery struct {
	command string
	currencyPair	string
}

func sendQuery(q APIQuery) (string, error) {
	if q.command == "returnTicker" || q.command == "return24Volume" {
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
		resp, err := http.Get("http://poloniex.com/public?command=" + q.command +
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
		resp, err := http.Get("http://poloniex.com/public?command=" + "returnTradeHistory" +
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
	return "", error("Code not yet implemented")
}

func returnTicker() (string, error) {
	var q APIQuery
	q.command = "returnTicker"

	return sendQuery(q)
}

