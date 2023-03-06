package config

type FlashArray struct {
        Address       string    `yaml:"address"`
        ApiToken      string    `yaml:"api_token"`
}

type FlashArrayList map[string]FlashArray

func (f *FlashArrayList) GetApiToken(addr string) string {
        for _, a := range *f {
                if a.Address == addr {
                        return a.ApiToken
                }
        }
        return ""
}
