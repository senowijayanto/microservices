package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
	"time"

	_userController "user/controller"
	_userRepo "user/repository"
	_userService "user/service"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// Setup database
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Setup Echo
	e := echo.New()

	// Setup middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//middleware := _userMiddleware.InitMiddleware()
	//e.Use(middleware.CORS)

	// Setup User Repository
	userRepo := _userRepo.NewUserRepository(db)

	// Setup User Service
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userService := _userService.NewUserService(userRepo, timeoutContext)

	// Setup User Controller
	_userController.NewUserController(e, userService)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
