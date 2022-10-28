package mongodb

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// username, pass, ip, port
const uriTemplate = "mongodb://%s:%s@%s:%d/?maxPoolSize=20&w=majority"

type Config struct {
	// MaxPoolSize 最大连接数
	MaxPoolSize uint64   `json:"max_pool_size"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Database    string   `json:"database"`
	Addr        []string `json:"addr"`
}

type Conn struct {
	cfg *Config
	cli *mongo.Client
	db  *mongo.Database
}

func (c *Conn) DB() *mongo.Database {
	return c.db
}

func NewMongoConn(cfg *Config) (*Conn, func()) {
	ctx := context.Background()

	option := options.Client()
	option.Hosts = cfg.Addr
	option.Auth = &options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	option.SetMaxPoolSize(cfg.MaxPoolSize)
	//option.SetReplicaSet("myRS")

	cli, err := mongo.Connect(ctx, option)
	if err != nil {
		panic(any(err))
	}
	if err := cli.Ping(ctx, readpref.Primary()); err != nil {
		panic(any(errors.Wrapf(err, "mongodb ping fail")))
	}

	db := cli.Database(cfg.Database)

	closeFunc := func() {
		if err := cli.Disconnect(context.Background()); err != nil {
			panic(any(err))
		}
	}

	return &Conn{
		cfg: cfg,
		cli: cli,
		db:  db,
	}, closeFunc
}
