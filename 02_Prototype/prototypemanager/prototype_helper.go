package prototypemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/jinzhu/copier"
)

var once sync.Once

type SwordConfigs []SwordConfig

type SwordConfig struct {
	Type      string `json:"type"`
	Weight    int    `json:"weight"`
	Length    int    `json:"length"`
	Hitpoints int    `json:"hitpoints"`
	Materals  string `json:"materals"`
	Color     string `json:"color"`
}

type Sword struct {
	Type      string
	Weight    int
	Length    int
	Hitpoints int
	Materals  string
	Color     string
}

var swords []Sword

func InitPrototype() {
	once.Do(func() { // <-- atomic, does not allow repeating
		if swords == nil {

			jsonFile, err := os.Open("config.json")
			if err != nil {
				fmt.Println(err)
			}
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var configs SwordConfigs
			json.Unmarshal(byteValue, &configs)
			swords = make([]Sword, len(configs))
			for i := 0; i < len(configs); i++ {
				swords[i] = Sword{
					Type:      configs[i].Type,
					Weight:    configs[i].Weight,
					Length:    configs[i].Length,
					Hitpoints: configs[i].Hitpoints,
					Materals:  configs[i].Materals,
					Color:     configs[i].Color,
				}

			}
			defer jsonFile.Close()
			fmt.Println(len(swords))
		}

	})
}

func GetInstance(swordtype string) Sword {
	fmt.Println(len(swords))

	for _, sword := range swords {
		if sword.Type == swordtype {
			newSword := Sword{}
			copier.Copy(&newSword, &sword)
			return newSword
		}
	}
	return Sword{}
}
