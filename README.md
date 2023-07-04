# Sample codes of KubeStellar mb-informer

This repository contains sample usages of KubeStellar mb-informer. You may also find sample usages in https://github.com/kubestellar/kubestellar/blob/main/cmd/inform-syncer-configs/main.go.

## Example usage for gathering ConfigMap objects of each mailbox workspace
1. Go to `root:espw` workspace
    ```
    kubectl ws root:espw
    ```
1. Run `./cmd/configmap/main.go` in GO
    ```
    go run ./cmd/configmap/main.go

    I0704 20:58:52.549842   99017 main.go:119] "Running"
    I0704 20:58:52.552996   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="local-path-storage" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553065   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-node-lease" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553072   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kubestellar-syncer-florin-2mg2dcem" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553084   99017 main.go:131] "Notified" action="add" name="kyverno-metrics" namespace="kyverno" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553089   99017 main.go:131] "Notified" action="add" name="kyverno" namespace="kyverno" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553094   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-public" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553098   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-system" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553102   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="default" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553106   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kyverno" cluster="1yrbrny6lrx8pzoz"
    I0704 20:58:52.553109   99017 main.go:131] "Notified" action="add" name="kyverno-metrics" namespace="kyverno" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553113   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-public" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553124   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="default" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553332   99017 main.go:131] "Notified" action="add" name="kyverno" namespace="kyverno" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553361   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-system" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553394   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kyverno" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553416   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kubestellar-syncer-guilder-1eu09n3i" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553426   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="local-path-storage" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:58:52.553430   99017 main.go:131] "Notified" action="add" name="kube-root-ca.crt" namespace="kube-node-lease" cluster="sgzrfvwvxjjpxr9x"
    ```

## Example usage for gathering PolicyReport objects of each mailbox workspace
1. Go to `root:espw` workspace
    ```
    kubectl ws root:espw
    ```
1. Run `./cmd/policyreport/main.go` in GO
    ```
    go run ./cmd/policyreport/main.go

    I0704 20:56:46.334875   98879 main.go:114] "Running"
    I0704 20:56:46.339750   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-node-lease" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339817   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-public" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339824   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-system" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339830   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kubestellar-syncer-florin-2mg2dcem" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339840   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kyverno" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339844   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="local-path-storage" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339848   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="default" cluster="1yrbrny6lrx8pzoz"
    I0704 20:56:46.339878   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kubestellar-syncer-guilder-1eu09n3i" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339882   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kyverno" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339887   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="local-path-storage" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339891   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="default" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339897   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-node-lease" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339901   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-public" cluster="sgzrfvwvxjjpxr9x"
    I0704 20:56:46.339905   98879 main.go:126] "Notified" action="add" name="cpol-sample-cluster-policy" namespace="kube-system" cluster="sgzrfvwvxjjpxr9x"
    ```

## Example usage for gathering any resource objects of each mailbox workspace
1. Go to `root:espw` workspace
    ```
    kubectl ws root:espw
    ```
1. Run `./cmd/anyresource/main.go` in GO
    ```
    go run ./cmd/anyresource/main.go --group wgpolicyk8s.io --version v1alpha2 --resource policyreports --kind PolicyReport

    I0704 21:12:44.327708     380 main.go:121] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Running"
    I0704 21:12:44.330754     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kyverno cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.330946     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=local-path-storage cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.330957     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=default cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.330963     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-node-lease cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.331182     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-public cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.331258     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-system cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.331292     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kubestellar-syncer-guilder-1eu09n3i cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:44.331392     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=default cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331404     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-node-lease cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331412     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-public cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331418     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kube-system cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331424     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kubestellar-syncer-florin-2mg2dcem cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331701     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=kyverno cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:44.331710     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=add name=cpol-sample-cluster-policy namespace=local-path-storage cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:46.741979     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=default cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:47.352673     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-node-lease cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:47.961449     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-public cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:48.285865     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-system cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:48.647108     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kubestellar-syncer-guilder-1eu09n3i cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:48.980984     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kyverno cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:49.592911     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=local-path-storage cluster=sgzrfvwvxjjpxr9x"
    I0704 21:12:49.937297     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=default cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:50.555536     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-node-lease cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:51.168072     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-public cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:51.487570     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kube-system cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:51.859920     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kubestellar-syncer-florin-2mg2dcem cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:52.471935     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=kyverno cluster=1yrbrny6lrx8pzoz"
    I0704 21:12:53.084350     380 main.go:134] "GVR=wgpolicyk8s.io/v1alpha2/policyreports: Notified action=update name=cpol-sample-cluster-policy namespace=local-path-storage cluster=1yrbrny6lrx8pzoz"
    ```