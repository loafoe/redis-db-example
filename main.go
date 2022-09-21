package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/go-redis/redis/v8"
	_ "github.com/philips-software/gautocloud-connectors/hsdp"
)

func main() {
	var client *redis.ClusterClient

	err := gautocloud.InjectFromId("hsdp:redis-db", &client)

	if err != nil {
		fmt.Printf("Cannot find bound hsdp-redis-db instance\n")
		os.Exit(1)
	}

	res := client.Keys(context.Background(), "*")
	fmt.Printf("Keys: %+v\n", res)
}
