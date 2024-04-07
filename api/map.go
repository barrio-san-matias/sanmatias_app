package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github/jfatta/smbot/localization"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/joeshaw/envdecode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const drivingPatternGoogle = "https://www.google.com/maps/dir/?api=1&destination=%v,%v&travelmode=driving"
const drivingPatternWaze = "https://www.waze.com/ul?ll=%v,%v&navigate=yes&zoom=17"
const drivingPatternApple = "http://maps.apple.com/?daddr=%v,%v"

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

	trace(r.Context(), loteParam, poiParam, mapType, coords)
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

func getMongoClient(ctx context.Context) (*mongo.Client, error) {
	var cfg struct {
		MongoUrl string `env:"MONGODB_URI,required"`
	}

	if err := envdecode.StrictDecode(&cfg); err != nil {
		return nil, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func trace(ctx context.Context, lote string, poi string, mapType string, ll localization.LatLng) {
	c, err := getMongoClient(ctx)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	defer c.Disconnect(ctx)
	apiDB := c.Database("api")
	events := apiDB.Collection("events")

	loc := lote
	if loc == "" {
		loc = poi
	}
	_, err = events.InsertOne(ctx, bson.D{
		{Key: "loc", Value: loc},
		{Key: "map_type", Value: mapType},
		{Key: "create_time", Value: time.Now().UTC()},
		{Key: "geopoint", Value: bson.M{
			"type":        "Point",
			"coordinates": []float64{ll.Longitude, ll.Latitude},
		}},
	})
	if err != nil {
		log.Fatalf("couldn't send to mongoDB: %v", err)
	}
}
