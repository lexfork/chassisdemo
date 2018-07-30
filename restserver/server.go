package restserver

import (
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-chassis/go-chassis/core/lager"
	"net/http"
	"github.com/go-chassis/go-chassis"
)


type RestFulHello struct {}



func (r *RestFulHello) Sayhello(b *restful.Context) {
	lager.Logger.Infof("Request:%s", b.ReadPathParameter("userid"))
	b.Write([]byte("Hello " + b.ReadPathParameter("userid") + " from rest server"))
}

func (s *RestFulHello) URLPatterns() []restful.Route {
	return []restful.Route{
		{http.MethodGet, "/server/{userid}", "Sayhello"},
	}
}

func Run()  {
	chassis.RegisterSchema("rest", &RestFulHello{})
}

