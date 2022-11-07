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

package shared

import (
	"time"

	"github.com/henderiw-nephio/ipam/internal/allocator"
	"github.com/henderiw-nephio/ipam/internal/injectors"
	"github.com/henderiw-nephio/ipam/internal/ipam2"
	"github.com/henderiw-nephio/ipam/pkg/alloc/allocpb"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

type Options struct {
	PorchClient client.Client
	AllocClient allocpb.AllocationClient
	Allocator   allocator.Allocator
	Poll        time.Duration
	Copts       controller.Options
	Ipam        ipam2.Ipam
	Injectors   injectors.Injectors
}
