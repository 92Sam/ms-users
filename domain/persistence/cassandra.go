package persistence

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gocql/gocql"
)

type CassandraDbContext struct {
	*gocql.Session
}

type CassandraTable int

// TABLES Enum
const (
	CQL_USERS CassandraTable = iota
	CQL_PRODUCTS
)

func (tableIndx *CassandraTable) GetTableNamePtr() *string {
	return aws.String([]string{
		"users",
		"products",
	}[*tableIndx])
}

func InitCassandraDb() *CassandraDbContext {
	// Init Cluster Session Cassandra
	cluster := gocql.NewCluster(os.Getenv("DB_CASSANDRA_HOST"))
	cluster.Keyspace = os.Getenv("DB_CASSANDRA_NAME")
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("unable to load Cassandra config, %v", err)
	}

	defer session.Close()

	return &CassandraDbContext{session}
}
