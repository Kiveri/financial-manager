package transaction_usecase

import (
	"context"
	"fmt"

	"github.com/Kiveri/financial-manager/internal/domain/model"
)

func (u *UseCase) ListTransactions(ctx context.Context) ([]*model.Transaction, error) {
	transactions, err := u.transactionRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("transactionRepo.FindAll: %w", err)
	}

	return transactions, nil
}
