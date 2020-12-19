// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	versioned "superedge/pkg/application-grid-controller/generated/clientset/versioned"
	internalinterfaces "superedge/pkg/application-grid-controller/generated/informers/externalversions/internalinterfaces"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	superedgeiov1 "superedge/pkg/application-grid-controller/apis/superedge.io/v1"
	v1 "superedge/pkg/application-grid-controller/generated/listers/superedge.io/v1"
)

// DeploymentGridInformer provides access to a shared informer and lister for
// DeploymentGrids.
type DeploymentGridInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DeploymentGridLister
}

type deploymentGridInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeploymentGridInformer constructs a new informer for DeploymentGrid type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentGridInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeploymentGridInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeploymentGridInformer constructs a new informer for DeploymentGrid type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeploymentGridInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SuperedgeV1().DeploymentGrids(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SuperedgeV1().DeploymentGrids(namespace).Watch(context.TODO(), options)
			},
		},
		&superedgeiov1.DeploymentGrid{},
		resyncPeriod,
		indexers,
	)
}

func (f *deploymentGridInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeploymentGridInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deploymentGridInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&superedgeiov1.DeploymentGrid{}, f.defaultInformer)
}

func (f *deploymentGridInformer) Lister() v1.DeploymentGridLister {
	return v1.NewDeploymentGridLister(f.Informer().GetIndexer())
}
