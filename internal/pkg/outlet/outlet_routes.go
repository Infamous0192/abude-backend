package outlet

import "abude-backend/internal/common"

func LoadRoutes(r *common.Router) {
	outletService := NewService(r.DB)
	outletHandler := NewController(r.Controller, outletService)

	r.Router.Get("/outlet/count", r.Auth(1), outletHandler.GetCount)
	r.Router.Get("/outlet", r.Auth(0), outletHandler.All)
	r.Router.Get("/outlet/:id", r.Auth(0), outletHandler.One)
	r.Router.Post("/outlet", r.Auth(1), outletHandler.Create)
	r.Router.Put("/outlet/:id", r.Auth(1), outletHandler.Update)
	r.Router.Delete("/outlet/:id", r.Auth(1), outletHandler.Delete)

	r.Router.Get("/outlet/:outletId/employee", r.Auth(1), outletHandler.GetEmployees)
	r.Router.Get("/outlet/:outletId/employee/:employeeId", r.Auth(1), outletHandler.GetEmployee)
	r.Router.Put("/outlet/:outletId/employee/:employeeId", r.Auth(2), outletHandler.AddEmployee)
	r.Router.Get("/outlet/:outletId/employee/:employeeId", r.Auth(2), outletHandler.RemoveEmployee)
}
