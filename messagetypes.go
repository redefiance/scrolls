package scrolls

type mActiveGame struct {
}

type mAchievementTypes struct {
	AchievementTypes []AchievementTypeInfo
}

type mAchievementUnlocked struct {
	TypeID AchievementTypeID
}

type mAvatarTypes struct {
	Types []AvatarTypeInfo
}

type mBuyStoreItemResponse struct {
	Cards       []CardInstanceInfo
	DeckInfos   []interface{}
	Avatars     []interface{}
	AvatarParts []interface{}
	Idols       []interface{}
}

type mCardTypes struct {
	CardTypes []CardTypeInfo
}

type mFail struct {
	Op   string
	Info string
}

type mFatalFail struct {
	Info string
}

type mFirstConnect struct {
	AccessToken struct {
		AccessToken       string
		ClientToken       string
		SelectedProfile   struct{ Id, Name string }
		AvailableProfiles []struct{ Id, Name string }
	}
}

type mFriendRequestUpdate struct {
	Request struct {
		From struct {
			Profile PlayerInfo
		}
		To struct {
			Profile     PlayerInfo
			OnlineState string
		}
		Request struct {
			ID               string
			RequestingUserID string
			TargetUserID     string
			Status           string
		}
	}
}

type mFriendUpdate struct {
	Friend       PlayerInfo
	OnlineStatus string
}

type mGetFriends struct {
	Friends []struct {
		Profile PlayerInfo
	}
	OnlineState string
}

type mGetFriendRequests struct {
	Requests []struct {
		From struct {
			Profile PlayerInfo
		}
		To struct {
			Profile     PlayerInfo
			OnlineState string
		}
		Request struct {
			ID               string
			RequestingUserID string
			TargetUserID     string
			Status           string
		}
	}
}

type mGetBlockedPersons struct {
	// TODO
}

type mGetStoreItems struct {
	CardSellbackGold []int
	Items            []StoreItemInfo
}

type mIdolTypes struct {
	Types []IdolTypeInfo
}

type mJoinLobby struct{}

type mLibraryView struct {
	ProfileID PlayerID
	Cards     []CardInstanceInfo
}

type mLobbyLookup struct {
	IP   string
	Port int
}

type mMappedStrings struct {
	Strings []struct {
		Key   string
		Value string
	}
}

type mMarketplaceAvailableOffersListView struct {
	Available []struct {
		Price int
		Level int
		Type  CardTypeID
	}
}

type mMarketplaceOffersSearchView struct {
	TypeID        CardTypeID
	CopiesForSale int
	Offer         struct {
		Price    int
		ID       MarketplaceOfferID
		SellerID PlayerID
		Card     CardInstanceInfo
	}
}

type mMarketplaceOffersView struct {
	ProfileID    int
	MaxNumOffers int
	Offers       []struct {
		ID       MarketplaceOfferID
		SellerID PlayerID
		Price    int
		Card     CardInstanceInfo
	}
}

type mMarketplaceSoldListView struct {
	Sold []struct {
		TransactionID int
		Level         int
		SellPrice     int
		Fee           int
		Claimed       bool
		CardTypeID    CardTypeID
		CardType      CardTypeID // ???
	}
}

type mMessage struct {
	Type string
}

type mOk struct {
	Op string
}

type mPing struct {
	Time int64
}

type mProfileDataInfo struct {
	ProfileData ProfileInfo
}

type mProfileInfo struct {
	Profile PlayerInfo
}

type mRoomChatMessage struct {
	RoomName RoomName
	From     PlayerName
	Text     string
}

type mRoomEnter struct {
	RoomName RoomName
}

type mRoomInfo struct {
	RoomName RoomName
	Reset    bool
	Updated  []PlayerChatInfo
	Removed  []struct {
		Name PlayerName
	}
}

type mServerInfo struct {
	Version  string
	AssetURL string
	NewsURL  string
	Roles    string
}

type mTradeInviteForward struct {
	Inviter PlayerInfo
}

type mTradeResponse struct {
	From   PlayerInfo
	To     PlayerInfo
	Status string
}

type mTradeView struct {
	From struct {
		Profile     PlayerInfo
		CardTypeIDs []CardInstanceID
		Gold        int
		Accepted    bool
	}
	To struct {
		Profile     PlayerInfo
		CardTypeIDs []CardInstanceID
		Gold        int
		Accepted    bool
	}
	Modified bool
}

type mWhisper struct {
	ToProfileName PlayerName
	From          PlayerName
	Text          string
}
