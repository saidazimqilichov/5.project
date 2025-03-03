package connections

import (
	kafkaconsumer "book/brokers/consumer"
	hotelservice "book/clients/hotel"
	notification17 "book/clients/notification"
	userservice "book/clients/user"
	"book/config"
	"book/database/methods"
	interface17 "book/interface"
	interfaceservices "book/interface/services"
	"book/service/adjsut"
	grpcmethods "book/service/methods"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interface17.Booking {
	c := config.Configuration()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &methods.Database{Db: db}
}

func NewService() *interfaceservices.Database {
	a := NewDatabase()
	return &interfaceservices.Database{D: a}
}

func NewAdjust() interface17.BookingAdjust {
	a := NewService()
	user := userservice.UserClinet()
	hotel := hotelservice.Hotel()
	n := notification17.Hotel()
	return &adjsut.Adjust{S: a, User: user, Hotel: hotel, N: n}
}

func NewAdjus() *interfaceservices.AdjustDatabase {
	a := NewAdjust()
	return &interfaceservices.AdjustDatabase{A: a}
}

func NewGrpc() *grpcmethods.Grpc {
	a := NewAdjus()
	return &grpcmethods.Grpc{A: a}
}

func NewConsumer() *kafkaconsumer.Consumer17 {
	a := NewAdjus()
	ctx := context.Background()
	return &kafkaconsumer.Consumer17{A: a, Ctx: ctx}
}
