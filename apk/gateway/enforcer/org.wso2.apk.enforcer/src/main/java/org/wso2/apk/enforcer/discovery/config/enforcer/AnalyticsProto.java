// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/analytics.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

public final class AnalyticsProto {
  private AnalyticsProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_Analytics_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_Analytics_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_Analytics_PropertiesEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_Analytics_PropertiesEntry_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n.wso2/discovery/config/enforcer/analyti" +
      "cs.proto\022\036wso2.discovery.config.enforcer" +
      "\032,wso2/discovery/config/enforcer/service" +
      ".proto\0328wso2/discovery/config/enforcer/a" +
      "nalytics_publisher.proto\"\251\002\n\tAnalytics\022\017" +
      "\n\007enabled\030\001 \001(\010\022O\n\023analytics_publisher\030\002" +
      " \003(\01322.wso2.discovery.config.enforcer.An" +
      "alyticsPublisher\0228\n\007service\030\003 \001(\0132\'.wso2" +
      ".discovery.config.enforcer.Service\022M\n\npr" +
      "operties\030\004 \003(\01329.wso2.discovery.config.e" +
      "nforcer.Analytics.PropertiesEntry\0321\n\017Pro" +
      "pertiesEntry\022\013\n\003key\030\001 \001(\t\022\r\n\005value\030\002 \001(\t" +
      ":\0028\001B\223\001\n/org.wso2.apk.enforcer.discovery" +
      ".config.enforcerB\016AnalyticsProtoP\001ZNgith" +
      "ub.com/envoyproxy/go-control-plane/wso2/" +
      "discovery/config/enforcer;enforcerb\006prot" +
      "o3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.wso2.apk.enforcer.discovery.config.enforcer.ServiceProto.getDescriptor(),
          org.wso2.apk.enforcer.discovery.config.enforcer.AnalyticsPublisherProto.getDescriptor(),
        });
    internal_static_wso2_discovery_config_enforcer_Analytics_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_Analytics_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_Analytics_descriptor,
        new java.lang.String[] { "Enabled", "AnalyticsPublisher", "Service", "Properties", });
    internal_static_wso2_discovery_config_enforcer_Analytics_PropertiesEntry_descriptor =
      internal_static_wso2_discovery_config_enforcer_Analytics_descriptor.getNestedTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_Analytics_PropertiesEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_Analytics_PropertiesEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    org.wso2.apk.enforcer.discovery.config.enforcer.ServiceProto.getDescriptor();
    org.wso2.apk.enforcer.discovery.config.enforcer.AnalyticsPublisherProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}