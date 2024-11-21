
package controllers

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type FruitController struct{}

func (c *FruitController) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log.Printf("Reconciling %s/%s", req.Namespace, req.Name)
	fruit := &corev1.ConfigMap{}
	if err := ctrl.GetClient().Get(ctx, req.NamespacedName, fruit); err != nil {
		return ctrl.Result{}, err
	}
	log.Printf("Processing Fruit: %s", fruit.Data["name"])
	return ctrl.Result{}, nil
}
