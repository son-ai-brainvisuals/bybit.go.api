package bybit_connector

import (
	"context"
	"net/url"
	"strings"
)

func (c *Client) SendCustomSignRequest(ctx context.Context, method, endpoint string, query *url.Values, bodyRequest []byte, opts ...RequestOption) (*ServerResponse, error) {
	r := &request{
		method:   method,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	if query != nil {
		r.query = *query
	}
	if len(bodyRequest) > 0 {
		r.body = strings.NewReader(string(bodyRequest))
	}
	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GetServerResponse(err, data)

}
