package handler

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github/jfatta/smbot/localization"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/joeshaw/envdecode"
	redis "github.com/redis/go-redis/v9"
)

const drivingPatternGoogle = "https://www.google.com/maps/dir/?api=1&destination=%v,%v&travelmode=driving"
const drivingPatternWaze = "https://www.waze.com/ul?ll=%v,%v&navigate=yes&zoom=17"
const drivingPatternApple = "http://maps.apple.com/?daddr=%v,%v"

var once sync.Once
var redisClient *redis.Client

type MapResponse struct {
	Coords localization.LatLng
	MapURL string
}

type Event struct {
	Id         string    `redis:"id"`
	Lote       string    `redis:"lote"`
	POI        string    `redis:"poi"`
	MapType    string    `redis:"map_type"`
	CreateTime time.Time `redis:"create_time"`
}

// http://localhost:3000/api/map?lote=636
func MapHandler(w http.ResponseWriter, r *http.Request) {
	loteParam := r.URL.Query().Get("lote")
	poiParam := r.URL.Query().Get("poi")
	mapType := r.URL.Query().Get("map-type")

	trace(r.Context(), loteParam, poiParam, mapType)

	if loteParam == "" && poiParam == "" {
		writeError(w, "parametro lote o poi es obligatorio", http.StatusBadRequest)
		return
	}

	response := &MapResponse{}

	var coords localization.LatLng

	if poiParam != "" {
		coords = localization.GetPOICoords(poiParam)
		if coords == (localization.LatLng{}) {
			writeError(w, "punto de interes no encontrado", http.StatusNotFound)
			return
		}
	}

	if loteParam != "" {
		numLote, err := strconv.ParseInt(loteParam, 10, 16)
		if err != nil {
			writeError(w, "parametro lote debe ser un numero valido", http.StatusBadRequest)
			return
		}

		coords = localization.GetCoords(int16(numLote))
		if coords == (localization.LatLng{}) {
			writeError(w, "lote no encontrado", http.StatusNotFound)
			return
		}

	}

	response.Coords = coords
	response.MapURL = fmt.Sprintf(drivingPatternGoogle, coords.Latitude, coords.Longitude)

	response.Coords = coords
	switch strings.ToLower(mapType) {
	case "waze":
		response.MapURL = fmt.Sprintf(drivingPatternWaze, coords.Latitude, coords.Longitude)
	case "apple":
		response.MapURL = fmt.Sprintf(drivingPatternApple, coords.Latitude, coords.Longitude)
	default:
		response.MapURL = fmt.Sprintf(drivingPatternGoogle, coords.Latitude, coords.Longitude)
	}

	writeResponse(w, response)
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

func getRedisClient() *redis.Client {
	once.Do(func() {
		var cfg struct {
			KVUrl string `env:"KV_URL,required"`
		}

		if err := envdecode.StrictDecode(&cfg); err != nil {
			log.Fatal(err)
			return
		}

		opt, err := redis.ParseURL(cfg.KVUrl)
		if err != nil {
			log.Fatal(err)
			return
		}

		opt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		redisClient = redis.NewClient(opt)
	})

	return redisClient
}

func trace(ctx context.Context, lote string, poi string, mapType string) {
	c := getRedisClient()
	e := &Event{
		Id:         uuid.New().String(),
		Lote:       lote,
		POI:        poi,
		MapType:    mapType,
		CreateTime: time.Now().UTC(),
	}
	if c != nil {

		c.HSet(ctx, e.Id, e)
		return
	}

	log.Printf("couldn't send to redis: +%v", e)

}
