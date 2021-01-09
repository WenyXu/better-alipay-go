/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 4:22
*/

/*
	Module for setup AppConfig, ClientConfig
*/
package options

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/WenyXu/better-alipay-go/cert"
	"github.com/WenyXu/better-alipay-go/config"
	"github.com/WenyXu/better-alipay-go/global"
	"github.com/WenyXu/better-alipay-go/logger"
	m "github.com/WenyXu/better-alipay-go/m"
)

var (
	DefaultTransport                 = http.DefaultTransport
	DefaultMakeReqFunc MakeReqFunc   = NewDefaultMakeReqFunc
	DefaultDecFunc     DecFunc       = NewDefaultDecFunc
	DefaultLocation                  = NewDefaultLocation()
	DefaultLogger      logger.Logger = logger.StdLogger
)

// SetDefaultTransport Set Global Default Transport
func SetDefaultTransport(f http.RoundTripper) {
	DefaultTransport = f
}

// SetDefaultMakeReqFunc Set Global Default MakeReqFunc
func SetDefaultMakeReqFunc(f MakeReqFunc) {
	DefaultMakeReqFunc = f
}

// SetDefaultDecFunc Set Global Default DecFunc
func SetDefaultDecFunc(f DecFunc) {
	DefaultDecFunc = f
}

// SetDefaultLocation Set Global Default Location
func SetDefaultLocation(f Option) {
	DefaultLocation = f
}

// SetDefaultLogger Set Global Default Logger
func SetDefaultLogger(f logger.Logger) {
	DefaultLogger = f
}

type Option func(*Options)

// ReqFunc a hook Before the request started
type ReqFunc func(context context.Context, req *http.Request) context.Context

// ResponseFunc a hook After the request finished
type ResponseFunc func(context context.Context, response *http.Response) context.Context

// MakeReqFunc make request
type MakeReqFunc func(context context.Context, method string, m m.M, config config.Config) (*http.Request, error)

// DecFunc convert response body into a struct or map
type DecFunc func(context.Context, interface{}, interface{}) error

// Options for alipay sdk
type Options struct {
	Config config.Config

	// hooks
	Before []ReqFunc
	After  []ResponseFunc

	Dec     DecFunc
	MakeReq MakeReqFunc

	// other custom options
	// can be stored in a context
	Context context.Context

	Transport http.RoundTripper

	Logger logger.Logger
}

// Copy create a current Options' Copy
func (o Options) Copy() Options {
	cp := Options{
		Config:    o.Config,
		Before:    []ReqFunc{},
		After:     []ResponseFunc{},
		Dec:       o.Dec,
		MakeReq:   o.MakeReq,
		Context:   o.Context,
		Transport: o.Transport,
		Logger:    o.Logger,
	}
	if len(o.After) != 0 {
		cp.After = append(cp.After, o.After...)
	}
	if len(o.Before) != 0 {
		cp.Before = append(cp.Before, o.Before...)
	}
	return o
}

// NewDefaultMakeReqFunc default make request func
func NewDefaultMakeReqFunc(ctx context.Context, method string, param m.M, c config.Config) (*http.Request, error) {

	body, err := m.CombineMakeMapEndpointFunc(
		config.SetMethod(method),
		config.SetPublicParam(c),
		config.SetOptionalParam(c),
		config.SetBizContent(param),
		config.SignParam(c),
	)

	reader := strings.NewReader(m.FormatURLParam(body))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Url(), reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", global.ContentType)
	return req, nil
}

// NewDefaultDecFunc encode json into target struct
func NewDefaultDecFunc(ctx context.Context, input interface{}, result interface{}) (err error) {
	resp := input.(*http.Response)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Request Error, StatusCode = %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Transport: DefaultTransport,
		Context:   context.Background(),
		MakeReq:   DefaultMakeReqFunc,
		Dec:       DefaultDecFunc,
		Logger:    DefaultLogger,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// DefaultOptions set DefaultAfterFunc, DefaultBeforeFunc, DefaultLogger
func DefaultOptions(opts ...Option) Options {
	return newOptions(
		append(
			append(
				[]Option{},
				SetAfterFunc(DefaultAfterFunc),
				SetBeforeFunc(DefaultBeforeFunc),
				SetLogger(DefaultLogger),
			),
			opts...,
		)...,
	)
}

// NewOptions
func NewOptions(opts ...Option) Options {
	return newOptions(opts...)
}

// Transport set Transport
func Transport(transport http.RoundTripper) Option {
	return func(options *Options) {
		options.Transport = transport
	}
}

// DefaultBeforeFunc log the request body
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

// DefaultAfterFunc log the response body
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

// EmptyBeforeFunc do nothing
func EmptyBeforeFunc(ctx context.Context, req *http.Request) context.Context {
	return ctx
}

// EmptyAfterFunc do nothing
func EmptyAfterFunc(ctx context.Context, resp *http.Response) context.Context {
	return ctx
}

// AppendAfterFunc append a ResponseFunc after current list
func AppendAfterFunc(f ...ResponseFunc) Option {
	return func(options *Options) {
		if options.After == nil {
			options.After = []ResponseFunc{}
		}
		options.After = append(options.After, f...)
	}
}

// AppendBeforeFunc append a ReqFunc after current list
func AppendBeforeFunc(f ...ReqFunc) Option {
	return func(options *Options) {
		if options.Before == nil {
			options.Before = []ReqFunc{}
		}
		options.Before = append(options.Before, f...)
	}
}

// SetAfterFunc set ResponseFunc
func SetAfterFunc(f ...ResponseFunc) Option {
	return func(options *Options) {
		if options.After == nil {
			options.After = []ResponseFunc{}
		}
		options.After = f
	}
}

// SetBeforeFunc set ReqFunc
func SetBeforeFunc(f ...ReqFunc) Option {
	return func(options *Options) {
		if options.Before == nil {
			options.Before = []ReqFunc{}
		}
		options.Before = f
	}
}

// DisableLog set NullLogger
func DisableLog() Option {
	return SetLogger(logger.NullLogger)
}

// SetLogger set logger
func SetLogger(Logger logger.Logger) Option {
	return func(options *Options) {
		options.Logger = Logger
	}
}

// DefaultCharset utf-8
func DefaultCharset() Option {
	return func(options *Options) {
		options.Config.Charset = "utf-8"
	}
}

// DefaultVersion 1.0
func DefaultVersion() Option {
	return func(options *Options) {
		options.Config.Version = "1.0"
	}
}

// DefaultVersion JSON
func DefaultFormat() Option {
	return func(options *Options) {
		options.Config.Format = "JSON"
	}
}

// SetLocation time.Location
func SetLocation(loc *time.Location) Option {
	return func(options *Options) {
		options.Config.Loc = loc
	}
}

// NewDefaultLocation set location as Asia/Shanghai
func NewDefaultLocation() Option {
	return func(options *Options) {
		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			options.Logger.Error(err.Error())
			options.Logger.Infof("using UTC")
			return
		}
		SetLocation(loc)(options)
	}
}

// AppId set AppId
func AppId(appId string) Option {
	return func(options *Options) {
		options.Config.AppId = appId
	}
}

// PrivateKey set PrivateKey string
func PrivateKey(privateKey string) Option {
	return func(options *Options) {
		options.Config.PrivateKey = privateKey
	}
}

// PrivateKey set PrivateKeyType
func PrivateKeyType(privateKeyType string) Option {
	return func(options *Options) {
		options.Config.PrivateKeyType = privateKeyType
	}
}

// ReturnUrl set ReturnUrl
func ReturnUrl(returnUrl string) Option {
	return func(options *Options) {
		options.Config.ReturnUrl = returnUrl
	}
}

// AppAuthToken set AppAuthToken
func AppAuthToken(appAuthToken string) Option {
	return func(options *Options) {
		options.Config.AppAuthToken = appAuthToken
	}
}

// NotifyUrl set NotifyUrl
func NotifyUrl(notifyUrl string) Option {
	return func(options *Options) {
		options.Config.NotifyUrl = notifyUrl
	}
}

// SignType set SignType
func SignType(signType string) Option {
	return func(options *Options) {
		options.Config.SignType = signType
	}
}

// AuthToken set AuthToken
func AuthToken(authToken string) Option {
	return func(options *Options) {
		options.Config.AuthToken = authToken
	}
}

// AppCertPath set AppCertSn form path
func AppCertPath(path string) Option {
	return func(options *Options) {
		sn, err := cert.LoadCertSN(path)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		options.Config.AppCertSN = sn
	}
}

// AppCertBytes set AppCertSn form bytes
func AppCertBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := cert.LoadCertSN(data)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		AppCertSn(sn)(options)
	}
}

// AppCertSn set AppCertSn
func AppCertSn(sn string) Option {
	return func(options *Options) {
		options.Config.AppCertSN = sn
	}
}

// RootCertPath set RootCertSn from path
func RootCertPath(path string) Option {
	return func(options *Options) {
		sn, err := cert.LoadRootCertSN(path)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		RootCertSn(sn)(options)
	}
}

// RootCertBytes set RootCertSn from bytes
func RootCertBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := cert.LoadRootCertSN(data)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		RootCertSn(sn)(options)
	}
}

// RootCertSn set RootCertSn
func RootCertSn(sn string) Option {
	return func(options *Options) {
		options.Config.AliPayRootCertSN = sn
	}
}

// PublicCertPath set PublicCertSn from path
func PublicCertPath(path string) Option {
	return func(options *Options) {
		sn, err := cert.LoadCertSN(path)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		PublicCertSn(sn)(options)
	}
}

// PublicCertBytes set PublicCertSn from bytes
func PublicCertBytes(data []byte) Option {
	return func(options *Options) {
		sn, err := cert.LoadCertSN(data)
		if err != nil {
			options.Logger.Error(err.Error())
		}
		PublicCertSn(sn)(options)
	}
}

// PublicCertSn set PublicCertSn
func PublicCertSn(sn string) Option {
	return func(options *Options) {
		options.Config.AliPayPublicCertSN = sn
	}
}

// Production set Production flag
func Production(input bool) Option {
	return func(options *Options) {
		options.Config.Production = input
	}
}
