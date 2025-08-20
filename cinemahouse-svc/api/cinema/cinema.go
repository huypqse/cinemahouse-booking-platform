// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cinema

import (
	"context"

	"cinemahouse-svc/api/cinema/cinema"
)

type ICinemaCinema interface {
	CinemaGetCinemas(ctx context.Context, req *cinema.CinemaGetCinemasReq) (res *cinema.CinemaGetCinemasRes, err error)
	CinemaGetCinema(ctx context.Context, req *cinema.CinemaGetCinemaReq) (res *cinema.CinemaGetCinemaRes, err error)
	CinemaCreateCinema(ctx context.Context, req *cinema.CinemaCreateCinemaReq) (res *cinema.CinemaCreateCinemaRes, err error)
	CinemaUpdateCinema(ctx context.Context, req *cinema.CinemaUpdateCinemaReq) (res *cinema.CinemaUpdateCinemaRes, err error)
	CinemaDeleteCinema(ctx context.Context, req *cinema.CinemaDeleteCinemaReq) (res *cinema.CinemaDeleteCinemaRes, err error)
	CinemaGetRoomTypes(ctx context.Context, req *cinema.CinemaGetRoomTypesReq) (res *cinema.CinemaGetRoomTypesRes, err error)
	CinemaCreateRoomType(ctx context.Context, req *cinema.CinemaCreateRoomTypeReq) (res *cinema.CinemaCreateRoomTypeRes, err error)
	CinemaUpdateRoomType(ctx context.Context, req *cinema.CinemaUpdateRoomTypeReq) (res *cinema.CinemaUpdateRoomTypeRes, err error)
	CinemaDeleteRoomType(ctx context.Context, req *cinema.CinemaDeleteRoomTypeReq) (res *cinema.CinemaDeleteRoomTypeRes, err error)
	CinemaGetScreeningRooms(ctx context.Context, req *cinema.CinemaGetScreeningRoomsReq) (res *cinema.CinemaGetScreeningRoomsRes, err error)
	CinemaGetScreeningRoom(ctx context.Context, req *cinema.CinemaGetScreeningRoomReq) (res *cinema.CinemaGetScreeningRoomRes, err error)
	CinemaCreateScreeningRoom(ctx context.Context, req *cinema.CinemaCreateScreeningRoomReq) (res *cinema.CinemaCreateScreeningRoomRes, err error)
	CinemaUpdateScreeningRoom(ctx context.Context, req *cinema.CinemaUpdateScreeningRoomReq) (res *cinema.CinemaUpdateScreeningRoomRes, err error)
	CinemaDeleteScreeningRoom(ctx context.Context, req *cinema.CinemaDeleteScreeningRoomReq) (res *cinema.CinemaDeleteScreeningRoomRes, err error)
}
