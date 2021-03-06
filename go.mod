module indeed.com/devops-incubation/harbor-container-webhook

go 1.13

require (
	github.com/containers/image/v5 v5.7.0
	github.com/goharbor/harbor/src v0.0.0-20200930121830-5293c8ff4b67
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/peterhellberg/link v1.1.0
	github.com/stretchr/testify v1.6.1
	github.com/vektra/mockery/v2 v2.3.0
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	gomodules.xyz/jsonpatch/v2 v2.0.1
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	oss.indeed.com/go/go-groups v1.1.3
	sigs.k8s.io/controller-runtime v0.6.3
)
