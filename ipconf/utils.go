package ipconf

import (
	"github.com/hardcore-os/plato/ipconf/domain"
)

func top5Endports(eds []*domain.Endport) []*domain.Endport {
	if len(eds) < 5 {
		return eds
	}
	return eds[:5]
}

func packRes(ed []*domain.Endport) Response {
	return Response{
		Message: "ok",
		Code:    0,
		Data:    ed,
	}
}
