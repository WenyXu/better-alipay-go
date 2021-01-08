/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 4:22
*/

package alipay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ReqFunc func(context context.Context, req *http.Request) context.Context
type ResponseFunc func(context context.Context, response *http.Response) context.Context
type MakeReqFunc func(context context.Context, method string, m M, config Config) (*http.Request, error)
type DecFunc func(context.Context, interface{}, interface{}) error

// Options for alipay sdk
type Options struct {
	config Config

	// hooks
	before []ReqFunc
	after  []ResponseFunc

	dec     DecFunc
	makeReq MakeReqFunc

	// other custom options
	// can be stored in a context
	context context.Context

	transport http.RoundTripper

	logger Logger
}

func (o Options) Copy() Options {
	cp := Options{
		config:    o.config,
		before:    []ReqFunc{},
		after:     []ResponseFunc{},
		dec:       o.dec,
		makeReq:   o.makeReq,
		context:   o.context,
		transport: o.transport,
		logger:    o.logger,
	}
	if len(o.after) != 0 {
		cp.after = append(cp.after, o.after...)
	}
	if len(o.before) != 0 {
		cp.before = append(cp.before, o.before...)
	}
	return o
}

func DefaultMakeReqFunc(ctx context.Context, method string, m M, config Config) (*http.Request, error) {

	body, err := CombineMakeMapEndpointFunc(
		SetMethod(method),
		SetPublicParam(config),
		SetOptionsParam(config),
		SetBizContent(m),
		SignParam(config),
	)

	buf := strings.NewReader(FormatURLParam(body))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.Url(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", ContentType)
	return req, nil
}

func DefaultDecFunc(ctx context.Context, input interface{}, result interface{}) (err error) {
	resp := input.(*http.Response)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func newOptions(opts ...Option) Options {
	opt := Options{
		transport: http.DefaultTransport,
		context:   context.Background(),
		makeReq:   DefaultMakeReqFunc,
		dec:       DefaultDecFunc,
		logger:    StdLogger,
		before:    newDefaultBeforeFunc(),
		after:     newDefaultAfterFunc(),
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func DefaultBeforeFunc(ctx context.Context, req *http.Request) context.Context {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Read Request body with error: %s", err.Error())
		return ctx
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println(string(body))
	return ctx
}

func newDefaultBeforeFunc() []ReqFunc {
	var f []ReqFunc
	f = append(f, DefaultBeforeFunc)
	return f
}

func DefaultAfterFunc(ctx context.Context, resp *http.Response) context.Context {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read Response body with error: %s", err.Error())
		return ctx
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println(string(body))
	return ctx
}

func EmptyBeforeFunc(ctx context.Context, req *http.Request) context.Context {
	return ctx
}

func EmptyAfterFunc(ctx context.Context, resp *http.Response) context.Context {
	return ctx
}

func newDefaultAfterFunc() []ResponseFunc {
	var f []ResponseFunc
	f = append(f, DefaultAfterFunc)
	return f
}

func AppendAfterFunc(f ...ResponseFunc) Option {
	return func(options *Options) {
		if options.after == nil {
			options.after = []ResponseFunc{}
		}
		options.after = append(options.after, f...)
	}
}

func AppendBeforeFunc(f ...ReqFunc) Option {
	return func(options *Options) {
		if options.before == nil {
			options.before = []ReqFunc{}
		}
		options.before = append(options.before, f...)
	}
}

func SetAfterFunc(f ...ResponseFunc) Option {
	return func(options *Options) {
		if options.after == nil {
			options.after = []ResponseFunc{}
		}
		options.after = f
	}
}

func SetBeforeFunc(f ...ReqFunc) Option {
	return func(options *Options) {
		if options.before == nil {
			options.before = []ReqFunc{}
		}
		options.before = f
	}
}

func DisableLog() Option {
	return SetLogger(NullLogger)
}

func SetLogger(logger Logger) Option {
	return func(options *Options) {
		options.logger = logger
	}
}

func DefaultCharset() Option {
	return func(options *Options) {
		options.config.charset = "utf-8"
	}
}

func DefaultVersion() Option {
	return func(options *Options) {
		options.config.version = "1.0"
	}
}

func DefaultFormat() Option {
	return func(options *Options) {
		options.config.format = "JSON"
	}
}

func SetLocation(loc *time.Location) Option {
	return func(options *Options) {
		options.config.loc = loc
	}
}

func DefaultLocation() Option {
	return func(options *Options) {
		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			options.logger.Error(err.Error())
			options.logger.Infof("using UTC")
			return
		}
		SetLocation(loc)(options)
	}
}

func AppId(appId string) Option {
	return func(options *Options) {
		options.config.appId = appId
	}
}

func PrivateKey(privateKey string) Option {
	return func(options *Options) {
		options.config.privateKey = privateKey
	}
}

func PrivateKeyType(privateKeyType string) Option {
	return func(options *Options) {
		options.config.privateKeyType = privateKeyType
	}
}

func ReturnUrl(returnUrl string) Option {
	return func(options *Options) {
		options.config.returnUrl = returnUrl
	}
}

func AppAuthToken(appAuthToken string) Option {
	return func(options *Options) {
		options.config.appAuthToken = appAuthToken
	}
}

func NotifyUrl(notifyUrl string) Option {
	return func(options *Options) {
		options.config.notifyUrl = notifyUrl
	}
}

func SignType(signType string) Option {
	return func(options *Options) {
		options.config.signType = signType
	}
}

func AuthToken(authToken string) Option {
	return func(options *Options) {
		options.config.authToken = authToken
	}
}

func AppCertSnPath(path string) Option {
	return func(options *Options) {
		sn, err := loadCertSN(path)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.appCertSN = sn
	}
}

func AppCertSnBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := loadCertSN(data)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.appCertSN = sn
	}
}

func RootCertSnPath(path string) Option {
	return func(options *Options) {
		sn, err := loadRootCertSN(path)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.aliPayRootCertSN = sn
	}
}

func RootCertSnBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := loadRootCertSN(data)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.aliPayRootCertSN = sn
	}
}

func PublicCertSnPath(path string) Option {
	return func(options *Options) {
		sn, err := loadCertSN(path)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.aliPayPublicCertSN = sn
	}
}

func PublicCertSnBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := loadCertSN(data)
		if err != nil {
			options.logger.Error(err.Error())
		}
		options.config.aliPayPublicCertSN = sn
	}
}
func Production(input bool) Option {
	return func(options *Options) {
		options.config.production = input
	}
}
