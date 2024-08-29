package connections

import (
	"database/sql"
	"fmt"
	"hotel/config"
	"hotel/databases/methods"
	interface17 "hotel/interface"
	"hotel/interface/services"
	adjustservice "hotel/service/adjust"
	grpcmethod "hotel/service/method"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interface17.Hotel {
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

func NewService() *services.Database {
	a := NewDatabase()
	return &services.Database{S: a}
}

func NewAdjust() interface17.Adjust {
	a := NewService()
	return &adjustservice.Adjust{S: a}
}

func NewAdjustService() *services.Adjust {
	a := NewAdjust()
	return &services.Adjust{A: a}
}

func NewGrpc() *grpcmethod.GrpcService {
	a := NewAdjustService()
	return &grpcmethod.GrpcService{A: a}
}
