package main

import (
	"net/http"

	"github.com/phftavares/example-modular-monolith/internal/customer"
	"github.com/phftavares/example-modular-monolith/internal/vehicle"
	"github.com/phftavares/example-modular-monolith/pkg/web"
)

func main() {
	hdlCustomer := customer.NewCustomerHandler(customer.NewCustomerService())
	hdlVehicle := vehicle.NewVehicleHandler()

	r := web.New()

	r.Post("/customers", hdlCustomer.Create)
	r.Put("/customers/{personal_id}", hdlCustomer.Update)
	r.Get("/customers/{personal_id}", hdlCustomer.GetByPersonalID)

	r.Post("/vehicles", hdlVehicle.Create)
	r.Get("/vehicles", hdlVehicle.GetWith)
	r.Put("/vehicles/{vehicle_id}", hdlVehicle.Update)
	r.Get("/vehicles/{vehicle_id}", hdlVehicle.GetByID)

	http.ListenAndServe(":8080", r)
}
