package schemas

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	id          string
	originalUrl string
}
