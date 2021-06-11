// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/wyzksp/workflow/pkg/apis/workflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CronWorkflowLister helps list CronWorkflows.
type CronWorkflowLister interface {
	// List lists all CronWorkflows in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CronWorkflow, err error)
	// CronWorkflows returns an object that can list and get CronWorkflows.
	CronWorkflows(namespace string) CronWorkflowNamespaceLister
	CronWorkflowListerExpansion
}

// cronWorkflowLister implements the CronWorkflowLister interface.
type cronWorkflowLister struct {
	indexer cache.Indexer
}

// NewCronWorkflowLister returns a new CronWorkflowLister.
func NewCronWorkflowLister(indexer cache.Indexer) CronWorkflowLister {
	return &cronWorkflowLister{indexer: indexer}
}

// List lists all CronWorkflows in the indexer.
func (s *cronWorkflowLister) List(selector labels.Selector) (ret []*v1alpha1.CronWorkflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronWorkflow))
	})
	return ret, err
}

// CronWorkflows returns an object that can list and get CronWorkflows.
func (s *cronWorkflowLister) CronWorkflows(namespace string) CronWorkflowNamespaceLister {
	return cronWorkflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronWorkflowNamespaceLister helps list and get CronWorkflows.
type CronWorkflowNamespaceLister interface {
	// List lists all CronWorkflows in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CronWorkflow, err error)
	// Get retrieves the CronWorkflow from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CronWorkflow, error)
	CronWorkflowNamespaceListerExpansion
}

// cronWorkflowNamespaceLister implements the CronWorkflowNamespaceLister
// interface.
type cronWorkflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronWorkflows in the indexer for a given namespace.
func (s cronWorkflowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CronWorkflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronWorkflow))
	})
	return ret, err
}

// Get retrieves the CronWorkflow from the indexer for a given namespace and name.
func (s cronWorkflowNamespaceLister) Get(name string) (*v1alpha1.CronWorkflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cronworkflow"), name)
	}
	return obj.(*v1alpha1.CronWorkflow), nil
}
