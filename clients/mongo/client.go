package gwm_mongo

import (
	"fmt"
	"log/slog"
	"net/url"
	"sync"
	"time"

	"github.com/go-logr/logr"
	gwm_app "github.com/slhmy/go-webmods/app"
	"github.com/slhmy/go-webmods/internal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	host               string
	username           string
	password           string
	database           string
	slowQueryThreshold time.Duration
	enableSeedList     bool
)

func init() {
	host = gwm_app.Config().GetString(internal.ConfigKeyMongoHost)
	username = gwm_app.Config().GetString(internal.ConfigKeyMongoUsername)
	password = gwm_app.Config().GetString(internal.ConfigKeyMongoPassword)
	database = gwm_app.Config().GetString(internal.ConfigKeyMongoDatabase)
	slowQueryThreshold = gwm_app.Config().GetDuration(internal.ConfigKeyMongoSlowQueryThreshold)
	if slowQueryThreshold == 0 {
		slowQueryThreshold = 100 * time.Millisecond
	}
	enableSeedList = gwm_app.Config().GetBool(internal.ConfigKeyMongoEnableSeedList)
}

var (
	initMutex sync.Mutex
	client    *mongo.Client
)

func GetClient() *mongo.Client {
	if client == nil {
		initMutex.Lock()
		defer initMutex.Unlock()
		if client != nil {
			return client
		}

		if len(username) == 0 || len(password) == 0 {
			host = fmt.Sprintf("mongodb://%s", host)
		} else {
			if enableSeedList {
				println("mongo seedlist enabled")
				host = fmt.Sprintf("mongodb+srv://%s:%s@%s", username, url.PathEscape(password), host)
			} else {
				host = fmt.Sprintf("mongodb://%s:%s@%s", username, url.PathEscape(password), host)
			}
		}

		clientOptions := options.Client().ApplyURI(host).SetBSONOptions(&options.BSONOptions{
			ObjectIDAsHexString: true,
		})
		registry := bson.NewRegistry()
		gwm_app.RegisterIDToBsonRegistry(registry)
		clientOptions.SetRegistry(registry)
		sink := logr.FromSlogHandler(slog.Default().Handler()).GetSink()
		loggerOptions := options.Logger().SetSink(sink).SetMaxDocumentLength(25).
			SetComponentLevel(options.LogComponentCommand, options.LogLevelInfo)
		clientOptions.SetLoggerOptions(loggerOptions)
		monitor := NewMongoCommandMonitor()
		clientOptions.SetMonitor(monitor)

		var err error
		client, err = mongo.Connect(clientOptions)
		if err != nil {
			panic(err)
		}
	}
	return client
}

func GetDatabase() *mongo.Database {
	client := GetClient()
	return client.Database(database)
}

func GetCollection(
	collectionName string,
) *mongo.Collection {
	client := GetClient()
	return client.Database(database).Collection(collectionName)
}
