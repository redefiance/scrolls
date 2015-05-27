package scrolls

import (
	"fmt"
	"time"
)

// Minimal information (via UpdateMarketInfo)
type MarketInfo struct {
	Price int
	Level int
	Date  time.Time
}

// Detailed information (via UpdateMarketDetails)
type MarketDetails struct {
	MarketInfo

	ID            MarketplaceOfferID
	CopiesForSale int
	SellerID      PlayerID
	Card          CardInstanceInfo
	Date          time.Time
}

// MarketplaceOfferID is a unique identifier for a Sell Order in the Marketplace
type MarketplaceOfferID int

// A SellOrder put up by the player (cancelable)
type SellOrder struct {
	ID    MarketplaceOfferID
	Card  CardInstanceInfo
	Price int
}

// UpdateMarketInfo updates the cached MarketInfo for all card types
func (c *Connection) UpdateMarketInfo() {
	c.request(nil, func(m mMarketplaceAvailableOffersListView) {
		for _, info := range m.Available {
			details := c.Market.Offers[info.Type]
			details.MarketInfo.Date = time.Now()
			details.MarketInfo.Price = info.Price
			details.MarketInfo.Level = info.Level
			c.Market.Offers[info.Type] = details
		}
	})
}

// UpdateMarketDetails updates the cached MarketDetails for the given typeID
func (c *Connection) UpdateMarketDetails(ID CardTypeID) {
	c.request(req{"cardTypeId": ID}, func(m mMarketplaceOffersSearchView) {
		fmt.Println(m)
	})
}
