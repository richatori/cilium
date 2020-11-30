// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new daemon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for daemon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteClusterNodesNeigh(params *DeleteClusterNodesNeighParams) (*DeleteClusterNodesNeighOK, *DeleteClusterNodesNeighErrors, error)

	GetClusterNodes(params *GetClusterNodesParams) (*GetClusterNodesOK, error)

	GetConfig(params *GetConfigParams) (*GetConfigOK, error)

	GetDebuginfo(params *GetDebuginfoParams) (*GetDebuginfoOK, error)

	GetHealthz(params *GetHealthzParams) (*GetHealthzOK, error)

	GetMap(params *GetMapParams) (*GetMapOK, error)

	GetMapName(params *GetMapNameParams) (*GetMapNameOK, error)

	PatchConfig(params *PatchConfigParams) (*PatchConfigOK, error)

	PutClusterNodesNeigh(params *PutClusterNodesNeighParams) (*PutClusterNodesNeighCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteClusterNodesNeigh removes node as a neighbor from cluster

  Removes a node as a neighbor from the cluster. This operation removes
the node from all other nodes' neighbor tables.

*/
func (a *Client) DeleteClusterNodesNeigh(params *DeleteClusterNodesNeighParams) (*DeleteClusterNodesNeighOK, *DeleteClusterNodesNeighErrors, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterNodesNeighParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterNodesNeigh",
		Method:             "DELETE",
		PathPattern:        "/cluster/nodes/neigh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterNodesNeighReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteClusterNodesNeighOK:
		return value, nil, nil
	case *DeleteClusterNodesNeighErrors:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for daemon: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClusterNodes gets nodes information stored in the cilium agent
*/
func (a *Client) GetClusterNodes(params *GetClusterNodesParams) (*GetClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterNodes",
		Method:             "GET",
		PathPattern:        "/cluster/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterNodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConfig gets configuration of cilium daemon

  Returns the configuration of the Cilium daemon.

*/
func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfig",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDebuginfo retrieves information about the agent and evironment for debugging
*/
func (a *Client) GetDebuginfo(params *GetDebuginfoParams) (*GetDebuginfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDebuginfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDebuginfo",
		Method:             "GET",
		PathPattern:        "/debuginfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDebuginfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDebuginfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDebuginfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetHealthz gets health of cilium daemon

  Returns health and status information of the Cilium daemon and related
components such as the local container runtime, connected datastore,
Kubernetes integration and Hubble.

*/
func (a *Client) GetHealthz(params *GetHealthzParams) (*GetHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthzParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHealthz",
		Method:             "GET",
		PathPattern:        "/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHealthzOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealthz: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMap lists all open maps
*/
func (a *Client) GetMap(params *GetMapParams) (*GetMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMap",
		Method:             "GET",
		PathPattern:        "/map",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMapReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMapOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMap: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMapName retrieves contents of b p f map
*/
func (a *Client) GetMapName(params *GetMapNameParams) (*GetMapNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMapNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMapName",
		Method:             "GET",
		PathPattern:        "/map/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMapNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMapNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMapName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchConfig modifies daemon configuration

  Updates the daemon configuration by applying the provided
ConfigurationMap and regenerates & recompiles all required datapath
components.

*/
func (a *Client) PatchConfig(params *PatchConfigParams) (*PatchConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchConfig",
		Method:             "PATCH",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutClusterNodesNeigh inserts node as a neighbor into cluster

  Inserts a node as a neighbor into the cluster, rather than a full node
(running Cilium). This operation informs all full nodes of this new
neighbor.

*/
func (a *Client) PutClusterNodesNeigh(params *PutClusterNodesNeighParams) (*PutClusterNodesNeighCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterNodesNeighParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterNodesNeigh",
		Method:             "PUT",
		PathPattern:        "/cluster/nodes/neigh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterNodesNeighReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterNodesNeighCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutClusterNodesNeigh: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
