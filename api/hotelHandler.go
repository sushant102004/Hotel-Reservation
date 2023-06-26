package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/Hotel-Reservation-System/db"
)

type HotelHandler struct {
	hs db.HotelStore
	rs db.RoomStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{hs: hs, rs: rs}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hs.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}
