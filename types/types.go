package types

import (
	"github.com/YenXXXW/FileDuplicateFinderService/genproto/fileduppb"
	grpc "google.golang.org/grpc"
)

type FileDuppbService interface {
	DuplicateFinder(grpc.ClientStreamingServer[fileduppb.FileData, fileduppb.DuplicateResult]) error
}
