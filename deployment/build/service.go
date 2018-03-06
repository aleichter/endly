package build

import (
	"fmt"
	"github.com/viant/endly"
	"github.com/viant/endly/deployment/deploy"
	"github.com/viant/endly/deployment/sdk"
	"github.com/viant/endly/system/exec"
	"github.com/viant/endly/system/storage"
	"github.com/viant/endly/workflow"
	"github.com/viant/toolbox/data"
	"github.com/viant/toolbox/url"
	"sync"
)

const (
	//ServiceID represent build service id
	ServiceID = "build"
)

type service struct {
	*endly.AbstractService
	mutex    *sync.RWMutex
	registry map[string]*Meta
}

func (s *service) getMeta(context *endly.Context, request *Request) (*Meta, error) {
	s.mutex.RLock()
	result, hasMeta := s.registry[request.BuildSpec.Name]
	s.mutex.RUnlock()
	var state = context.State()
	if !hasMeta {
		var metaURL = request.MetaURL
		if metaURL == "" {
			service, err := context.Service(workflow.ServiceID)
			if err != nil {
				return nil, err
			}
			if workflowService, ok := service.(*workflow.Service); ok {
				workflowResource, err := workflowService.Dao.NewRepoResource(state, fmt.Sprintf("meta/build/%v.json", request.BuildSpec.Name))
				if err != nil {
					return nil, err
				}
				metaURL = workflowResource.URL
			}
		}
		var credential = ""
		mainWorkflow := context.Workflow()
		if mainWorkflow != nil {
			credential = mainWorkflow.Source.Credential
		}
		response, err := s.loadMeta(context, &LoadMetaRequest{
			Source: url.NewResource(metaURL, credential),
		})
		if err != nil {
			return nil, err
		}
		result = response.Meta
	}
	return result, nil
}

func (s *service) loadMeta(context *endly.Context, request *LoadMetaRequest) (*LoadMetaResponse, error) {
	source, err := context.ExpandResource(request.Source)
	if err != nil {
		return nil, err
	}
	meta := &Meta{}
	err = source.JSONDecode(meta)
	if err != nil {
		return nil, fmt.Errorf("unable to decode: %v, %v", source.URL, err)
	}

	meta.goalsIndex = make(map[string]*Goal)
	indexBuildGoals(meta.Goals, meta.goalsIndex)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.registry[meta.Name] = meta
	return &LoadMetaResponse{
		Meta: meta,
	}, nil
}

func (s *service) deployDependencyIfNeeded(context *endly.Context, meta *Meta, spec *Spec, target *url.Resource) error {
	if len(meta.Dependencies) == 0 {
		return nil
	}
	deploymentService, err := context.Service(deploy.ServiceID)
	if err != nil {
		return err
	}
	for _, dependency := range meta.Dependencies {
		var app = context.Expand(dependency.Name)
		var version = context.Expand(dependency.Version)
		response := deploymentService.Run(context, &deploy.Request{
			AppName: app,
			Version: version,
			Target:  target,
		})
		if response.Err != nil {
			return err
		}
	}
	return nil
}

func indexBuildGoals(goals []*Goal, index map[string]*Goal) {
	if len(goals) == 0 {
		return
	}
	for _, goal := range goals {
		index[goal.Name] = goal
	}
}

func (s *service) setSdkIfNeeded(context *endly.Context, request *Request) error {
	if request.BuildSpec.Sdk == "" {
		return nil
	}
	sdkService, err := context.Service(sdk.ServiceID)
	if err != nil {
		return err
	}
	serviceResponse := sdkService.Run(context, &sdk.SetRequest{
		Target:  request.Target,
		Sdk:     request.BuildSpec.Sdk,
		Version: request.BuildSpec.SdkVersion,
		Env:     request.Env,
	})
	return serviceResponse.Err
}

func (s *service) build(context *endly.Context, request *Request) (*Response, error) {
	var result = &Response{}
	state := context.State()
	target, err := context.ExpandResource(request.Target)
	if err != nil {
		return nil, err
	}
	meta, err := s.getMeta(context, request)
	if err != nil {
		return nil, err
	}
	buildSpec := request.BuildSpec
	goal, has := meta.goalsIndex[buildSpec.Goal]
	if !has {
		return nil, fmt.Errorf("failed to lookup build %v goal: %v", buildSpec.Name, buildSpec.Goal)
	}

	buildState, err := newBuildState(buildSpec, target, request, context)
	if err != nil {
		return nil, err
	}
	state.Put("buildSpec", buildState)

	err = s.setSdkIfNeeded(context, request)
	if err != nil {
		return nil, err
	}

	err = s.deployDependencyIfNeeded(context, meta, buildSpec, target)
	if err != nil {
		return nil, err
	}
	if goal.InitTransfers != nil {
		_, err = storage.Copy(context, goal.InitTransfers.Transfers...)
		if err != nil {
			return nil, err
		}
	}

	if len(request.Credentials) > 0 {
		for _, execution := range goal.Command.Executions {
			if execution.MatchOutput != "" {
				execution.Credentials = request.Credentials
			}
		}
	}
	commandInfo, err := exec.Execute(context, target, goal.Command)
	if err != nil {
		return nil, err
	}
	result.CommandInfo = commandInfo

	if goal.PostTransfers != nil {
		_, err = storage.Copy(context, goal.PostTransfers.Transfers...)
		if err != nil {
			return nil, err
		}
	}

	if goal.VerificationCommand != nil {
		_, err = exec.Execute(context, target, goal.VerificationCommand)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
func newBuildState(buildSepc *Spec, target *url.Resource, request *Request, context *endly.Context) (data.Map, error) {
	target, err := context.ExpandResource(request.Target)
	if err != nil {
		return nil, err
	}
	build := data.NewMap()
	build.Put("name", buildSepc.Name)
	build.Put("version", buildSepc.Version)
	build.Put("args", buildSepc.Args)
	build.Put("goal", buildSepc.BuildGoal)
	build.Put("path", target.ParsedURL.Path)
	build.Put("host", target.ParsedURL.Host)
	build.Put("credential", target.Credential)
	build.Put("target", target)
	build.Put("sdk", buildSepc.Sdk)
	build.Put("sdkVersion", buildSepc.SdkVersion)
	return build, nil
}

const (
	buildGoBuildExample = `{
	"Spec": {
		"Name": "go",
		"Goal": "build",
		"Goal": "build",
		"Args": " -o echo",
		"Sdk": "go",
		"SdkVersion": "1.8"
	},
	"Env": {
		"GOOS": "linux"
	},
	"Target": {
		"URL": "scp://127.0.0.1/tmp/app/echo",
		"Credential": "${env.HOME}/.secret/localhost.json"
	}
}
`
	buildJavaBuildExample = `{
  "Spec": {
    "Name": "maven",
    "Version": "3.5",
    "Goal": "build",
    "Goal": "install",
    "Args": " -f pom.xml -am -pl server -DskipTest",
    "Sdk": "jdk",
    "SdkVersion": "1.7"
  },
 "Target": {
    "URL": "scp://127.0.0.1/tmp/app/server/",
    "Credential": "${env.HOME}/.secret/scp.json"
  }
}
`
)

func (s *service) registerRoutes() {
	s.Register(&endly.ServiceActionRoute{
		Action: "build",
		RequestInfo: &endly.ActionInfo{
			Description: "build app with supplied specification",
			Examples: []*endly.ExampleUseCase{
				{
					UseCase: "go app build",
					Data:    buildGoBuildExample,
				},
				{
					UseCase: "java app build",
					Data:    buildJavaBuildExample,
				},
			},
		},
		RequestProvider: func() interface{} {
			return &Request{}
		},
		ResponseProvider: func() interface{} {
			return &Response{}
		},
		Handler: func(context *endly.Context, request interface{}) (interface{}, error) {
			if req, ok := request.(*Request); ok {
				return s.build(context, req)
			}
			return nil, fmt.Errorf("unsupported request type: %T", request)
		},
	})

	s.Register(&endly.ServiceActionRoute{
		Action: "load",
		RequestInfo: &endly.ActionInfo{
			Description: "load build meta instruction",
		},
		RequestProvider: func() interface{} {
			return &LoadMetaRequest{}
		},
		ResponseProvider: func() interface{} {
			return &LoadMetaResponse{}
		},
		Handler: func(context *endly.Context, request interface{}) (interface{}, error) {
			if req, ok := request.(*LoadMetaRequest); ok {
				return s.loadMeta(context, req)
			}
			return nil, fmt.Errorf("unsupported request type: %T", request)
		},
	})

}

//New creates a new build service
func New() endly.Service {
	var result = &service{
		registry:        make(map[string]*Meta),
		mutex:           &sync.RWMutex{},
		AbstractService: endly.NewAbstractService(ServiceID),
	}
	result.AbstractService.Service = result
	result.registerRoutes()
	return result
}