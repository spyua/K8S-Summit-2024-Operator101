package main

import (
	"context"
	myresourcev1alpha1 "deploy-operator/pkg/apis/myresource/v1alpha1"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var scheme = runtime.NewScheme()

func init() {
	log.SetLogger(zap.New())
	clientgoscheme.AddToScheme(scheme)
	myresourcev1alpha1.AddToScheme(scheme)
}

func main() {

	mgr, err := manager.New(
		config.GetConfigOrDie(),
		manager.Options{
			Scheme: scheme,
		},
	)

	if err != nil {
		panic(err)
	}

	err = builder.
		ControllerManagedBy(mgr).
		For(&myresourcev1alpha1.MyResource{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&batchv1.Job{}).
		Complete(&MyReconciler{
			client: mgr.GetClient(),
			scheme: mgr.GetScheme(),
		})

	if err != nil {
		panic(err)
	}

	err = mgr.Start(context.Background())

	if err != nil {
		panic(err)
	}
}
