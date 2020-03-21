package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_articleHttpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	_articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	_articleUsecase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
	"github.com/bxcodec/go-clean-arch/middleware"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbConn := connectDatabase()

	e := echo.New()
	mdw := middleware.InitMiddleware()
	e.Use(mdw.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUsecase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDeliver.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func connectDatabase() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetInt(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	pgConnString := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbPort, dbHost, dbUser, dbPass, dbName)

	dbConn, err := sql.Open(`postgres`, pgConnString)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Database connection successful.")

	return dbConn
}