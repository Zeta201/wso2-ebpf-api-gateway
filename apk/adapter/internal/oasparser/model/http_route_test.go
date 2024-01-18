/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	dpv1alpha1 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	dpv1alpha2 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"
)

func TestConcatRateLimitPolicies(t *testing.T) {
	type testItem struct {
		schemeUpSpec   dpv1alpha1.RateLimitPolicySpec
		schemeDownSpec dpv1alpha1.RateLimitPolicySpec
		result         dpv1alpha1.RateLimitPolicySpec
		message        string
	}

	schemeUp := &dpv1alpha1.RateLimitPolicy{}
	schemeDown := &dpv1alpha1.RateLimitPolicy{}
	resultScheme := &dpv1alpha1.RateLimitPolicy{}

	dataItems := []testItem{
		{
			schemeUpSpec: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			schemeDownSpec: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			result: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			message: "When API level override and Resource level override policies both provided",
		},
		{
			schemeUpSpec: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			schemeDownSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			result: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			message: "When API level override and Resource level default policies both provided",
		},
		{
			schemeUpSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			schemeDownSpec: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			result: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			message: "When API level default and Resource level override policies both provided",
		},
		{
			schemeUpSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
			},
			schemeDownSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			result: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Day",
					},
				},
			},
			message: "When API level default and Resource level default policies both provided",
		},
		{
			schemeUpSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 10,
						Unit:            "Minute",
					},
				},
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Second",
					},
				},
			},
			schemeDownSpec: dpv1alpha1.RateLimitPolicySpec{
				Default: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 30,
						Unit:            "Day",
					},
				},
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 40,
						Unit:            "Hour",
					},
				},
			},
			result: dpv1alpha1.RateLimitPolicySpec{
				Override: &dpv1alpha1.RateLimitAPIPolicy{
					API: &dpv1alpha1.APIRateLimitPolicy{
						RequestsPerUnit: 20,
						Unit:            "Second",
					},
				},
			},
			message: "When both API level and Resource level both override and default policies provided",
		},
	}

	for _, item := range dataItems {
		schemeUp.Spec = item.schemeUpSpec
		schemeDown.Spec = item.schemeDownSpec
		resultScheme.Spec = item.result
		actualResult := concatRateLimitPolicies(schemeUp, schemeDown)
		assert.Equal(t, resultScheme, actualResult, item.message)
	}
}

func TestConcatAPIPolicies(t *testing.T) {

	type testItem struct {
		schemeUpSpec   dpv1alpha2.APIPolicySpec
		schemeDownSpec dpv1alpha2.APIPolicySpec
		result         dpv1alpha2.APIPolicySpec
		message        string
	}

	schemeUp := &dpv1alpha2.APIPolicy{}
	schemeDown := &dpv1alpha2.APIPolicy{}
	resultScheme := &dpv1alpha2.APIPolicy{}

	dataItems := []testItem{
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			schemeDownSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i2"},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			message: "only schemeUp override policies should be provided",
		},
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Default: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			schemeDownSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i2"},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i2"},
					},
				},
			},
			message: "only schemeDown override policies should be provided",
		},
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Default: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			schemeDownSpec: dpv1alpha2.APIPolicySpec{
				Default: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i2"},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i2"},
					},
				},
			},
			message: "only schemeDown default policies should be provided",
		},
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			message: "only schemeUp override policies is provided",
		},
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Default: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{Name: "i1"},
					},
				},
			},
			message: "only schemeUp default policies is provided",
		},
		{
			schemeUpSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{
							Name: "up-request-interceptor",
						},
					},
				},
			},
			schemeDownSpec: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{
							Name: "down-request-interceptor",
						},
					},
				},
				Default: &dpv1alpha2.PolicySpec{
					ResponseInterceptors: []dpv1alpha2.InterceptorReference{
						{
							Name: "down-response-interceptor",
						},
					},
				},
			},
			result: dpv1alpha2.APIPolicySpec{
				Override: &dpv1alpha2.PolicySpec{
					RequestInterceptors: []dpv1alpha2.InterceptorReference{
						{
							Name: "up-request-interceptor",
						},
					},
					ResponseInterceptors: []dpv1alpha2.InterceptorReference{
						{
							Name: "down-response-interceptor",
						},
					},
				},
			},
			message: `not expected result for request interceptor or response interceptor`,
		},
	}
	for _, item := range dataItems {
		schemeUp.Spec = item.schemeUpSpec
		schemeDown.Spec = item.schemeDownSpec
		resultScheme.Spec = item.result
		actualResult := concatAPIPolicies(schemeUp, schemeDown)
		assert.Equal(t, resultScheme, actualResult, item.message)
	}
}
