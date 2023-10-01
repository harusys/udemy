package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	// メタデータから取得
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// log.Printf("metadata: %v¥n", md)
		// gRPC Gateway 用（HTTP Server）
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		if clientIps := md.Get(xForwardedForHeader); len(clientIps) > 0 {
			mtdt.ClientIp = clientIps[0]
		}

		// gRPC 用（gRPC Server）
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
	}

	// ピア情報から取得
	if p, ok := peer.FromContext(ctx); ok {
		// log.Printf("peer information: %v¥n", p)
		// gRPC 用（gRPC Server）
		mtdt.ClientIp = p.Addr.String()
	}

	return mtdt
}
