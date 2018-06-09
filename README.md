Operator for replicating [Secrets](https://kubernetes.io/docs/concepts/configuration/secret/) to every namespace in a Kubernetes cluster.

This can be useful, for example, for a shared TLS certificate or Docker registry
credential.

To deploy, deploy the manifests:

```
kubectl apply -f deploy/rbac.yaml
kubectl apply -f deploy/crd.yaml
kubectl apply -f deploy/operator.yaml
```

Create a `ClusterSecret` called `example.yaml`. A `ClusterSecret` has all of the
same fields as a `Secret`, but no namespace.

```
apiVersion: "clustersecret.codesink.net/v1alpha1"
kind: "ClusterSecret"
metadata:
  name: "my-certificate"
data:
  tls.crt: RGF0YQ==
  tls.key: RGF0YQ==
type: kubernetes.io/tls
```

Deploy it:

```
kubectl apply -f example.yaml
```

The secret `cluster-my-certificate` will now be available in every namespace.
