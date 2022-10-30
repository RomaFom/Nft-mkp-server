package main

import (
	"app"
	"app/pkg/handler"
	"app/pkg/repository"
	"app/pkg/service"
	"app/pkg/smart-contracts"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	confErr := initConfig()
	if confErr != nil {
		logrus.Fatalf("Error with configs %s", confErr.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error with envs %s", err.Error())
	}

	db, dbErr := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PWD"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	if dbErr != nil {
		logrus.Fatalf("Failed to init db %s", dbErr.Error())
	}

	ethersClient, etherErr := ethclient.Dial("http://127.0.0.1:8545")
	if etherErr != nil {
		logrus.Fatalf("Failed to init ethers connection %s", etherErr.Error())
	}

	mkp, err := sc_api.NewScApi(common.HexToAddress("0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0"), ethersClient)
	if err != nil {
		panic(err)
	}

	//test, err := mkp.ItemCount(&bind.CallOpts{})
	//fmt.Println(test)

	repos := repository.NewRepository(db, mkp)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(app.Server)
	go func() {
		err := server.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("Error running server %s", err.Error())
		}
		server.EnableCors()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("Error when server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error when database shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
