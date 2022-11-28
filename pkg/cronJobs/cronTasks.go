package cronJobs

import (
	"app/pkg/service"
)

type Cron struct {
	services *service.Service
}

//type Handler struct {
//	services *service.Service
//}

func NewCronRunner(services *service.Service) *Cron {
	return &Cron{services: services}
}

func (c *Cron) RunCronJobs() error {
	return c.services.Marketplace.ValidateSCItems()

}
