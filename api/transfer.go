package api

import (
	`database/sql`
	`errors`
	`fmt`
	`net/http`

	db `github.com/divin3circle/sba/db/sqlc`
	`github.com/gin-gonic/gin`
)

type createTransferRequest struct {
	Currency      string `json:"currency" binding:"required,currency`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	FromAccountID int64  `json:"fromAccountId" binding:"required,min=1"`
	ToAccountID   int64  `json:"toAccountId" binding:"required,min=1"`
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	if !s.validAccountAndCurrency(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !s.validAccountAndCurrency(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxnParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := s.store.TransferTxn(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (s *Server) validAccountAndCurrency(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("invalid currency. sending %s to %s", currency, account.Currency)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return false
	}
	return true
}
