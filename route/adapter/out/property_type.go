package out

import (
	"github.com/KaySar12/NextZen-MessageBus/codegen"
	"github.com/KaySar12/NextZen-MessageBus/model"
)

func PropertyTypeAdapter(propertyType model.PropertyType) codegen.PropertyType {
	return codegen.PropertyType{
		Name: propertyType.Name,
	}
}
