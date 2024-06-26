package containers

import (
	"context"
	"log"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:	"redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:	wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started: true,
	})

	if err != nil {
		log.Fatalf("Could not start redis %s", err)
	}

	endpoint, err := redisC.Endpoint(ctx, "")
	log.Printf("The endpoint for redis is %s", endpoint)
	if err != nil {
		t.Error(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	err = client.Set(ctx, "name", "kato", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	log.Printf("The value for name is: %s", val)

	defer func() {
		if err:= redisC.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop redis %s", err)
		}
	}()
}
