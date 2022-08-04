package vcsclient

import (
	"context"
	bitbucketv1 "github.com/gfleury/go-bitbucket-v1"
	"github.com/hubmat00/froggit-go/vcsutils"
	"net/http"
)

type OpenBitbucketServerClient struct {
	BitbucketServerClient
}

func NewOpenBitbucketServerClient(vcsInfo VcsInfo) (*OpenBitbucketServerClient, error) {
	client, err := NewBitbucketServerClient(vcsInfo)
	if err != nil {
		return nil, err
	}
	advancedClient := OpenBitbucketServerClient{}
	advancedClient.BitbucketServerClient = *client
	return &advancedClient, nil
}

func (client *OpenBitbucketServerClient) BuildBitbucketClient(ctx context.Context) (*bitbucketv1.DefaultApiService, error) {
	return client.buildBitbucketClient(ctx)
}

func (client *OpenBitbucketServerClient) BuildHTTPClient(ctx context.Context) *http.Client {
	return client.buildHTTPClient(ctx)
}

func (client *OpenBitbucketServerClient) ListProjects(bitbucketClient *bitbucketv1.DefaultApiService) ([]string, error) {
	return client.listProjects(bitbucketClient)
}

func CreatePaginationOptions(nextPageStart int) map[string]interface{} {
	return createPaginationOptions(nextPageStart)
}

func UnmarshalAPIResponseValues(response *bitbucketv1.APIResponse, target interface{}) error {
	return unmarshalAPIResponseValues(response, target)
}

func GetBitbucketServerWebhookID(r *bitbucketv1.APIResponse) (string, error) {
	return getBitbucketServerWebhookID(r)
}

func CreateBitbucketServerHook(token, payloadURL string, webhookEvents ...vcsutils.WebhookEvent) *map[string]interface{} {
	return createBitbucketServerHook(token, payloadURL, webhookEvents...)
}

func GetBitbucketServerWebhookEvents(webhookEvents ...vcsutils.WebhookEvent) []string {
	return getBitbucketServerWebhookEvents(webhookEvents...)
}

func (client *OpenBitbucketServerClient) MapBitbucketServerCommitToCommitInfo(commit bitbucketv1.Commit, owner, repo string) CommitInfo {
	return client.mapBitbucketServerCommitToCommitInfo(commit, owner, repo)
}
