package routes

import (
	"hotelapi/handlers"
	"hotelapi/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /booking", handlers.GetAllBookings())
	mux.HandleFunc("GET /guest", handlers.GetAllGuests())
	mux.HandleFunc("GET /room", handlers.GetAllRooms())
	mux.HandleFunc("GET /room/available", handlers.GetAvailableRooms())
	mux.HandleFunc("GET /booking/byguest", handlers.GetBookingByGuest())
	mux.HandleFunc("GET /guest/byid", handlers.GetGuestById())
	mux.HandleFunc("GET /booking/byid", handlers.GetBookingById())
	mux.HandleFunc("GET /room/byid", handlers.GetRoomByID())
	mux.HandleFunc("GET /room/bystatus", handlers.GetRoomByStatus())
	mux.HandleFunc("GET /room/bytype", handlers.GetRoomByType())
	mux.Handle("POST /room/addroom", middleware.HotelMiddleware("allowpost")(handlers.AddRoom()))
	mux.Handle("POST /guest/createguest", middleware.HotelMiddleware("allowpost")(handlers.CreateGuest()))
	mux.Handle("POST /booking/createbooking", middleware.HotelMiddleware("allowpost")(handlers.CreateBooking()))
	mux.Handle("PUT /guest/updateguest", middleware.HotelMiddleware("allowput")(handlers.UpdateGuest()))
	mux.Handle("PUT /room/updateroom", middleware.HotelMiddleware("allowput")(handlers.UpdateRoom()))
	mux.Handle("DELETE /guest/deleteguest", middleware.HotelMiddleware("allowdelete")(handlers.DeleteGuest()))
	mux.Handle("DELETE /room/deleteroom", middleware.HotelMiddleware("allowdelete")(handlers.DeleteRoom()))
	mux.Handle("DELETE /booking/cancelbooking", middleware.HotelMiddleware("allowdelete")(handlers.CancelBooking()))
	mux.Handle("PUT /booking/checkin", middleware.HotelMiddleware("allowput")(handlers.CheckIn()))
	mux.Handle("PUT /booking/checkout", middleware.HotelMiddleware("allowput")(handlers.CheckOut()))

}
