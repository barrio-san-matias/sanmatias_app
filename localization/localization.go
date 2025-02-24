package localization

import (
	"context"
	"log"
	"time"

	"github.com/joeshaw/envdecode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lotes map[int16]lote
var pois map[string]poi

type LatLng struct {
	Longitude float64 `bson:"longitude"`
	Latitude  float64 `bson:"latitude"`
}

type lote struct {
	Area   int16  `bson:"area"`
	Number int16  `bson:"number"`
	Coords LatLng `bson:"coords"`
}

type poi struct {
	Name   string `bson:"name"`
	Coords LatLng `bson:"coords"`
}

func init() {
	initMongoData()
}

func initMongoData() {
	ctx := context.Background()

	var cfg struct {
		MongoUrl string `env:"MONGODB_URI,required"`
	}

	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer client.Disconnect(ctx)

	db := client.Database("api")

	loadLotes(ctx, db.Collection("lotes"))
	loadPOIs(ctx, db.Collection("poi"))
}

func loadLotes(ctx context.Context, collection *mongo.Collection) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Error obteniendo lotes: %v", err)
	}
	defer cursor.Close(ctx)

	var list []lote
	if err := cursor.All(ctx, &list); err != nil {
		log.Fatalf("Error decodificando lotes: %v", err)
	}

	lotes = make(map[int16]lote, len(list))
	for _, l := range list {
		lotes[l.Number] = l
	}
}

func loadPOIs(ctx context.Context, collection *mongo.Collection) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Error obteniendo POIs: %v", err)
	}
	defer cursor.Close(ctx)

	var list []poi
	if err := cursor.All(ctx, &list); err != nil {
		log.Fatalf("Error decodificando POIs: %v", err)
	}

	pois = make(map[string]poi, len(list))
	for _, p := range list {
		pois[p.Name] = p
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
