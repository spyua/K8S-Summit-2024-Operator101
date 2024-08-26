# 初始化 Operator


```bash
cat << EOF > main.go
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
		For(&webv1.MyWeb{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Service{}).
		Owns(&appsv1.Deployment{}).
		Complete(&MyReconciler{})

	if err != nil {
		panic(err)
	}

	err = mgr.Start(context.Background())

	if err != nil {
		panic(err)
	}
}
EOF
```



```bash
cat << EOF > go.mod
module operator

go 1.22

require (
	k8s.io/api v0.29.2
	k8s.io/apimachinery v0.29.2
	k8s.io/client-go v0.29.2
	sigs.k8s.io/controller-runtime v0.17.6
)
EOF
```


```bash 
go mod tidy
go build
```