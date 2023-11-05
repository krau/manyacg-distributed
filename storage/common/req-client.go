package common

import (
	"time"

	"github.com/imroc/req/v3"
)

var ReqCilent *req.Client

func init() {
	c := req.C().ImpersonateChrome()
	c.SetCommonRetryCount(2)
	c.TLSHandshakeTimeout = time.Second * 3
	c.SetTimeout(time.Second * 5)
	ReqCilent = c
}
