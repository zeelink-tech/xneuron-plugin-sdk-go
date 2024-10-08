package plugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/zeelink-tech/xneuron-plugin-sdk-go/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "DRIVER_PLUGIN",
	MagicCookieValue: "ON",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"driver_grpc": &DriverGRPCPlugin{},
}

type Response struct {
	Data      string `json:"data"`
	RequestID string `json:"requestID"`
}

type Request struct {
	BrokerID  uint32 `json:"brokerID"`
	Req       string `json:"req"`
	RequestID string `json:"requestID"`
}

type Driver interface {
	// SetConfig 配置驱动，用户自定义实现
	SetConfig(req *Request) (*Response, error)
	// Start 驱动采集启动，用户自定义实现
	Start(req *Request) (*Response, error)
	// Restart 驱动重启，用户自定义实现
	Restart(req *Request) (*Response, error)
	// Stop 驱动停止，用户自定义实现
	Stop(req *Request) (*Response, error)
	// Get 召测，用户自定义实现
	Get(req *Request) (*Response, error)
	// Set 置数，用户自定义实现
	Set(req *Request) (*Response, error)
}

// Report 驱动 --> 宿主
type Report interface {
	Post(req *Request) (*Response, error)
	State(req *Request) (*Response, error)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type DriverGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Driver
}

func (p *DriverGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterDriverServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *DriverGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewDriverClient(c)}, nil
}
