package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_orderController "order/controller"
	_orderRepo "order/repository"
	_orderService "order/service"
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
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Setup Echo
	e := echo.New()

	// Setup middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup Order Repository
	orderRepo := _orderRepo.NewOrderRepository(dbConn)

	// Setup Order Service
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	orderService := _orderService.NewOrderService(orderRepo, timeoutContext)

	// Setup Order Controller
	_orderController.NewOrderController(e, orderService)

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	log.Fatal(e.Start(viper.GetString("server.address")))

	// TODO
	//fmt.Println("Order Service")
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//defer conn.Close()
	//
	//fmt.Println("Successfully connected!")
	//
	//ch, err := conn.Channel()
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//defer ch.Close()
	//
	//q, err := ch.QueueDeclare("OrderQueue", false, false, false, false, nil)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println(q)
	//
	//type data struct {
	//	id int
	//	qty int
	//}
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//err = ch.Publish("", "OrderQueue", false, false,
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body: []byte("Order Queue"),
	//	})
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println("Successfully Published Message to Queue")
}