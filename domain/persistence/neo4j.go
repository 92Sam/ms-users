package persistence

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jDbContext struct {
	Driver neo4j.Driver
	neo4j.Session
}

type Neo4jTag int
type Neo4jRelationTag int

// TABLES Enum
const (
	PRODUCTS Neo4jTag = iota
)

func (tagIndx Neo4jTag) GetTableNamePtr() string {
	return []string{
		"Product",
	}[tagIndx]
}

func InitNeo4jDb() *Neo4jDbContext {
	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	dbUri := os.Getenv("DB_NEO4J_HOST")
	authCredentials := neo4j.NoAuth()
	if os.Getenv("DB_NEO4J_USERNAME") != "" {
		authCredentials = neo4j.BasicAuth(
			os.Getenv("DB_NEO4J_USERNAME"),
			os.Getenv("DB_NEO4J_PASSWORD"),
			"",
		)
	}

	driver, err := neo4j.NewDriver(dbUri, authCredentials)
	if err != nil {
		panic(err)
	}
	// defer driver.Close()

	err = driver.VerifyConnectivity()
	if err != nil {
		panic(err)
	}
	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application

	neo4jConfig := neo4j.SessionConfig{
		DatabaseName: os.Getenv("DB_NEO4J_DATABASE"),
	}

	session := driver.NewSession(neo4jConfig)
	x := &Neo4jDbContext{
		driver,
		session,
	}

	return x
}

// func insertItem(driverSession neo4j.Session) (*Item, error) {
// 	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
// 	// per request in your web application. Make sure to call Close on the session when done.
// 	// For multi-database support, set sessionConfig.DatabaseName to requested database
// 	// Session config will default to write mode, if only reads are to be used configure session for
// 	// read mode.

// 	session := driverSession
// 	defer session.Close()
// 	records, err := session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
// 		"id":   1,
// 		"name": "Item 1",
// 	})
// 	// In face of driver native errors, make sure to return them directly.
// 	// Depending on the error, the driver may try to execute the function again.
// 	if err != nil {
// 		return nil, err
// 	}
// 	record, err := records.Single()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
// 	return &Item{
// 		Id:   record.Values[0].(int64),
// 		Name: record.Values[1].(string),
// 	}, nil
// }

// func insertItem2(driverSession neo4j.Session) (*Item, error) {
// 	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
// 	// per request in your web application. Make sure to call Close on the session when done.
// 	// For multi-database support, set sessionConfig.DatabaseName to requested database
// 	// Session config will default to write mode, if only reads are to be used configure session for
// 	// read mode.

// 	session := driverSession
// 	defer session.Close()
// 	records, err := session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
// 		"id":   1,
// 		"name": "Item 1",
// 	})
// 	// In face of driver native errors, make sure to return them directly.
// 	// Depending on the error, the driver may try to execute the function again.
// 	if err != nil {
// 		return nil, err
// 	}
// 	record, err := records.Single()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
// 	return &Item{
// 		Id:   record.Values[0].(int64),
// 		Name: record.Values[1].(string),
// 	}, nil
// }

// type Item struct {
// 	Id   int64
// 	Name string
// }
