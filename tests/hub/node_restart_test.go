/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
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

package hub_test

import (
	"context"
	"os"
	"time"

	hubv1alpha1 "github.com/kubeslice/apis/pkg/controller/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
)

var _ = Describe("NodeRestart Test Suite", func() {
	var node1, node2, internalNode1, internalNode2 *v1.Node
	var ns *v1.Namespace
	var cluster *hubv1alpha1.Cluster
	var nsmconfig *v1.ConfigMap
	var nsmgrMock *appsv1.DaemonSet

	Context("With kubeslice node restarting", func() {
		BeforeEach(func() {
			//create 2 kubeslice gateway nodes
			node1 = &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "kubeslice-gw-node-1",
					Labels: map[string]string{
						"topology.kubernetes.io/region": "us-east-1",
						"kubeslice.io/node-type":        "gateway",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "gce://demo",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    v1.NodeExternalIP,
							Address: "35.235.10.1",
						},
					},
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, node1)).Should(Succeed())
			node2 = &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "kubeslice-gw-node-2",
					Labels: map[string]string{
						"topology.kubernetes.io/region": "us-east-1",
						"kubeslice.io/node-type":        "gateway",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "gce://demo",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    v1.NodeExternalIP,
							Address: "35.235.10.2",
						},
					},
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			}

			internalNode1 = &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "kubeslice-int-node-1",
					Labels: map[string]string{
						"topology.kubernetes.io/region": "us-east-1",
						"kubeslice.io/node-type":        "gateway",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "gce://demo",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    v1.NodeExternalIP,
							Address: "35.235.10.3",
						},
					},
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			}
			internalNode2 = &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "kubeslice-int-node-2",
					Labels: map[string]string{
						"topology.kubernetes.io/region": "us-east-1",
						"kubeslice.io/node-type":        "gateway",
					},
				},
				Spec: v1.NodeSpec{
					ProviderID: "gce://demo",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    v1.NodeExternalIP,
							Address: "35.235.10.4",
						},
					},
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			}
			// create project namespace (simulate controller cluster behaviour)
			ns = &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: PROJECT_NS,
				},
			}
			err := k8sClient.Get(ctx, types.NamespacedName{Name: ns.Name}, ns)
			if errors.IsNotFound(err) {
				Expect(k8sClient.Create(ctx, ns)).Should(Succeed())
			}
			// nsmgr daemonset (isNetworkPresent)
			nsmgrMock = &appsv1.DaemonSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nsmgr",
					Namespace: CONTROL_PLANE_NS,
					Labels:    map[string]string{"app": "nsmgr"},
				},
				Spec: appsv1.DaemonSetSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"app": "nsmgr"},
					},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "nsmgr"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name:  "nsmgr",
									Image: "containerImage:tag",
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, nsmgrMock)).Should(Succeed())
			// create cluster CR under project namespace
			cluster = &hubv1alpha1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "cluster-node-restart",
					Namespace: PROJECT_NS,
				},
				Spec: hubv1alpha1.ClusterSpec{
					ClusterProperty: hubv1alpha1.ClusterProperty{
						Monitoring: hubv1alpha1.Monitoring{
							KubernetesDashboard: hubv1alpha1.KubernetesDashboard{
								Enabled: true,
							},
						},
					},
				},
				Status: hubv1alpha1.ClusterStatus{},
			}

			err = k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: PROJECT_NS}, cluster)
			if errors.IsNotFound(err) {
				Expect(k8sClient.Create(ctx, cluster)).Should(Succeed())
			}
			nsmconfig = configMap("nsm-config", "kubeslice-system", `
Prefixes:
 - 192.168.0.0/16
 - 10.96.0.0/12`)
			err = k8sClient.Get(ctx, types.NamespacedName{Name: nsmconfig.Name, Namespace: nsmconfig.Namespace}, nsmconfig)
			if errors.IsNotFound(err) {
				Expect(k8sClient.Create(ctx, nsmconfig)).Should(Succeed())
			}

			DeferCleanup(func() {
				// Delete cluster object
				Expect(k8sClient.Delete(ctx, cluster)).Should(Succeed())
				err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
					err := k8sClient.Get(ctx, types.NamespacedName{
						Name: cluster.Name, Namespace: cluster.Namespace,
					}, cluster)
					if err != nil {
						return err
					}
					// remove finalizer from cluster CR
					cluster.ObjectMeta.SetFinalizers([]string{})
					return k8sClient.Update(ctx, cluster)
				})
				Expect(err).To(BeNil())
				// delete nsmgr
				Expect(k8sClient.Delete(ctx, nsmgrMock)).Should(Succeed())
				Eventually(func() bool {
					err := k8sClient.Delete(ctx, node1)
					return errors.IsNotFound(err)
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
				Eventually(func() bool {
					err := k8sClient.Delete(ctx, node2)
					return errors.IsNotFound(err)
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			})
		})
		It("should update cluster CR with new node IP", func() {
			os.Setenv("CLUSTER_NAME", cluster.Name)
			os.Setenv("HUB_PROJECT_NAMESPACE", PROJECT_NS)
			ctx := context.Background()

			//get the cluster object
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())

			// verify if new node IP is updated on cluster CR
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				if err != nil {
					return false
				}
				return len(cluster.Status.NodeIPs) > 0
			}, time.Second*120, time.Millisecond*250).Should(BeTrue())

			//create another kubeslice node
			Expect(k8sClient.Create(ctx, node2)).Should(Succeed())
			// delete the node whose IP was selected to replicate node failure
			Expect(k8sClient.Delete(ctx, node1)).Should(Succeed())

			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				if err != nil {
					return false
				}
				return cluster.Status.NodeIPs[0] == "35.235.10.2"
			}, time.Second*240, time.Millisecond*250).Should(BeTrue())
		})

		It("should update cluster CR with new node IP", func() {
			Expect(k8sClient.Create(ctx, internalNode1)).Should(Succeed())
			os.Setenv("CLUSTER_NAME", cluster.Name)
			os.Setenv("HUB_PROJECT_NAMESPACE", PROJECT_NS)
			ctx := context.Background()

			//get the cluster object
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())

			// verify if new node IP is updated on cluster CR
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				if err != nil {
					return false
				}
				return len(cluster.Status.NodeIPs) > 0
			}, time.Second*120, time.Millisecond*250).Should(BeTrue())

			//create another kubeslice node
			Expect(k8sClient.Create(ctx, internalNode2)).Should(Succeed())
			// delete the node whose IP was selected to replicate node failure
			Expect(k8sClient.Delete(ctx, internalNode1)).Should(Succeed())
			// verify if new node IP is updated on cluster CR
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, cluster)
				if err != nil {
					return false
				}
				return cluster.Status.NodeIPs[0] == "35.235.10.4" || cluster.Status.NodeIPs[1] == "35.235.10.4"
			}, time.Second*240, time.Millisecond*250).Should(BeTrue())
		})
	})
})

func configMap(name, namespace, data string) *v1.ConfigMap {
	configMapData := make(map[string]string)
	configMapData["excluded_prefixes_output.yaml"] = data
	configMap := v1.ConfigMap{
		// TypeMeta: metav1.TypeMeta{
		// 	Kind:       "ConfigMap",
		// 	APIVersion: "v1",
		// },
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: configMapData,
	}
	return &configMap
}
