package rid

import (
	"fmt"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/teris-io/shortid"

	"github.com/VEuPathDB/util-exporter-server/internal/except"
)

const (
	RIDKey = "request-id"

	errFailedToGen = "failed to generate a request id: %s"
)

type RID string

func GenerateRID() (RID, error) {
	tmp, err := shortid.Generate()
	if err != nil {
		return "", except.NewServerError(fmt.Sprintf(errFailedToGen, err))
	}

	return RID(tmp), nil
}

func GetRID(request midl.Request) RID {
	return RID(request.RawRequest().Header[RIDKey][0])
}
