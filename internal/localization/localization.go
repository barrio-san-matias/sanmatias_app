package localization

import (
	_ "embed"
	"encoding/json"
)

//go:embed lotes.json
var lotesData []byte

var lotes map[int16]lote

type LatLng struct {
	Longitude float64
	Latitude  float64
}

type lote struct {
	Area   int16
	Number int16
	Coords LatLng
}

func init() {
	var list []lote
	if err := json.Unmarshal(lotesData, &list); err != nil {
		panic(err)
	}

	lotes = make(map[int16]lote, len(list))
	for _, l := range list {
		lotes[l.Number] = l
	}
}

func GetCoords(num int16) LatLng {
	if lote, ok := lotes[num]; ok {
		return lote.Coords
	}

	return LatLng{}
}
