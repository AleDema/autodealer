package dealer

import (
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/account"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"github.com/thrasher-corp/gocryptotrader/exchanges/orderbook"
	"github.com/thrasher-corp/gocryptotrader/exchanges/stream"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
)

// Strategy is an interface and defines all function needed for a user defined strategy. The RootStrategy provides a way to create and
// use strategies and action according to given strategies.
// Strategy interface defines all the functions that must be implemented in order for the strategy to operate correctly.
// It has "On" functions and a data type that comprises keys and values that may be strings, interfaces, or anything else.
// While this may seem to be a constraint on the execution of your approach, you may use the OnUnrecognized function to implement anything at an experimental stage.
type Strategy interface {
	Init(d *Dealer, e exchange.IBotExchange) error
	OnFunding(d *Dealer, e exchange.IBotExchange, x stream.FundingData) error
	OnPrice(d *Dealer, e exchange.IBotExchange, x ticker.Price) error
	OnKline(d *Dealer, e exchange.IBotExchange, x stream.KlineData) error
	OnOrderBook(d *Dealer, e exchange.IBotExchange, x orderbook.Base) error
	OnOrder(d *Dealer, e exchange.IBotExchange, x order.Detail) error
	OnModify(d *Dealer, e exchange.IBotExchange, x order.Modify) error
	OnBalanceChange(d *Dealer, e exchange.IBotExchange, x account.Change) error
	OnUnrecognized(d *Dealer, e exchange.IBotExchange, x interface{}) error
	Deinit(d *Dealer, e exchange.IBotExchange) error
}
