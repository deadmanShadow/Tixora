package event

import (
	"errors"
	"tixora/pkg/domain/event/dto"
)

var ErrInvalidTicketCount = errors.New("cannot reduce total tickets below booked ticket count")

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateEvent(req dto.CreateRequest) (*dto.Response, error) {
	event := Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartsAt:         req.StartsAt,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}

	if err := s.repo.Create(&event); err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) GetEvents() ([]dto.Response, error) {
	events, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	var responses []dto.Response

	for _, event := range events {
		responses = append(responses, *event.ToResponse())
	}

	return responses, nil
}

func (s *service) GetEventByID(eventId uint) (*dto.Response, error) {
	event, err := s.repo.GetByID(eventId)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) UpdateEvent(eventId uint, req dto.UpdateRequest) (*dto.Response, error) {
	event, err := s.repo.GetByID(eventId) // getting existing event by the id first
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		event.Title = req.Title
	}

	if req.Description != "" {
		event.Description = req.Description
	}

	if req.Location != "" {
		event.Location = req.Location
	}

	if !req.StartsAt.IsZero() {
		event.StartsAt = req.StartsAt
	}

	if req.Price != nil {
		event.Price = *req.Price
	}

	if req.TotalTickets != nil {
		diff := *req.TotalTickets - event.TotalTickets
		if event.AvailableTickets+diff < 0 {
			return nil, ErrInvalidTicketCount
		}
		event.TotalTickets = *req.TotalTickets
		event.AvailableTickets += diff
	}

	if err := s.repo.Update(event); err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) DeleteEvent(eventId uint) error {
	return s.repo.Delete(eventId)
}