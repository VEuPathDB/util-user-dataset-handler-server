package middle

import (
	"encoding/json"
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
	"io"
	"net/http"
)

func NewBinaryAdaptor() midl.Adapter {

}

type binaryAdaptor struct {
	handlers []midl.Middleware
}

func (b binaryAdaptor) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	req, err := midl.NewRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write(simpleInternalError(err.Error()))
		return
	}

	var res midl.Response
	for _, mid := range b.handlers {
		res = mid.Handle(req)
		if res != nil {
			break
		}
	}

	if res == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write(simpleInternalError("Invalid state: no response."))
		return
	}

	writer.WriteHeader(res.Code())

	if res.Code() != http.StatusOK {
		bytes, _ := json.Marshal(res.Body())
		_, _ = writer.Write(bytes)
		return
	}

	pipe, ok := res.Body().(io.ReadCloser)
	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write(simpleInternalError("Invalid state: bad response body."))
		return
	}
	defer pipe.Close()

	writer.Header().Add(xhttp.HeaderContentType, "application/binary")
	size := 8 * util.SizeKibibyte
	buff := make([]byte, size)

	for true {
		n, err := pipe.Read(buff)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write(simpleInternalError("Failed to write to output buffer"))
			return
		}
		_, _ = writer.Write(buff[0:n])
		if n < size {
			break
		}
	}
}

func (b binaryAdaptor) EmptyHandler(handler midl.EmptyHandler) midl.Adapter {
	panic("implement me")
}

func (b binaryAdaptor) ContentType(s string) midl.Adapter {
	panic("do not call this method")
}

func (b binaryAdaptor) ErrorSerializer(serializer midl.ErrorSerializer) midl.Adapter {
	panic("implement me")
}

func (b binaryAdaptor) Serializer(serializer midl.Serializer) midl.Adapter {
	panic("implement me")
}

func (b binaryAdaptor) AddHandlers(middleware ...midl.Middleware) midl.Adapter {
	panic("implement me")
}

func (b binaryAdaptor) SetHandlers(middleware ...midl.Middleware) midl.Adapter {
	panic("implement me")
}

const simpleErrFmt = `{"status": "server-error", "message": "%s"}`

func simpleInternalError(msg string) []byte {
	return []byte(fmt.Sprintf(simpleErrFmt, msg))
}
