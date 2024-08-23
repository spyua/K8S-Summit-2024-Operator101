install-client-gen:
	go install k8s.io/code-generator/cmd/client-gen@v0.24.4

install-deepcopy-gen:
	go install k8s.io/code-generator/cmd/deepcopy-gen@v0.24.4

install-register-gen:
	go install k8s.io/code-generator/cmd/register-gen@v0.24.4


install-generator: install-client-gen install-deepcopy-gen install-register-gen

generate-deepcopy: install-deepcopy-gen
	deepcopy-gen \
	--input-dirs crd-sample/pkg/apis/myresource/v1alpha1 \
	-O zz_generated.deepcopy \
	--output-base .. \
	--go-header-file \
	./hack/boilerplate.go.txt

generate-clientset: install-client-gen
	client-gen \
	--clientset-name clientset \
	--input-base "" \
	--input crd-sample/pkg/apis/myresource/v1alpha1 \
	--output-package crd-sample/pkg/ \
	--output-base .. \
	--go-header-file ./hack/boilerplate.go.txt


generate-register: install-register-gen
	register-gen \
    -O zz_generated.register \
    --go-header-file ./hack/boilerplate.go.txt \
	--input-dirs crd-sample/pkg/apis/myresource/v1alpha1 \
    --output-base ..


auto-generate: generate-deepcopy generate-register generate-clientset


remove-generated:
	rm -rf pkg/clientset
	rm -rf pkg/apis/myresource/v1alpha1/zz_generated.deepcopy.go
	rm -rf pkg/apis/myresource/v1alpha1/zz_generated.register.go