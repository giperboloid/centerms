package entities

var (
	DefaultFridgeConfig = FridgeConfig{
		TurnedOn:    true,
		CollectFreq: 1000,
		SendFreq:    2000,
	}
)

type Fridge struct {
	Data   FridgeData   `json:"data"`
	Config FridgeConfig `json:"config"`
	Meta   DevMeta      `json:"meta"`
}

type FridgeData struct {
	TopCompart map[int64]float32 `json:"topCompart"`
	BotCompart map[int64]float32 `json:"botCompart"`
}

type FridgeConfig struct {
	TurnedOn    bool  `json:"turnedOn"`
	CollectFreq int64 `json:"collectFreq"`
	SendFreq    int64 `json:"sendFreq"`
}
