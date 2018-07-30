package api

import (
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"net/http"
	"github.com/go-chassis/go-chassis"
	"github.com/tomlee0201/chassisdemo/protobuf"
	"context"
)

type RestFulApi struct {}

func (r *RestFulApi) SayRestHello(b *restful.Context) {
	lager.Logger.Infof("Request:%s", b.ReadPathParameter("userid"))

	req, _ := rest.NewRequest("GET", "cse://RestServer/server/" + b.ReadPathParameter("userid"))
	defer req.Close()
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		b.WriteHeader(http.StatusServiceUnavailable)
		b.Write([]byte("Server internal error"))

		lager.Logger.Error("error", err)
		return
	}
	defer resp.Close()
	responseBody := resp.ReadBody()
	lager.Logger.Info(string(responseBody))

	if resp.GetStatusCode() != 200 {
		b.WriteHeader(resp.GetStatusCode())
		b.Write([]byte("Server internal error"))
	} else {
		b.Write(responseBody)
	}
}


func (r *RestFulApi) SayRPCHello(b *restful.Context) {
	lager.Logger.Infof("Request:%s", b.ReadPathParameter("userid"))

	//declare reply struct
	reply := &protobuf.HelloReply{}
	//Invoke with microservice name, schema ID and operation ID
	if err := core.NewRPCInvoker().Invoke(context.Background(), "RpcServer", "HelloService", "SayHello", &protobuf.HelloRequest{Name: b.ReadPathParameter("userid")}, reply); err != nil {
		lager.Logger.Error("error", err)
		b.WriteHeader(http.StatusInternalServerError)
		b.Write([]byte("Server internal error"))
	} else {
		b.Write([]byte(reply.Message))
	}
}


func (s *RestFulApi) URLPatterns() []restful.Route {
	return []restful.Route{
		{http.MethodGet, "/sayresthello/{userid}", "SayRestHello"},
		{http.MethodGet, "/sayrpchello/{userid}", "SayRPCHello"},
	}
}

func Run() {
	//start all server you register in server/schemas.
	chassis.RegisterSchema("rest", &RestFulApi{})

}

