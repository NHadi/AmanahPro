package services

import (
	"context"
	"fmt"
	"log"

	pb "github.com/NHadi/AmanahPro-common/protos"
)

// GrpcSphService implements the gRPC service for SPH
type GrpcSphService struct {
	sphService *SphService
	pb.UnimplementedSphServiceServer
}

// NewGrpcSphService creates a new instance of GrpcSphService
func NewGrpcSphService(sphService *SphService) *GrpcSphService {
	return &GrpcSphService{sphService: sphService}
}

// GetSphDetails handles the gRPC call to fetch SPH sections and details
func (s *GrpcSphService) GetSphDetails(ctx context.Context, req *pb.GetSphDetailsRequest) (*pb.GetSphDetailsResponse, error) {
	log.Printf("Fetching SPH details for SphId: %d", req.SphId)

	// Fetch SPH by ID
	sph, err := s.sphService.GetSphByID(int(req.SphId))
	if err != nil {
		log.Printf("Error fetching SPH: %v", err)
		return nil, fmt.Errorf("failed to fetch SPH: %w", err)
	}
	if sph == nil {
		return nil, fmt.Errorf("SPH with ID %d not found", req.SphId)
	}

	// Map sections and details to gRPC response
	var sections []*pb.SphSection
	for _, section := range sph.Sections {
		grpcSection := &pb.SphSection{
			SphSectionId: int32(section.SphSectionId),
			SectionTitle: *section.SectionTitle,
		}

		for _, detail := range section.Details {
			grpcDetail := &pb.SphDetail{
				SphDetailId:     int32(detail.SphDetailId),
				ItemDescription: *detail.ItemDescription,
				Quantity:        *detail.Quantity,
				Unit:            *detail.Unit,
				UnitPrice:       *detail.UnitPrice,
				DiscountPrice:   *detail.DiscountPrice,
			}
			grpcSection.Details = append(grpcSection.Details, grpcDetail)
		}

		sections = append(sections, grpcSection)
	}

	return &pb.GetSphDetailsResponse{Sections: sections}, nil
}
