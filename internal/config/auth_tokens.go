package config

type FlashArray struct {
        Address       string    `yaml:"address"`
        ApiToken      string    `yaml:"api_token"`
}

type FlashArrayList map[string]FlashArray

func (f *FlashArrayList) GetArrayParams(fa string) (string, string) {
	for a_name, a := range *f {
                if a_name == fa {
			return a.Address, a.ApiToken
		}
	}
	return "", ""
}
