package buyer52

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestBidRequest(t *testing.T) {
	buf, err := ioutil.ReadFile("./testdata/bid_request_native.json")
	if err != nil {
		t.Fatal(err)
	}
	var b BidRequest
	if err := json.Unmarshal(buf, &b); err != nil {
		t.Fatal(err)
	}

	if len(b.Imp) == 0 {
		t.Errorf("wont: 2, got: %d", len(b.Imp))
	}

	for _, imp := range b.Imp {
		if imp.Native == nil {
			t.Errorf("wont: imp.Native not nil, got: imp.Native is nil")
		}

		if imp.Native.Request.PlcmtCnt != 1 {
			t.Errorf("wont: 1, got: %d", imp.Native.Request.PlcmtCnt)
		}

		if imp.Native.Request.PlcmtType != 2 {
			t.Errorf("wont: 2, got: %d", imp.Native.Request.PlcmtType)
		}

		if imp.Native.Request.Privacy != 1 {
			t.Errorf("wont: 1, got: %d", imp.Native.Request.Privacy)
		}

		if imp.Native.Request.Context != 1 {
			t.Errorf("wont: 1, got: %d", imp.Native.Request.Context)
		}

		if imp.Native.Request.ContextSubType != 12 {
			t.Errorf("wont: 12, got: %d", imp.Native.Request.ContextSubType)
		}

		if len(imp.Native.Request.Assets) != 4 {
			t.Errorf("wont: 4, got: %d", len(imp.Native.Request.Assets))
		}

		asset1 := imp.Native.Request.Assets[0]

	}
}
