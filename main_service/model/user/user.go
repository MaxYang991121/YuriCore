package user

type (
	UserNetInfo struct {
		ExternalIpAddress  uint32
		ExternalClientPort uint16
		ExternalServerPort uint16

		LocalIpAddress  uint32
		LocalClientPort uint16
		LocalServerPort uint16
	}

	UserInfo struct {
		UserID        uint32 `json:"userid" bson:"-"`
		UserName      string `json:"username" bson:"username"`
		NickName      string `json:"nickname" bson:"nickname"`
		Password      string `json:"password" bson:"password"`
		Level         uint8  `json:"level" bson:"level"`
		CurExp        uint64 `json:"curexp" bson:"curexp"`
		MaxExp        uint64 `json:"maxexp" bson:"maxexp"`
		Points        uint64 `json:"points" bson:"points"`
		PlayedMatches uint32 `json:"playedmatches" bson:"playedmatches"`
		Wins          uint32 `json:"wins" bson:"wins"`
		Kills         uint32 `json:"kills" bson:"kills"`
		Deaths        uint32 `json:"deaths" bson:"deaths"`
		Campaign      uint8  `json:"campaign" bson:"campaign"`
		Rank          uint32 `json:"rank" bson:"rank"`
		ChatTimes     uint8  `json:"chatimes" bson:"chatimes"`
		Options       []byte `json:"options" bson:"options"`

		UserInventory Inventory   `json:"inventory" bson:"inventory"`
		NetInfo       UserNetInfo `json:"netinfo" bson:"-"`
		Friends       []string    `json:"friends" bson:"friends"`
	}
)
