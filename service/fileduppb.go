package service

import (
	"github.com/YenXXXW/FileDuplicateFinderService/genproto/fileduppb"
	grpc "google.golang.org/grpc"
)

type FileDuppbService struct {
}

func (f *FileDuppbService) DuplicateFinder(grpc.ClientStreamingServer[fileduppb.FileData, fileduppb.DuplicateResult]) error {
	//prefrom hasing and check
	return nil
}
