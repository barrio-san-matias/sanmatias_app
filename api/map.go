package handler

import (
	"github/jfatta/smbot/localization"
	"net/http"
)

type MapResponse struct {
	Coords localization.LatLng
	MapURL string
}

// http://localhost:3000/api/map?lote=636
func MapHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, "no autorizado", http.StatusUnauthorized)
}

func writeError(w http.ResponseWriter, errorMsg string, code int) {
	http.Error(w, errorMsg, code)
}
