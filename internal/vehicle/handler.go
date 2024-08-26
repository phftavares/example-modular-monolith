package vehicle

import (
	"net/http"

	"github.com/phftavares/example-modular-monolith/pkg/web"
)

type handler struct {
}

func (hdl *handler) Create(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func (hdl *handler) Update(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func (hdl *handler) GetByID(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func (hdl *handler) GetWith(w http.ResponseWriter, r *http.Request) error {
	return web.ErrNotImplemented
}

func NewVehicleHandler() *handler {
	return &handler{}
}
