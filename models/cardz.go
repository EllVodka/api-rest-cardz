package models

type Cardz struct {
	Id         int         `json:"id" db:"id"`
	Name       string      `json:"name" db:"name"`
	Bio        string      `json:"biography" db:"biography"`
	Avatar     string      `json:"avatar" db:"avatar"`
	Animation  string      `json:"animation" db:"animation"`
	BacgroudId int         `json:"-" db:"backgroudId"`
	Background Back        `json:"background"`
	Links      []CardzLink `json:"links"`
}

type Back struct {
	Id        int    `json:"-" db:"id"`
	Image     string `json:"image" db:"image"`
	Animation string `json:"animation" db:"animation"`
}

type CardzLink struct {
	Id      int    `json:"-" db:"id"`
	Url     string `json:"url" db:"url"`
	Text    string `json:"text" db:"text"`
	Icon    string `json:"icon" db:"icon"`
	CardzId int    `json:"-" db:"idCardz"`
}
