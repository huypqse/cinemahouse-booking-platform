// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cinema

import (
	"context"

	"cinemahouse-svc/api/cinema/cinema"
)

type ICinemaV1 interface {
	// Cinema management
	GetCinemas(ctx context.Context, req *cinema.CinemaGetCinemasReq) (res *cinema.CinemaGetCinemasRes, err error)
	GetCinema(ctx context.Context, req *cinema.CinemaGetCinemaReq) (res *cinema.CinemaGetCinemaRes, err error)
	CreateCinema(ctx context.Context, req *cinema.CinemaCreateCinemaReq) (res *cinema.CinemaCreateCinemaRes, err error)
	UpdateCinema(ctx context.Context, req *cinema.CinemaUpdateCinemaReq) (res *cinema.CinemaUpdateCinemaRes, err error)
	DeleteCinema(ctx context.Context, req *cinema.CinemaDeleteCinemaReq) (res *cinema.CinemaDeleteCinemaRes, err error)
	
	// Screening room management
	GetScreeningRooms(ctx context.Context, req *cinema.CinemaGetScreeningRoomsReq) (res *cinema.CinemaGetScreeningRoomsRes, err error)
	GetScreeningRoom(ctx context.Context, req *cinema.CinemaGetScreeningRoomReq) (res *cinema.CinemaGetScreeningRoomRes, err error)
	CreateScreeningRoom(ctx context.Context, req *cinema.CinemaCreateScreeningRoomReq) (res *cinema.CinemaCreateScreeningRoomRes, err error)
	UpdateScreeningRoom(ctx context.Context, req *cinema.CinemaUpdateScreeningRoomReq) (res *cinema.CinemaUpdateScreeningRoomRes, err error)
	DeleteScreeningRoom(ctx context.Context, req *cinema.CinemaDeleteScreeningRoomReq) (res *cinema.CinemaDeleteScreeningRoomRes, err error)
	
	// Room type management
	GetRoomTypes(ctx context.Context, req *cinema.CinemaGetRoomTypesReq) (res *cinema.CinemaGetRoomTypesRes, err error)
	CreateRoomType(ctx context.Context, req *cinema.CinemaCreateRoomTypeReq) (res *cinema.CinemaCreateRoomTypeRes, err error)
	UpdateRoomType(ctx context.Context, req *cinema.CinemaUpdateRoomTypeReq) (res *cinema.CinemaUpdateRoomTypeRes, err error)
	DeleteRoomType(ctx context.Context, req *cinema.CinemaDeleteRoomTypeReq) (res *cinema.CinemaDeleteRoomTypeRes, err error)
}
