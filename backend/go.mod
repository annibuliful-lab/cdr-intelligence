module cdr-intelligence-backend

go 1.23.4

require github.com/redis/go-redis/v9 v9.7.0

require github.com/gorilla/websocket v1.4.1 // indirect

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/graph-gophers/graphql-go v1.5.1-0.20231220101041-a3a932cfa9a7
	github.com/graph-gophers/graphql-transport-ws v0.0.2
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/mununki/gqlmerge v0.2.15
	github.com/neo4j/neo4j-go-driver/v5 v5.27.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/rs/cors v1.11.1
)
