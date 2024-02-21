package out

import (
	"time"

	"github.com/KaySar12/NextZen-Common/utils"
	"github.com/KaySar12/NextZen-MessageBus/codegen"
	"github.com/KaySar12/NextZen-MessageBus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
