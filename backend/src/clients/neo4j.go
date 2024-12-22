package clients

import (
	"cdr-intelligence-backend/src/config"
	"sync"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	neo4jDriver neo4j.DriverWithContext
	neo4jOnce   sync.Once
)

// NewNeo4jClient ensures thread-safe initialization of the Neo4j driver.
func NewNeo4jClient() (neo4j.DriverWithContext, error) {
	var err error
	neo4jOnce.Do(func() {
		neo4jDriver, err = neo4j.NewDriverWithContext(
			config.GetEnv("NEO4J_URI", "bolt://localhost:7687"),
			neo4j.BasicAuth(
				config.GetEnv("NEO4J_USERNAME", "neo4j"),
				config.GetEnv("NEO4J_PASSWORD", "test"),
				"",
			),
		)
	})

	if err != nil {
		return nil, err
	}
	return neo4jDriver, nil
}
