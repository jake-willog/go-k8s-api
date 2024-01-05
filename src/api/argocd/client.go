package argocd

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/jake-willog/go-k8s-api.git/src/api/model"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Client struct {
	projectClient     project.ProjectServiceClient
	clusterClient     cluster.ClusterServiceClient
	applicationClient application.ApplicationServiceClient
}

func NewClient(c model.Connection) (*Client, error) {
	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: fmt.Sprintf(c.Address),
		Insecure:   true,
		AuthToken:  c.Token,
	})
	if err != nil {
		return nil, err
	}

	_, projectClient, err := apiClient.NewProjectClient()
	if err != nil {
		return nil, err
	}

	_, clusterClient, err := apiClient.NewClusterClient()
	if err != nil {
		return nil, err
	}

	_, applicationClient, err := apiClient.NewApplicationClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		projectClient:     projectClient,
		clusterClient:     clusterClient,
		applicationClient: applicationClient,
	}, nil
}

func (c *Client) GetClusters() ([]v1alpha1.Cluster, error) {
	cl, err := c.clusterClient.List(context.Background(), &cluster.ClusterQuery{})
	if err != nil {
		return nil, err
	}

	return cl.Items, nil
}

func (c *Client) CreateProject(name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Create(context.Background(), &project.ProjectCreateRequest{
		Project: &v1alpha1.AppProject{
			ObjectMeta: v1.ObjectMeta{
				Name: name,
			},
		},
	})
}

func (c *Client) GetAppliction() ([]v1alpha1.Application, error) {
	cl, err := c.applicationClient.List(context.Background(), &application.ApplicationQuery{})
	if err != nil {
		return nil, err
	}

	return cl.Items, nil
}

func (c *Client) CreateApplication(name string) (*v1alpha1.Application, error) {
	return c.applicationClient.Create(context.Background(), &application.ApplicationCreateRequest{
		Application: &v1alpha1.Application{
			ObjectMeta: v1.ObjectMeta{
				Name: name,
			},
		},
	})
}

// func (c *Client) GetProject(name string) (*v1alpha1.AppProject, error) {
// 	return c.projectClient.Get(context.Background(), &project.ProjectQuery{
// 		Name: name,
// 	})
// }

// func (c *Client) AddDestination(projectName, server, namespace, name string) error {
// 	p, err := c.GetProject(projectName)
// 	if err != nil {
// 		return err
// 	}

// 	p.Spec.Destinations = []v1alpha1.ApplicationDestination{
// 		{
// 			Server:    server,
// 			Namespace: namespace,
// 			Name:      name,
// 		},
// 	}

// 	_, err = c.projectClient.Update(context.Background(), &project.ProjectUpdateRequest{
// 		Project: p,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *Client) CreateProject(name string) (*v1alpha1.AppProject, error) {
// 	return c.projectClient.Create(context.Background(), &project.ProjectCreateRequest{
// 		Project: &v1alpha1.AppProject{
// 			ObjectMeta: v1.ObjectMeta{
// 				Name: name,
// 			},
// 		},
// 	})
// }

// func (c *Client) DeleteProject(name string) error {
// 	_, err := c.projectClient.Delete(context.Background(), &project.ProjectQuery{
// 		Name: name,
// 	})

// 	return err
// }
