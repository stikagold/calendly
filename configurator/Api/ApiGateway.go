package Api

import "fmt"

type Gateway struct {
	IsUsed bool   `yaml:"is_used"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

func (gt *Gateway) GetUrl() string {
	fmt.Println(gt)
	return fmt.Sprintf("%s:%s", gt.Host, gt.Port)
}
