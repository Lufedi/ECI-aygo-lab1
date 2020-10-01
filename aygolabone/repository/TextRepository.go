package repository

import "aygolabone/model"

type TextRepository interface {
	Save(text *model.Text) (*model.Text, error)
	GetRecent() ([]*model.Text, error)
}