package kubeinteraction

import (
	"testing"

	"github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/keys"
	apipac "github.com/openshift-pipelines/pipelines-as-code/pkg/apis/pipelinesascode/v1alpha1"
	"github.com/openshift-pipelines/pipelines-as-code/pkg/params/info"
	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"gotest.tools/v3/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAddLabelsAndAnnotations(t *testing.T) {
	event := info.NewEvent()
	event.Organization = "org"
	event.Repository = "repo"
	event.SHA = "sha"
	event.Sender = "sender"
	event.EventType = "pull_request"
	event.BaseBranch = "main"
	event.SHAURL = "https://url/sha"

	type args struct {
		event       *info.Event
		pipelineRun *tektonv1.PipelineRun
		repo        *apipac.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test label and annotation added to pr",
			args: args{
				event: event,
				pipelineRun: &tektonv1.PipelineRun{
					ObjectMeta: metav1.ObjectMeta{
						Labels:      map[string]string{},
						Annotations: map[string]string{},
					},
				},
				repo: &apipac.Repository{
					ObjectMeta: metav1.ObjectMeta{
						Name: "repo",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddLabelsAndAnnotations(tt.args.event, tt.args.pipelineRun, tt.args.repo, &info.ProviderConfig{})
			assert.Assert(t, tt.args.pipelineRun.Labels[keys.URLOrg] == tt.args.event.Organization, "'%s' != %s",
				tt.args.pipelineRun.Labels[keys.URLOrg], tt.args.event.Organization)
			assert.Assert(t, tt.args.pipelineRun.Annotations[keys.ShaURL] == tt.args.event.SHAURL)
		})
	}
}
