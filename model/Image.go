package model

//图片
type Image struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

func (Image) TableName() string {
	return "image"
}
