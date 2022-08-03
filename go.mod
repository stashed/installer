module stash.appscode.dev/installer

go 1.18

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/google/gofuzz v1.2.0
	github.com/spf13/pflag v1.0.5
	gomodules.xyz/semvers v0.0.0-20220316103017-cfbe8c37b63d
	k8s.io/api v0.24.3
	k8s.io/apimachinery v0.24.3
	kmodules.xyz/client-go v0.24.5
	kmodules.xyz/schema-checker v0.3.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/dustin/go-humanize v1.0.1-0.20220316001817-d5090ed65664 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	gomodules.xyz/encoding v0.0.5 // indirect
	gomodules.xyz/jsonpath v0.0.2 // indirect
	gomodules.xyz/mergo v0.3.13 // indirect
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gobuffalo/flect v0.2.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.3.0

replace github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8

replace sigs.k8s.io/controller-runtime => github.com/kmodules/controller-runtime v0.12.2-0.20220603144237-6cd001896bf3

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.24.2-rc.0.0.20220603191800-1c7484099dee

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20220603223637-59dad1716c43

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.25.0-alpha.0.0.20220603172133-1c9d09d1b3b8
