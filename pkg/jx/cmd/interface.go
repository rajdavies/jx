package cmd

import (
	"io"

	"github.com/heptio/sonobuoy/pkg/client"
	"github.com/jenkins-x/jx/pkg/table"

	"github.com/jenkins-x/golang-jenkins"
	"github.com/jenkins-x/jx/pkg/auth"
	"github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metricsclient "k8s.io/metrics/pkg/client/clientset_generated/clientset"

	// this is so that we load the auth plugins so we can connect to, say, GCP

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// Factory is the interface defined for jx interactions via the cli
//go:generate pegomock generate github.com/jenkins-x/jx/pkg/jx/cmd Factory -o mocks/factory.go
type Factory interface {
	CreateJenkinsClient(kubeClient kubernetes.Interface, ns string) (*gojenkins.Jenkins, error)

	GetJenkinsURL(kubeClient kubernetes.Interface, ns string) (string, error)

	CreateAuthConfigService(fileName string) (auth.AuthConfigService, error)

	CreateJenkinsAuthConfigService(kubernetes.Interface, string) (auth.AuthConfigService, error)

	CreateChartmuseumAuthConfigService() (auth.AuthConfigService, error)

	CreateIssueTrackerAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)

	CreateChatAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)

	CreateAddonAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)

	CreateClient() (kubernetes.Interface, string, error)

	CreateKubeConfig() (*rest.Config, error)

	CreateJXClient() (versioned.Interface, string, error)

	CreateApiExtensionsClient() (apiextensionsclientset.Interface, error)

	CreateMetricsClient() (*metricsclient.Clientset, error)

	CreateComplianceClient() (*client.SonobuoyClient, error)

	CreateTable(out io.Writer) table.Table

	SetBatch(batch bool)

	ImpersonateUser(user string) Factory

	IsInCluster() bool

	IsInCDPIpeline() bool

	AuthMergePipelineSecrets(config *auth.AuthConfig, secrets *corev1.SecretList, kind string, isCDPipeline bool) error
}
