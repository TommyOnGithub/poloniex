package poloniex

import (
	"testing"
	"regexp"
)

func ReturnTickerTest(t *testing.T) {
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
			t.Fatalf("HTTPS response does not match expected; got: %s", got)
		}
	}
}

func Return24hVolumeTest(t *testing.T) {}
