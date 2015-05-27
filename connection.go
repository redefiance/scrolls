package scrolls

// Connection describes a Scrolls account connected to the game server
type Connection struct {
	*messageHandler

	// A cache of the Black Market
	Market struct {
		Offers     map[CardTypeID]MarketDetails
		SellOrders map[MarketplaceOfferID]SellOrder
	}

	// Info about the connected Account
	Account struct {
		Name PlayerName
	}
}

// Connect establishes a connection to the Scrolls Game server
func Connect(email, password string) (*Connection, error) {
	conn := &Connection{}
	conn.Market.Offers = map[CardTypeID]MarketDetails{}
	conn.Market.SellOrders = map[MarketplaceOfferID]SellOrder{}

	var err error

	// establish a connection to the game server
	conn.messageHandler, err = newMessageHandler("54.208.22.193:8081")
	if err != nil {
		return nil, err
	}

	initGameInfo(conn)

	// a couple of messages that are sent on login, but we don't care about
	conn.receive(func(m mServerInfo) {})
	conn.receive(func(m mMappedStrings) {})
	conn.receive(func(m mActiveGame) {})

	conn.receive(func(m mProfileInfo) {
		conn.Account.Name = m.Profile.Name
	})

	mFirstConnect, err := authenticate(email, password)
	if err != nil {
		return nil, err
	}

	conn.confirm(mFirstConnect)

	return conn, nil
}
