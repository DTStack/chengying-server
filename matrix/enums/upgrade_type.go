package enums

type upgradeType struct {
	Upgrade  EnumValueType
	Rollback EnumValueType
}

var UpgradeType = upgradeType{
	Upgrade: EnumValueType{
		Code: 0,
		Desc: "升级",
	},
	Rollback: EnumValueType{
		Code: 1,
		Desc: "回滚",
	},
}
