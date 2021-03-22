// Copyright 2021 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"

	"github.com/redhat-marketplace/redhat-marketplace-operator/airgap/v2/apis/fileserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	fileserver.UnimplementedFileServerServer
}

func (*Server) UploadFile(stream fileserver.FileServer_UploadFileServer) error {
	return status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) ListFileMetadata(lis *fileserver.ListFileMetadataRequest, stream fileserver.FileServer_ListFileMetadataServer) error {
	return status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) GetFileMetadata(ctx context.Context, gfmr *fileserver.GetFileMetadataRequest) (*fileserver.GetFileMetadataResponse, error) {
	return nil, status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) DownloadFile(dfr *fileserver.DownloadFileRequest, stream fileserver.FileServer_DownloadFileServer) error {
	return status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) UpdateFileMetadata(ctx context.Context, ufmr *fileserver.UpdateFileMetadataRequest) (*fileserver.UpdateFileMetadataResponse, error) {
	return nil, status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) DeleteFile(ctx context.Context, dfr *fileserver.DeleteFileRequest) (*fileserver.DeleteFileResponse, error) {
	return nil, status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}

func (*Server) CleanTombstones(ctx context.Context, ctr *fileserver.CleanTombstonesRequest) (*fileserver.CleanTombstonesResponse, error) {
	return nil, status.Errorf(
		codes.Unimplemented,
		"Method Unimplemented",
	)
}
