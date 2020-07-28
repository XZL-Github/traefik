/*
The MIT License (MIT)

Copyright (c) 2016-2020 Containous SAS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IngressRouteUDPsGetter has a method to return a IngressRouteUDPInterface.
// A group's client should implement this interface.
type IngressRouteUDPsGetter interface {
	IngressRouteUDPs(namespace string) IngressRouteUDPInterface
}

// IngressRouteUDPInterface has methods to work with IngressRouteUDP resources.
type IngressRouteUDPInterface interface {
	Create(*v1alpha1.IngressRouteUDP) (*v1alpha1.IngressRouteUDP, error)
	Update(*v1alpha1.IngressRouteUDP) (*v1alpha1.IngressRouteUDP, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IngressRouteUDP, error)
	List(opts v1.ListOptions) (*v1alpha1.IngressRouteUDPList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IngressRouteUDP, err error)
	IngressRouteUDPExpansion
}

// ingressRouteUDPs implements IngressRouteUDPInterface
type ingressRouteUDPs struct {
	client rest.Interface
	ns     string
}

// newIngressRouteUDPs returns a IngressRouteUDPs
func newIngressRouteUDPs(c *TraefikV1alpha1Client, namespace string) *ingressRouteUDPs {
	return &ingressRouteUDPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ingressRouteUDP, and returns the corresponding ingressRouteUDP object, and an error if there is any.
func (c *ingressRouteUDPs) Get(name string, options v1.GetOptions) (result *v1alpha1.IngressRouteUDP, err error) {
	result = &v1alpha1.IngressRouteUDP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IngressRouteUDPs that match those selectors.
func (c *ingressRouteUDPs) List(opts v1.ListOptions) (result *v1alpha1.IngressRouteUDPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IngressRouteUDPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ingressRouteUDPs.
func (c *ingressRouteUDPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ingressRouteUDP and creates it.  Returns the server's representation of the ingressRouteUDP, and an error, if there is any.
func (c *ingressRouteUDPs) Create(ingressRouteUDP *v1alpha1.IngressRouteUDP) (result *v1alpha1.IngressRouteUDP, err error) {
	result = &v1alpha1.IngressRouteUDP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		Body(ingressRouteUDP).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ingressRouteUDP and updates it. Returns the server's representation of the ingressRouteUDP, and an error, if there is any.
func (c *ingressRouteUDPs) Update(ingressRouteUDP *v1alpha1.IngressRouteUDP) (result *v1alpha1.IngressRouteUDP, err error) {
	result = &v1alpha1.IngressRouteUDP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		Name(ingressRouteUDP.Name).
		Body(ingressRouteUDP).
		Do().
		Into(result)
	return
}

// Delete takes name of the ingressRouteUDP and deletes it. Returns an error if one occurs.
func (c *ingressRouteUDPs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ingressRouteUDPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingressrouteudps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ingressRouteUDP.
func (c *ingressRouteUDPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IngressRouteUDP, err error) {
	result = &v1alpha1.IngressRouteUDP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ingressrouteudps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}