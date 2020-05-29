package types

import (
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AutoMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Template  EmailTemplate      `bson:"template"`
	Type      string             `bson:"type"`
	StudyKey  string             `bson:"studyKey,omitempty"`
	Condition *ExpressionArg     `bson:"condition,omitempty"`
	NextTime  int64              `bson:"nextTime"`
	Period    int64              `bson:"period"`
}

func AutoMessageFromAPI(obj *api.AutoMessage) *AutoMessage {
	if obj == nil {
		return nil
	}
	_id, _ := primitive.ObjectIDFromHex(obj.Id)
	return &AutoMessage{
		ID:        _id,
		Template:  EmailTemplateFromAPI(obj.Template),
		Type:      obj.Type,
		StudyKey:  obj.StudyKey,
		Condition: ExpressionArgFromAPI(obj.Condition),
		NextTime:  obj.NextTime,
		Period:    obj.Period,
	}
}

func (obj *AutoMessage) ToAPI() *api.AutoMessage {
	if obj == nil {
		return nil
	}
	return &api.AutoMessage{
		Id:        obj.ID.Hex(),
		Template:  obj.Template.ToAPI(),
		Type:      obj.Type,
		StudyKey:  obj.StudyKey,
		Condition: obj.Condition.ToAPI(),
		NextTime:  obj.NextTime,
		Period:    obj.Period,
	}
}
