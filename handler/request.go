package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateDinoRequest struct {
	DinoName        string `json:"dinoName"`
	SteamId         int64  `json:"steamId"`
	LevelSpinner    int    `json:"levelSpinner"`
	ImprintSpinner  int    `json:"imprintSpinner"`
	NeuteredSpinner int    `json:"neuteredSpinner"`
	SexSpinner      int    `json:"sexSpinner"`
	HealthSpinner   int    `json:"healthSpinner"`
	StaminaSpinner  int    `json:"staminaSpinner"`
	OxygenSpinner   int    `json:"oxygenSpinner"`
	FoodSpinner     int    `json:"foodSpinner"`
	WeightSpinner   int    `json:"weightSpinner"`
	DamageSpinner   int    `json:"damageSpinner"`
	SpeedSpinner    int    `json:"speedSpinner"`
	TorporSpinner   int    `json:"torporSpinner"`
}

func (r *CreateDinoRequest) Validate() error {
	if r.DinoName == "" && r.SteamId == 0 && r.LevelSpinner == 0 && r.ImprintSpinner == 0 && r.NeuteredSpinner == 0 && r.SexSpinner == 0 && r.HealthSpinner == 0 && r.StaminaSpinner == 0 && r.OxygenSpinner == 0 && r.FoodSpinner == 0 && r.WeightSpinner == 0 && r.DamageSpinner == 0 && r.SpeedSpinner == 0 && r.TorporSpinner == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.DinoName == "" {
		return errParamIsRequired("dinoName", "string")
	}
	if r.SteamId < 0 {
		return errParamIsRequired("steamId", "int64")
	}
	if r.LevelSpinner < 0 {
		return errParamIsRequired("levelSpinner", "int")
	}
	if r.ImprintSpinner < 0 {
		return errParamIsRequired("imprintSpinner", "int")
	}
	if r.NeuteredSpinner < 0 {
		return errParamIsRequired("neuteredSpinner", "int")
	}
	if r.SexSpinner < 0 {
		return errParamIsRequired("sexSpinner", "int")
	}
	if r.HealthSpinner < 0 {
		return errParamIsRequired("healthSpinner", "int")
	}
	if r.StaminaSpinner < 0 {
		return errParamIsRequired("staminaSpinner", "int")
	}
	if r.OxygenSpinner < 0 {
		return errParamIsRequired("oxygenSpinner", "int")
	}
	if r.FoodSpinner < 0 {
		return errParamIsRequired("foodSpinner", "int")
	}
	if r.WeightSpinner < 0 {
		return errParamIsRequired("weightSpinner", "int")
	}
	if r.DamageSpinner < 0 {
		return errParamIsRequired("damageSpinner", "int")
	}
	if r.SpeedSpinner < 0 {
		return errParamIsRequired("speedSpinner", "int")
	}
	if r.TorporSpinner < 0 {
		return errParamIsRequired("torporSpinner", "int")
	}
	return nil
}

type CreateItemRequest struct {
	Item         string `json:"item"`
	Quantidade   int64  `json:"quantidade"`
	Qualidade    int64  `json:"qualidade"`
	BPOrSela     int64  `json:"bpOrSela"`
	Durabilidade int64  `json:"durabilidade"`
	Dano         int64  `json:"dano"`
	Armadura     int64  `json:"armadura"`
}

func (r *CreateItemRequest) Validate() error {
	if r.Item == "" && r.Quantidade == 0 && r.Qualidade == 0 && r.BPOrSela == 0 && r.Durabilidade == 0 && r.Dano == 0 && r.Armadura == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Item == "" {
		return errParamIsRequired("item", "string")
	}
	if r.Quantidade < 0 {
		return errParamIsRequired("quantidade", "int64")
	}
	if r.Qualidade < 0 {
		return errParamIsRequired("qualidade", "int64")
	}
	if r.BPOrSela < 0 {
		return errParamIsRequired("bpOrSela", "int64")
	}
	if r.Durabilidade < 0 {
		return errParamIsRequired("durabilidade", "int64")
	}
	if r.Dano < 0 {
		return errParamIsRequired("dano", "int64")
	}
	if r.Armadura < 0 {
		return errParamIsRequired("armadura", "int64")
	}
	return nil
}

type CreateChibiRequest struct {
	Dino string `json:"name"`
}

func (r *CreateChibiRequest) Validate() error {
	if r.Dino == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
