/*
Copyright 2023 Compose, Cosmicrocks SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	acidcosmicv1 "github.com/cosmicrocks/scdl8/pkg/apis/acid.cosmic.rocks/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostgresqls implements PostgresqlInterface
type FakePostgresqls struct {
	Fake *FakeAcidV1
	ns   string
}

var postgresqlsResource = schema.GroupVersionResource{Group: "acid.cosmic.rocks", Version: "v1", Resource: "postgresqls"}

var postgresqlsKind = schema.GroupVersionKind{Group: "acid.cosmic.rocks", Version: "v1", Kind: "Postgresql"}

// Get takes name of the postgresql, and returns the corresponding postgresql object, and an error if there is any.
func (c *FakePostgresqls) Get(ctx context.Context, name string, options v1.GetOptions) (result *acidcosmicv1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresqlsResource, c.ns, name), &acidcosmicv1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidcosmicv1.Postgresql), err
}

// List takes label and field selectors, and returns the list of Postgresqls that match those selectors.
func (c *FakePostgresqls) List(ctx context.Context, opts v1.ListOptions) (result *acidcosmicv1.PostgresqlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresqlsResource, postgresqlsKind, c.ns, opts), &acidcosmicv1.PostgresqlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &acidcosmicv1.PostgresqlList{ListMeta: obj.(*acidcosmicv1.PostgresqlList).ListMeta}
	for _, item := range obj.(*acidcosmicv1.PostgresqlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgresqls.
func (c *FakePostgresqls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresqlsResource, c.ns, opts))

}

// Create takes the representation of a postgresql and creates it.  Returns the server's representation of the postgresql, and an error, if there is any.
func (c *FakePostgresqls) Create(ctx context.Context, postgresql *acidcosmicv1.Postgresql, opts v1.CreateOptions) (result *acidcosmicv1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresqlsResource, c.ns, postgresql), &acidcosmicv1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidcosmicv1.Postgresql), err
}

// Update takes the representation of a postgresql and updates it. Returns the server's representation of the postgresql, and an error, if there is any.
func (c *FakePostgresqls) Update(ctx context.Context, postgresql *acidcosmicv1.Postgresql, opts v1.UpdateOptions) (result *acidcosmicv1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresqlsResource, c.ns, postgresql), &acidcosmicv1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidcosmicv1.Postgresql), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostgresqls) UpdateStatus(ctx context.Context, postgresql *acidcosmicv1.Postgresql, opts v1.UpdateOptions) (*acidcosmicv1.Postgresql, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postgresqlsResource, "status", c.ns, postgresql), &acidcosmicv1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidcosmicv1.Postgresql), err
}

// Delete takes name of the postgresql and deletes it. Returns an error if one occurs.
func (c *FakePostgresqls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(postgresqlsResource, c.ns, name, opts), &acidcosmicv1.Postgresql{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgresqls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresqlsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &acidcosmicv1.PostgresqlList{})
	return err
}

// Patch applies the patch and returns the patched postgresql.
func (c *FakePostgresqls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *acidcosmicv1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresqlsResource, c.ns, name, pt, data, subresources...), &acidcosmicv1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidcosmicv1.Postgresql), err
}
