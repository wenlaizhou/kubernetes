/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	"k8s.io/kubernetes/pkg/client/listers/policy/internalversion"
)

// PodDisruptionBudgetInformer provides access to a shared informer and lister for
// PodDisruptionBudgets.
type PodDisruptionBudgetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PodDisruptionBudgetLister
}

type podDisruptionBudgetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodDisruptionBudgetInformer constructs a new informer for PodDisruptionBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodDisruptionBudgetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodDisruptionBudgetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodDisruptionBudgetInformer constructs a new informer for PodDisruptionBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodDisruptionBudgetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Policy().PodDisruptionBudgets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Policy().PodDisruptionBudgets(namespace).Watch(options)
			},
		},
		&policy.PodDisruptionBudget{},
		resyncPeriod,
		indexers,
	)
}

func (f *podDisruptionBudgetInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodDisruptionBudgetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podDisruptionBudgetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policy.PodDisruptionBudget{}, f.defaultInformer)
}

func (f *podDisruptionBudgetInformer) Lister() internalversion.PodDisruptionBudgetLister {
	return internalversion.NewPodDisruptionBudgetLister(f.Informer().GetIndexer())
}
