/*
 *  Copyright (c) 2026 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
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
 */

package slicegateway

import (
	"context"
	"testing"

	gwsidecarpb "github.com/kubeslice/gateway-sidecar/pkg/sidecar/sidecarpb"
	kubeslicev1beta1 "github.com/kubeslice/worker-operator/api/v1beta1"
	"github.com/kubeslice/worker-operator/controllers"
	"github.com/kubeslice/worker-operator/pkg/router"
	webhook "github.com/kubeslice/worker-operator/pkg/webhook/pod"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// A restarted NSM connection broker is what makes every gateway on a node change address at
// once, so it has to be recognised as NSM infrastructure -- the reconciler is only woken for pod
// types it knows, and an unrecognised pod produces no reconcile at all.
func TestGetPodTypeRecognizesTheNsmConnectionBroker(t *testing.T) {
	for name, tc := range map[string]struct {
		labels map[string]string
		want   string
	}{
		"the nsm connection broker": {map[string]string{"app": "nsc-grpc-server"}, "nsm"},
		"nsmgr":                     {map[string]string{"app": "nsmgr-daemonset"}, "nsm"},
		"the nsm forwarder":         {map[string]string{"app": "nsm-kernel-plane"}, "nsm"},
		"a slice gateway":           {map[string]string{webhook.PodInjectLabelKey: "slicegateway"}, "slicegateway"},
		"a slice router":            {map[string]string{webhook.PodInjectLabelKey: "router"}, "router"},
		"an unrelated pod":          {map[string]string{"app": "postgres"}, ""},
		"no labels":                 {map[string]string{}, ""},
	} {
		if got := getPodType(tc.labels); got != tc.want {
			t.Errorf("%s: getPodType(%v) = %q, want %q", name, tc.labels, got, tc.want)
		}
	}
}

type recordingRouterClient struct {
	sent []*router.SliceRouterConnCtx
}

func (c *recordingRouterClient) GetClientConnectionInfo(ctx context.Context, addr string) ([]kubeslicev1beta1.AppPod, error) {
	return nil, nil
}

func (c *recordingRouterClient) SendConnectionContext(ctx context.Context, serverAddr string, connCtx *router.SliceRouterConnCtx) error {
	c.sent = append(c.sent, connCtx)
	return nil
}

func (c *recordingRouterClient) UpdateEcmpRoutes(ctx context.Context, serverAddr string, info *router.UpdateEcmpInfo) error {
	return nil
}

func readyGw(name, nsmIP string) *kubeslicev1beta1.GwPodInfo {
	return &kubeslicev1beta1.GwPodInfo{
		PodName:     name,
		PeerPodName: "peer-of-" + name,
		LocalNsmIP:  nsmIP,
		TunnelStatus: kubeslicev1beta1.TunnelStatus{
			Status: int32(gwsidecarpb.TunnelStatusType_GW_TUNNEL_STATE_UP),
		},
	}
}

// reattachingGw is a gateway between NSM addresses: its interface is gone and the new one has
// not arrived, which is the state every gateway on a node is in just after the connection broker
// restarts.
func reattachingGw(name string) *kubeslicev1beta1.GwPodInfo {
	return &kubeslicev1beta1.GwPodInfo{PodName: name, PeerPodName: "peer-of-" + name}
}

func TestSendConnectionContextToSliceRouter(t *testing.T) {
	for name, tc := range map[string]struct {
		gwPods       []*kubeslicev1beta1.GwPodInfo
		wantSent     [][]string
		wantRequeue  bool
		wantAfterSec float64
	}{
		"every gateway ready: send both, and do not come back early": {
			gwPods:      []*kubeslicev1beta1.GwPodInfo{readyGw("gw-0", "10.1.32.7"), readyGw("gw-1", "10.1.32.21")},
			wantSent:    [][]string{{"10.1.32.7", "10.1.32.21"}},
			wantRequeue: false,
		},
		"one gateway reattaching: route over the other, and come back for it": {
			gwPods:       []*kubeslicev1beta1.GwPodInfo{readyGw("gw-0", "10.1.32.7"), reattachingGw("gw-1")},
			wantSent:     [][]string{{"10.1.32.7"}},
			wantRequeue:  true,
			wantAfterSec: controllers.GatewaySettlingRequeueInterval.Seconds(),
		},
		// Sending an empty list is how the router is told to drop the subnet. Doing that here
		// would tear down a route that is seconds from being repairable, and discard the
		// nexthops the router caches to repair itself with.
		"every gateway reattaching: send nothing at all": {
			gwPods:       []*kubeslicev1beta1.GwPodInfo{reattachingGw("gw-0"), reattachingGw("gw-1")},
			wantSent:     nil,
			wantRequeue:  true,
			wantAfterSec: controllers.GatewaySettlingRequeueInterval.Seconds(),
		},
		// Note GW_TUNNEL_STATE_UP is the zero value of the enum, so a gateway that has never
		// reported a tunnel status counts as up. That is long standing behaviour and not
		// changed here; a gateway is counted as settling under exactly the conditions that
		// exclude it from the nexthop list, so the two cannot disagree.
		"a gateway whose tunnel is down does not carry traffic": {
			gwPods: []*kubeslicev1beta1.GwPodInfo{
				readyGw("gw-0", "10.1.32.7"),
				{
					PodName: "gw-1", PeerPodName: "peer-of-gw-1", LocalNsmIP: "10.1.32.21",
					TunnelStatus: kubeslicev1beta1.TunnelStatus{
						Status: int32(gwsidecarpb.TunnelStatusType_GW_TUNNEL_STATE_DOWN),
					},
				},
			},
			wantSent:     [][]string{{"10.1.32.7"}},
			wantRequeue:  true,
			wantAfterSec: controllers.GatewaySettlingRequeueInterval.Seconds(),
		},
	} {
		t.Run(name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			if err := clientgoscheme.AddToScheme(scheme); err != nil {
				t.Fatal(err)
			}
			if err := kubeslicev1beta1.AddToScheme(scheme); err != nil {
				t.Fatal(err)
			}

			sliceRouterPod := &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vl3-slice-router-red-0",
					Namespace: controllers.ControlPlaneNamespace,
					Labels:    map[string]string{"networkservicemesh.io/impl": "vl3-service-red"},
				},
				Status: corev1.PodStatus{Phase: corev1.PodRunning, PodIP: "192.168.0.10"},
			}
			sliceGw := &kubeslicev1beta1.SliceGateway{
				ObjectMeta: metav1.ObjectMeta{Name: "red-worker-1-worker-2", Namespace: controllers.ControlPlaneNamespace},
				Spec:       kubeslicev1beta1.SliceGatewaySpec{SliceName: "red"},
				Status: kubeslicev1beta1.SliceGatewayStatus{
					Config:           kubeslicev1beta1.SliceGatewayConfig{SliceGatewayRemoteSubnet: "10.1.16.0/20"},
					GatewayPodStatus: tc.gwPods,
				},
			}

			routerClient := &recordingRouterClient{}
			r := &SliceGwReconciler{
				Client:             fake.NewClientBuilder().WithScheme(scheme).WithObjects(sliceRouterPod, sliceGw).Build(),
				Log:                zap.New(zap.UseDevMode(true)),
				WorkerRouterClient: routerClient,
			}

			res, err, requeue := r.SendConnectionContextToSliceRouter(context.Background(), sliceGw)
			if err != nil {
				t.Fatalf("SendConnectionContextToSliceRouter: %v", err)
			}

			if len(routerClient.sent) != len(tc.wantSent) {
				t.Fatalf("sent %d connection contexts, want %d: %v", len(routerClient.sent), len(tc.wantSent), routerClient.sent)
			}
			for i, want := range tc.wantSent {
				got := routerClient.sent[i].LocalNsmGwPeerIPs
				if len(got) != len(want) {
					t.Fatalf("nexthops sent = %v, want %v", got, want)
				}
				for j := range want {
					if got[j] != want[j] {
						t.Errorf("nexthops sent = %v, want %v", got, want)
						break
					}
				}
			}

			if requeue != tc.wantRequeue {
				t.Errorf("requeue = %v, want %v", requeue, tc.wantRequeue)
			}
			if got := res.RequeueAfter.Seconds(); got != tc.wantAfterSec {
				t.Errorf("RequeueAfter = %vs, want %vs", got, tc.wantAfterSec)
			}
			if tc.wantRequeue && res.RequeueAfter >= controllers.SliceGatewayReconcileInterval {
				t.Errorf("RequeueAfter %v is no better than the steady state interval %v",
					res.RequeueAfter, controllers.SliceGatewayReconcileInterval)
			}
		})
	}
}

var _ WorkerRouterClientProvider = (*recordingRouterClient)(nil)
