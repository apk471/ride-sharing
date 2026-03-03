package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)



func main(){
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()

	fare := domain.RideFareModel{
		UserID: "42",
	}

	svc:= service.NewService(inmemRepo)

	t, err := svc.CreateTrip(ctx , &fare)

	if err != nil{
		log.Println(err)
	}

	log.Println(t)

	//TODO: Remove this
	for{
		time.Sleep(time.Second)
	}
}
