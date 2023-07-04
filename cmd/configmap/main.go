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
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	upstreamcache "k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"

	kcpscopedclient "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	kcpinformers "github.com/kcp-dev/kcp/pkg/client/informers/externalversions"

	clientopts "github.com/kubestellar/kubestellar/pkg/client-options"
	"github.com/kubestellar/kubestellar/pkg/mailboxwatch"

	kcpkubernetes "github.com/kcp-dev/client-go/kubernetes"
	corev1 "github.com/kcp-dev/client-go/kubernetes/typed/core/v1"
	coreapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8scoresv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func main() {
	namespace, ok := os.LookupEnv("NAMESPACE")
	if !ok {
		namespace = metav1.NamespaceAll
	}
	fs := pflag.NewFlagSet("inform-namespaced-resources", pflag.ExitOnError)
	klog.InitFlags(flag.CommandLine)
	fs.AddGoFlagSet(flag.CommandLine)

	espwClientOpts := clientopts.NewClientOpts("espw", "access to the edge service provider workspace")
	espwClientOpts.AddFlags(fs)

	allClientOpts := clientopts.NewClientOpts("all", "access to the SyncerConfig objects in all clusters")
	allClientOpts.SetDefaultCurrentContext("system:admin")
	allClientOpts.AddFlags(fs)

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
	allClientConfig.UserAgent = "inform-configmap"

	espwInformerFactory := kcpinformers.NewSharedScopedInformerFactory(espwClient, 0, "")
	mbPreInformer := espwInformerFactory.Tenancy().V1alpha1().Workspaces()

	// Configmap GVR/GVK(List)
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps",
	}
	gvk := schema.GroupVersionKind{
		Group:   gvr.Group,
		Version: gvr.Version,
		Kind:    "ConfigMapList",
	}

	// use mailboxwatch.FixNamespace
	kcpClientSet, err := kcpkubernetes.NewForConfig(allClientConfig)
	if err != nil {
		logger.Error(err, "Failed to create all-cluster dynamic client")
		os.Exit(30)
	}
	configmapSets := kcpClientSet.CoreV1().ConfigMaps()
	lw := mailboxwatch.FixNamespace[corev1.ConfigMapsNamespacer, k8scoresv1.ConfigMapInterface, *coreapi.ConfigMapList](configmapSets, namespace)
	informer := mailboxwatch.NewSharedInformer(ctx, gvk, mbPreInformer, lw, &coreapi.ConfigMap{}, 0, upstreamcache.Indexers{})
	informer.AddEventHandler(upstreamcache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj any) { logCm(logger, "add", obj) },
		UpdateFunc: func(oldObj, newObj any) { logCm(logger, "update", newObj) },
		DeleteFunc: func(obj any) { logCm(logger, "delete", obj) },
	})

	espwInformerFactory.Start(ctx.Done())
	upstreamcache.WaitForCacheSync(ctx.Done(), mbPreInformer.Informer().HasSynced)

	go informer.Run(ctx.Done())

	logger.Info("Running")

	<-ctx.Done()
}

var logmu sync.Mutex

func logCm(logger klog.Logger, action string, obj any) {
	logmu.Lock()
	defer logmu.Unlock()
	x, ok := obj.(*coreapi.ConfigMap)
	_ = ok
	logger.Info("Notified", "action", action, "name", x.GetName(), "namespace", x.GetNamespace(), "cluster", x.GetAnnotations()["kcp.io/cluster"])
}
