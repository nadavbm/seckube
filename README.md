# seckube
kubernetes operator for secret management inside kubernetes namespaces

created with kubebuilder:
```
kubebuilder init --domain eba.com --repo github.com/nadavbm/seckube
kubebuilder create api --group seckube --version v1alpha1 --kind Secret
make manifests
make generate
```