package buyer53

import (
	"github.com/mxmCherry/openrtb/native1"
	"github.com/mxmCherry/openrtb/native1/request"
)

type Request struct {
	Ver string `json:"ver,omitempty"`
	Layout native1.Layout `json:"layout,omitempty"`
	AdUnit native1.AdUnit `json:"adunit,omitempty"`
	PlcmtType native1.PlacementType `json:"plcmttype,omitempty"`
	PlcmtCnt int64 `json:"plcmtcnt,omitempty"`
	Seq int64 `json:"seq,omitempty"`
	Assets []request.Asset `json:"assets"`
	Privacy int8 `json:"privacy,omitempty"`
	Context native1.ContextType `json:"context,omitempty"`
	ContextSubType native1.ContextSubType `json:"contextsubtype,omitempty"`
	EventTrackers []request.EventTracker `json:"eventtrackers,omitempty"`

}