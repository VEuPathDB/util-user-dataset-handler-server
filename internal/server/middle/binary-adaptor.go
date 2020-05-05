package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
)

func NewBinaryAdaptor() midl.Adapter {
	return &binaryAdaptor{}
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

func (b binaryAdaptor) EmptyHandler(midl.EmptyHandler) midl.Adapter {
	panic("unused")
}

func (b binaryAdaptor) ContentType(string) midl.Adapter {
	panic("unused")
}

func (b binaryAdaptor) ErrorSerializer(midl.ErrorSerializer) midl.Adapter {
	panic("unused")
}

func (b binaryAdaptor) Serializer(midl.Serializer) midl.Adapter {
	panic("unused")
}

func (b *binaryAdaptor) AddHandlers(middleware ...midl.Middleware) midl.Adapter {
	b.handlers = append(b.handlers, middleware...)
	return b
}

func (b binaryAdaptor) SetHandlers(...midl.Middleware) midl.Adapter {
	panic("unused")
}

const simpleErrFmt = `{"status": "server-error", "message": "%s"}`

func simpleInternalError(msg string) []byte {
	return []byte(fmt.Sprintf(simpleErrFmt, msg))
}
