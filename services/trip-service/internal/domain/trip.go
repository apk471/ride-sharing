package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type TripModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"userID"`
	Status   string             `bson:"status"`
	RideFare *RideFareModel     `bson:"rideFare"`
	// Driver   *pb.TripDriver     `bson:"driver"`
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
	// SaveRideFare(ctx context.Context, f *RideFareModel) error
	// GetRideFareByID(ctx context.Context, id string) (*RideFareModel, error)
	// GetTripByID(ctx context.Context, id string) (*TripModel, error)
	// UpdateTrip(ctx context.Context, tripID string, status string, driver *pbd.Driver) error
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	// GetRoute(ctx context.Context, pickup, destination *types.Coordinate, useOsrmApi bool) (*tripTypes.OsrmApiResponse, error)
	// EstimatePackagesPriceWithRoute(route *tripTypes.OsrmApiResponse) []*RideFareModel
	// GenerateTripFares(
	// 	ctx context.Context,
	// 	fares []*RideFareModel,
	// 	userID string,
	// 	Route *tripTypes.OsrmApiResponse,
	// ) ([]*RideFareModel, error)
	// GetAndValidateFare(ctx context.Context, fareID, userID string) (*RideFareModel, error)
	// GetTripByID(ctx context.Context, id string) (*TripModel, error)
	// UpdateTrip(ctx context.Context, tripID string, status string, driver *pbd.Driver) error
}