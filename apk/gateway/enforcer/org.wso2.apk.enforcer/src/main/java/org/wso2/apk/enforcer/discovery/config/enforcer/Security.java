// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/security.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

/**
 * <pre>
 * Enforcer config model
 * </pre>
 *
 * Protobuf type {@code wso2.discovery.config.enforcer.Security}
 */
public final class Security extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:wso2.discovery.config.enforcer.Security)
    SecurityOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Security.newBuilder() to construct.
  private Security(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Security() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Security();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Security(
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
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder subBuilder = null;
            if (apiKey_ != null) {
              subBuilder = apiKey_.toBuilder();
            }
            apiKey_ = input.readMessage(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(apiKey_);
              apiKey_ = subBuilder.buildPartial();
            }

            break;
          }
          case 18: {
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder subBuilder = null;
            if (runtimeToken_ != null) {
              subBuilder = runtimeToken_.toBuilder();
            }
            runtimeToken_ = input.readMessage(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(runtimeToken_);
              runtimeToken_ = subBuilder.buildPartial();
            }

            break;
          }
          case 34: {
            org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder subBuilder = null;
            if (mutualSSL_ != null) {
              subBuilder = mutualSSL_.toBuilder();
            }
            mutualSSL_ = input.readMessage(org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(mutualSSL_);
              mutualSSL_ = subBuilder.buildPartial();
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
    return org.wso2.apk.enforcer.discovery.config.enforcer.SecurityProto.internal_static_wso2_discovery_config_enforcer_Security_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.wso2.apk.enforcer.discovery.config.enforcer.SecurityProto.internal_static_wso2_discovery_config_enforcer_Security_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.wso2.apk.enforcer.discovery.config.enforcer.Security.class, org.wso2.apk.enforcer.discovery.config.enforcer.Security.Builder.class);
  }

  public static final int APIKEY_FIELD_NUMBER = 1;
  private org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer apiKey_;
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
   * @return Whether the apiKey field is set.
   */
  @java.lang.Override
  public boolean hasApiKey() {
    return apiKey_ != null;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
   * @return The apiKey.
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer getApiKey() {
    return apiKey_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : apiKey_;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder getApiKeyOrBuilder() {
    return getApiKey();
  }

  public static final int RUNTIMETOKEN_FIELD_NUMBER = 2;
  private org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer runtimeToken_;
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
   * @return Whether the runtimeToken field is set.
   */
  @java.lang.Override
  public boolean hasRuntimeToken() {
    return runtimeToken_ != null;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
   * @return The runtimeToken.
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer getRuntimeToken() {
    return runtimeToken_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : runtimeToken_;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder getRuntimeTokenOrBuilder() {
    return getRuntimeToken();
  }

  public static final int MUTUALSSL_FIELD_NUMBER = 4;
  private org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL mutualSSL_;
  /**
   * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
   * @return Whether the mutualSSL field is set.
   */
  @java.lang.Override
  public boolean hasMutualSSL() {
    return mutualSSL_ != null;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
   * @return The mutualSSL.
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL getMutualSSL() {
    return mutualSSL_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.getDefaultInstance() : mutualSSL_;
  }
  /**
   * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
   */
  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSLOrBuilder getMutualSSLOrBuilder() {
    return getMutualSSL();
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
    if (apiKey_ != null) {
      output.writeMessage(1, getApiKey());
    }
    if (runtimeToken_ != null) {
      output.writeMessage(2, getRuntimeToken());
    }
    if (mutualSSL_ != null) {
      output.writeMessage(4, getMutualSSL());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (apiKey_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getApiKey());
    }
    if (runtimeToken_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getRuntimeToken());
    }
    if (mutualSSL_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(4, getMutualSSL());
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
    if (!(obj instanceof org.wso2.apk.enforcer.discovery.config.enforcer.Security)) {
      return super.equals(obj);
    }
    org.wso2.apk.enforcer.discovery.config.enforcer.Security other = (org.wso2.apk.enforcer.discovery.config.enforcer.Security) obj;

    if (hasApiKey() != other.hasApiKey()) return false;
    if (hasApiKey()) {
      if (!getApiKey()
          .equals(other.getApiKey())) return false;
    }
    if (hasRuntimeToken() != other.hasRuntimeToken()) return false;
    if (hasRuntimeToken()) {
      if (!getRuntimeToken()
          .equals(other.getRuntimeToken())) return false;
    }
    if (hasMutualSSL() != other.hasMutualSSL()) return false;
    if (hasMutualSSL()) {
      if (!getMutualSSL()
          .equals(other.getMutualSSL())) return false;
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
    if (hasApiKey()) {
      hash = (37 * hash) + APIKEY_FIELD_NUMBER;
      hash = (53 * hash) + getApiKey().hashCode();
    }
    if (hasRuntimeToken()) {
      hash = (37 * hash) + RUNTIMETOKEN_FIELD_NUMBER;
      hash = (53 * hash) + getRuntimeToken().hashCode();
    }
    if (hasMutualSSL()) {
      hash = (37 * hash) + MUTUALSSL_FIELD_NUMBER;
      hash = (53 * hash) + getMutualSSL().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security parseFrom(
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
  public static Builder newBuilder(org.wso2.apk.enforcer.discovery.config.enforcer.Security prototype) {
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
   * <pre>
   * Enforcer config model
   * </pre>
   *
   * Protobuf type {@code wso2.discovery.config.enforcer.Security}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:wso2.discovery.config.enforcer.Security)
      org.wso2.apk.enforcer.discovery.config.enforcer.SecurityOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.SecurityProto.internal_static_wso2_discovery_config_enforcer_Security_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.SecurityProto.internal_static_wso2_discovery_config_enforcer_Security_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.wso2.apk.enforcer.discovery.config.enforcer.Security.class, org.wso2.apk.enforcer.discovery.config.enforcer.Security.Builder.class);
    }

    // Construct using org.wso2.apk.enforcer.discovery.config.enforcer.Security.newBuilder()
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
      if (apiKeyBuilder_ == null) {
        apiKey_ = null;
      } else {
        apiKey_ = null;
        apiKeyBuilder_ = null;
      }
      if (runtimeTokenBuilder_ == null) {
        runtimeToken_ = null;
      } else {
        runtimeToken_ = null;
        runtimeTokenBuilder_ = null;
      }
      if (mutualSSLBuilder_ == null) {
        mutualSSL_ = null;
      } else {
        mutualSSL_ = null;
        mutualSSLBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.SecurityProto.internal_static_wso2_discovery_config_enforcer_Security_descriptor;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Security getDefaultInstanceForType() {
      return org.wso2.apk.enforcer.discovery.config.enforcer.Security.getDefaultInstance();
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Security build() {
      org.wso2.apk.enforcer.discovery.config.enforcer.Security result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.wso2.apk.enforcer.discovery.config.enforcer.Security buildPartial() {
      org.wso2.apk.enforcer.discovery.config.enforcer.Security result = new org.wso2.apk.enforcer.discovery.config.enforcer.Security(this);
      if (apiKeyBuilder_ == null) {
        result.apiKey_ = apiKey_;
      } else {
        result.apiKey_ = apiKeyBuilder_.build();
      }
      if (runtimeTokenBuilder_ == null) {
        result.runtimeToken_ = runtimeToken_;
      } else {
        result.runtimeToken_ = runtimeTokenBuilder_.build();
      }
      if (mutualSSLBuilder_ == null) {
        result.mutualSSL_ = mutualSSL_;
      } else {
        result.mutualSSL_ = mutualSSLBuilder_.build();
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
      if (other instanceof org.wso2.apk.enforcer.discovery.config.enforcer.Security) {
        return mergeFrom((org.wso2.apk.enforcer.discovery.config.enforcer.Security)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.wso2.apk.enforcer.discovery.config.enforcer.Security other) {
      if (other == org.wso2.apk.enforcer.discovery.config.enforcer.Security.getDefaultInstance()) return this;
      if (other.hasApiKey()) {
        mergeApiKey(other.getApiKey());
      }
      if (other.hasRuntimeToken()) {
        mergeRuntimeToken(other.getRuntimeToken());
      }
      if (other.hasMutualSSL()) {
        mergeMutualSSL(other.getMutualSSL());
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
      org.wso2.apk.enforcer.discovery.config.enforcer.Security parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.wso2.apk.enforcer.discovery.config.enforcer.Security) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer apiKey_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder> apiKeyBuilder_;
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     * @return Whether the apiKey field is set.
     */
    public boolean hasApiKey() {
      return apiKeyBuilder_ != null || apiKey_ != null;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     * @return The apiKey.
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer getApiKey() {
      if (apiKeyBuilder_ == null) {
        return apiKey_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : apiKey_;
      } else {
        return apiKeyBuilder_.getMessage();
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public Builder setApiKey(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer value) {
      if (apiKeyBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        apiKey_ = value;
        onChanged();
      } else {
        apiKeyBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public Builder setApiKey(
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder builderForValue) {
      if (apiKeyBuilder_ == null) {
        apiKey_ = builderForValue.build();
        onChanged();
      } else {
        apiKeyBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public Builder mergeApiKey(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer value) {
      if (apiKeyBuilder_ == null) {
        if (apiKey_ != null) {
          apiKey_ =
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.newBuilder(apiKey_).mergeFrom(value).buildPartial();
        } else {
          apiKey_ = value;
        }
        onChanged();
      } else {
        apiKeyBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public Builder clearApiKey() {
      if (apiKeyBuilder_ == null) {
        apiKey_ = null;
        onChanged();
      } else {
        apiKey_ = null;
        apiKeyBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder getApiKeyBuilder() {
      
      onChanged();
      return getApiKeyFieldBuilder().getBuilder();
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder getApiKeyOrBuilder() {
      if (apiKeyBuilder_ != null) {
        return apiKeyBuilder_.getMessageOrBuilder();
      } else {
        return apiKey_ == null ?
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : apiKey_;
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer apiKey = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder> 
        getApiKeyFieldBuilder() {
      if (apiKeyBuilder_ == null) {
        apiKeyBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder>(
                getApiKey(),
                getParentForChildren(),
                isClean());
        apiKey_ = null;
      }
      return apiKeyBuilder_;
    }

    private org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer runtimeToken_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder> runtimeTokenBuilder_;
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     * @return Whether the runtimeToken field is set.
     */
    public boolean hasRuntimeToken() {
      return runtimeTokenBuilder_ != null || runtimeToken_ != null;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     * @return The runtimeToken.
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer getRuntimeToken() {
      if (runtimeTokenBuilder_ == null) {
        return runtimeToken_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : runtimeToken_;
      } else {
        return runtimeTokenBuilder_.getMessage();
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public Builder setRuntimeToken(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer value) {
      if (runtimeTokenBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        runtimeToken_ = value;
        onChanged();
      } else {
        runtimeTokenBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public Builder setRuntimeToken(
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder builderForValue) {
      if (runtimeTokenBuilder_ == null) {
        runtimeToken_ = builderForValue.build();
        onChanged();
      } else {
        runtimeTokenBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public Builder mergeRuntimeToken(org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer value) {
      if (runtimeTokenBuilder_ == null) {
        if (runtimeToken_ != null) {
          runtimeToken_ =
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.newBuilder(runtimeToken_).mergeFrom(value).buildPartial();
        } else {
          runtimeToken_ = value;
        }
        onChanged();
      } else {
        runtimeTokenBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public Builder clearRuntimeToken() {
      if (runtimeTokenBuilder_ == null) {
        runtimeToken_ = null;
        onChanged();
      } else {
        runtimeToken_ = null;
        runtimeTokenBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder getRuntimeTokenBuilder() {
      
      onChanged();
      return getRuntimeTokenFieldBuilder().getBuilder();
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder getRuntimeTokenOrBuilder() {
      if (runtimeTokenBuilder_ != null) {
        return runtimeTokenBuilder_.getMessageOrBuilder();
      } else {
        return runtimeToken_ == null ?
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.getDefaultInstance() : runtimeToken_;
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.APIKeyEnforcer runtimeToken = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder> 
        getRuntimeTokenFieldBuilder() {
      if (runtimeTokenBuilder_ == null) {
        runtimeTokenBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcer.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.APIKeyEnforcerOrBuilder>(
                getRuntimeToken(),
                getParentForChildren(),
                isClean());
        runtimeToken_ = null;
      }
      return runtimeTokenBuilder_;
    }

    private org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL mutualSSL_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSLOrBuilder> mutualSSLBuilder_;
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     * @return Whether the mutualSSL field is set.
     */
    public boolean hasMutualSSL() {
      return mutualSSLBuilder_ != null || mutualSSL_ != null;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     * @return The mutualSSL.
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL getMutualSSL() {
      if (mutualSSLBuilder_ == null) {
        return mutualSSL_ == null ? org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.getDefaultInstance() : mutualSSL_;
      } else {
        return mutualSSLBuilder_.getMessage();
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public Builder setMutualSSL(org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL value) {
      if (mutualSSLBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        mutualSSL_ = value;
        onChanged();
      } else {
        mutualSSLBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public Builder setMutualSSL(
        org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder builderForValue) {
      if (mutualSSLBuilder_ == null) {
        mutualSSL_ = builderForValue.build();
        onChanged();
      } else {
        mutualSSLBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public Builder mergeMutualSSL(org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL value) {
      if (mutualSSLBuilder_ == null) {
        if (mutualSSL_ != null) {
          mutualSSL_ =
            org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.newBuilder(mutualSSL_).mergeFrom(value).buildPartial();
        } else {
          mutualSSL_ = value;
        }
        onChanged();
      } else {
        mutualSSLBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public Builder clearMutualSSL() {
      if (mutualSSLBuilder_ == null) {
        mutualSSL_ = null;
        onChanged();
      } else {
        mutualSSL_ = null;
        mutualSSLBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder getMutualSSLBuilder() {
      
      onChanged();
      return getMutualSSLFieldBuilder().getBuilder();
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    public org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSLOrBuilder getMutualSSLOrBuilder() {
      if (mutualSSLBuilder_ != null) {
        return mutualSSLBuilder_.getMessageOrBuilder();
      } else {
        return mutualSSL_ == null ?
            org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.getDefaultInstance() : mutualSSL_;
      }
    }
    /**
     * <code>.wso2.discovery.config.enforcer.MutualSSL mutualSSL = 4;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSLOrBuilder> 
        getMutualSSLFieldBuilder() {
      if (mutualSSLBuilder_ == null) {
        mutualSSLBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSL.Builder, org.wso2.apk.enforcer.discovery.config.enforcer.MutualSSLOrBuilder>(
                getMutualSSL(),
                getParentForChildren(),
                isClean());
        mutualSSL_ = null;
      }
      return mutualSSLBuilder_;
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


    // @@protoc_insertion_point(builder_scope:wso2.discovery.config.enforcer.Security)
  }

  // @@protoc_insertion_point(class_scope:wso2.discovery.config.enforcer.Security)
  private static final org.wso2.apk.enforcer.discovery.config.enforcer.Security DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.wso2.apk.enforcer.discovery.config.enforcer.Security();
  }

  public static org.wso2.apk.enforcer.discovery.config.enforcer.Security getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Security>
      PARSER = new com.google.protobuf.AbstractParser<Security>() {
    @java.lang.Override
    public Security parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Security(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Security> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Security> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.wso2.apk.enforcer.discovery.config.enforcer.Security getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
