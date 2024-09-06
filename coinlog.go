package migration

import (
	"domain/account"
	"domain/balance"
	"domain/betting"
	"domain/recharge"
	"domain/withdraw"
	"fmt"
	"github.com/shopspring/decimal"
	"migration/pkg/model"
	"pkg/errorx"
	"pkg/types"
)

type CoinLog interface {
	WithBetSlip(in *WinBetSlips) *WinCoinLog
	WithDeposit(in *model.WinCoinDepositRecord) *WinCoinLog
	WithWithdrawal(in *model.WinCoinWithdrawalRecord) *WinCoinLog
	Validate() errorx.Error
	Adjust() errorx.Error
	Deposit() errorx.Error
	Withdraw() errorx.Error
	PlaceBet() errorx.Error
	Payout() errorx.Error
	Cancel() errorx.Error
}

type WinCoinLog struct {
	coinLog    *model.WinCoinLog              `bson:"coinLog"`
	user       *WinUser                       `bson:"user"`
	deposit    *model.WinCoinDepositRecord    `bson:"deposit,omitempty"`
	withdrawal *model.WinCoinWithdrawalRecord `bson:"withdrawal,omitempty"`
	betSlip    *WinBetSlips                   `bson:"betSlip,omitempty"`
	betDetails *model.WinBetslipsDetails      `bson:"betDetails,omitempty"`
}

func NewWinCoinLog(user *WinUser, coinLog *model.WinCoinLog) CoinLog {
	return &WinCoinLog{user: user, coinLog: coinLog}
}

func (c *WinCoinLog) WithBetSlip(in *WinBetSlips) *WinCoinLog {
	c.betSlip = in
	return c
}

func (c *WinCoinLog) WithDeposit(in *model.WinCoinDepositRecord) *WinCoinLog {
	c.deposit = in
	return c
}

func (c *WinCoinLog) WithWithdrawal(in *model.WinCoinWithdrawalRecord) *WinCoinLog {
	c.withdrawal = in
	return c
}

func (c *WinCoinLog) Validate() errorx.Error {
	b, err := balance.GetUserCurrentBalanceService.Request(newDDDCtx(c.user.ID), c.user.ID)
	if err != nil {
		return err
	}
	coinBefore := decimal.NewFromFloat(c.coinLog.CoinBefore)
	dif := coinBefore.Sub(b.Available)
	if dif.IsZero() {
		return nil
	}
	{
		_, err := balance.RecordService.Request(newDDDCtx(c.user.ID), balance.RecordCommand{
			Action: balance.ActionAdjust,
			Amount: dif,
			Force:  true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *WinCoinLog) Adjust() errorx.Error {
	_, err := balance.RecordService.Request(newDDDCtx(c.user.ID), balance.RecordCommand{
		Action: balance.ActionAdjust,
		Amount: decimal.NewFromFloat(c.coinLog.Coin),
		Force:  true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *WinCoinLog) Deposit() errorx.Error {
	id := types.NewID()
	{
		cmd := recharge.CreateCommand{
			UserID:       c.user.ID,
			ChannelName:  "gcash",
			ProviderName: "dev",
			Currency:     types.PHP,
			Amount:       decimal.NewFromFloat(c.coinLog.Coin),
		}
		_, err := recharge.CreateService.Request(newDDDCtx(id), cmd)
		if err != nil {
			return err
		}
	}
	{
		cmd := recharge.CompleteCommand{}
		_, err := recharge.CompleteService.Request(newDDDCtx(id), cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *WinCoinLog) Withdraw() errorx.Error {
	id := types.NewID()
	{
		cmd := withdraw.CreateCommand{
			Account: account.Account{
				UserID:      c.user.ID,
				ChannelName: "gcash",
			},
			Amount: decimal.NewFromFloat(c.coinLog.Coin),
		}
		_, err := withdraw.CreateService.Request(newDDDCtx(id), cmd)
		if err != nil {
			return err
		}
	}
	{
		cmd := withdraw.AgreeCommand{
			ProviderName: "dev",
		}
		_, err := withdraw.AgreeService.Request(newDDDCtx(id), cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *WinCoinLog) PlaceBet() errorx.Error {
	if c.betSlip == nil {
		return errorx.New("betSlip is nil")
	}
	_, err := betting.PlaceBetService.Request(newDDDCtx(c.betSlip.ID), betting.BetCommand{
		BetDetailWithChildren: betting.BetDetailWithChildren{
			BetDetail: betting.BetDetail{
				BetID:     fmt.Sprintf("%v", c.betSlip.Bet.ID),
				BetAt:     c.betSlip.CreatedAt,
				BetAmount: decimal.NewFromFloat(c.betSlip.Bet.Stake),
				BetContent: func() string {
					game1, ok := games[c.betSlip.Bet.GameListID]
					if !ok {
						return ""
					}
					return game1.Name
				}(),
			},
		},
		UserID: c.user.ID,
		Category: func() betting.GameCategory {
			category1, ok := gameTypes[c.betSlip.Bet.GameTypeID]
			if !ok {
				return ""
			}
			return category1
		}(),
		ProviderName: func() string {
			provider1, ok := brands[c.betSlip.Bet.GameProviderID]
			if !ok {
				return ""
			}
			return provider1
		}(),
		GameName: func() string {
			game1, ok := games[c.betSlip.Bet.GameListID]
			if !ok {
				return ""
			}
			return game1.Name
		}(),
		BrandName: func() string {
			provider1, ok := brands[c.betSlip.Bet.GameProviderID]
			if !ok {
				return ""
			}
			return provider1
		}(),
		Currency: types.PHP,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *WinCoinLog) Payout() errorx.Error {
	if c.betSlip == nil {
		return errorx.New("betSlip is nil")
	}
	_, err := betting.SettleService.Request(newDDDCtx(c.betSlip.ID), betting.SettleCommand{
		SettleDetailWithChildren: betting.SettleDetailWithChildren{
			SettleDetail: betting.SettleDetail{
				SettleAmount: decimal.NewFromFloat(c.betSlip.Bet.Payout),
				BetOutcome:   fmt.Sprintf("bet=%v,win=%v", c.betSlip.Bet.Stake, c.betSlip.Bet.Payout),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *WinCoinLog) Cancel() errorx.Error {
	if c.betSlip == nil {
		return errorx.New("betSlip is nil")
	}
	_, err := betting.CancelService.Request(newDDDCtx(c.betSlip.ID), betting.CancelCommand{
		Reason: "cancel",
	})
	if err != nil {
		return err
	}
	return nil
}
