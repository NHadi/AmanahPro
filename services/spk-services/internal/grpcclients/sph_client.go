package grpcclients

import (
	pb "AmanahPro/services/sph-services/protos"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SphGrpcClient struct {
	client pb.SphServiceClient
	conn   *grpc.ClientConn
}

// NewSphGrpcClient initializes the SPH gRPC client
func NewSphGrpcClient(address string) (*SphGrpcClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to SPH service: %v", err)
		return nil, err
	}

	client := pb.NewSphServiceClient(conn)
	return &SphGrpcClient{
		client: client,
		conn:   conn,
	}, nil
}

// Close closes the gRPC connection
func (c *SphGrpcClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// GetSphDetails fetches SPH details
func (c *SphGrpcClient) GetSphDetails(sphId int32) (*pb.GetSphDetailsResponse, error) {
	req := &pb.GetSphDetailsRequest{SphId: sphId}
	return c.client.GetSphDetails(context.Background(), req)
}
