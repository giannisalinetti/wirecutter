package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional), absolute path for the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path for the kubeconfig file")
	}
	flag.Parse()

	// Generate config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get the services list
	svcList, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// get the services and selector list
	for _, svc := range svcList.Items {
		fmt.Printf("Service name: %s\n", svc.Name)
		for k, v := range svc.Spec.Selector {
			fmt.Printf("\t%s: %s\n", k, v)
		}
	}

}
