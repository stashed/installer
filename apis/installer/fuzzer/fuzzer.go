/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzzer

import (
	"stash.appscode.dev/installer/apis/installer/v1alpha1"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/randfill"
)

// Funcs returns the fuzzer functions for this api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []any {
	return []any{
		func(s *v1alpha1.StashCommunity, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.StashEnterprise, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha1.Stash, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
	}
}
