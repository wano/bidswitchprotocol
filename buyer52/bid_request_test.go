package buyer52

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestBidRequest_Other(t *testing.T) {
	fileNames := []string{
		"./testdata/bid_request_banner.json",
		"./testdata/bid_request_video.json",
		"./testdata/bid_request_inapp.json",
		"./testdata/bid_request_private_deal.json",
		"./testdata/bid_request_audio.json",
		"./testdata/bid_request_tv.json",
		"./testdata/bid_request_dooh.json",

	}
	for i, fileName := range fileNames {
		buf, err := ioutil.ReadFile(fileName)
		if err != nil {
			t.Errorf("file_num: %d, file_name: %s, error: %#v", i, fileName, err)
		}
		//unmarshal only
		var b BidRequest
		if err := json.Unmarshal(buf, &b); err != nil {
			t.Errorf("file_num: %d, file_name: %s, error: %#v", i, fileName, err)
		}
	}
}

func TestBidRequest_Native(t *testing.T) {
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

		if len(imp.Native.Request.Assets) != 6 {
			t.Errorf("wont: 6, got: %d", len(imp.Native.Request.Assets))
		}

		asset1 := imp.Native.Request.Assets[0]
		if asset1.ID != 1 {
			t.Errorf("wont: 1, got: %d", asset1.ID)
		}

		if asset1.Data == nil {
			t.Errorf("wont: not nil, got: nil")
		}

		if asset1.Data.Type != 12 {
			t.Errorf("wont: 12, got: %d", asset1.Data.Type)
		}

		if asset1.Required != 1 {
			t.Errorf("wont: 1, got: %d", asset1.Required)
		}

		asset2 := imp.Native.Request.Assets[1]
		if asset2.ID != 2 {
			t.Errorf("wont: 2, got: %d", asset2.ID)
		}

		if asset2.Title == nil {
			t.Errorf("wont: not nil, got: nil")
		}

		if asset2.Title.Len != 50 {
			t.Errorf("wont: 50, got: %d", asset2.Title.Len)
		}

		if asset2.Required != 1 {
			t.Errorf("wont: 1, got: %d", asset2.Required)
		}

		asset3 := imp.Native.Request.Assets[2]
		if asset3.ID != 3 {
			t.Errorf("wont: 3, got: %d", asset3.ID)
		}

		if asset3.Img == nil {
			t.Errorf("wont: not nil, got: nil")
		}

		if asset3.Img.W != 80 {
			t.Errorf("wont: 80, got: %d", asset3.Img.W)
		}

		if asset3.Img.H != 80 {
			t.Errorf("wont: 80, got: %d", asset3.Img.H)
		}

		if asset3.Img.Type != 1 {
			t.Errorf("wont: 1, got: %d", asset3.Img.Type)
		}

		if asset3.Required != 1 {
			t.Errorf("wont: 1, got: %d", asset3.Required)
		}

		asset4 := imp.Native.Request.Assets[3]
		if asset4.ID != 4 {
			t.Errorf("wont: 4, got: %d", asset4.ID)
		}

		if asset4.Img == nil {
			t.Errorf("wont: not nil, got nil")
		}

		if asset4.Img.W != 1200 {
			t.Errorf("wont: 1200, got: %d", asset4.Img.W)
		}

		if asset4.Img.H != 627 {
			t.Errorf("wont: 627, got: %d", asset4.Img.H)
		}

		if asset4.Img.Type != 3 {
			t.Errorf("wont: 3, got: %d", asset4.Img.Type)
		}

		if asset4.Required != 1 {
			t.Errorf("wont: 1, got: %d", asset4.Required)
		}

		asset5 := imp.Native.Request.Assets[4]
		if asset5.ID != 5 {
			t.Errorf("wont: 5, got: %d", asset5.ID)
		}

		if asset5.Data == nil {
			t.Errorf("wont: not nil, got: nil")
		}

		if asset5.Data.Type != 3 {
			t.Errorf("wont: 3, got: %d", asset5.Data.Type)
		}

		if asset5.Required != 0 {
			t.Errorf("wont: 0, got: %d", asset5.Required)
		}

		asset6 := imp.Native.Request.Assets[5]
		if asset6.ID != 6 {
			t.Errorf("wont: 6, got: %d", asset6.ID)
		}

		if asset6.Data == nil {
			t.Errorf("wont: not nil, got: nil")
		}

		if asset6.Data.Len != 100 {
			t.Errorf("wont: 100, got: %d", asset6.Data.Len)
		}

		if asset6.Data.Type != 2 {
			t.Errorf("wont: 2, got: %d", asset6.Data.Type)
		}

		if asset6.Required != 1 {
			t.Errorf("wont: 1, got: %d", asset6.Required)
		}

	}
}
