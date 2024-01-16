package api

import (
	"github.com/zRich/go-backend/internal/server"
)

type APIServerConfig struct{}

type APIServer struct {
	endpoints []server.Endpoint
}

func (api *APIServer) Initialize() {
	// blockNumber
	// blockNumberendpoint := &GetBlockCountEndpoint{}

	// api.endpoints = append(api.endpoints, blockNumberendpoint)

	// // transactionCount
	// transactionCountendpoint := &GetTransactionCountEndpoint{}
	// api.endpoints = append(api.endpoints, transactionCountendpoint)
}

func (api *APIServer) GetEndpoints() []server.Endpoint {
	return api.endpoints
}

func NewAPIServer(config *APIServerConfig) *APIServer {
	server := &APIServer{}
	server.Initialize()
	return server
}
