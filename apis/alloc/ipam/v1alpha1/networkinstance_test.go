/*
Copyright 2022 Nokia.

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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNiGetNameFromNetworkInstancePrefix(t *testing.T) {
	tests := map[string]struct {
		input       *NetworkInstance
		inputPrefix string
		want        string
	}{
		"OK": {
			input: &NetworkInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "a",
					Namespace: "b",
				},
			},
			inputPrefix: "10.0.0.1/24",
			want:        "a-aggregate-10.0.0.1-24",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetNameFromNetworkInstancePrefix(tc.inputPrefix)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}

		})
	}
}

func TestNiGetNetworkInstanceNameSpace(t *testing.T) {
	tests := map[string]struct {
		input *NetworkInstance
		want  string
	}{
		"NamespaceNAme": {
			input: &NetworkInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "a",
					Namespace: "b",
				},
			},
			want: "b-a",
		},
		"NoNamespace": {
			input: &NetworkInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "a",
				},
			},
			want: "a",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetGenericNamespacedName()

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestNiGetNetworkInstanceNameSpace: -want, +got:\n%s", diff)
			}

		})
	}
}