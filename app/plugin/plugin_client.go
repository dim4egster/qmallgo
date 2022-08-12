// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package plugin

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pluginpb "github.com/dim4egster/avalanchego/proto/pb/plugin"
)

type Client struct {
	client pluginpb.NodeClient
}

// NewServer returns an app instance connected to a remote app instance
func NewClient(node pluginpb.NodeClient) *Client {
	return &Client{
		client: node,
	}
}

func (c *Client) Start() error {
	_, err := c.client.Start(context.Background(), &emptypb.Empty{})
	return err
}

func (c *Client) Stop() error {
	_, err := c.client.Stop(context.Background(), &emptypb.Empty{})
	return err
}

func (c *Client) ExitCode() (int, error) {
	resp, err := c.client.ExitCode(context.Background(), &emptypb.Empty{})
	if err != nil {
		return 0, err
	}
	return int(resp.ExitCode), nil
}
