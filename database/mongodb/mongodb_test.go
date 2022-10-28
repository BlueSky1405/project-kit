package mongodb

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestNewMongoConn(t *testing.T) {
	cfg := &Config{
		MaxPoolSize: 10,
		Username:    "root",
		Password:    "root",
		Database:    "j23_local",
		Addr:        []string{"43.138.178.12:27017"},
	}
	conn, f := NewMongoConn(cfg)
	defer f()

	doc := bson.D{
		{"created_at", time.Now().Unix()},
		{"report_id", "2022-12-32"},
	}
	coll := conn.DB().Collection("i_am_test")
	_, err := coll.InsertOne(context.Background(), doc)
	require.Nil(t, err)
}
