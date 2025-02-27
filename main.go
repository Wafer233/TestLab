package main

import (
	"github.com/gin-gonic/gin"
	"gos/gins"
)

func main() {
	r := gin.Default()

	//gins.Init(r)
	//gins.Req(r)
	//gins.ReqAny(r)
	//gins.ReqMany(r)
	//gins.URIs(r)
	//gins.URIv2(r)
	//gins.Group(r)
	//gins.GetReq(r)
	//gins.GetReqV2(r)
	//gins.GetArrayV1(r)
	//gins.GetArrayV2(r)
	//gins.GetMapV1(r)
	//gins.PostParm(r)
	//gins.PostParmV2(r)
	//gins.PostParmV3(r)
	//gins.PostParmV4(r)
	//gins.Response(r)
	//gins.Temp(r)
	//gins.Cookie(r)
	//gins.Session(r)
	gins.SessionMany(r)
	r.Run(":8080")
}
