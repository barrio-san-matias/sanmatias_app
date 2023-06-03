package main

import (
	"context"
	"crypto/tls"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	kvurl := ""

	opt, err := redis.ParseURL(kvurl)
	if err != nil {
		panic(err)
	}

	opt.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		//Certificates: []tls.Certificate{cert}
	}

	rdb := redis.NewClient(opt)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

}
