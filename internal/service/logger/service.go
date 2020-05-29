package logger

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/service/rid"
	"time"

	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

const (
	lifespan = 30 * time.Minute
	cleanup  = 5 * time.Minute
	ridKey   = "request-id"
)

var logCache = cache.New(lifespan, cleanup)

func Get(id rid.RID) *logrus.Entry {
	if tmp, ok := logCache.Get(string(id)); ok {
		return tmp.(*logrus.Entry)
	}

	tmp := log.Logger().WithField(ridKey, id)
	logCache.SetDefault(string(id), tmp)

	return tmp
}

func ByRequest(req midl.Request) *logrus.Entry {
	return Get(rid.GetRID(req))
}

func AddFields(id rid.RID, fields map[string]interface{}) *logrus.Entry {
	tmp := Get(id).WithFields(fields)
	_, left, _ := logCache.GetWithExpiration(string(id))
	logCache.Set(string(id), tmp, time.Until(left))
	return tmp
}

func Delete(id rid.RID) {
	logCache.Delete(string(id))
}
