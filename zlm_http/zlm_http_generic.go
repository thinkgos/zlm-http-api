package zlm_http

import "context"

type responder interface {
	inspectError() error
}

func genericGet[T any, R responder](c *Client, path string, ctx context.Context, req *T, opts ...CallOption) (*R, error) {
	var resp R

	err := c.Get(ctx, path, req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}

func genericPost[T any, R responder](c *Client, path string, ctx context.Context, req *T, opts ...CallOption) (*R, error) {
	var resp R

	err := c.Post(ctx, path, req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
