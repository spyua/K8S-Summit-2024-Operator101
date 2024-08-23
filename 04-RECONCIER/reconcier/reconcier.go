package main

import (
	"context"
	"fmt"
	myresourcev1alpha1 "reconcier/pkg/apis/myresource/v1alpha1"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type MyReconciler struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *MyReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	log := log.FromContext(ctx)

	sample := &myresourcev1alpha1.MyResource{}
	err := r.client.Get(ctx, req.NamespacedName, sample)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("MyResource resource not found. Ignoring since object must be deleted")
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get sample")
		return reconcile.Result{}, err
	}

	foundJob := &batchv1.Job{}
	err = r.client.Get(ctx, types.NamespacedName{Name: sample.Name, Namespace: sample.Namespace}, foundJob)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.newJob(sample)
		log.Info("Creating a new Job", "Job.Namespace", dep.Namespace, "Job.Name", dep.Name)
		err = r.client.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new Job", "Job.Namespace", dep.Namespace, "Job.Name", dep.Name)
			return reconcile.Result{}, err
		}
		// Deployment created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Job")
		return reconcile.Result{}, err
	}

	foundCM := &corev1.ConfigMap{}
	err = r.client.Get(ctx, types.NamespacedName{Name: sample.Name, Namespace: sample.Namespace}, foundCM)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.newConfigMap(sample)
		log.Info("Creating a new ConfigMap", "ConfigMap.Namespace", dep.Namespace, "ConfigMap.Name", dep.Name)
		err = r.client.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new ConfigMap", "ConfigMap.Namespace", dep.Namespace, "ConfigMap.Name", dep.Name)
			return reconcile.Result{}, err
		}
		// Deployment created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get ConfigMap")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil

}

func (r *MyReconciler) newConfigMap(s *myresourcev1alpha1.MyResource) *corev1.ConfigMap {

	cm := &corev1.ConfigMap{
		Data: map[string]string{s.Spec.Key: s.Spec.Value},
	}
	cm.Name = s.Name
	cm.Namespace = s.Namespace

	ctrl.SetControllerReference(s, cm, r.scheme)
	return cm
}

func (r *MyReconciler) newJob(s *myresourcev1alpha1.MyResource) *batchv1.Job {

	job := &batchv1.Job{
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: s.Name,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   s.Spec.Image,
						Name:    s.Name,
						Command: []string{"cat", fmt.Sprintf("/tmp/cm/%s", s.Spec.Key)},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      s.Name,
								MountPath: "/tmp/cm",
							},
						},
					}},
					Volumes: []corev1.Volume{
						{
							Name: s.Name,
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: s.Name,
									},
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
	}
	job.Name = s.Name
	job.Namespace = s.Namespace

	ctrl.SetControllerReference(s, job, r.scheme)
	return job
}
