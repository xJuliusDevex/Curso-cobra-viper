package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cfg Config
)

func init() {
	LoadConfig()
}
func GetConnection() (client *mongo.Client, err error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=Cluster0", cfg.DB_user, cfg.DB_password, cfg.DB_host, cfg.DB_name)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, err
}
func GetCollection(name string) (col *mongo.Collection, err error) {
	client, err := GetConnection()
	if err != nil {
		return nil, err
	}
	return client.Database("task-admi").Collection(name), nil
}

type Config struct {
	DB_name     string `mapstructure: "DB_NAME"`
	DB_host     string `mapstructure: "DB_HOST"`
	DB_user     string `mapstructure: "DB_USER"`
	DB_password string `mapstructure: "DB_PASSWORD"`
}

func LoadConfig() {
	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(path)
	v1 := viper.New()
	v1.AddConfigPath(path + "/conec")
	v1.SetConfigType("env")
	v1.SetConfigName("app")
	v1.AutomaticEnv()
	err = v1.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("ReadConfig: ", v1.ConfigFileUsed())
	}
	err = v1.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(cfg)
}
