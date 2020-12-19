// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	scheme "superedge/pkg/application-grid-controller/generated/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "superedge/pkg/application-grid-controller/apis/superedge.io/v1"
)

// ServiceGridsGetter has a method to return a ServiceGridInterface.
// A group's client should implement this interface.
type ServiceGridsGetter interface {
	ServiceGrids(namespace string) ServiceGridInterface
}

// ServiceGridInterface has methods to work with ServiceGrid resources.
type ServiceGridInterface interface {
	Create(ctx context.Context, serviceGrid *v1.ServiceGrid, opts metav1.CreateOptions) (*v1.ServiceGrid, error)
	Update(ctx context.Context, serviceGrid *v1.ServiceGrid, opts metav1.UpdateOptions) (*v1.ServiceGrid, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ServiceGrid, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceGridList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceGrid, err error)
	ServiceGridExpansion
}

// serviceGrids implements ServiceGridInterface
type serviceGrids struct {
	client rest.Interface
	ns     string
}

// newServiceGrids returns a ServiceGrids
func newServiceGrids(c *SuperedgeV1Client, namespace string) *serviceGrids {
	return &serviceGrids{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceGrid, and returns the corresponding serviceGrid object, and an error if there is any.
func (c *serviceGrids) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ServiceGrid, err error) {
	result = &v1.ServiceGrid{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicegrids").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceGrids that match those selectors.
func (c *serviceGrids) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServiceGridList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ServiceGridList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicegrids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceGrids.
func (c *serviceGrids) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicegrids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceGrid and creates it.  Returns the server's representation of the serviceGrid, and an error, if there is any.
func (c *serviceGrids) Create(ctx context.Context, serviceGrid *v1.ServiceGrid, opts metav1.CreateOptions) (result *v1.ServiceGrid, err error) {
	result = &v1.ServiceGrid{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicegrids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceGrid).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceGrid and updates it. Returns the server's representation of the serviceGrid, and an error, if there is any.
func (c *serviceGrids) Update(ctx context.Context, serviceGrid *v1.ServiceGrid, opts metav1.UpdateOptions) (result *v1.ServiceGrid, err error) {
	result = &v1.ServiceGrid{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicegrids").
		Name(serviceGrid.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceGrid).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceGrid and deletes it. Returns an error if one occurs.
func (c *serviceGrids) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicegrids").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceGrids) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicegrids").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceGrid.
func (c *serviceGrids) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceGrid, err error) {
	result = &v1.ServiceGrid{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicegrids").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
