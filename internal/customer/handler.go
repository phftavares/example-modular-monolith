package customer

import (
	"context"
	"net/http"

	"github.com/phftavares/example-modular-monolith/pkg/logger"
	"github.com/phftavares/example-modular-monolith/pkg/web"
)

type CustomerService interface {
	Create(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	FindByPersonalID(ctx context.Context, personalID string) (*Customer, error)
}

type handler struct {
	service CustomerService
}

func (hdl *handler) Create(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var customer Customer
	if err := web.DecodeJSON(r, &customer); err != nil {
		logger.Error(ctx, "fail decode json", logger.Err(err))
		return err
	}

	err := hdl.service.Create(ctx, &customer)
	if err != nil {
		logger.Error(ctx, "fail create customer", logger.Any("customer", customer), logger.Err(err))
		return err
	}

	return web.EncodeJSON(w, customer, http.StatusCreated)
}

func (hdl *handler) Update(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func (hdl *handler) GetByPersonalID(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func NewCustomerHandler(service CustomerService) *handler {
	return &handler{
		service: service,
	}
}
