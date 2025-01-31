// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	apiv1alpha1 "github.com/solo-io/gloo/projects/gateway2/api/applyconfiguration/api/v1alpha1"
	v1alpha1 "github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	scheme "github.com/solo-io/gloo/projects/gateway2/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// HttpListenerPoliciesGetter has a method to return a HttpListenerPolicyInterface.
// A group's client should implement this interface.
type HttpListenerPoliciesGetter interface {
	HttpListenerPolicies(namespace string) HttpListenerPolicyInterface
}

// HttpListenerPolicyInterface has methods to work with HttpListenerPolicy resources.
type HttpListenerPolicyInterface interface {
	Create(ctx context.Context, httpListenerPolicy *v1alpha1.HttpListenerPolicy, opts v1.CreateOptions) (*v1alpha1.HttpListenerPolicy, error)
	Update(ctx context.Context, httpListenerPolicy *v1alpha1.HttpListenerPolicy, opts v1.UpdateOptions) (*v1alpha1.HttpListenerPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, httpListenerPolicy *v1alpha1.HttpListenerPolicy, opts v1.UpdateOptions) (*v1alpha1.HttpListenerPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HttpListenerPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HttpListenerPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HttpListenerPolicy, err error)
	Apply(ctx context.Context, httpListenerPolicy *apiv1alpha1.HttpListenerPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HttpListenerPolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, httpListenerPolicy *apiv1alpha1.HttpListenerPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HttpListenerPolicy, err error)
	HttpListenerPolicyExpansion
}

// httpListenerPolicies implements HttpListenerPolicyInterface
type httpListenerPolicies struct {
	*gentype.ClientWithListAndApply[*v1alpha1.HttpListenerPolicy, *v1alpha1.HttpListenerPolicyList, *apiv1alpha1.HttpListenerPolicyApplyConfiguration]
}

// newHttpListenerPolicies returns a HttpListenerPolicies
func newHttpListenerPolicies(c *GatewayV1alpha1Client, namespace string) *httpListenerPolicies {
	return &httpListenerPolicies{
		gentype.NewClientWithListAndApply[*v1alpha1.HttpListenerPolicy, *v1alpha1.HttpListenerPolicyList, *apiv1alpha1.HttpListenerPolicyApplyConfiguration](
			"httplistenerpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.HttpListenerPolicy { return &v1alpha1.HttpListenerPolicy{} },
			func() *v1alpha1.HttpListenerPolicyList { return &v1alpha1.HttpListenerPolicyList{} }),
	}
}
