package booking

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) (Booking, error) {
	return s.store.Book(b)
}

func (s *Service) ListBooking(movieID string) ([]Booking, error) {
	return s.store.ListBooking(movieID), nil
}
