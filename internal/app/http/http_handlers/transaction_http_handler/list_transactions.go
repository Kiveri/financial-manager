package transaction_http_handler

import (
	"fmt"
	"net/http"

	"github.com/Kiveri/financial-manager/internal/app/http/http_helpers"
)

func (h *Handler) ListTransactions(rw http.ResponseWriter, r *http.Request) {
	rh := http_helpers.NewResponseHandler(rw)

	transactions, err := h.transactionUseCase.ListTransactions(r.Context())
	if err != nil {
		rh.ErrorResponse(
			http.StatusInternalServerError,
			fmt.Errorf("transactionUseCase.ListTransactions: %w", err).Error(),
			http_helpers.CodeInternalError,
		)

		return
	}

	rh.SuccessResponse(http.StatusOK, transactions)
}
