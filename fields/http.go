package fields

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	headerPrefix = "header-"
	protoKey     = "proto"
	methodKey    = "method"
	fromKey      = "from"
	urlKey       = "url"
	bodyKey      = "body"
	statusKey    = "status"
)

func HTTPHeader(h http.Header) logrus.Fields {
	f := make(logrus.Fields, len(h)+20)
	for k, vs := range h {
		f[headerPrefix+k] = strings.Join(vs, ",")
	}
	return f
}

func HTTPRequest(req *http.Request, body []byte) logrus.Fields {
	f := HTTPHeader(req.Header)
	f[protoKey] = req.Proto
	f[methodKey] = req.Method
	f[fromKey] = req.RemoteAddr
	f[urlKey] = req.URL.String()
	if body != nil {
		f[bodyKey] = string(body)
	}
	return f
}

func HTTPResponse(resp *http.Response, body []byte) logrus.Fields {
	f := HTTPHeader(resp.Header)
	f[protoKey] = resp.Proto
	f[statusKey] = resp.Status
	if body != nil {
		f[bodyKey] = string(body)
	}
	return f
}
