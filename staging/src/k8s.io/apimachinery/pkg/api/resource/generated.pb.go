/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource/generated.proto
// DO NOT EDIT!

/*
	Package resource is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource/generated.proto

	It has these top-level messages:
		Quantity
*/
package resource

import "github.com/gogo/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Quantity) Reset()                    { *m = Quantity{} }
func (*Quantity) ProtoMessage()               {}
func (*Quantity) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.apimachinery.pkg.api.resource.Quantity")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xa1, 0x4e, 0x03, 0x41,
	0x10, 0x86, 0x77, 0x0d, 0x29, 0x95, 0x0d, 0x21, 0xa4, 0x62, 0xaf, 0x21, 0x08, 0x0c, 0x3b, 0x02,
	0xd3, 0x20, 0xf1, 0x08, 0x90, 0xb8, 0xbb, 0xeb, 0xb0, 0xdd, 0x1c, 0xdd, 0xbd, 0xcc, 0xce, 0x92,
	0xd4, 0x55, 0x22, 0x2b, 0x91, 0xbd, 0xb7, 0xa9, 0xac, 0xac, 0x40, 0x70, 0xcb, 0x8b, 0x90, 0x5e,
	0xdb, 0x84, 0x90, 0xe0, 0xe6, 0xfb, 0x27, 0xdf, 0xe4, 0x9f, 0xfe, 0x43, 0x35, 0x0e, 0xda, 0x7a,
	0xa8, 0x62, 0x81, 0xe4, 0x90, 0x31, 0xc0, 0x1b, 0xba, 0x89, 0x27, 0x38, 0x2c, 0xf2, 0xda, 0xce,
	0xf2, 0x72, 0x6a, 0x1d, 0xd2, 0x1c, 0xea, 0xca, 0xec, 0x02, 0x20, 0x0c, 0x3e, 0x52, 0x89, 0x60,
	0xd0, 0x21, 0xe5, 0x8c, 0x13, 0x5d, 0x93, 0x67, 0x3f, 0xb8, 0xda, 0x5b, 0xfa, 0xb7, 0xa5, 0xeb,
	0xca, 0xec, 0x02, 0x7d, 0xb4, 0x86, 0x37, 0xc6, 0xf2, 0x34, 0x16, 0xba, 0xf4, 0x33, 0x30, 0xde,
	0x78, 0xe8, 0xe4, 0x22, 0xbe, 0x74, 0xd4, 0x41, 0x37, 0xed, 0x8f, 0x0e, 0x6f, 0xff, 0xab, 0x12,
	0xd9, 0xbe, 0x82, 0x75, 0x1c, 0x98, 0xfe, 0x36, 0xb9, 0x1c, 0xf7, 0x7b, 0x8f, 0x31, 0x77, 0x6c,
	0x79, 0x3e, 0x38, 0xef, 0x9f, 0x04, 0x26, 0xeb, 0xcc, 0x85, 0x1c, 0xc9, 0xeb, 0xd3, 0xa7, 0x03,
	0xdd, 0x9d, 0x7d, 0xac, 0x32, 0xf1, 0xde, 0x64, 0x62, 0xd9, 0x64, 0x62, 0xd5, 0x64, 0x62, 0xf1,
	0x39, 0x12, 0xf7, 0x7a, 0xdd, 0x2a, 0xb1, 0x69, 0x95, 0xd8, 0xb6, 0x4a, 0x2c, 0x92, 0x92, 0xeb,
	0xa4, 0xe4, 0x26, 0x29, 0xb9, 0x4d, 0x4a, 0x7e, 0x25, 0x25, 0x97, 0xdf, 0x4a, 0x3c, 0xf7, 0x8e,
	0xdf, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x5e, 0xda, 0xf9, 0x43, 0x01, 0x00, 0x00,
}
