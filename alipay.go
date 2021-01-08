/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 4:20
*/

package alipay

import (
	"context"
)

type Service interface {
	Options() Options
	Request(method string, m M, response interface{}, opts ...Option) (err error)
	MakeParam(method string, m M, opts ...Option) (data M, err error)
}

type Option func(*Options)

type service struct {
	opts Options
}

func (s service) MakeParam(method string, m M, opts ...Option) (data M, err error) {
	copyOpts := s.opts.Copy()
	// setup options
	for _, o := range opts {
		o(&copyOpts)
	}
	return CombineMakeMapEndpointFunc(
		SetMethod(method),
		SetPublicParam(copyOpts.config),
		SetOptionsParam(copyOpts.config),
		SetBizContent(m),
		SignParam(copyOpts.config),
	)
}

func (s service) Request(method string, m M, response interface{}, opts ...Option) (err error) {
	copyOpts := s.opts.Copy()
	ctx, cancel := context.WithCancel(copyOpts.context)
	defer cancel()
	// setup options
	for _, o := range opts {
		o(&copyOpts)
	}

	// make request
	req, err := copyOpts.makeReq(ctx, method, m, copyOpts.config)
	if err != nil {
		cancel()
		return err
	}

	// before hooks
	for _, f := range copyOpts.before {
		ctx = f(ctx, req)
	}

	// do request
	resp, err := copyOpts.transport.RoundTrip(req)
	if err != nil {
		cancel()
		return err
	}
	defer resp.Body.Close()

	// after hooks
	for _, f := range copyOpts.after {
		ctx = f(ctx, resp)
	}

	// dec response
	err = copyOpts.dec(ctx, resp, response)
	if err != nil {
		cancel()
		return err
	}
	return
}

func (s service) Options() Options {
	return s.opts
}

func NewService(opts ...Option) Service {
	return newService(opts...)
}

func DefaultClient(opts ...Option) Service {
	var newOpts []Option
	newOpts = append(newOpts, DefaultLocation(), DefaultCharset(), DefaultFormat(), DefaultVersion())
	if opts != nil {
		newOpts = append(newOpts, opts...)
	}
	return newService(newOpts...)
}

func newService(opts ...Option) Service {
	service := service{}
	options := newOptions(opts...)
	service.opts = options
	return service
}
