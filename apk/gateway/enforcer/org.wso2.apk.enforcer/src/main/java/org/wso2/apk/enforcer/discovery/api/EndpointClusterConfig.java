// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/api/endpoint_cluster.proto

package org.wso2.apk.enforcer.discovery.api;

/**
 * Protobuf type {@code wso2.discovery.api.EndpointClusterConfig}
 */
public final class EndpointClusterConfig extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:wso2.discovery.api.EndpointClusterConfig)
    EndpointClusterConfigOrBuilder {
private static final long serialVersionUID = 0L;
  // Use EndpointClusterConfig.newBuilder() to construct.
  private EndpointClusterConfig(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private EndpointClusterConfig() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new EndpointClusterConfig();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private EndpointClusterConfig(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder subBuilder = null;
            if (retryConfig_ != null) {
              subBuilder = retryConfig_.toBuilder();
            }
            retryConfig_ = input.readMessage(org.wso2.apk.enforcer.discovery.api.RetryConfig.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(retryConfig_);
              retryConfig_ = subBuilder.buildPartial();
            }

            break;
          }
          case 18: {
            org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder subBuilder = null;
            if (timeoutConfig_ != null) {
              subBuilder = timeoutConfig_.toBuilder();
            }
            timeoutConfig_ = input.readMessage(org.wso2.apk.enforcer.discovery.api.TimeoutConfig.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(timeoutConfig_);
              timeoutConfig_ = subBuilder.buildPartial();
            }

            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.wso2.apk.enforcer.discovery.api.EndpointClusterProto.internal_static_wso2_discovery_api_EndpointClusterConfig_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.wso2.apk.enforcer.discovery.api.EndpointClusterProto.internal_static_wso2_discovery_api_EndpointClusterConfig_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.class, org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.Builder.class);
  }

  public static final int RETRYCONFIG_FIELD_NUMBER = 1;
  private org.wso2.apk.enforcer.discovery.api.RetryConfig retryConfig_;
  /**
   * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
   * @return Whether the retryConfig field is set.
   */
  @java.lang.Override
  public boolean hasRetryConfig() {
    return retryConfig_ != null;
  }
  /**
   * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
   * @return The retryConfig.
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.api.RetryConfig getRetryConfig() {
    return retryConfig_ == null ? org.wso2.apk.enforcer.discovery.api.RetryConfig.getDefaultInstance() : retryConfig_;
  }
  /**
   * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.api.RetryConfigOrBuilder getRetryConfigOrBuilder() {
    return getRetryConfig();
  }

  public static final int TIMEOUTCONFIG_FIELD_NUMBER = 2;
  private org.wso2.apk.enforcer.discovery.api.TimeoutConfig timeoutConfig_;
  /**
   * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
   * @return Whether the timeoutConfig field is set.
   */
  @java.lang.Override
  public boolean hasTimeoutConfig() {
    return timeoutConfig_ != null;
  }
  /**
   * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
   * @return The timeoutConfig.
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.api.TimeoutConfig getTimeoutConfig() {
    return timeoutConfig_ == null ? org.wso2.apk.enforcer.discovery.api.TimeoutConfig.getDefaultInstance() : timeoutConfig_;
  }
  /**
   * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.api.TimeoutConfigOrBuilder getTimeoutConfigOrBuilder() {
    return getTimeoutConfig();
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (retryConfig_ != null) {
      output.writeMessage(1, getRetryConfig());
    }
    if (timeoutConfig_ != null) {
      output.writeMessage(2, getTimeoutConfig());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (retryConfig_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getRetryConfig());
    }
    if (timeoutConfig_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getTimeoutConfig());
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig)) {
      return super.equals(obj);
    }
    org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig other = (org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig) obj;

    if (hasRetryConfig() != other.hasRetryConfig()) return false;
    if (hasRetryConfig()) {
      if (!getRetryConfig()
          .equals(other.getRetryConfig())) return false;
    }
    if (hasTimeoutConfig() != other.hasTimeoutConfig()) return false;
    if (hasTimeoutConfig()) {
      if (!getTimeoutConfig()
          .equals(other.getTimeoutConfig())) return false;
    }
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasRetryConfig()) {
      hash = (37 * hash) + RETRYCONFIG_FIELD_NUMBER;
      hash = (53 * hash) + getRetryConfig().hashCode();
    }
    if (hasTimeoutConfig()) {
      hash = (37 * hash) + TIMEOUTCONFIG_FIELD_NUMBER;
      hash = (53 * hash) + getTimeoutConfig().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code wso2.discovery.api.EndpointClusterConfig}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:wso2.discovery.api.EndpointClusterConfig)
      org.wso2.apk.enforcer.discovery.api.EndpointClusterConfigOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.wso2.apk.enforcer.discovery.api.EndpointClusterProto.internal_static_wso2_discovery_api_EndpointClusterConfig_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.wso2.apk.enforcer.discovery.api.EndpointClusterProto.internal_static_wso2_discovery_api_EndpointClusterConfig_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.class, org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.Builder.class);
    }

    // Construct using org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (retryConfigBuilder_ == null) {
        retryConfig_ = null;
      } else {
        retryConfig_ = null;
        retryConfigBuilder_ = null;
      }
      if (timeoutConfigBuilder_ == null) {
        timeoutConfig_ = null;
      } else {
        timeoutConfig_ = null;
        timeoutConfigBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.wso2.apk.enforcer.discovery.api.EndpointClusterProto.internal_static_wso2_discovery_api_EndpointClusterConfig_descriptor;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig getDefaultInstanceForType() {
      return org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.getDefaultInstance();
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig build() {
      org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig buildPartial() {
      org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig result = new org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig(this);
      if (retryConfigBuilder_ == null) {
        result.retryConfig_ = retryConfig_;
      } else {
        result.retryConfig_ = retryConfigBuilder_.build();
      }
      if (timeoutConfigBuilder_ == null) {
        result.timeoutConfig_ = timeoutConfig_;
      } else {
        result.timeoutConfig_ = timeoutConfigBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig) {
        return mergeFrom((org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig other) {
      if (other == org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig.getDefaultInstance()) return this;
      if (other.hasRetryConfig()) {
        mergeRetryConfig(other.getRetryConfig());
      }
      if (other.hasTimeoutConfig()) {
        mergeTimeoutConfig(other.getTimeoutConfig());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.wso2.apk.enforcer.discovery.api.RetryConfig retryConfig_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.api.RetryConfig, org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder, org.wso2.apk.enforcer.discovery.api.RetryConfigOrBuilder> retryConfigBuilder_;
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     * @return Whether the retryConfig field is set.
     */
    public boolean hasRetryConfig() {
      return retryConfigBuilder_ != null || retryConfig_ != null;
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     * @return The retryConfig.
     */
    public org.wso2.apk.enforcer.discovery.api.RetryConfig getRetryConfig() {
      if (retryConfigBuilder_ == null) {
        return retryConfig_ == null ? org.wso2.apk.enforcer.discovery.api.RetryConfig.getDefaultInstance() : retryConfig_;
      } else {
        return retryConfigBuilder_.getMessage();
      }
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public Builder setRetryConfig(org.wso2.apk.enforcer.discovery.api.RetryConfig value) {
      if (retryConfigBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        retryConfig_ = value;
        onChanged();
      } else {
        retryConfigBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public Builder setRetryConfig(
        org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder builderForValue) {
      if (retryConfigBuilder_ == null) {
        retryConfig_ = builderForValue.build();
        onChanged();
      } else {
        retryConfigBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public Builder mergeRetryConfig(org.wso2.apk.enforcer.discovery.api.RetryConfig value) {
      if (retryConfigBuilder_ == null) {
        if (retryConfig_ != null) {
          retryConfig_ =
            org.wso2.apk.enforcer.discovery.api.RetryConfig.newBuilder(retryConfig_).mergeFrom(value).buildPartial();
        } else {
          retryConfig_ = value;
        }
        onChanged();
      } else {
        retryConfigBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public Builder clearRetryConfig() {
      if (retryConfigBuilder_ == null) {
        retryConfig_ = null;
        onChanged();
      } else {
        retryConfig_ = null;
        retryConfigBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder getRetryConfigBuilder() {
      
      onChanged();
      return getRetryConfigFieldBuilder().getBuilder();
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    public org.wso2.apk.enforcer.discovery.api.RetryConfigOrBuilder getRetryConfigOrBuilder() {
      if (retryConfigBuilder_ != null) {
        return retryConfigBuilder_.getMessageOrBuilder();
      } else {
        return retryConfig_ == null ?
            org.wso2.apk.enforcer.discovery.api.RetryConfig.getDefaultInstance() : retryConfig_;
      }
    }
    /**
     * <code>.wso2.discovery.api.RetryConfig retryConfig = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.api.RetryConfig, org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder, org.wso2.apk.enforcer.discovery.api.RetryConfigOrBuilder> 
        getRetryConfigFieldBuilder() {
      if (retryConfigBuilder_ == null) {
        retryConfigBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.wso2.apk.enforcer.discovery.api.RetryConfig, org.wso2.apk.enforcer.discovery.api.RetryConfig.Builder, org.wso2.apk.enforcer.discovery.api.RetryConfigOrBuilder>(
                getRetryConfig(),
                getParentForChildren(),
                isClean());
        retryConfig_ = null;
      }
      return retryConfigBuilder_;
    }

    private org.wso2.apk.enforcer.discovery.api.TimeoutConfig timeoutConfig_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.api.TimeoutConfig, org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder, org.wso2.apk.enforcer.discovery.api.TimeoutConfigOrBuilder> timeoutConfigBuilder_;
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     * @return Whether the timeoutConfig field is set.
     */
    public boolean hasTimeoutConfig() {
      return timeoutConfigBuilder_ != null || timeoutConfig_ != null;
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     * @return The timeoutConfig.
     */
    public org.wso2.apk.enforcer.discovery.api.TimeoutConfig getTimeoutConfig() {
      if (timeoutConfigBuilder_ == null) {
        return timeoutConfig_ == null ? org.wso2.apk.enforcer.discovery.api.TimeoutConfig.getDefaultInstance() : timeoutConfig_;
      } else {
        return timeoutConfigBuilder_.getMessage();
      }
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public Builder setTimeoutConfig(org.wso2.apk.enforcer.discovery.api.TimeoutConfig value) {
      if (timeoutConfigBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        timeoutConfig_ = value;
        onChanged();
      } else {
        timeoutConfigBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public Builder setTimeoutConfig(
        org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder builderForValue) {
      if (timeoutConfigBuilder_ == null) {
        timeoutConfig_ = builderForValue.build();
        onChanged();
      } else {
        timeoutConfigBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public Builder mergeTimeoutConfig(org.wso2.apk.enforcer.discovery.api.TimeoutConfig value) {
      if (timeoutConfigBuilder_ == null) {
        if (timeoutConfig_ != null) {
          timeoutConfig_ =
            org.wso2.apk.enforcer.discovery.api.TimeoutConfig.newBuilder(timeoutConfig_).mergeFrom(value).buildPartial();
        } else {
          timeoutConfig_ = value;
        }
        onChanged();
      } else {
        timeoutConfigBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public Builder clearTimeoutConfig() {
      if (timeoutConfigBuilder_ == null) {
        timeoutConfig_ = null;
        onChanged();
      } else {
        timeoutConfig_ = null;
        timeoutConfigBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder getTimeoutConfigBuilder() {
      
      onChanged();
      return getTimeoutConfigFieldBuilder().getBuilder();
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    public org.wso2.apk.enforcer.discovery.api.TimeoutConfigOrBuilder getTimeoutConfigOrBuilder() {
      if (timeoutConfigBuilder_ != null) {
        return timeoutConfigBuilder_.getMessageOrBuilder();
      } else {
        return timeoutConfig_ == null ?
            org.wso2.apk.enforcer.discovery.api.TimeoutConfig.getDefaultInstance() : timeoutConfig_;
      }
    }
    /**
     * <code>.wso2.discovery.api.TimeoutConfig timeoutConfig = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.api.TimeoutConfig, org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder, org.wso2.apk.enforcer.discovery.api.TimeoutConfigOrBuilder> 
        getTimeoutConfigFieldBuilder() {
      if (timeoutConfigBuilder_ == null) {
        timeoutConfigBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.wso2.apk.enforcer.discovery.api.TimeoutConfig, org.wso2.apk.enforcer.discovery.api.TimeoutConfig.Builder, org.wso2.apk.enforcer.discovery.api.TimeoutConfigOrBuilder>(
                getTimeoutConfig(),
                getParentForChildren(),
                isClean());
        timeoutConfig_ = null;
      }
      return timeoutConfigBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:wso2.discovery.api.EndpointClusterConfig)
  }

  // @@protoc_insertion_point(class_scope:wso2.discovery.api.EndpointClusterConfig)
  private static final org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig();
  }

  public static org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<EndpointClusterConfig>
      PARSER = new com.google.protobuf.AbstractParser<EndpointClusterConfig>() {
    @java.lang.Override
    public EndpointClusterConfig parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new EndpointClusterConfig(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<EndpointClusterConfig> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<EndpointClusterConfig> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.api.EndpointClusterConfig getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

