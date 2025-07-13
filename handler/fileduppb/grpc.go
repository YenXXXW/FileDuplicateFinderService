package handler

import (
	"github.com/YenXXXW/FileDuplicateFinderService/genproto/fileduppb"
	"github.com/YenXXXW/FileDuplicateFinderService/types"
	"google.golang.org/grpc"
)

type FileDuppbHandler struct {
	fileDuupbService types.FileDuppbService
	fileduppb.UnimplementedDuplicateFinderServer
}

func NewFileDuppbService(grpc *grpc.Server, fileduppbService types.FileDuppbService) {
	grpcHandler := &FileDuppbHandler{
		fileDuupbService: fileduppbService,
	}

	fileduppb.RegisterDuplicateFinderServer(grpc, grpcHandler)
}

func (h *FileDuppbHandler) DuplicateFinder(grpc.ClientStreamingServer[fileduppb.FileData, fileduppb.DuplicateResult]) error {
	return nil
}
