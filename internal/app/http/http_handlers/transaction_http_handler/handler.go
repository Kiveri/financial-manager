package transaction_http_handler

import (
	"github.com/Kiveri/financial-manager/internal/service/usecases/transaction_usecase"
)

type Handler struct {
	transactionUseCase *transaction_usecase.UseCase
}

func NewHandler(transactionUseCase *transaction_usecase.UseCase) *Handler {
	return &Handler{
		transactionUseCase: transactionUseCase,
	}
}
