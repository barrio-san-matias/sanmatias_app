package handler

import (
	"encoding/json"
	"fmt"
	"github/jfatta/smbot/localization"
	"net/http"
	"strconv"
	"strings"
)

const drivingPatternGoogle = "https://www.google.com/maps/dir/?api=1&destination=%v,%v&travelmode=driving"
const drivingPatternWaze = "https://www.waze.com/ul?ll=%v,%v&navigate=yes&zoom=17"
const drivingPatternApple = "http://maps.apple.com/?daddr=%v,%v"

type MapResponse struct {
	Coords localization.LatLng
	MapURL string
}

// http://localhost:3000/api/map?lote=636
func MapHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, "no autorizado", http.StatusUnauthorized)
	return
}

func writeError(w http.ResponseWriter, errorMsg string, code int) {
	http.Error(w, errorMsg, code)
}

func writeResponse(w http.ResponseWriter, response any) {
	res, err := json.Marshal(response)
	if err != nil {
		writeError(w, "error al procesar respuesta", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
