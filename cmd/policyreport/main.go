/*
Copyright 2015 The Kubernetes Authors.
Modifications Copyright 2022 The KCP Authors.

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

package main

import (
	"context"
	"flag"
	"os"
	"sync"

	"github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	upstreamcache "k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"

	kcpscopedclient "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	kcpinformers "github.com/kcp-dev/kcp/pkg/client/informers/externalversions"

	clusterdynamic "github.com/kcp-dev/client-go/dynamic"
	clientopts "github.com/kubestellar/kubestellar/pkg/client-options"
	"github.com/kubestellar/kubestellar/pkg/mailboxwatch"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func main() {
	var useClusterClient bool
	fs := pflag.NewFlagSet("inform-namespaced-resources", pflag.ExitOnError)
	klog.InitFlags(flag.CommandLine)
	fs.AddGoFlagSet(flag.CommandLine)

	espwClientOpts := clientopts.NewClientOpts("espw", "access to the edge service provider workspace")
	espwClientOpts.AddFlags(fs)

	allClientOpts := clientopts.NewClientOpts("all", "access to the SyncerConfig objects in all clusters")
	allClientOpts.SetDefaultCurrentContext("system:admin")
	allClientOpts.AddFlags(fs)
	fs.BoolVar(&useClusterClient, "use-cluster-client", false, "If true, use ClusterClient")
	fs.Parse(os.Args)

	fs.Parse(os.Args[1:])

	ctx := context.Background()
	logger := klog.Background()
	ctx = klog.NewContext(ctx, logger)

	espwClientConfig, err := espwClientOpts.ToRESTConfig()
	if err != nil {
		logger.Error(err, "failed to make ESPW client config")
		os.Exit(2)
	}
	espwClientConfig.UserAgent = "inform-namespaced-resources"

	espwClient := kcpscopedclient.NewForConfigOrDie(espwClientConfig)

	allClientConfig, err := allClientOpts.ToRESTConfig()
	if err != nil {
		logger.Error(err, "failed to make all-cluster client config")
		os.Exit(2)
	}
	allClientConfig.UserAgent = "inform-policyreport"

	espwInformerFactory := kcpinformers.NewSharedScopedInformerFactory(espwClient, 0, "")
	mbPreInformer := espwInformerFactory.Tenancy().V1alpha1().Workspaces()

	// PolicyReport GVR/GVK(List)
	gvr := schema.GroupVersionResource{
		Group:    "wgpolicyk8s.io",
		Version:  "v1alpha2",
		Resource: "policyreports",
	}
	gvk := schema.GroupVersionKind{
		Group:   gvr.Group,
		Version: gvr.Version,
		Kind:    "PolicyReportList",
	}

	// arbitrary resource using dynamic client
	dynamicClusterClient, err := clusterdynamic.NewForConfig(allClientConfig)
	if err != nil {
		logger.Error(err, "Failed to create all-cluster dynamic client")
		os.Exit(30)
	}
	crClusterClient := dynamicClusterClient.Resource(gvr)
	informer := mailboxwatch.NewSharedInformer[dynamic.NamespaceableResourceInterface, *unstructured.UnstructuredList](ctx, gvk, mbPreInformer, crClusterClient, &unstructured.Unstructured{}, 0, upstreamcache.Indexers{})
	informer.AddEventHandler(upstreamcache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj any) { logUnst(logger, "add", obj) },
		UpdateFunc: func(oldObj, newObj any) { logUnst(logger, "update", newObj) },
		DeleteFunc: func(obj any) { logUnst(logger, "delete", obj) },
	})

	espwInformerFactory.Start(ctx.Done())
	upstreamcache.WaitForCacheSync(ctx.Done(), mbPreInformer.Informer().HasSynced)
	go informer.Run(ctx.Done())

	logger.Info("Running")

	<-ctx.Done()
}

var logmu sync.Mutex

func logUnst(logger klog.Logger, action string, obj any) {
	logmu.Lock()
	defer logmu.Unlock()
	x, ok := obj.(*unstructured.Unstructured)
	_ = ok
	logger.Info("Notified", "action", action, "name", x.GetName(), "namespace", x.GetNamespace(), "cluster", x.GetAnnotations()["kcp.io/cluster"])
}
