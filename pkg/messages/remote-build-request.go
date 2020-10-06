package messages

const (
	RequestTypePR                = "PR"
	RequestTypeBranch            = "BRANCH"
	RequestTypeTag               = "TAG"
	RequestParameterKind         = "KIND"
	RequestParameterTarget       = "TARGET"
	RequestParameterRunScript    = "RUN_SCRIPT"
	RequestParameterRcvQueueName = "RCV_QUEUE_NAME"
)

type RemoteBuildRequestMessageParameters map[string]string

type RemoteBuildRequestMessage struct {
	Project   string                                `json:"project"`
	Token     string                                `json:"token"`
	Parameter []RemoteBuildRequestMessageParameters `json:"parameter"`
}

func NewRemoteBuildRequestMessage(project, token, repoURL, kind, target, runscript, recieveQueueName string) *RemoteBuildRequestMessage {
	r := &RemoteBuildRequestMessage{
		Project: project,
		Token:   token,
	}
	r.AddParameter("REPO_URL", repoURL)
	r.AddParameter(RequestParameterKind, kind)
	r.AddParameter(RequestParameterTarget, target)
	r.AddParameter(RequestParameterRunScript, runscript)
	r.AddParameter(RequestParameterRcvQueueName, recieveQueueName)
	return r
}

func (rbrm *RemoteBuildRequestMessage) AddParameter(name, value string) {
	p := make(RemoteBuildRequestMessageParameters)
	p["name"] = name
	p["value"] = value
	rbrm.Parameter = append(rbrm.Parameter, p)
}
