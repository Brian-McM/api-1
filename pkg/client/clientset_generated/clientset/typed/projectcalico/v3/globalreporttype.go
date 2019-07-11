// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalReportTypesGetter has a method to return a GlobalReportTypeInterface.
// A group's client should implement this interface.
type GlobalReportTypesGetter interface {
	GlobalReportTypes(namespace string) GlobalReportTypeInterface
}

// GlobalReportTypeInterface has methods to work with GlobalReportType resources.
type GlobalReportTypeInterface interface {
	Create(*v3.GlobalReportType) (*v3.GlobalReportType, error)
	Update(*v3.GlobalReportType) (*v3.GlobalReportType, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v3.GlobalReportType, error)
	List(opts v1.ListOptions) (*v3.GlobalReportTypeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.GlobalReportType, err error)
	GlobalReportTypeExpansion
}

// globalReportTypes implements GlobalReportTypeInterface
type globalReportTypes struct {
	client rest.Interface
	ns     string
}

// newGlobalReportTypes returns a GlobalReportTypes
func newGlobalReportTypes(c *ProjectcalicoV3Client, namespace string) *globalReportTypes {
	return &globalReportTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the globalReportType, and returns the corresponding globalReportType object, and an error if there is any.
func (c *globalReportTypes) Get(name string, options v1.GetOptions) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalreporttypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalReportTypes that match those selectors.
func (c *globalReportTypes) List(opts v1.ListOptions) (result *v3.GlobalReportTypeList, err error) {
	result = &v3.GlobalReportTypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalreporttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalReportTypes.
func (c *globalReportTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("globalreporttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a globalReportType and creates it.  Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *globalReportTypes) Create(globalReportType *v3.GlobalReportType) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("globalreporttypes").
		Body(globalReportType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a globalReportType and updates it. Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *globalReportTypes) Update(globalReportType *v3.GlobalReportType) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalreporttypes").
		Name(globalReportType.Name).
		Body(globalReportType).
		Do().
		Into(result)
	return
}

// Delete takes name of the globalReportType and deletes it. Returns an error if one occurs.
func (c *globalReportTypes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalreporttypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalReportTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalreporttypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched globalReportType.
func (c *globalReportTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.GlobalReportType, err error) {
	result = &v3.GlobalReportType{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("globalreporttypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
