package schemas

type Dino struct {
	dinoName        string
	steamId         int64
	levelSpinner    int64
	imprintSpinner  int64
	neuteredSpinner int64
	sexSpinner      int64
	healthSpinner   int64
	staminaSpinner  int64
	oxygenSpinner   int64
	foodSpinner     int64
	weightSpinner   int64
	damageSpinner   int64
	speedSpinner    int64
	torporSpinner   int64
}

type DinoResponse struct {
	ID              uint   `json:"id"`
	DinoName        string `json:"dinoName"`
	SteamId         int64  `json:"steamId"`
	LevelSpinner    int64  `json:"levelSpinner"`
	ImprintSpinner  int64  `json:"imprintSpinner"`
	NeuteredSpinner int64  `json:"neuteredSpinner"`
	SexSpinner      int64  `json:"sexSpinner"`
	HealthSpinner   int64  `json:"healthSpinner"`
	StaminaSpinner  int64  `json:"staminaSpinner"`
	OxygenSpinner   int64  `json:"oxygenSpinner"`
	FoodSpinner     int64  `json:"foodSpinner"`
	WeightSpinner   int64  `json:"weightSpinner"`
	DamageSpinner   int64  `json:"damageSpinner"`
	SpeedSpinner    int64  `json:"speedSpinner"`
	TorporSpinner   int64  `json:"torporSpinner"`
}
