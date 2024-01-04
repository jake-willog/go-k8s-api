package controller

import (
	"context"
	"fmt"

	"net/http"

	"github.com/jake-willog/go-k8s-api.git/src/api/model"
	"github.com/jake-willog/go-k8s-api.git/src/config"
	"github.com/jake-willog/go-k8s-api.git/src/logger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "user/v1/pod/create POST started")

	clientset, err := config.LoadKubeConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get kubeconfig %s", err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	resp := newOkResponse(w, reqId, "Ok")
	writeResponse(reqId, w, resp)
}

func GetPodInfo(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "user/v1/pod/info GET started")

	clientset, err := config.LoadKubeConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get kubeconfig %s", err)
	}

	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})

	podItems := make([]model.Pod, 0)
	for _, item := range pods.Items {
		podItems = append(podItems, model.Pod{
			Name:      item.GetName(),
			NameSpace: item.GetNamespace(),
			NodeName:  item.Spec.NodeName,
		})
	}

	resp := newOkResponse(w, reqId, "Ok")
	resp.Pods = &podItems
	writeResponse(reqId, w, resp)
}

func int32Ptr(i int32) *int32 { return &i }
