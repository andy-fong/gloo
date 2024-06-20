// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionHttpListenerOptionFunc func(original, desired *HttpListenerOption) (bool, error)

type HttpListenerOptionReconciler interface {
	Reconcile(namespace string, desiredResources HttpListenerOptionList, transition TransitionHttpListenerOptionFunc, opts clients.ListOpts) error
}

func httpListenerOptionsToResources(list HttpListenerOptionList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, httpListenerOption := range list {
		resourceList = append(resourceList, httpListenerOption)
	}
	return resourceList
}

func NewHttpListenerOptionReconciler(client HttpListenerOptionClient, statusSetter resources.StatusSetter) HttpListenerOptionReconciler {
	return &httpListenerOptionReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type httpListenerOptionReconciler struct {
	base reconcile.Reconciler
}

func (r *httpListenerOptionReconciler) Reconcile(namespace string, desiredResources HttpListenerOptionList, transition TransitionHttpListenerOptionFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "httpListenerOption_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*HttpListenerOption), desired.(*HttpListenerOption))
		}
	}
	return r.base.Reconcile(namespace, httpListenerOptionsToResources(desiredResources), transitionResources, opts)
}