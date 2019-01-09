package secdb

import (
	"context"
	"fmt"
	secdbv1beta1 "github.com/sanjid133/secdb/pkg/apis/secdb/v1beta1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	controllerclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileSecDb) upsert(ctx context.Context, secdb *secdbv1beta1.SecDb) error {
	currentSecrets := make(map[string]string, 0)
	for _, secret := range secdb.Spec.Entities {
		found := &core.Secret{}
		err := r.Get(ctx, types.NamespacedName{Name: secret.Name, Namespace: secdb.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			klog.Infof("Creating Secret in namespace %v, name %v", secdb.Namespace, secret.Name)
			if found, err = r.create(secdb, &secret); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		if !metav1.IsControlledBy(found, secdb) {
			msg := fmt.Sprintf("Resource %q already exists and is not managed by Something", secret.Name)
			return fmt.Errorf(msg)
		}
		currentSecrets[secret.Name] = found.Labels[secdbv1beta1.SecDbLabel]
		if isDataupdated(found.Data, secret.Data) {
			fmt.Println("data is updated...")
			found.Data = convertData(secret.Data)
			err = r.Update(ctx, found)
			if err != nil {
				return err
			}
		}
	}

	seclist := &core.SecretList{}
	err := r.List(ctx, &controllerclient.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{
			secdbv1beta1.SecDbLabel: secdb.Name,
		}),
	}, seclist)
	if err != nil {
		return err
	}

	for _, rs := range seclist.Items {
		if _, found := currentSecrets[rs.Name]; !found && rs.Labels[secdbv1beta1.SecDbLabel] == secdb.Name {
			if err = r.Delete(ctx, &rs); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ReconcileSecDb) drop(secdb *secdbv1beta1.SecDb) error {
	fmt.Println("Drop called")
	fmt.Println(secdb.Spec)
	return nil
}

func (r *ReconcileSecDb) create(secdb *secdbv1beta1.SecDb, secret *secdbv1beta1.EntitySpec) (*core.Secret, error) {
	fmt.Println("Creating secret....")
	sec := r.newSecret(secdb.Spec.Type, secdb.Namespace, secret)
	if err := controllerutil.SetControllerReference(secdb, sec, r.scheme); err != nil {
		return nil, err
	}
	sec.Labels = map[string]string{
		secdbv1beta1.SecDbLabel: secdb.Name,
	}

	if err := r.Create(context.TODO(), sec); err != nil {
		return nil, err
	}
	s := &core.Secret{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: secret.Name, Namespace: secdb.Namespace}, s)
	return s, err
}

func (r *ReconcileSecDb) newSecret(tp, namespace string, entity *secdbv1beta1.EntitySpec) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      entity.Name,
			Namespace: namespace,
		},
		Type: core.SecretType(tp),
		Data: convertData(entity.Data),
	}
}

func convertData(data map[string]string) map[string][]byte {
	retdata := make(map[string][]byte)
	for k, v := range data {
		retdata[k] = []byte(v)
	}
	return retdata
}

func isDataupdated(oldData map[string][]byte, newData map[string]string) bool {
	for k, v := range oldData {
		if string(v) != newData[k] {
			return true
		}
	}
	return false
}
