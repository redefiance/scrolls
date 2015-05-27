package scrolls

type AchievementTypeID int
type AchievementTypeInfo struct {
	ID          AchievementTypeID
	Name        string
	Description string
	GoldReward  int
	Group       int
	SortID      int
	Icon        string
}

type AvatarTypeID int
type AvatarTypeInfo struct {
	ID       AvatarTypeID
	Type     string
	Part     string
	Filename string
	Set      string
}

// CardInstanceID describes a card that owned by a player
type CardInstanceID int
type CardInstanceInfo struct {
	ID       CardInstanceID
	TypeID   CardTypeID
	Tradable bool
	IsToken  bool
	Level    int
}

// CardTypeID describes an abstract card type
type CardTypeID int
type CardTypeInfo struct {
	ID                    CardTypeID
	Name                  string
	Description           string
	SubTypesStr           string
	Kind                  string
	Rarity                int
	CostGrowth            int
	CostOrder             int
	CostEnergy            int
	CostDecay             int
	Ap                    int
	Ac                    int
	Hp                    int
	TargetArea            string
	Set                   int
	Flavor                string
	Sound                 string
	Available             bool
	AnimationPreviewImage int
	CardImage             int
	AnimationPreviewInfo  string
	AnimationBundle       int

	Abilities []struct {
		ID          string
		Name        string
		Description string
		Cost        struct {
			Decay  int
			Order  int
			Energy int
			Growth int
		}
	}
	RulesList    []string
	PassiveRules []struct {
		DisplayName string
		Description string
	}
}

type IdolTypeID int
type IdolTypeInfo struct {
	ID       IdolTypeID
	Name     string
	Type     string
	Filename string
}

type PlayerName string
type PlayerID int

// Information about a player provided by Friend and Trade requests
type PlayerInfo struct {
	ID          PlayerID
	Name        PlayerName
	AdminRole   string
	FeatureType string
}

// Information about a player provided by the chat room system
type PlayerChatInfo struct {
	ProfileID        PlayerID
	Name             PlayerName
	AcceptChallenges bool
	AcceptTrades     bool
	AdminRole        string
	FeatureType      string
}

// ProfileInfo is only available for connected player account
type ProfileInfo struct {
	Gold                   int
	Shards                 int
	Rating                 int
	SelectedPreconstructed int
	AcceptChallenges       bool
	AcceptTrades           bool
	SpectatePermission     string
}

type RoomName string

type StoreItemID int
type StoreItemInfo struct {
	// Information provided for every Item
	ItemID      StoreItemID
	IsPurchased bool
	CostGold    int
	CostShards  int
	IsPublic    bool
	Description string
	ItemType    string
	ItemName    string

	// Extra Information for Avatar Items
	Avatar struct {
		ID          int
		Name        string
		Image       string
		Description string

		Head     int
		Body     int
		Leg      int
		ArmBack  int
		ArmFront int
	}

	// Extra Information for Custom Decks
	DeckName        string
	CardTypeIDs     []CardTypeID
	DeckDescription string

	// Extra Information for Custom Idols
	IdolID IdolTypeID

	// Extra Information for "Just for you" cards
	CardTypeID CardTypeID
	Expires    string
}
