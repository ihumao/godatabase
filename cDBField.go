package cdatabase

import (
	"github.com/ihumao/gopublic"
)

func (this *TField) AsString() string {
	if this == nil {
		return ``
	}
	return public.ToString(this.Variant)
}

func (this *TField) AsInt() int {
	if this == nil {
		return 0
	}
	return public.ToInt(this.Variant)
}

func (this *TField) AsFloat32() float32 {
	if this == nil {
		return 0
	}
	return public.ToFloat32(this.Variant)
}

func (this *TField) AsFloat64() float64 {
	if this == nil {
		return 0
	}
	return public.ToFloat64(this.Variant)
}
