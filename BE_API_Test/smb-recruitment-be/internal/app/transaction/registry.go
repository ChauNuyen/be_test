package transaction

import (
	"github.com/tunaiku/mobilebanking/internal/app/transaction/service"
	"log"

	"github.com/go-chi/chi"
	"github.com/tunaiku/mobilebanking/internal/app/domain"
	"github.com/tunaiku/mobilebanking/internal/app/transaction/handler"
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(func() service.ITransactionService {
		return service.NewTransactionService()
	})

	container.Provide(func(userSessionHelper domain.UserSessionHelper, transactionService service.ITransactionService) *handler.TransactionEndpoint {
		return handler.NewTransactionEndpoint(userSessionHelper, transactionService)
	})
}

func Invoke(container *dig.Container) {
	err := container.Invoke(func(router chi.Router, endpoint *handler.TransactionEndpoint) {
		log.Println("invoke transaction startup ...")
		endpoint.BindRoutes(router)
	})
	if err != nil {
		log.Fatal(err)
	}
}
