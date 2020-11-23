package gql

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"goUpdateTaskQty/completion"
	"goUpdateTaskQty/gqldocs"

	"github.com/machinebox/graphql"
)

//HTTPClientSettings HTTPClientSettings
type HTTPClientSettings struct {
    Connect          time.Duration
    ConnKeepAlive    time.Duration
    ExpectContinue   time.Duration
    IdleConn         time.Duration
    MaxAllIdleConns  int
    MaxHostIdleConns int
    ResponseHeader   time.Duration
    TLSHandshake     time.Duration
}
//SetupClient SetupClient
func SetupClient(httpSettings *HTTPClientSettings) *http.Client {
    tr := &http.Transport{
        ResponseHeaderTimeout: httpSettings.ResponseHeader,
        Proxy:                 http.ProxyFromEnvironment,
        DialContext: (&net.Dialer{
            KeepAlive: httpSettings.ConnKeepAlive,
            DualStack: true,
            Timeout:   httpSettings.Connect,
        }).DialContext,
        MaxIdleConns:          httpSettings.MaxAllIdleConns,
        IdleConnTimeout:       httpSettings.IdleConn,
        TLSHandshakeTimeout:   httpSettings.TLSHandshake,
        MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
        ExpectContinueTimeout: httpSettings.ExpectContinue,
    }

    return &http.Client{
        Transport: tr,
    }
}

// UpdateTaskQty update Tasks table row with qty
func UpdateTaskQty(c *completion.Completion) error {

	gqlclient := graphql.NewClient(os.Getenv("GQL_SERVER_URL"), 
		graphql.WithHTTPClient(SetupClient(&HTTPClientSettings{
    		Connect:          5 * time.Second,
    		ExpectContinue:   1 * time.Second,
    		IdleConn:         90 * time.Second,
    		ConnKeepAlive:    30 * time.Second,
    		MaxAllIdleConns:  100,
    		MaxHostIdleConns: 10,
    		ResponseHeader:   5 * time.Second,
    		TLSHandshake:     5 * time.Second,
		})),
	)

	request := graphql.NewRequest(gqldocs.UpdateOperationCompletedQty)
	request.Var("internalid", c.WorkorderID+c.OperationSequence)
	request.Var("completedQty", c.CompletedQty)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-hasura-admin-secret", os.Getenv("GQL_SERVER_SECRET"))

	resp := &gqldocs.UpdatedTaskByPk{}
	
	reqerr := gqlclient.Run(context.Background(), request, &resp)
	if reqerr != nil {
		return reqerr
	}
	
	if resp.CompletedQty != c.CompletedQty {
		return errors.New("Qty did not update")
	}
	return nil
}
	
