package controller

import (
	"math/rand"

	"github.com/nadavbm/seckube/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$&*"

func buildSecret(ns string, s *v1alpha1.Secret) *v1.Secret {
	controller := true
	ownerRef := []metav1.OwnerReference{
		{
			APIVersion: s.APIVersion,
			Kind:       s.Kind,
			Name:       s.Name,
			UID:        s.UID,
			Controller: &controller,
		},
	}

	secret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
	}
	secret.ObjectMeta = buildObjectMetadata(s.Name, ns, ownerRef)
	secret.StringData = generateSecretData(s)

	return secret
}

//
// ------------------------------------------------------------------------------ secret generator helpers -----------------------------------------------------------------------------
//

func buildObjectMetadata(name, namespace string, ownerRef []metav1.OwnerReference) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:            name,
		Namespace:       namespace,
		OwnerReferences: ownerRef,
	}
}

func buildSecretObjectMetadata(name, namespace string, s *v1alpha1.Secret) metav1.ObjectMeta {
	controller := true
	return metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: s.APIVersion,
				Kind:       s.Kind,
				Name:       name,
				UID:        s.UID,
				Controller: &controller,
			},
		},
	}
}

func generateSecretData(secret *v1alpha1.Secret) map[string]string {
	m := make(map[string]string)
	for k, v := range secret.Spec.Data {
		if v == "" {
			v = randStringBytes(12)
		}
		m[k] = v
	}

	return m
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
