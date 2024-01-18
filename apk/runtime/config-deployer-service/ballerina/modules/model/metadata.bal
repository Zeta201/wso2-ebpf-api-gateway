//
// Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
//
// WSO2 LLC. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
public type Metadata record {
    string name;
    string namespace?;
    string uid?;
    string creationTimestamp?;
    string selfLink?;
    string resourceVersion?;
    OwnerReference[] ownerReferences?;
    ManagedFieldsEntry[] managedFields?;
    int generation?;
    string generateName?;
    map<string> labels?;
    map<string> annotations?;
};

public type OwnerReference record {
    string apiVersion;
    boolean blockOwnerDeletion?;
    boolean controller?;
    string kind;
    string name;
    string uid;
};

public type ManagedFieldsEntry record {
};
