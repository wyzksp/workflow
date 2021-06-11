package apiclient

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/wyzksp/workflow/persist/sqldb"
	"github.com/wyzksp/workflow/pkg/apiclient/clusterworkflowtemplate"
	"github.com/wyzksp/workflow/pkg/apiclient/cronworkflow"
	infopkg "github.com/wyzksp/workflow/pkg/apiclient/info"
	workflowpkg "github.com/wyzksp/workflow/pkg/apiclient/workflow"
	workflowarchivepkg "github.com/wyzksp/workflow/pkg/apiclient/workflowarchive"
	"github.com/wyzksp/workflow/pkg/apiclient/workflowtemplate"
	"github.com/wyzksp/workflow/pkg/client/clientset/versioned"
	"github.com/wyzksp/workflow/server/auth"
	clusterworkflowtmplserver "github.com/wyzksp/workflow/server/clusterworkflowtemplate"
	cronworkflowserver "github.com/wyzksp/workflow/server/cronworkflow"
	workflowserver "github.com/wyzksp/workflow/server/workflow"
	workflowtemplateserver "github.com/wyzksp/workflow/server/workflowtemplate"
	"github.com/wyzksp/workflow/util/help"
	"github.com/wyzksp/workflow/util/instanceid"
)

var argoKubeOffloadNodeStatusRepo = sqldb.ExplosiveOffloadNodeStatusRepo
var NoArgoServerErr = fmt.Errorf("this is impossible if you are not using the Argo Server, see " + help.CLI)

type argoKubeClient struct {
	instanceIDService instanceid.Service
}

func newArgoKubeClient(clientConfig clientcmd.ClientConfig, instanceIDService instanceid.Service) (context.Context, Client, error) {
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, nil, err
	}
	wfClient, err := versioned.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, err
	}
	gatekeeper, err := auth.NewGatekeeper(auth.Modes{auth.Server: true}, wfClient, kubeClient, restConfig, nil, auth.DefaultClientForAuthorization, "unused")
	if err != nil {
		return nil, nil, err
	}
	ctx, err := gatekeeper.Context(context.Background())
	if err != nil {
		return nil, nil, err
	}
	return ctx, &argoKubeClient{instanceIDService}, nil
}

func (a *argoKubeClient) NewWorkflowServiceClient() workflowpkg.WorkflowServiceClient {
	return &errorTranslatingWorkflowServiceClient{&argoKubeWorkflowServiceClient{workflowserver.NewWorkflowServer(a.instanceIDService, argoKubeOffloadNodeStatusRepo)}}
}

func (a *argoKubeClient) NewCronWorkflowServiceClient() cronworkflow.CronWorkflowServiceClient {
	return &errorTranslatingCronWorkflowServiceClient{&argoKubeCronWorkflowServiceClient{cronworkflowserver.NewCronWorkflowServer(a.instanceIDService)}}
}

func (a *argoKubeClient) NewWorkflowTemplateServiceClient() workflowtemplate.WorkflowTemplateServiceClient {
	return &errorTranslatingWorkflowTemplateServiceClient{&argoKubeWorkflowTemplateServiceClient{workflowtemplateserver.NewWorkflowTemplateServer(a.instanceIDService)}}
}

func (a *argoKubeClient) NewArchivedWorkflowServiceClient() (workflowarchivepkg.ArchivedWorkflowServiceClient, error) {
	return nil, NoArgoServerErr
}

func (a *argoKubeClient) NewInfoServiceClient() (infopkg.InfoServiceClient, error) {
	return nil, NoArgoServerErr
}

func (a *argoKubeClient) NewClusterWorkflowTemplateServiceClient() clusterworkflowtemplate.ClusterWorkflowTemplateServiceClient {
	return &errorTranslatingWorkflowClusterTemplateServiceClient{&argoKubeWorkflowClusterTemplateServiceClient{clusterworkflowtmplserver.NewClusterWorkflowTemplateServer(a.instanceIDService)}}
}
