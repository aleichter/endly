package v1beta2



import (
	"fmt"
"errors"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	vvc "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"	
)

/*autogenerated contract adapter*/

//ReplicaSetCreateRequest represents request
type ReplicaSetCreateRequest struct {
  service_ v1beta2.ReplicaSetInterface
   *vvc.ReplicaSet
}

//ReplicaSetUpdateRequest represents request
type ReplicaSetUpdateRequest struct {
  service_ v1beta2.ReplicaSetInterface
   *vvc.ReplicaSet
}

//ReplicaSetUpdateStatusRequest represents request
type ReplicaSetUpdateStatusRequest struct {
  service_ v1beta2.ReplicaSetInterface
   *vvc.ReplicaSet
}

//ReplicaSetDeleteRequest represents request
type ReplicaSetDeleteRequest struct {
  service_ v1beta2.ReplicaSetInterface
  Name string
   *v1.DeleteOptions
}

//ReplicaSetDeleteCollectionRequest represents request
type ReplicaSetDeleteCollectionRequest struct {
  service_ v1beta2.ReplicaSetInterface
   *v1.DeleteOptions
  ListOptions v1.ListOptions
}

//ReplicaSetGetRequest represents request
type ReplicaSetGetRequest struct {
  service_ v1beta2.ReplicaSetInterface
  Name string
   v1.GetOptions
}

//ReplicaSetListRequest represents request
type ReplicaSetListRequest struct {
  service_ v1beta2.ReplicaSetInterface
   v1.ListOptions
}

//ReplicaSetWatchRequest represents request
type ReplicaSetWatchRequest struct {
  service_ v1beta2.ReplicaSetInterface
   v1.ListOptions
}

//ReplicaSetPatchRequest represents request
type ReplicaSetPatchRequest struct {
  service_ v1beta2.ReplicaSetInterface
  Name string
  Pt types.PatchType
  Data []byte
  Subresources []string
}


func init() {
	register(&ReplicaSetCreateRequest{})
	register(&ReplicaSetUpdateRequest{})
	register(&ReplicaSetUpdateStatusRequest{})
	register(&ReplicaSetDeleteRequest{})
	register(&ReplicaSetDeleteCollectionRequest{})
	register(&ReplicaSetGetRequest{})
	register(&ReplicaSetListRequest{})
	register(&ReplicaSetWatchRequest{})
	register(&ReplicaSetPatchRequest{})
}


func (r * ReplicaSetCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.ReplicaSet)
	return result, err	
}

func (r * ReplicaSetCreateRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Create";	
}

func (r * ReplicaSetUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.ReplicaSet)
	return result, err	
}

func (r * ReplicaSetUpdateRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Update";	
}

func (r * ReplicaSetUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.ReplicaSet)
	return result, err	
}

func (r * ReplicaSetUpdateStatusRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.UpdateStatus";	
}

func (r * ReplicaSetDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name,r.DeleteOptions)
	return result, err	
}

func (r * ReplicaSetDeleteRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Delete";	
}

func (r * ReplicaSetDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions,r.ListOptions)
	return result, err	
}

func (r * ReplicaSetDeleteCollectionRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.DeleteCollection";	
}

func (r * ReplicaSetGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name,r.GetOptions)
	return result, err	
}

func (r * ReplicaSetGetRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Get";	
}

func (r * ReplicaSetListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err	
}

func (r * ReplicaSetListRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.List";	
}

func (r * ReplicaSetWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err	
}

func (r * ReplicaSetWatchRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Watch";	
}

func (r * ReplicaSetPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta2.ReplicaSetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta2.ReplicaSetInterface", service)
	}
	return nil
}

func (r * ReplicaSetPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name,r.Pt,r.Data,r.Subresources...)
	return result, err	
}

func (r * ReplicaSetPatchRequest) GetId() string {
	return "apps/v1beta2.ReplicaSet.Patch";	
}
