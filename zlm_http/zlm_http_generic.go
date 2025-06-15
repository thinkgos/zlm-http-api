package zlm_http

import "context"

type Responder interface {
	InspectError() error
}

func GenericGet[T any, R Responder](c *Client, path string, ctx context.Context, req *T, opts ...CallOption) (*R, error) {
	var resp R

	err := c.Get(ctx, path, req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.InspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}

func GenericPost[T any, R Responder](c *Client, path string, ctx context.Context, req *T, opts ...CallOption) (*R, error) {
	var resp R

	err := c.Post(ctx, path, req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.InspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
