package domain

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type IpConfConext struct {
	Ctx       *context.Context
	AppCtx    *app.RequestContext
	ClinetCtx *ClientConext
}

type ClientConext struct {
	IP string `json:"ip"`
}

func BuildIpConfContext(c *context.Context, ctx *app.RequestContext) *IpConfConext {
	ipConfConext := &IpConfConext{
		Ctx:       c,
		AppCtx:    ctx,
		ClinetCtx: &ClientConext{},
	}
	return ipConfConext
}
