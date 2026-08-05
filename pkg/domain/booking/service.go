package booking

import (
	"tixora/pkg/domain/booking/dto"
	"tixora/pkg/domain/event"

	"github.com/google/uuid"
)

type service struct {
	bookingRepo Repository
	eventRepo   event.Repository
}

func NewService(bookingRepo Repository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func generateBookingCode() string {
	return "GT-" + uuid.New().String()
}

func (s *service) CreateBooking(userId uint, req dto.CreateRequest) (*dto.Response, error) {
	booking, err := s.bookingRepo.CreateWithTicketUpdate(userId, req.EventID, req.Quantity)
	if err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}

func (s *service) GetMyBookings(userId uint) ([]*dto.Response, error) {
	bookings, err := s.bookingRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.Response, len(bookings))

	for i, b := range bookings {
		responses[i] = b.ToResponse()
	}

	return responses, nil
}

func (s *service) CancelBooking(bookingId uint, userId uint) (*dto.Response, error) {
	booking, err := s.bookingRepo.Cancel(bookingId, userId)
	if err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}
