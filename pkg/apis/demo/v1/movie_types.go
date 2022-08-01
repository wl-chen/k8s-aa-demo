
/*
Copyright 2022.

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

package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Movie
// +k8s:openapi-gen=true
type Movie struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MovieSpec   `json:"spec,omitempty"`
	Status MovieStatus `json:"status,omitempty"`
}

// MovieList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MovieList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Movie `json:"items"`
}

// MovieSpec defines the desired state of Movie
type MovieSpec struct {
}

var _ resource.Object = &Movie{}
var _ resourcestrategy.Validater = &Movie{}

func (in *Movie) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Movie) NamespaceScoped() bool {
	return false
}

func (in *Movie) New() runtime.Object {
	return &Movie{}
}

func (in *Movie) NewList() runtime.Object {
	return &MovieList{}
}

func (in *Movie) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "demo.demo",
		Version:  "v1",
		Resource: "movies",
	}
}

func (in *Movie) IsStorageVersion() bool {
	return true
}

func (in *Movie) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &MovieList{}

func (in *MovieList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}
// MovieStatus defines the observed state of Movie
type MovieStatus struct {
}

func (in MovieStatus) SubResourceName() string {
	return "status"
}

// Movie implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &Movie{}

func (in *Movie) GetStatus() resource.StatusSubResource {
	return in.Status
}

// MovieStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &MovieStatus{}

func (in MovieStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*Movie).Status = in
}
