package main

import (
	"database/sql"

	"github.com/dghubble/sling"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v5"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/insomnius/example-cli-in-go/command"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()

	gin.SetMode(gin.ReleaseMode)
	apiEngine := gin.Default()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	serviceA := sling.New().Base(os.Getenv("SERVICE_A_HOST"))
	serviceB := sling.New().Base(os.Getenv("SERVICE_B_HOST")).SetBasicAuth(os.Getenv("SERVICE_B_BASIC_USERNAME"), os.Getenv("SERVICE_B_BASIC_PASSWORD"))

	rootCmd := command.NewRoot()
	rootCmd.AddCommand(
		command.NewRunAPI(apiEngine),
		command.NewPingServiceA(serviceA),
		command.NewPingServiceB(serviceB),
		command.NewPingDB(db),
		command.NewPingRedis(rdb),
		command.NewScriptAddToCart(),
	)

	_ = rootCmd.Execute()
}
