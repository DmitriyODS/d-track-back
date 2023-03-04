package endpoints

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

// TransactionMiddleware - промежуточное ПО для обработки транзакций к БД
func TransactionMiddleware(rClient global.RepoClient) global.Middleware {
	return func(next global.Endpoint) global.Endpoint {
		return func(ctx context.Context, request interface{}) (res global.ResponseData, err error) {
			// создаём новую транзакцию
			txCtx, err := rClient.NewTxContext(ctx)
			if err != nil {
				log.Println("TransactionMiddleware err:", err)
				return global.NewErrResponseData(global.InternalServerErr), global.InternalServerErr
			}

			// закрываем транзакцию
			defer func() {
				if err != nil {
					rClient.RollbackTx(txCtx)
					return
				}

				err = rClient.CommitTx(txCtx)
				if err != nil {
					log.Println("TransactionMiddleware commit err:", err)
					err = global.InternalServerErr
				}
			}()

			return next(txCtx, request)
		}
	}
}
