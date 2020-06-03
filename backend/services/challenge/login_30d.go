package challenge

import (
	"crawlab/model"
	"github.com/globalsign/mgo/bson"
)

type Login30dService struct {
	UserId bson.ObjectId
}

func (s *Login30dService) Check() (bool, error) {
	_, err := model.GetVisitDays(s.UserId)
	if err != nil {
		return false, err
	}
	return false, nil
	//return days >= 30, nil
}
