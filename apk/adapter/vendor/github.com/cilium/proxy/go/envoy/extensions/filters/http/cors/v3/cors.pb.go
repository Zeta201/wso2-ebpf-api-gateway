// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.12
// source: envoy/extensions/filters/http/cors/v3/cors.proto

package corsv3

import (
	v31 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Cors filter config. Set this in
// ref:`http_filters <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.http_filters>`
// to enable the CORS filter.
//
// Please note that the :ref:`CorsPolicy <envoy_v3_api_msg_extensions.filters.http.cors.v3.CorsPolicy>`
// must be configured in the “RouteConfiguration“ as “typed_per_filter_config“ at some level to make the filter work.
type Cors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cors) Reset() {
	*x = Cors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cors) ProtoMessage() {}

func (x *Cors) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cors.ProtoReflect.Descriptor instead.
func (*Cors) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescGZIP(), []int{0}
}

// Per route configuration for the CORS filter. This configuration should be configured in the “RouteConfiguration“ as “typed_per_filter_config“ at some level to
// make the filter work.
// [#next-free-field: 10]
type CorsPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies string patterns that match allowed origins. An origin is allowed if any of the
	// string matchers match.
	AllowOriginStringMatch []*v3.StringMatcher `protobuf:"bytes,1,rep,name=allow_origin_string_match,json=allowOriginStringMatch,proto3" json:"allow_origin_string_match,omitempty"`
	// Specifies the content for the “access-control-allow-methods“ header.
	AllowMethods string `protobuf:"bytes,2,opt,name=allow_methods,json=allowMethods,proto3" json:"allow_methods,omitempty"`
	// Specifies the content for the “access-control-allow-headers“ header.
	AllowHeaders string `protobuf:"bytes,3,opt,name=allow_headers,json=allowHeaders,proto3" json:"allow_headers,omitempty"`
	// Specifies the content for the “access-control-expose-headers“ header.
	ExposeHeaders string `protobuf:"bytes,4,opt,name=expose_headers,json=exposeHeaders,proto3" json:"expose_headers,omitempty"`
	// Specifies the content for the “access-control-max-age“ header.
	MaxAge string `protobuf:"bytes,5,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Specifies whether the resource allows credentials.
	AllowCredentials *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	// Specifies the % of requests for which the CORS filter is enabled.
	//
	// If neither “filter_enabled“, nor “shadow_enabled“ are specified, the CORS
	// filter will be enabled for 100% of the requests.
	//
	// If :ref:`runtime_key <envoy_v3_api_field_config.core.v3.RuntimeFractionalPercent.runtime_key>` is
	// specified, Envoy will lookup the runtime key to get the percentage of requests to filter.
	FilterEnabled *v31.RuntimeFractionalPercent `protobuf:"bytes,7,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies the % of requests for which the CORS policies will be evaluated and tracked, but not
	// enforced.
	//
	// This field is intended to be used when “filter_enabled“ is off. That field have to explicitly disable
	// the filter in order for this setting to take effect.
	//
	// If :ref:`runtime_key <envoy_v3_api_field_config.core.v3.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests for which it will evaluate
	// and track the request's “Origin“ to determine if it's valid but will not enforce any policies.
	ShadowEnabled *v31.RuntimeFractionalPercent `protobuf:"bytes,8,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	// Specify whether allow requests whose target server's IP address is more private than that from
	// which the request initiator was fetched.
	//
	// More details refer to https://developer.chrome.com/blog/private-network-access-preflight.
	AllowPrivateNetworkAccess *wrapperspb.BoolValue `protobuf:"bytes,9,opt,name=allow_private_network_access,json=allowPrivateNetworkAccess,proto3" json:"allow_private_network_access,omitempty"`
}

func (x *CorsPolicy) Reset() {
	*x = CorsPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorsPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorsPolicy) ProtoMessage() {}

func (x *CorsPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorsPolicy.ProtoReflect.Descriptor instead.
func (*CorsPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescGZIP(), []int{1}
}

func (x *CorsPolicy) GetAllowOriginStringMatch() []*v3.StringMatcher {
	if x != nil {
		return x.AllowOriginStringMatch
	}
	return nil
}

func (x *CorsPolicy) GetAllowMethods() string {
	if x != nil {
		return x.AllowMethods
	}
	return ""
}

func (x *CorsPolicy) GetAllowHeaders() string {
	if x != nil {
		return x.AllowHeaders
	}
	return ""
}

func (x *CorsPolicy) GetExposeHeaders() string {
	if x != nil {
		return x.ExposeHeaders
	}
	return ""
}

func (x *CorsPolicy) GetMaxAge() string {
	if x != nil {
		return x.MaxAge
	}
	return ""
}

func (x *CorsPolicy) GetAllowCredentials() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowCredentials
	}
	return nil
}

func (x *CorsPolicy) GetFilterEnabled() *v31.RuntimeFractionalPercent {
	if x != nil {
		return x.FilterEnabled
	}
	return nil
}

func (x *CorsPolicy) GetShadowEnabled() *v31.RuntimeFractionalPercent {
	if x != nil {
		return x.ShadowEnabled
	}
	return nil
}

func (x *CorsPolicy) GetAllowPrivateNetworkAccess() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowPrivateNetworkAccess
	}
	return nil
}

var File_envoy_extensions_filters_http_cors_v3_cors_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x73, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a,
	0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x72, 0x73, 0x22, 0xcb, 0x04, 0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x73, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5f, 0x0a, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x16,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x12, 0x47, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x55, 0x0a, 0x0e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x55, 0x0a, 0x0e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x5b, 0x0a, 0x1c, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x9f, 0x01, 0x0a, 0x33, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x43, 0x6f,
	0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x72, 0x73, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescData = file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDesc
)

func file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDescData
}

var file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_cors_v3_cors_proto_goTypes = []interface{}{
	(*Cors)(nil),                         // 0: envoy.extensions.filters.http.cors.v3.Cors
	(*CorsPolicy)(nil),                   // 1: envoy.extensions.filters.http.cors.v3.CorsPolicy
	(*v3.StringMatcher)(nil),             // 2: envoy.type.matcher.v3.StringMatcher
	(*wrapperspb.BoolValue)(nil),         // 3: google.protobuf.BoolValue
	(*v31.RuntimeFractionalPercent)(nil), // 4: envoy.config.core.v3.RuntimeFractionalPercent
}
var file_envoy_extensions_filters_http_cors_v3_cors_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.cors.v3.CorsPolicy.allow_origin_string_match:type_name -> envoy.type.matcher.v3.StringMatcher
	3, // 1: envoy.extensions.filters.http.cors.v3.CorsPolicy.allow_credentials:type_name -> google.protobuf.BoolValue
	4, // 2: envoy.extensions.filters.http.cors.v3.CorsPolicy.filter_enabled:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	4, // 3: envoy.extensions.filters.http.cors.v3.CorsPolicy.shadow_enabled:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	3, // 4: envoy.extensions.filters.http.cors.v3.CorsPolicy.allow_private_network_access:type_name -> google.protobuf.BoolValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_cors_v3_cors_proto_init() }
func file_envoy_extensions_filters_http_cors_v3_cors_proto_init() {
	if File_envoy_extensions_filters_http_cors_v3_cors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cors); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorsPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_cors_v3_cors_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_cors_v3_cors_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_cors_v3_cors_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_cors_v3_cors_proto = out.File
	file_envoy_extensions_filters_http_cors_v3_cors_proto_rawDesc = nil
	file_envoy_extensions_filters_http_cors_v3_cors_proto_goTypes = nil
	file_envoy_extensions_filters_http_cors_v3_cors_proto_depIdxs = nil
}