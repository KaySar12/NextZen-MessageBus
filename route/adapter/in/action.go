package in

import (
	"github.com/KaySar12/NextZen-MessageBus/codegen"
	"github.com/KaySar12/NextZen-MessageBus/model"
)

func ActionAdapter(action codegen.Action) model.Action {
	var timestamp int64
	if action.Timestamp != nil {
		timestamp = action.Timestamp.Unix()
	}

	return model.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  timestamp,
	}
}
