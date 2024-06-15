package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name      string `json:"name"`
	Category  string `json:"category"`
	Level1    int    `json:"level1"`
	Level2    int    `json:"level2"`
	Level3    int    `json:"level3"`
	Level4    int    `json:"level4"`
	Level5    int    `json:"level5"`
	Position1 int    `json:"position1"`
	Position2 int    `json:"position2"`
	Position3 int    `json:"position3"`
	MaxPoint  int    `json:"max_point"`
}
