package v1beta1



import (
	"fmt"
"errors"
	"k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	vvc "k8s.io/api/policy/v1beta1"	
)

/*autogenerated contract adapter*/

//PodDisruptionBudgetCreateRequest represents request
type PodDisruptionBudgetCreateRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   *vvc.PodDisruptionBudget
}

//PodDisruptionBudgetUpdateRequest represents request
type PodDisruptionBudgetUpdateRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   *vvc.PodDisruptionBudget
}

//PodDisruptionBudgetUpdateStatusRequest represents request
type PodDisruptionBudgetUpdateStatusRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   *vvc.PodDisruptionBudget
}

//PodDisruptionBudgetDeleteRequest represents request
type PodDisruptionBudgetDeleteRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
  Name string
   *v1.DeleteOptions
}

//PodDisruptionBudgetDeleteCollectionRequest represents request
type PodDisruptionBudgetDeleteCollectionRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   *v1.DeleteOptions
  ListOptions v1.ListOptions
}

//PodDisruptionBudgetGetRequest represents request
type PodDisruptionBudgetGetRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
  Name string
   v1.GetOptions
}

//PodDisruptionBudgetListRequest represents request
type PodDisruptionBudgetListRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   v1.ListOptions
}

//PodDisruptionBudgetWatchRequest represents request
type PodDisruptionBudgetWatchRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
   v1.ListOptions
}

//PodDisruptionBudgetPatchRequest represents request
type PodDisruptionBudgetPatchRequest struct {
  service_ v1beta1.PodDisruptionBudgetInterface
  Name string
  Pt types.PatchType
  Data []byte
  Subresources []string
}


func init() {
	register(&PodDisruptionBudgetCreateRequest{})
	register(&PodDisruptionBudgetUpdateRequest{})
	register(&PodDisruptionBudgetUpdateStatusRequest{})
	register(&PodDisruptionBudgetDeleteRequest{})
	register(&PodDisruptionBudgetDeleteCollectionRequest{})
	register(&PodDisruptionBudgetGetRequest{})
	register(&PodDisruptionBudgetListRequest{})
	register(&PodDisruptionBudgetWatchRequest{})
	register(&PodDisruptionBudgetPatchRequest{})
}


func (r * PodDisruptionBudgetCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.PodDisruptionBudget)
	return result, err	
}

func (r * PodDisruptionBudgetCreateRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Create";	
}

func (r * PodDisruptionBudgetUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.PodDisruptionBudget)
	return result, err	
}

func (r * PodDisruptionBudgetUpdateRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Update";	
}

func (r * PodDisruptionBudgetUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.PodDisruptionBudget)
	return result, err	
}

func (r * PodDisruptionBudgetUpdateStatusRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.UpdateStatus";	
}

func (r * PodDisruptionBudgetDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name,r.DeleteOptions)
	return result, err	
}

func (r * PodDisruptionBudgetDeleteRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Delete";	
}

func (r * PodDisruptionBudgetDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions,r.ListOptions)
	return result, err	
}

func (r * PodDisruptionBudgetDeleteCollectionRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.DeleteCollection";	
}

func (r * PodDisruptionBudgetGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name,r.GetOptions)
	return result, err	
}

func (r * PodDisruptionBudgetGetRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Get";	
}

func (r * PodDisruptionBudgetListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err	
}

func (r * PodDisruptionBudgetListRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.List";	
}

func (r * PodDisruptionBudgetWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err	
}

func (r * PodDisruptionBudgetWatchRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Watch";	
}

func (r * PodDisruptionBudgetPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodDisruptionBudgetInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodDisruptionBudgetInterface", service)
	}
	return nil
}

func (r * PodDisruptionBudgetPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name,r.Pt,r.Data,r.Subresources...)
	return result, err	
}

func (r * PodDisruptionBudgetPatchRequest) GetId() string {
	return "policy/v1beta1.PodDisruptionBudget.Patch";	
}
