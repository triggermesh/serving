/*
Copyright 2019 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/serving/pkg/apis/networking/v1alpha1"
)

// FakeServerlessServices implements ServerlessServiceInterface
type FakeServerlessServices struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var serverlessservicesResource = schema.GroupVersionResource{Group: "networking.internal.knative.dev", Version: "v1alpha1", Resource: "serverlessservices"}

var serverlessservicesKind = schema.GroupVersionKind{Group: "networking.internal.knative.dev", Version: "v1alpha1", Kind: "ServerlessService"}

// Get takes name of the serverlessService, and returns the corresponding serverlessService object, and an error if there is any.
func (c *FakeServerlessServices) Get(name string, options v1.GetOptions) (result *v1alpha1.ServerlessService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serverlessservicesResource, c.ns, name), &v1alpha1.ServerlessService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerlessService), err
}

// List takes label and field selectors, and returns the list of ServerlessServices that match those selectors.
func (c *FakeServerlessServices) List(opts v1.ListOptions) (result *v1alpha1.ServerlessServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serverlessservicesResource, serverlessservicesKind, c.ns, opts), &v1alpha1.ServerlessServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServerlessServiceList{ListMeta: obj.(*v1alpha1.ServerlessServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServerlessServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serverlessServices.
func (c *FakeServerlessServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serverlessservicesResource, c.ns, opts))

}

// Create takes the representation of a serverlessService and creates it.  Returns the server's representation of the serverlessService, and an error, if there is any.
func (c *FakeServerlessServices) Create(serverlessService *v1alpha1.ServerlessService) (result *v1alpha1.ServerlessService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serverlessservicesResource, c.ns, serverlessService), &v1alpha1.ServerlessService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerlessService), err
}

// Update takes the representation of a serverlessService and updates it. Returns the server's representation of the serverlessService, and an error, if there is any.
func (c *FakeServerlessServices) Update(serverlessService *v1alpha1.ServerlessService) (result *v1alpha1.ServerlessService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serverlessservicesResource, c.ns, serverlessService), &v1alpha1.ServerlessService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerlessService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServerlessServices) UpdateStatus(serverlessService *v1alpha1.ServerlessService) (*v1alpha1.ServerlessService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serverlessservicesResource, "status", c.ns, serverlessService), &v1alpha1.ServerlessService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerlessService), err
}

// Delete takes name of the serverlessService and deletes it. Returns an error if one occurs.
func (c *FakeServerlessServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serverlessservicesResource, c.ns, name), &v1alpha1.ServerlessService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServerlessServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serverlessservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServerlessServiceList{})
	return err
}

// Patch applies the patch and returns the patched serverlessService.
func (c *FakeServerlessServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServerlessService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serverlessservicesResource, c.ns, name, data, subresources...), &v1alpha1.ServerlessService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerlessService), err
}
