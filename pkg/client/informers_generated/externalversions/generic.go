// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	projectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=projectcalico.org, Version=v3
	case v3.SchemeGroupVersion.WithResource("globalreporttypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().GlobalReportTypes().Informer()}, nil

		// Group=projectcalico.org, Version=projectcalico
	case projectcalico.SchemeGroupVersion.WithResource("globalreporttypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().Projectcalico().GlobalReportTypes().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
