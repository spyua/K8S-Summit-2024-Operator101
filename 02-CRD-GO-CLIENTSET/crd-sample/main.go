package main

import (
	"context"
	"crd-sample/pkg/clientset"
	"os"
	"path"

	"crd-sample/pkg/apis/myresource/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	if err != nil {
		panic(err.Error())
	}

	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	myresource := &v1alpha1.MyResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myresource",
		},
		Spec: v1alpha1.MyResourceSpec{
			Image: "nginx",
			Key:   "key",
			Value: "value",
		},
	}
	clientset.MygroupV1alpha1().MyResources("default").Create(context.Background(), myresource, metav1.CreateOptions{})

	_, err = clientset.MygroupV1alpha1().MyResources("default").Get(context.Background(), "myresource", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

}
