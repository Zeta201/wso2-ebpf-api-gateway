// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/mutual_ssl.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

public final class MutualSSLProto {
  private MutualSSLProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_MutualSSL_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_MutualSSL_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n/wso2/discovery/config/enforcer/mutual_" +
      "ssl.proto\022\036wso2.discovery.config.enforce" +
      "r\"\220\001\n\tMutualSSL\022\031\n\021certificateHeader\030\001 \001" +
      "(\t\022\036\n\026enableClientValidation\030\002 \001(\010\022\037\n\027cl" +
      "ientCertificateEncode\030\003 \001(\010\022\'\n\037enableOut" +
      "boundCertificateHeader\030\004 \001(\010B\223\001\n/org.wso" +
      "2.apk.enforcer.discovery.config.enforcer" +
      "B\016MutualSSLProtoP\001ZNgithub.com/envoyprox" +
      "y/go-control-plane/wso2/discovery/config" +
      "/enforcer;enforcerb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_wso2_discovery_config_enforcer_MutualSSL_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_MutualSSL_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_MutualSSL_descriptor,
        new java.lang.String[] { "CertificateHeader", "EnableClientValidation", "ClientCertificateEncode", "EnableOutboundCertificateHeader", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
