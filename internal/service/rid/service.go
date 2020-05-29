package rid

import (
	"fmt"
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/except"
	"github.com/teris-io/shortid"
)

const (
	RIDKey = "request-id"

	errFailedToGen = "failed to generate a request id: %s"
)

type RID string

func GenerateRID() (RID, error) {
	if tmp, err := shortid.Generate(); err != nil {
		return "", except.NewServerError(fmt.Sprintf(errFailedToGen, err))
	} else {
		return RID(tmp), nil
	}
}

func AssignRID(request midl.Request) (RID, error) {
	rid, err := GenerateRID()
	if err != nil {
		return "", err
	}

	request.RawRequest().Header[RIDKey] = []string{string(rid)}
	return rid, nil
}

func GetRID(request midl.Request) RID {
	return RID(request.RawRequest().Header[RIDKey][0])
}
