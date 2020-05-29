package rid

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/except"
	"github.com/teris-io/shortid"
)

const RIDKey = "request-id"

type RID string

func GenerateRID() (RID, error) {
	if tmp, err := shortid.Generate(); err != nil {
		return "", except.NewServerError(err.Error())
	} else {
		return RID(tmp), nil
	}
}

func AssignRID(request midl.Request) (RID, error) {
	rid, err := GenerateRID()
	if err != nil {
		return "", err
	}

	request.AdditionalContext()[RIDKey] = rid
	return rid, nil
}

func GetRID(request midl.Request) RID {
	return request.AdditionalContext()[RIDKey].(RID)
}
