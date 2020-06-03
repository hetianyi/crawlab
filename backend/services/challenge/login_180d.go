package challenge

import (
	"crawlab/model"
	"github.com/globalsign/mgo/bson"
)

type Login180dService struct {
	UserId bson.ObjectId
}

func (s *Login180dService) Check() (bool, error) {
	_, err := model.GetVisitDays(s.UserId)
	if err != nil {
		return false, err
	}
	return true, nil
	// return days >= 180, nil
}
