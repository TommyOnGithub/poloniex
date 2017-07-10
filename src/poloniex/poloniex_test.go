package poloniex

import (
	"testing"
	"regexp"
)

func TextReturnTicker(t *testing.T) {
	wantRegexp := regexp.MustCompile("^\\{(\"\\w{7,10}\":\\{\"id\":[0-9]+,\"last\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"lowestAsk\":\"[0-9]+\\.[0-9]{8}\",\"highestBid\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"percentChange\":\"-?[0-9]+\\.[0-9]{8}\",\"baseVolume\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"quoteVolume\":\"[0-9]+\\.[0-9]{8}\",\"isFrozen\":\"[01]\",\"high24hr\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"low24hr\":\"[0-9]+\\.[0-9]{8}\"\\},)+\"\\w{7,10}\":\\{\"id\":[0-9]+,\"last\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"lowestAsk\":\"[0-9]+\\.[0-9]{8}\",\"highestBid\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"percentChange\":\"-?[0-9]+\\.[0-9]{8}\",\"baseVolume\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"quoteVolume\":\"[0-9]+\\.[0-9]{8}\",\"isFrozen\":\"[01]\",\"high24hr\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"low24hr\":\"[0-9]+\\.[0-9]{8}\"\\}}")
	got, err := ReturnTicker()
	if err != nil {
		t.Error(err)
	} else {
		matched, _ := regexp.MatchString(wantRegexp.String(), got)
		if !matched {
			t.Fatalf("HTTPS response does not match expected; got: %s\n", got)
		}
	}
}

func TestReturn24hVolume(t *testing.T) {
	wantRegexp := regexp.MustCompile("^\\{(\\w{7,10}\":\\{\"[A-Z]{3,5}\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"[A-Z]{3,5}\":\"[0-9]+\\.[0-9]{8}\",\\},)+\"totalBTC\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"totalETH\":\"[0-9]+\\.[0-9]{8}\",\"totalUSDT\":\"[0-9]+\\.[0-9]{8}\",\"totalXMR\":\"[0-9]+\\.[0-9]{8}\"," +
		"\"totalXUSD\":\"[0-9]+\\.[0-9]{8}\"\\}")
	got, err := Return24hVolume()
	if err != nil {
		t.Error(err)
	} else {
		matched, _ := regexp.MatchString(wantRegexp.String(), got)
		if !matched {
			t.Fatalf("HTTPS response does not match expected; got: %s\n", got)
		}
	}
}

func TestReturnOrderBook(t *testing.T) {
	wantRegexp := regexp.MustCompile("\\{\"asks\":\\[(\\[\"[0-9]+\\.[0-9]{8}\",[0-9]+\\.?[0-9\\We]{0,8}\\]," +
		")+\\[\"[0-9]+\\.[0-9]{8}\",[0-9]+\\.[0-9]{8}\\]\\],\"isFrozen\":\"0\",\"seq\":[0-9]{9}\\}")
	got, err := ReturnOrderBook("USDT_ETH")
	if err != nil {
		t.Error(err)
	} else {
		matched, _ := regexp.MatchString(wantRegexp.String(), got)
		if !matched {
			t.Fatalf("HTTPS response does not match expected; got: %s\n", got)
		}
	}
}