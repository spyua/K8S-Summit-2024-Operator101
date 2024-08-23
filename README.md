# K8S-Summit-2024-Operator101

## Create Folder Structure

```bash
mkdir crd-sample
cd crd-sample
go mod init crd-sample

mkdir hack
touch hack/boilerplate.go.txt

mkdir -p pkg/apis/myresource/v1alpha1
```

## Create types.go

```bash
vim pkg/apis/myresource/v1alpha1/types.go
```

```go
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyResourceSpec   `json:"spec"`
	Status MyResourceStatus `json:"status"`
}

type MyResourceSpec struct {
	Image string `json:"image"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MyResourceStatus struct {
	Completed bool `json:"completed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MyResource `json:"items"`
}
```

## Create doc.go

```bash
vim pkg/apis/mygroup.example.com/v1alpha1/doc.go
```

```go
// +k8s:deepcopy-gen=package
// +groupName=mygroup.example.com
package v1alpha1
```


## Install code generators

```bash
make install-generator
```


## Get Dependencies

```bash
go mod tidy
```

## Auto Generate Code

```bash
make auto-generate
```

## Build

```bash
go build main.go
```


## Reference
1. [最不厌其烦的 K8s 代码生成教程](https://www.zeng.dev/post/2023-k8s-api-codegen/)
2. [code-generator简单介绍](https://juejin.cn/post/7096484178128011277)