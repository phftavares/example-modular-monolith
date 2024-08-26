package customer

import "context"

type service struct {
}

func (s *service) Create(ctx context.Context, c *Customer) error {
	return nil
}

func (s *service) Update(ctx context.Context, c *Customer) error {
	return nil
}

func (s *service) FindByPersonalID(ctx context.Context, personalID string) (*Customer, error) {
	return nil, nil
}

func NewCustomerService() *service {
	return &service{}
}
