package setting

import (
	"encoding/json"
	"github.com/shieldnet/gobit/strategy"
	"io/ioutil"
	"log"
	"os"
)

// Gloabl Variables
const CinfoFileName = "result.json"


var Cinfo map[string]CandleSet = nil

// Implementation
type CandleSet struct {
	Bc int `json:"bc"`
	Sc int `json:"sc"`
}

func LoadTradeCandleInfo() {
	if Cinfo == nil {
		Cinfo = map[string]CandleSet{}
	}

	jsonFile, err := os.Open(CinfoFileName)
	if err != nil {
		log.Panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &Cinfo)
	if err != nil {
		log.Panic(string(byteValue), err)
	}
}

func GetTradeCandleInfo() map[string]CandleSet {
	if Cinfo != nil {
		LoadTradeCandleInfo()
	}
	return Cinfo
}

func RefreshCandleInfo(strategies *[]*strategy.Strategy) {
	LoadTradeCandleInfo()
	for _, st := range *strategies {
		st.BuyCandleNum = Cinfo[st.Market].Bc
		st.SellCandleNum = Cinfo[st.Market].Sc
	}
}
