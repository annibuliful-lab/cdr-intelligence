package main

import (
	"cdr-intelligence-backend/src/config"
	uploadmiddleware "cdr-intelligence-backend/src/graphql/middleware/upload"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
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

	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.MaxParallelism(20),
		graphql.UseStringDescriptions(),
		graphql.RestrictIntrospection(func(context.Context) bool {
			return false
		}),
		graphql.Directives(),
	}

	type query struct{}

	// init graphQL schema
	schema, err := graphql.ParseSchema(string(mergedSchema[:]), &query{}, opts...)
	if err != nil {
		panic(err)
	}

	// graphQL handler
	graphQLHandler := corsMiddleware.Handler(
		uploadmiddleware.Handler(&relay.Handler{Schema: schema}),
		// graphqlws.NewHandlerFunc(
		// 	schema,
		// 	// auth.GraphqlContext(uploadmiddleware.Handler(&relay.Handler{Schema: schema})),
		// 	// graphqlws.WithContextGenerator(
		// 	// 	graphqlws.ContextGeneratorFunc(auth.WebsocketGraphqlContext),
		// 	// ),
		// ),
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
