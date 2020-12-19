// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	superedgeiov1 "superedge/pkg/application-grid-controller/apis/superedge.io/v1"
)

// FakeServiceGrids implements ServiceGridInterface
type FakeServiceGrids struct {
	Fake *FakeSuperedgeV1
	ns   string
}

var servicegridsResource = schema.GroupVersionResource{Group: "superedge.io", Version: "v1", Resource: "servicegrids"}

var servicegridsKind = schema.GroupVersionKind{Group: "superedge.io", Version: "v1", Kind: "ServiceGrid"}

// Get takes name of the serviceGrid, and returns the corresponding serviceGrid object, and an error if there is any.
func (c *FakeServiceGrids) Get(ctx context.Context, name string, options v1.GetOptions) (result *superedgeiov1.ServiceGrid, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicegridsResource, c.ns, name), &superedgeiov1.ServiceGrid{})

	if obj == nil {
		return nil, err
	}
	return obj.(*superedgeiov1.ServiceGrid), err
}

// List takes label and field selectors, and returns the list of ServiceGrids that match those selectors.
func (c *FakeServiceGrids) List(ctx context.Context, opts v1.ListOptions) (result *superedgeiov1.ServiceGridList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicegridsResource, servicegridsKind, c.ns, opts), &superedgeiov1.ServiceGridList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &superedgeiov1.ServiceGridList{ListMeta: obj.(*superedgeiov1.ServiceGridList).ListMeta}
	for _, item := range obj.(*superedgeiov1.ServiceGridList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceGrids.
func (c *FakeServiceGrids) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicegridsResource, c.ns, opts))

}

// Create takes the representation of a serviceGrid and creates it.  Returns the server's representation of the serviceGrid, and an error, if there is any.
func (c *FakeServiceGrids) Create(ctx context.Context, serviceGrid *superedgeiov1.ServiceGrid, opts v1.CreateOptions) (result *superedgeiov1.ServiceGrid, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicegridsResource, c.ns, serviceGrid), &superedgeiov1.ServiceGrid{})

	if obj == nil {
		return nil, err
	}
	return obj.(*superedgeiov1.ServiceGrid), err
}

// Update takes the representation of a serviceGrid and updates it. Returns the server's representation of the serviceGrid, and an error, if there is any.
func (c *FakeServiceGrids) Update(ctx context.Context, serviceGrid *superedgeiov1.ServiceGrid, opts v1.UpdateOptions) (result *superedgeiov1.ServiceGrid, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicegridsResource, c.ns, serviceGrid), &superedgeiov1.ServiceGrid{})

	if obj == nil {
		return nil, err
	}
	return obj.(*superedgeiov1.ServiceGrid), err
}

// Delete takes name of the serviceGrid and deletes it. Returns an error if one occurs.
func (c *FakeServiceGrids) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicegridsResource, c.ns, name), &superedgeiov1.ServiceGrid{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceGrids) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicegridsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &superedgeiov1.ServiceGridList{})
	return err
}

// Patch applies the patch and returns the patched serviceGrid.
func (c *FakeServiceGrids) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *superedgeiov1.ServiceGrid, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicegridsResource, c.ns, name, pt, data, subresources...), &superedgeiov1.ServiceGrid{})

	if obj == nil {
		return nil, err
	}
	return obj.(*superedgeiov1.ServiceGrid), err
}
