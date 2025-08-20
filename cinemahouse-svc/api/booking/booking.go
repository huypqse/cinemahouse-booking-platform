// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package booking

import (
	"context"

	"cinemahouse-svc/api/booking/booking"
)

type IBookingBooking interface {
	BookingGetBills(ctx context.Context, req *booking.BookingGetBillsReq) (res *booking.BookingGetBillsRes, err error)
	BookingGetBill(ctx context.Context, req *booking.BookingGetBillReq) (res *booking.BookingGetBillRes, err error)
	BookingUpdateBillStatus(ctx context.Context, req *booking.BookingUpdateBillStatusReq) (res *booking.BookingUpdateBillStatusRes, err error)
	BookingGetScreeningSessions(ctx context.Context, req *booking.BookingGetScreeningSessionsReq) (res *booking.BookingGetScreeningSessionsRes, err error)
	BookingGetScreeningSession(ctx context.Context, req *booking.BookingGetScreeningSessionReq) (res *booking.BookingGetScreeningSessionRes, err error)
	BookingCreateScreeningSession(ctx context.Context, req *booking.BookingCreateScreeningSessionReq) (res *booking.BookingCreateScreeningSessionRes, err error)
	BookingUpdateScreeningSession(ctx context.Context, req *booking.BookingUpdateScreeningSessionReq) (res *booking.BookingUpdateScreeningSessionRes, err error)
	BookingDeleteScreeningSession(ctx context.Context, req *booking.BookingDeleteScreeningSessionReq) (res *booking.BookingDeleteScreeningSessionRes, err error)
	BookingGetSeats(ctx context.Context, req *booking.BookingGetSeatsReq) (res *booking.BookingGetSeatsRes, err error)
	BookingGetSeatsBySession(ctx context.Context, req *booking.BookingGetSeatsBySessionReq) (res *booking.BookingGetSeatsBySessionRes, err error)
	BookingCreateSeat(ctx context.Context, req *booking.BookingCreateSeatReq) (res *booking.BookingCreateSeatRes, err error)
	BookingUpdateSeat(ctx context.Context, req *booking.BookingUpdateSeatReq) (res *booking.BookingUpdateSeatRes, err error)
	BookingDeleteSeat(ctx context.Context, req *booking.BookingDeleteSeatReq) (res *booking.BookingDeleteSeatRes, err error)
	BookingSeatTypes(ctx context.Context, req *booking.BookingSeatTypesReq) (res *booking.BookingSeatTypesRes, err error)
	BookingCreateSeatType(ctx context.Context, req *booking.BookingCreateSeatTypeReq) (res *booking.BookingCreateSeatTypeRes, err error)
	BookingUpdateSeatType(ctx context.Context, req *booking.BookingUpdateSeatTypeReq) (res *booking.BookingUpdateSeatTypeRes, err error)
	BookingDeleteSeatType(ctx context.Context, req *booking.BookingDeleteSeatTypeReq) (res *booking.BookingDeleteSeatTypeRes, err error)
	BookingBookTickets(ctx context.Context, req *booking.BookingBookTicketsReq) (res *booking.BookingBookTicketsRes, err error)
	BookingGetTickets(ctx context.Context, req *booking.BookingGetTicketsReq) (res *booking.BookingGetTicketsRes, err error)
	BookingGetTicket(ctx context.Context, req *booking.BookingGetTicketReq) (res *booking.BookingGetTicketRes, err error)
	BookingCancelTicket(ctx context.Context, req *booking.BookingCancelTicketReq) (res *booking.BookingCancelTicketRes, err error)
}
