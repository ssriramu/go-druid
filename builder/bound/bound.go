package bound

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Aggregator, error) {
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var b builder.Bound
	switch t.Typ {
	case "polygon":
		b = NewPolygon()
	case "radius":
		b = NewRadius()
	case "rectangular":
		b = NewRectangular()
	}
	return b, json.Unmarshal(data, &b)
}
