package user

type (
	UserLoadout struct {
		MainWeapon   uint16 `json:"mainweapon" bson:"mainweapon"`
		SecondWeapon uint16 `json:"secondweapon" bson:"secondweapon"`
		Knife        uint16 `json:"knife" bson:"knife"`
		Grenade      uint16 `json:"grenade" bson:"grenade"`
	}

	UserBuyMenu struct {
		PistolsTR     [9]uint16 `json:"pistolstr" bson:"pistolstr"`         // 0
		PistolsCT     [9]uint16 `json:"pistolsct" bson:"pistolsct"`         // 7
		ShotgunsTR    [9]uint16 `json:"shotgunstr" bson:"shotgunstr"`       // 1
		ShotgunsCT    [9]uint16 `json:"shotgunsct" bson:"shotgunsct"`       // 8
		SmgsTR        [9]uint16 `json:"smgstr" bson:"smgstr"`               // 2
		SmgsCT        [9]uint16 `json:"smgsct" bson:"smgsct"`               // 9
		RiflesTR      [9]uint16 `json:"riflestr" bson:"riflestr"`           // 3
		RiflesCT      [9]uint16 `json:"riflesct" bson:"riflesct"`           // 10
		MachinegunsTR [9]uint16 `json:"machinegunstr" bson:"machinegunstr"` // 4
		MachinegunsCT [9]uint16 `json:"machinegunsct" bson:"machinegunsct"` // 11
		MeleesTR      [9]uint16 `json:"meleestr" bson:"meleestr"`           // 14
		MeleesCT      [9]uint16 `json:"meleesct" bson:"meleesct"`           // 15
		EquipmentTR   [9]uint16 `json:"equipmenttr" bson:"equipmenttr"`     // 5
		EquipmentCT   [9]uint16 `json:"equipmentct" bson:"equipmentct"`     // 12
		ClassesTR     [9]uint16 `json:"classestr" bson:"classestr"`         // 6
		ClassesCT     [9]uint16 `json:"classesct" bson:"classesct"`         // 13
	}

	UserInventoryItem struct {
		Id      uint16 `json:"id" bson:"id"`       //物品id
		Count   uint16 `json:"count" bson:"count"` //数量
		Existed uint8  `json:"existed" bson:"existed"`
		Type    uint8  `json:"type" bson:"type"`
		Time    uint32 `json:"time" bson:"time"`
	}

	UserCosmetics struct {
		CosmeticsName  string `json:"cosmeticsname" bson:"cosmeticsname"`
		MainWeapon     uint16 `json:"mainweapon" bson:"mainweapon"`
		MainBullet     uint16 `json:"mainbullet" bson:"mainbullet"`
		SecondWeapon   uint16 `json:"secondweapon" bson:"secondweapon"`
		SecondBullet   uint16 `json:"secondBullet" bson:"secondBullet"`
		FlashbangNum   uint16 `json:"flashbangnum" bson:"flashbangnum"`
		GrenadeID      uint16 `json:"grenadeid" bson:"grenadeid"`
		SmokeNum       uint16 `json:"smokenum" bson:"smokenum"`
		DefuserNum     uint16 `json:"defusernum" bson:"defusernum"`
		TelescopeNum   uint16 `json:"telescopenum" bson:"telescopenum"`
		BulletproofNum uint16 `json:"Bulletproofnum" bson:"Bulletproofnum"`
		KnifeID        uint16 `json:"knifeid" bson:"knifeid"`
	}

	Inventory struct {
		Items     []UserInventoryItem `json:"items" bson:"items"`     //物品
		BuyMenu   UserBuyMenu         `json:"buymenu" bson:"buymenu"` //购买菜单
		Cosmetics [5]UserCosmetics    `json:"cosmetics" bson:"cosmetics"`
		Loadouts  [3]UserLoadout      `json:"loadouts" bson:"loadouts"`
	}
)
