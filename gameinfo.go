package scrolls

// Info in these maps will be available after the first call to Connect()
var (
	AchievementTypes = map[AchievementTypeID]AchievementTypeInfo{}
	AvatarTypes      = map[AvatarTypeID]AvatarTypeInfo{}
	CardTypes        = map[CardTypeID]CardTypeInfo{}
	IdolTypes        = map[IdolTypeID]IdolTypeInfo{}
)

func initGameInfo(c *Connection) {
	if len(AchievementTypes) > 0 { // already initialized?
		return
	}

	c.receive(func(m mAchievementTypes) {
		for _, t := range m.AchievementTypes {
			AchievementTypes[t.ID] = t
		}
	})
	c.receive(func(m mAvatarTypes) {
		for _, t := range m.Types {
			AvatarTypes[t.ID] = t
		}
	})
	c.receive(func(m mCardTypes) {
		for _, t := range m.CardTypes {
			CardTypes[t.ID] = t
		}
	})
	c.receive(func(m mIdolTypes) {
		for _, t := range m.Types {
			IdolTypes[t.ID] = t
		}
	})
}
