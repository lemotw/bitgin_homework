package trade

import (
	"BitginHomework/model"
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type EventFunc func(u model.User, trade *model.TradeRecord, cost float64) error

type TradeContext struct {
	EventList   []EventFunc
	tradeRecord model.TradeRecord
}

// apply every event
func (tc *TradeContext) Apply(ctx context.Context, u model.User, balance model.UserBalance, cost float64, point int) error {
	// init var
	record := model.TradeRecord{
		UserID:      u.ID,
		BalanceDiff: cost,
		PointDiff:   point,
		Balance:     balance.Balance,
		Point:       balance.Point,
	}
	// event loop
	for _, event := range tc.EventList {
		if err := event(u, &record, cost); err != nil {
			return err
		}
	}

	// balance broken
	if record.BalanceDiff+balance.Balance < 0 {
		return errors.New("balance not enough")
	}

	// point broken
	if record.PointDiff+balance.Point < 0 {
		return errors.New("point not enough")
	}

	record.Balance = balance.Balance + record.BalanceDiff
	record.Point = balance.Point + record.PointDiff

	tc.tradeRecord = record
	return nil
}

// update to db
func (tc *TradeContext) UpdateBalance(ctx context.Context, balance model.UserBalance, tx sqlx.Tx) error {
	// store record
	err := tc.tradeRecord.Insert(ctx, tx)
	if err != nil {
		return err
	}

	// store user balance
	balance.Balance = tc.tradeRecord.Balance
	err = balance.Update(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
