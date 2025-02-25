//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package clusterclient

import (
	"fmt"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned"
	"github.com/kcp-dev/logicalcluster"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"

	examplev1client "github.com/kcp-dev/code-generator/examples/pkg/clusterclient/typed/example/v1"
	examplev1 "github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned/typed/example/v1"
)

// NewForConfig creates a new ClusterClient for the given config.
// It uses a custom round tripper that wraps the given client's
// endpoint. The clientset returned from NewForConfig is kcp
// cluster-aware.
func NewForConfig(config *rest.Config) (*ClusterClient, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}

	clusterRoundTripper := kcp.NewClusterRoundTripper(client.Transport)
	client.Transport = clusterRoundTripper

	delegate, err := versioned.NewForConfigAndClient(config, client)
	if err != nil {
		return nil, fmt.Errorf("error creating delegate clientset: %w", err)
	}

	return &ClusterClient{
		delegate: delegate,
	}, nil
}

// ClusterClient wraps the underlying interface.
type ClusterClient struct {
	delegate versioned.Interface
}

// Cluster returns a wrapped interface scoped to a particular cluster.
func (c *ClusterClient) Cluster(cluster logicalcluster.Name) versioned.Interface {
	return &wrappedInterface{
		cluster:  cluster,
		delegate: c.delegate,
	}
}

type wrappedInterface struct {
	cluster  logicalcluster.Name
	delegate versioned.Interface
}

// Discovery retrieves the DiscoveryClient.
func (w *wrappedInterface) Discovery() discovery.DiscoveryInterface {
	return w.delegate.Discovery()
}

// ExampleV1 retrieves the ExampleV1Client.
func (w *wrappedInterface) ExampleV1() examplev1.ExampleV1Interface {
	return examplev1client.New(w.cluster, w.delegate.ExampleV1())
}
