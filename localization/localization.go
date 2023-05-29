package localization

import (
	_ "embed"
	"encoding/json"
)

//go:embed lotes.json
var lotesData []byte

//go:embed poi.json
var poisData []byte

var lotes map[int16]lote
var pois map[string]poi

type LatLng struct {
	Longitude float64
	Latitude  float64
}

type lote struct {
	Area   int16
	Number int16
	Coords LatLng
}

type poi struct {
	Name   string
	Coords LatLng
}

func init() {
	initLotes()
	initPOIs()
}

func initLotes() {
	var list []lote
	if err := json.Unmarshal(lotesData, &list); err != nil {
		panic(err)
	}

	lotes = make(map[int16]lote, len(list))
	for _, l := range list {
		lotes[l.Number] = l
	}
}

func initPOIs() {
	var list []poi
	if err := json.Unmarshal(poisData, &list); err != nil {
		panic(err)
	}

	pois = make(map[string]poi, len(list))
	for _, poi := range list {
		pois[poi.Name] = poi
	}
}

func GetCoords(num int16) LatLng {
	if lote, ok := lotes[num]; ok {
		return lote.Coords
	}

	return LatLng{}
}

func GetPOICoords(poiName string) LatLng {
	if poi, ok := pois[poiName]; ok {
		return poi.Coords
	}

	return LatLng{}
}
