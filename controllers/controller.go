/*
 Copyright 2023 The Nephio Authors.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/nokia/k8s-ipam/controllers/ipamallocation"
	"github.com/nokia/k8s-ipam/controllers/ipamnetworkinstance"
	"github.com/nokia/k8s-ipam/controllers/ipamprefix"
	"github.com/nokia/k8s-ipam/controllers/specializers/config"
	ipamspec "github.com/nokia/k8s-ipam/controllers/specializers/ipam"
	vlanspec "github.com/nokia/k8s-ipam/controllers/specializers/vlan"
	"github.com/nokia/k8s-ipam/controllers/vlanallocation"
	"github.com/nokia/k8s-ipam/controllers/vlandatabase"
	"github.com/nokia/k8s-ipam/controllers/vlanvlan"
	"github.com/nokia/k8s-ipam/internal/shared"
	"github.com/nokia/k8s-ipam/pkg/proxy/clientproxy"
	"github.com/nokia/k8s-ipam/pkg/proxy/clientproxy/ipam"
	"github.com/nokia/k8s-ipam/pkg/proxy/clientproxy/vlan"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

// Setup package controllers.
func Setup(ctx context.Context, mgr ctrl.Manager, opts *shared.Options) error {
	ipamproxyClient := ipam.New(ctx, clientproxy.Config{
		Address: opts.Address,
	})
	opts.IpamClientProxy = ipamproxyClient
	opts.VlanClientProxy = vlan.New(ctx, clientproxy.Config{
		Address: opts.Address,
	})

	eventChs := map[schema.GroupVersionKind]chan event.GenericEvent{}
	for _, setup := range []func(ctrl.Manager, *shared.Options) (schema.GroupVersionKind, chan event.GenericEvent, error){
		ipamnetworkinstance.Setup,
		vlandatabase.Setup,
		ipamprefix.Setup,
		vlanvlan.Setup,
		ipamallocation.Setup,
		vlanallocation.Setup,
	} {
		gvk, geCh, err := setup(mgr, opts)
		if err != nil {
			return err
		}
		eventChs[gvk] = geCh
	}

	/*
		for _, setup := range []func(ctrl.Manager, *shared.Options) error{
			ipamspecializer.Setup,
			vlanspecializer.Setup,
		} {
			if err := setup(mgr, opts); err != nil {
				return err
			}
		}
	*/
	for _, setup := range []func(ctx context.Context, mgr ctrl.Manager, cfg config.SpecializerControllerConfig) error{
		ipamspec.Setup,
		vlanspec.Setup,
	} {
		if err := setup(ctx, mgr, config.SpecializerControllerConfig{
			PorchClient: opts.PorchClient,
		}); err != nil {
			return err
		}
	}

	ipamproxyClient.AddEventChs(eventChs)

	return nil
}
