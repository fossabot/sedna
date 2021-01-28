// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	sednav1alpha1 "github.com/kubeedge/sedna/pkg/apis/sedna/v1alpha1"
	versioned "github.com/kubeedge/sedna/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeedge/sedna/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeedge/sedna/pkg/client/listers/sedna/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// JointInferenceServiceInformer provides access to a shared informer and lister for
// JointInferenceServices.
type JointInferenceServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.JointInferenceServiceLister
}

type jointInferenceServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewJointInferenceServiceInformer constructs a new informer for JointInferenceService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJointInferenceServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredJointInferenceServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJointInferenceServiceInformer constructs a new informer for JointInferenceService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJointInferenceServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SednaV1alpha1().JointInferenceServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SednaV1alpha1().JointInferenceServices(namespace).Watch(context.TODO(), options)
			},
		},
		&sednav1alpha1.JointInferenceService{},
		resyncPeriod,
		indexers,
	)
}

func (f *jointInferenceServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJointInferenceServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jointInferenceServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sednav1alpha1.JointInferenceService{}, f.defaultInformer)
}

func (f *jointInferenceServiceInformer) Lister() v1alpha1.JointInferenceServiceLister {
	return v1alpha1.NewJointInferenceServiceLister(f.Informer().GetIndexer())
}