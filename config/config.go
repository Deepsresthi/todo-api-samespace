package config

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

type ScyllaDBConfig struct {
	Hosts       []string
	Keyspace    string
	Consistency gocql.Consistency
}

func InitScyllaDB(config ScyllaDBConfig) {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = config.Consistency
	cluster.Timeout = 5 * time.Second

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to ScyllaDB:", err)
	}

	log.Println("Connected to ScyllaDB")
}

func CloseScyllaDB() {
	if Session != nil {
		Session.Close()
		log.Println("ScyllaDB connection closed")
	}
}
