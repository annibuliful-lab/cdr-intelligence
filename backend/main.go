package main

import (
	"backend/src/clients"
	"backend/src/config"
	"backend/src/graphql"
	uploadmiddleware "backend/src/graphql/middleware/upload"
	"backend/src/graphql/subscription/graphqlws"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	gql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	gqlMerge "github.com/mununki/gqlmerge/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/cors"
)

func mergeGql() {
	schema := gqlMerge.Merge(" ", "./src")
	err := os.WriteFile("./generated.graphql", []byte(*schema), 0777)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mergeGql()

	mergedSchema, err := os.ReadFile("generated.graphql")

	if err != nil {
		log.Fatal("Error loading graphql file")
	}

	isDevelopment := config.GetEnv("ENV", "development") == "development"

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            isDevelopment,
	})

	opts := []gql.SchemaOpt{
		gql.UseFieldResolvers(),
		gql.MaxParallelism(20),
		gql.UseStringDescriptions(),
		gql.RestrictIntrospection(func(context.Context) bool {
			return isDevelopment
		}),
		gql.Directives(),
	}

	db, err := clients.NewPostgreSQLClient()
	if err != nil {
		panic(err)
	}
	redis, err := clients.NewRedisClient()
	if err != nil {
		panic(err)
	}

	rabbitmq, err := clients.NewRabbitMQClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	neo4jDriver, err := clients.NewNeo4jClient()
	neo4jSession := neo4jDriver.NewSession(ctx, neo4j.SessionConfig{})
	if err != nil {
		panic(err)
	}

	schema, err := gql.ParseSchema(string(mergedSchema[:]), graphql.GraphqlResolver(graphql.GraphqlResolverParams{
		Db:       db,
		Redis:    redis,
		Rabbitmq: rabbitmq,
		Neo4j:    &neo4jSession,
	}), opts...)

	if err != nil {
		panic(err)
	}

	// graphQL handler
	graphQLHandler := corsMiddleware.Handler(
		graphqlws.NewHandlerFunc(
			schema,
			// auth.GraphqlContext(uploadmiddleware.Handler(&relay.Handler{Schema: schema})),
			uploadmiddleware.Handler(&relay.Handler{Schema: schema}),
		),
	)

	http.Handle("/graphql", graphQLHandler)

	var listenAddress = flag.String("listen", config.GetEnv("BACKEND_PORT", ":3000"), "Listen address.")

	log.Printf("Listening at http://%s", *listenAddress)

	httpServer := http.Server{
		Addr: *listenAddress,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		db.Close()
		redis.Close()
		rabbitmq.Close()
		neo4jSession.Close(ctx)
		neo4jDriver.Close(ctx)

		close(idleConnectionsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-idleConnectionsClosed
}
