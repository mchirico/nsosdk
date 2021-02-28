/*
Copyright 2021.

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

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	topv1alpha1 "github.com/mchirico/ns-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	//topv1alpha1 "github.com/mchirico/ns-operator/api/v1alpha1"
)

// NamespecialReconciler reconciles a Namespecial object
type NamespecialReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// TODO-mmc: rbac:groups=*,resources=*,verbs=* is too much.

//+kubebuilder:rbac:groups=top.cwxstat.com,resources=namespecials,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=top.cwxstat.com,resources=namespecials/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=top.cwxstat.com,resources=namespecials/finalizers,verbs=update
// +kubebuilder:rbac:groups=*,resources=*,verbs=*

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Namespecial object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
// +kubebuilder:rbac:groups=*,resources=*,verbs=*
func (r *NamespecialReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("namespecial", req.NamespacedName)

	// TODO-fix: get this to work
	// Sample of creating an event
	// operator-sdk create api --group top --version v1alpha1 --kind Namespecial --resource --controller
	// instance := &topv1alpha1.Namespecial{}

	//if err := r.Get(context.TODO(), req.NamespacedName, instance); err != nil {
	//	r.recorder.Event(instance, "Reconcile", "Loop", fmt.Sprintf("Sample msg %s/%s", "one", "two"))
	//}

	var instance topv1alpha1.Namespecial
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {

		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.

		// return ctrl.Result{}, client.IgnoreNotFound(err)
		fmt.Printf("here2")

	}

	r.Log.Info("Start Reconcile:", "test", "one")
	// your logic here
	// Can we do this?
	ns := &corev1.Namespace{}
	ns.Name = "stuff"
	if err := r.Client.Get(context.TODO(), types.NamespacedName{Name: ns.Name}, ns); err != nil {
		fmt.Printf("Return: %s, err: %s\n", ns.Name, err)

		log.V(1).Info("here sample logging")

		if strings.Contains(err.Error(), "not found") {
			fmt.Printf("We create\n")
			err = r.Client.Create(context.TODO(), ns)
			if err != nil {
				fmt.Printf("E")
				return ctrl.Result{}, err
			}
			fmt.Printf("requeue\n\n")
			return ctrl.Result{}, nil
			//return ctrl.Result{Requeue: true}, nil
		}
	}

	fmt.Printf("\n")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NamespecialReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Namespace{}).
		Complete(r)
}
