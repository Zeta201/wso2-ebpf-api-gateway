/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Package model contains the implementation of DTOs to convert OpenAPI/Swagger files
// and create a common model which can represent both types.
package model

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/spec"
	"github.com/google/uuid"
	"github.com/wso2/apk/adapter/config"
	"github.com/wso2/apk/adapter/internal/interceptor"
	"github.com/wso2/apk/adapter/internal/oasparser/constants"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/api"
	dpv1alpha1 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
)

// Operation type object holds data about each http method in the REST API.
type Operation struct {
	iD     string
	method string
	//security map of security scheme names -> list of scopes
	scopes           []string
	auth             *Authentication
	tier             string
	disableSecurity  bool
	vendorExtensions map[string]interface{}
	policies         OperationPolicies
	mockedAPIConfig  *api.MockedApiConfig
	//todo(amali) refactor all vars to private/public vars
	RateLimitPolicy *RateLimitPolicy
}

// Authentication holds authentication related configurations
type Authentication struct {
	Disabled       bool
	JWT            *JWT
	APIKey         []APIKey
	TestConsoleKey *TestConsoleKey
}

// JWT holds JWT related configurations
type JWT struct {
	Header              string
	SendTokenToUpstream bool
}

// TestConsoleKey holds testkey related configurations
type TestConsoleKey struct {
	Header              string
	SendTokenToUpstream bool
}

// APIKey holds API Key related configurations
type APIKey struct {
	In                  string
	Name                string
	SendTokenToUpstream bool
}

// SetAuthentication set authentication configurations
func (operation *Operation) SetAuthentication(authentication *Authentication) {
	operation.auth = authentication
}

// GetAuthentication get authentication configurations
func (operation *Operation) GetAuthentication() *Authentication {
	return operation.auth
}

// SetMockedAPIConfigOAS3 generate mock impl endpoint configurations
func (operation *Operation) SetMockedAPIConfigOAS3(openAPIOperation *openapi3.Operation) {
	if len(openAPIOperation.Responses) > 0 {
		mockedAPIConfig := &api.MockedApiConfig{
			Responses: make([]*api.MockedResponseConfig, 0),
		}
		for responseCode, responseRef := range openAPIOperation.Responses {
			code := strings.ToLower(responseCode)
			if matched, _ := regexp.MatchString("^[0-9x]*", code); (matched && len(code) == 3) || code == "default" {
				mockedResponse := &api.MockedResponseConfig{
					Code:    code,
					Content: make([]*api.MockedContentConfig, 0),
				}
				if responseRef != nil && responseRef.Value != nil {
					for mediaType, content := range responseRef.Value.Content {
						example, err := convertToJSON(content.Example)
						if err == nil {
							mockedResponse.Content = append(mockedResponse.Content, &api.MockedContentConfig{
								ContentType: mediaType,
								Examples:    []*api.MockedContentExample{{Ref: "", Body: example}},
							})
						} else if len(content.Examples) > 0 {
							mockedContentExamples := make([]*api.MockedContentExample, 0)
							for ref, exampleVal := range content.Examples {
								if exampleVal != nil && exampleVal.Value != nil {
									example, err = convertToJSON(exampleVal.Value.Value)
									if err == nil {
										mockedContentExamples = append(mockedContentExamples, &api.MockedContentExample{
											Ref:  ref,
											Body: example,
										})
									}
								}

							}
							mockedResponse.Content = append(mockedResponse.Content,
								&api.MockedContentConfig{
									ContentType: mediaType,
									Examples:    mockedContentExamples,
								})
						}
					}
					for headerName, headerValues := range responseRef.Value.Headers {
						example, err := convertToJSON(headerValues.Value.Example)
						if err == nil {
							mockedResponse.Headers = append(mockedResponse.Headers, &api.MockedHeaderConfig{
								Name:  headerName,
								Value: example,
							})
						}
					}
				}
				if len(mockedResponse.Content) > 0 {
					mockedAPIConfig.Responses = append(mockedAPIConfig.Responses, mockedResponse)
				}
			}
		}
		if len(mockedAPIConfig.Responses) > 0 {
			operation.mockedAPIConfig = mockedAPIConfig
		}
	}
}

// SetMockedAPIConfigOAS2 generate mock impl endpoint configurations
func (operation *Operation) SetMockedAPIConfigOAS2(openAPIOperation *spec.Operation) {
	if openAPIOperation.Responses != nil && len(openAPIOperation.Responses.StatusCodeResponses) > 0 {
		mockedAPIConfig := &api.MockedApiConfig{
			Responses: make([]*api.MockedResponseConfig, 0),
		}
		// get response codes
		for responseCode, responseRef := range openAPIOperation.Responses.StatusCodeResponses {
			mockedResponse := &api.MockedResponseConfig{
				Code:    strconv.Itoa(responseCode),
				Content: make([]*api.MockedContentConfig, 0),
			}
			for mediaType, content := range responseRef.ResponseProps.Examples {
				//todo(amali) xml payload gen
				example, err := convertToJSON(content)
				if err == nil {
					mockedResponse.Content = append(mockedResponse.Content, &api.MockedContentConfig{
						ContentType: mediaType,
						Examples:    []*api.MockedContentExample{{Ref: "", Body: example}},
					})
				}
			}
			// swagger does not support header example/examples
			if len(mockedResponse.Content) > 0 {
				mockedAPIConfig.Responses = append(mockedAPIConfig.Responses, mockedResponse)
			}
		}
		// get default response examples
		if openAPIOperation.Responses.Default != nil && len(openAPIOperation.Responses.Default.Examples) > 0 {
			mockedResponse := &api.MockedResponseConfig{
				Code:    "default",
				Content: make([]*api.MockedContentConfig, 0),
			}
			for mediaType, content := range openAPIOperation.Responses.Default.Examples {
				example, err := convertToJSON(content)
				if err == nil {
					mockedResponse.Content = append(mockedResponse.Content, &api.MockedContentConfig{
						ContentType: mediaType,
						Examples:    []*api.MockedContentExample{{Ref: "", Body: example}},
					})
				}
			}
			// swagger does not support header example/examples
			if len(mockedResponse.Content) > 0 {
				mockedAPIConfig.Responses = append(mockedAPIConfig.Responses, mockedResponse)
			}
		}
		if len(mockedAPIConfig.Responses) > 0 {
			operation.mockedAPIConfig = mockedAPIConfig
		}
	}
}

// convertToJSON parse interface to JSON string. returns error if a null value has passed
func convertToJSON(data interface{}) (string, error) {
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", errors.New("null object passed")
}

// GetMethod returns the http method name of the give API operation
func (operation *Operation) GetMethod() string {
	return operation.method
}

// GetPolicies returns if the resouce is secured.
func (operation *Operation) GetPolicies() *OperationPolicies {
	return &operation.policies
}

// GetScopes returns the security schemas defined for the http opeartion
func (operation *Operation) GetScopes() []string {
	return operation.scopes
}

// SetScopes sets the security schemas for the http opeartion
func (operation *Operation) SetScopes(scopes []string) {
	operation.scopes = scopes
}

// GetTier returns the operation level throttling tier
func (operation *Operation) GetTier() string {
	return operation.tier
}

// GetMockedAPIConfig returns the operation level mocked API implementation configs
func (operation *Operation) GetMockedAPIConfig() *api.MockedApiConfig {
	return operation.mockedAPIConfig
}

// GetVendorExtensions returns vendor extensions which are explicitly defined under
// a given resource.
func (operation *Operation) GetVendorExtensions() map[string]interface{} {
	return operation.vendorExtensions
}

// GetID returns the id of a given resource.
// This is a randomly generated UUID
func (operation *Operation) GetID() string {
	return operation.iD
}

// GetCallInterceptorService returns the interceptor configs for a given operation.
func (operation *Operation) GetCallInterceptorService(isIn bool) InterceptEndpoint {
	var policies []Policy
	if isIn {
		policies = operation.policies.Request
	} else {
		policies = operation.policies.Response
	}
	if len(policies) > 0 {
		for _, policy := range policies {
			if strings.EqualFold(constants.ActionInterceptorService, policy.Action) {
				if paramMap, isMap := policy.Parameters.(map[string]interface{}); isMap {
					endpoints, endpointsFound := paramMap[constants.InterceptorEndpoints]
					includesValue, includesFound := paramMap[constants.InterceptorServiceIncludes]
					if endpointsFound {
						endpoints, isEndpoints := endpoints.([]Endpoint)
						if isEndpoints {
							conf := config.ReadConfigs()
							clusterTimeoutV := conf.Envoy.ClusterTimeoutInSeconds
							requestTimeoutV := conf.Envoy.ClusterTimeoutInSeconds
							includesV := &interceptor.RequestInclusions{}
							if includesFound {
								includes, ok := includesValue.([]dpv1alpha1.InterceptorInclusion)
								if ok {
									includesV = GenerateInterceptorIncludes(includes)
								}
							}
							return InterceptEndpoint{
								Enable:          true,
								EndpointCluster: EndpointCluster{Endpoints: endpoints},
								ClusterTimeout:  clusterTimeoutV,
								RequestTimeout:  requestTimeoutV,
								Includes:        includesV,
								Level:           constants.OperationLevelInterceptor,
							}
						}
					}
				}
			}
		}
	}
	return InterceptEndpoint{}
}

// NewOperation Creates and returns operation type object
func NewOperation(method string, security []string, extensions map[string]interface{}) *Operation {
	tier := ResolveThrottlingTier(extensions)
	disableSecurity := ResolveDisableSecurity(extensions)
	id := uuid.New().String()
	return &Operation{id, method, security, nil, tier, disableSecurity, extensions, OperationPolicies{}, &api.MockedApiConfig{}, nil}
}

// NewOperationWithPolicies Creates and returns operation with given method and policies
func NewOperationWithPolicies(method string, policies OperationPolicies) *Operation {
	return &Operation{iD: uuid.New().String(), method: method, policies: policies}
}
