package main

import (
	"app"
	"app/pkg/MkpSc"
	nft_api "app/pkg/NftSc"
	cronJobs "app/pkg/cronJobs"
	"app/pkg/handler"
	"app/pkg/repository"
	"app/pkg/service"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title NFT Marketplace API
// @version 1.0
// @description API Server for NFT's marketplace.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	mkp, err := mkp_api.NewMkpApi(common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"), ethersClient)
	if err != nil {
		logrus.Fatalf("Failed to init ethers connection %s", err.Error())
	}

	nft, err := nft_api.NewNftApi(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), ethersClient)
	if err != nil {
		logrus.Fatalf("Failed to init ethers connection %s", err.Error())
	}

	repos := repository.NewRepository(db, mkp, nft)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(app.Server)
	cron_jobs := cronJobs.NewCronRunner(services)

	appCron := cron.New()
	appCron.AddFunc("@every 1m", func() {
		cron_jobs.RunCronJobs()
	})

	go func() {

		go cron_jobs.RunCronJobs()
		appCron.Start()
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
