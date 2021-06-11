// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/wyzksp/workflow/pkg/apis/workflow/v1alpha1"
	scheme "github.com/wyzksp/workflow/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterWorkflowTemplatesGetter has a method to return a ClusterWorkflowTemplateInterface.
// A group's client should implement this interface.
type ClusterWorkflowTemplatesGetter interface {
	ClusterWorkflowTemplates() ClusterWorkflowTemplateInterface
}

// ClusterWorkflowTemplateInterface has methods to work with ClusterWorkflowTemplate resources.
type ClusterWorkflowTemplateInterface interface {
	Create(*v1alpha1.ClusterWorkflowTemplate) (*v1alpha1.ClusterWorkflowTemplate, error)
	Update(*v1alpha1.ClusterWorkflowTemplate) (*v1alpha1.ClusterWorkflowTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterWorkflowTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterWorkflowTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterWorkflowTemplate, err error)
	ClusterWorkflowTemplateExpansion
}

// clusterWorkflowTemplates implements ClusterWorkflowTemplateInterface
type clusterWorkflowTemplates struct {
	client rest.Interface
}

// newClusterWorkflowTemplates returns a ClusterWorkflowTemplates
func newClusterWorkflowTemplates(c *ArgoprojV1alpha1Client) *clusterWorkflowTemplates {
	return &clusterWorkflowTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterWorkflowTemplate, and returns the corresponding clusterWorkflowTemplate object, and an error if there is any.
func (c *clusterWorkflowTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	result = &v1alpha1.ClusterWorkflowTemplate{}
	err = c.client.Get().
		Resource("clusterworkflowtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterWorkflowTemplates that match those selectors.
func (c *clusterWorkflowTemplates) List(opts v1.ListOptions) (result *v1alpha1.ClusterWorkflowTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterWorkflowTemplateList{}
	err = c.client.Get().
		Resource("clusterworkflowtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterWorkflowTemplates.
func (c *clusterWorkflowTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterworkflowtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterWorkflowTemplate and creates it.  Returns the server's representation of the clusterWorkflowTemplate, and an error, if there is any.
func (c *clusterWorkflowTemplates) Create(clusterWorkflowTemplate *v1alpha1.ClusterWorkflowTemplate) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	result = &v1alpha1.ClusterWorkflowTemplate{}
	err = c.client.Post().
		Resource("clusterworkflowtemplates").
		Body(clusterWorkflowTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterWorkflowTemplate and updates it. Returns the server's representation of the clusterWorkflowTemplate, and an error, if there is any.
func (c *clusterWorkflowTemplates) Update(clusterWorkflowTemplate *v1alpha1.ClusterWorkflowTemplate) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	result = &v1alpha1.ClusterWorkflowTemplate{}
	err = c.client.Put().
		Resource("clusterworkflowtemplates").
		Name(clusterWorkflowTemplate.Name).
		Body(clusterWorkflowTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterWorkflowTemplate and deletes it. Returns an error if one occurs.
func (c *clusterWorkflowTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterworkflowtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterWorkflowTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterworkflowtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterWorkflowTemplate.
func (c *clusterWorkflowTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	result = &v1alpha1.ClusterWorkflowTemplate{}
	err = c.client.Patch(pt).
		Resource("clusterworkflowtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
