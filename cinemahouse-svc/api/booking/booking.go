// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package booking

import (
	"context"

	"cinemahouse-svc/api/booking/booking"
)

type IBookingV1 interface {
	// Screening session management
	GetScreeningSessions(ctx context.Context, req *booking.BookingGetScreeningSessionsReq) (res *booking.BookingGetScreeningSessionsRes, err error)
	GetScreeningSession(ctx context.Context, req *booking.BookingGetScreeningSessionReq) (res *booking.BookingGetScreeningSessionRes, err error)
	CreateScreeningSession(ctx context.Context, req *booking.BookingCreateScreeningSessionReq) (res *booking.BookingCreateScreeningSessionRes, err error)
	UpdateScreeningSession(ctx context.Context, req *booking.BookingUpdateScreeningSessionReq) (res *booking.BookingUpdateScreeningSessionRes, err error)
	DeleteScreeningSession(ctx context.Context, req *booking.BookingDeleteScreeningSessionReq) (res *booking.BookingDeleteScreeningSessionRes, err error)
	
	// Seat management
	GetSeats(ctx context.Context, req *booking.BookingGetSeatsReq) (res *booking.BookingGetSeatsRes, err error)
	GetSeatsBySession(ctx context.Context, req *booking.BookingGetSeatsBySessionReq) (res *booking.BookingGetSeatsBySessionRes, err error)
	CreateSeat(ctx context.Context, req *booking.BookingCreateSeatReq) (res *booking.BookingCreateSeatRes, err error)
	UpdateSeat(ctx context.Context, req *booking.BookingUpdateSeatReq) (res *booking.BookingUpdateSeatRes, err error)
	DeleteSeat(ctx context.Context, req *booking.BookingDeleteSeatReq) (res *booking.BookingDeleteSeatRes, err error)
	
	// Seat type management
	GetSeatTypes(ctx context.Context, req *booking.BookingSeatTypesReq) (res *booking.BookingSeatTypesRes, err error)
	CreateSeatType(ctx context.Context, req *booking.BookingCreateSeatTypeReq) (res *booking.BookingCreateSeatTypeRes, err error)
	UpdateSeatType(ctx context.Context, req *booking.BookingUpdateSeatTypeReq) (res *booking.BookingUpdateSeatTypeRes, err error)
	DeleteSeatType(ctx context.Context, req *booking.BookingDeleteSeatTypeReq) (res *booking.BookingDeleteSeatTypeRes, err error)
	
	// Ticket booking
	BookTickets(ctx context.Context, req *booking.BookingBookTicketsReq) (res *booking.BookingBookTicketsRes, err error)
	GetTickets(ctx context.Context, req *booking.BookingGetTicketsReq) (res *booking.BookingGetTicketsRes, err error)
	GetTicket(ctx context.Context, req *booking.BookingGetTicketReq) (res *booking.BookingGetTicketRes, err error)
	CancelTicket(ctx context.Context, req *booking.BookingCancelTicketReq) (res *booking.BookingCancelTicketRes, err error)
	
	// Bill management
	GetBills(ctx context.Context, req *booking.BookingGetBillsReq) (res *booking.BookingGetBillsRes, err error)
	GetBill(ctx context.Context, req *booking.BookingGetBillReq) (res *booking.BookingGetBillRes, err error)
	UpdateBillStatus(ctx context.Context, req *booking.BookingUpdateBillStatusReq) (res *booking.BookingUpdateBillStatusRes, err error)
}
