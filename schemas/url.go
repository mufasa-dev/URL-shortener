package schemas

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortId     string
}
