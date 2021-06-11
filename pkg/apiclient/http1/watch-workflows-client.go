package http1

import (
	workflowpkg "github.com/wyzksp/workflow/pkg/apiclient/workflow"
)

type watchWorkflowsClient struct{ serverSentEventsClient }

func (f watchWorkflowsClient) Recv() (*workflowpkg.WorkflowWatchEvent, error) {
	v := &workflowpkg.WorkflowWatchEvent{}
	return v, f.RecvEvent(v)
}
