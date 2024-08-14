package core

type UnitResource struct {
	Idle string
	Walk string
	Die  string
}

var unitResources = map[string]UnitResource{
	"blue-soldier": {
		Idle: "soldier-idle",
		Walk: "soldier-walk",
		Die:  "soldier-die",
	},
	"red-soldier": {
		Idle: "soldier-idle",
		Walk: "soldier-walk",
		Die:  "soldier-die",
	},
	"red-archer": {
		Idle: "archer-idle",
		Walk: "archer-idle",
		Die:  "archer-idle",
	},
	"red-knight": {
		Idle: "horse",
		Walk: "horse",
		Die:  "horse",
	},
	"thug": {
		Idle: "thug",
		Walk: "thug",
		Die:  "thug",
	},
}
