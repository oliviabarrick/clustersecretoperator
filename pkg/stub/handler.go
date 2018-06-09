package stub

import (
	"context"

	"github.com/justinbarrick/clustersecretoperator/pkg/apis/clustersecret/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.ClusterSecret:
    namespaceList := &corev1.NamespaceList{
        TypeMeta: metav1.TypeMeta{
            Kind: "Namespace",
            APIVersion: "v1",
        },
    }

    err := sdk.List("", namespaceList)
    if err != nil {
        return err
    }

    for _, namespace := range namespaceList.Items {
		    err = sdk.Create(newClusterSecret(o, namespace.Name))
		    if err != nil && !errors.IsAlreadyExists(err) {
		        logrus.Errorf("Failed to create secret: %v", err)
		        return err
		    }
    }
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
func newClusterSecret(cr *v1alpha1.ClusterSecret, namespace string) *corev1.Secret {
	logrus.Infof("creating cluster secret %s in %s", cr.Name, namespace)
	labels := map[string]string{
		"clustersecret.name": cr.Name,
		"clustersecret.namespace": cr.Namespace,
	}
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cluster-" + cr.Name,
			Namespace: namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "ClusterSecret",
				}),
			},
			Labels: labels,
		},
		Data: cr.Data,
		StringData: cr.StringData,
		Type: cr.Type,
	}
}
