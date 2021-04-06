// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/package/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PackageLister helps list Packages.
// All objects returned here must be treated as read-only.
type PackageLister interface {
	// List lists all Packages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Package, err error)
	// Get retrieves the Package from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Package, error)
	PackageListerExpansion
}

// packageLister implements the PackageLister interface.
type packageLister struct {
	indexer cache.Indexer
}

// NewPackageLister returns a new PackageLister.
func NewPackageLister(indexer cache.Indexer) PackageLister {
	return &packageLister{indexer: indexer}
}

// List lists all Packages in the indexer.
func (s *packageLister) List(selector labels.Selector) (ret []*v1alpha1.Package, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Package))
	})
	return ret, err
}

// Get retrieves the Package from the index for a given name.
func (s *packageLister) Get(name string) (*v1alpha1.Package, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("package"), name)
	}
	return obj.(*v1alpha1.Package), nil
}
