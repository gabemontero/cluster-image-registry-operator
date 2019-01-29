// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/cluster-image-registry-operator/pkg/apis/imageregistry/v1"
	scheme "github.com/openshift/cluster-image-registry-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigsGetter has a method to return a ConfigInterface.
// A group's client should implement this interface.
type ConfigsGetter interface {
	Configs() ConfigInterface
}

// ConfigInterface has methods to work with RegistryConfigs resources.
type ConfigInterface interface {
	Create(*v1.Config) (*v1.Config, error)
	Update(*v1.Config) (*v1.Config, error)
	UpdateStatus(*v1.Config) (*v1.Config, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Config, error)
	List(opts metav1.ListOptions) (*v1.ConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Config, err error)
	ConfigExpansion
}

// configs implements ConfigInterface
type configs struct {
	client rest.Interface
}

// newConfigs returns a Configs
func newConfigs(c *ImageregistryV1Client) *configs {
	return &configs{
		client: c.RESTClient(),
	}
}

// Get takes name of the config, and returns the corresponding config object, and an error if there is any.
func (c *configs) Get(name string, options metav1.GetOptions) (result *v1.Config, err error) {
	result = &v1.Config{}
	err = c.client.Get().
		Resource("configs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Configs that match those selectors.
func (c *configs) List(opts metav1.ListOptions) (result *v1.ConfigList, err error) {
	result = &v1.ConfigList{}
	err = c.client.Get().
		Resource("configs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configs.
func (c *configs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("configs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a config and creates it.  Returns the server's representation of the config, and an error, if there is any.
func (c *configs) Create(config *v1.Config) (result *v1.Config, err error) {
	result = &v1.Config{}
	err = c.client.Post().
		Resource("configs").
		Body(config).
		Do().
		Into(result)
	return
}

// Update takes the representation of a config and updates it. Returns the server's representation of the config, and an error, if there is any.
func (c *configs) Update(config *v1.Config) (result *v1.Config, err error) {
	result = &v1.Config{}
	err = c.client.Put().
		Resource("configs").
		Name(config.Name).
		Body(config).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configs) UpdateStatus(config *v1.Config) (result *v1.Config, err error) {
	result = &v1.Config{}
	err = c.client.Put().
		Resource("configs").
		Name(config.Name).
		SubResource("status").
		Body(config).
		Do().
		Into(result)
	return
}

// Delete takes name of the config and deletes it. Returns an error if one occurs.
func (c *configs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("configs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Resource("configs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched config.
func (c *configs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Config, err error) {
	result = &v1.Config{}
	err = c.client.Patch(pt).
		Resource("configs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
