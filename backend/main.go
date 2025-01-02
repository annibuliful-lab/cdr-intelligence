package main

import (
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

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Debug:            true,
	})

	opts := []gql.SchemaOpt{
		gql.UseFieldResolvers(),
		gql.MaxParallelism(20),
		gql.UseStringDescriptions(),
		gql.RestrictIntrospection(func(context.Context) bool {
			return false
		}),
		gql.Directives(),
	}

	schema, err := gql.ParseSchema(string(mergedSchema[:]), graphql.GraphqlResolver(), opts...)
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
		// db.GetPrimaryClient().Close()
		// db.GetRedisClient().Close()

		close(idleConnectionsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-idleConnectionsClosed
}
