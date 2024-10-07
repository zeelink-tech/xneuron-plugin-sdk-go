package plugin

import (
	"context"
	"github.com/zeelink-tech/xneuron-plugin-sdk-go/proto"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client proto.DriverClient }

func (m *GRPCClient) SetConfig(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (m GRPCClient) Start(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (m GRPCClient) Restart(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (m GRPCClient) Stop(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (m GRPCClient) Get(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (m GRPCClient) Set(req *Request) (*Response, error) {
	res, err := m.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Driver
}

func (m *GRPCServer) SetConfig(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	res, err := m.Impl.SetConfig(&Request{
		Req: req.Request,
	})
	if err != nil {
		return nil, err
	}

	return &proto.ResponseResult{Data: res.Data}, nil
}

func (m *GRPCServer) Start(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := m.Impl.Start(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (m *GRPCServer) Restart(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := m.Impl.Restart(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (m *GRPCServer) Stop(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := m.Impl.Stop(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (m *GRPCServer) Get(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := m.Impl.Get(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (m *GRPCServer) Set(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := m.Impl.Set(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}
