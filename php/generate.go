// MIT License
//
// Copyright (c) 2022 - present Open Swoole Group
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package php

import (
	desc "google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func Generate(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := &plugin.CodeGeneratorResponse{}

	for _, file := range req.ProtoFile {
		for _, service := range file.Service {
			resp.File = append(resp.File, generateInterface(req, file, service))
			resp.File = append(resp.File, generateService(req, file, service))
			resp.File = append(resp.File, generateClient(req, file, service))
		}
	}

	return resp
}

func generateInterface(
	req *plugin.CodeGeneratorRequest,
	file *desc.FileDescriptorProto,
	service *desc.ServiceDescriptorProto,
) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    str(interfacefilename(file, service.Name)),
		Content: str(interfacebody(req, file, service)),
	}
}

func generateService(
	req *plugin.CodeGeneratorRequest,
	file *desc.FileDescriptorProto,
	service *desc.ServiceDescriptorProto,
) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    str(servicefilename(file, service.Name)),
		Content: str(servicebody(req, file, service)),
	}
}

func generateClient(
	req *plugin.CodeGeneratorRequest,
	file *desc.FileDescriptorProto,
	service *desc.ServiceDescriptorProto,
) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    str(clientfilename(file, service.Name)),
		Content: str(clientbody(req, file, service)),
	}
}

func str(str string) *string {
	return &str
}
