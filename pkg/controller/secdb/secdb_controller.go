/*
Copyright 2019 The Kubernetes Authors.

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

package secdb

import (
	"context"
	"fmt"
	"github.com/sanjid133/secdb/util"
	"k8s.io/klog"

	//"reflect"

	secdbv1beta1 "github.com/sanjid133/secdb/pkg/apis/secdb/v1beta1"
	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SecDb Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSecDb{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("secdb-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	fmt.Println("Watching secdb")
	// Watch for changes to SecDb
	err = c.Watch(&source.Kind{Type: &secdbv1beta1.SecDb{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &secdbv1beta1.SecDb{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSecDb{}

// ReconcileSecDb reconciles a SecDb object
type ReconcileSecDb struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SecDb object and makes changes based on the state read
// and what is in the SecDb.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=secdb.k8s.io,resources=secdbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=secdb.k8s.io,resources=secdbs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets/status,verbs=get;update;patch
func (r *ReconcileSecDb) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	fmt.Println("reconciling.")

	ctx := context.TODO()
	// Fetch the SecDb instance
	instance := &secdbv1beta1.SecDb{}
	err := r.Get(ctx, request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Implement controller logic here
	name := instance.Name
	klog.Infof("Running reconcile SecDb for %s\n", name)

	// If object hasn't been deleted and doesn't have a finalizer, add one
	// Add a finalizer to newly created objects.
	if instance.ObjectMeta.DeletionTimestamp.IsZero() &&
		!util.Contains(instance.ObjectMeta.Finalizers, secdbv1beta1.SecDbFinalizer) {
		instance.Finalizers = append(instance.Finalizers, secdbv1beta1.SecDbFinalizer)
		if err = r.Client.Update(ctx, instance); err != nil {
			klog.Infof("failed to add finalizer to machine object %v due to error %v.", name, err)
			return reconcile.Result{}, err
		}
	}

	if !instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// no-op if finalizer has been removed.
		if !util.Contains(instance.ObjectMeta.Finalizers, secdbv1beta1.SecDbFinalizer) {
			klog.Infof("reconciling secdb object %v causes a no-op as there is no finalizer.", name)
			return reconcile.Result{}, nil
		}

		klog.Infof("reconciling secdb object %v triggers delete.", name)
		if err := r.drop(instance); err != nil {
			klog.Errorf("Error deleting secdb object %v; %v", name, err)
			// requee ???
			return reconcile.Result{}, err
		}

		// Remove finalizer on successful deletion.
		klog.Infof("secdb object %v deletion successful, removing finalizer.", name)
		instance.ObjectMeta.Finalizers = util.Filter(instance.ObjectMeta.Finalizers, secdbv1beta1.SecDbFinalizer)
		if err := r.Client.Update(context.Background(), instance); err != nil {
			klog.Errorf("Error removing finalizer from secdb object %v; %v", name, err)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	klog.Infof("Reconciling secdb object %v triggers idempotent create.", instance.ObjectMeta.Name)
	if err := r.upsert(ctx, instance); err != nil {
		klog.Warningf("unable to create secdb %v: %v", name, err)
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}
