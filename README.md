# Run minimal example

## Install the operator

Create a secret in your cluster to allow the controller to perform changes on spacelift backend.
It's simpler to configure a token the dedicated space on spacelift preprod because everything is configured.
The secret should be created in the same namespace as Stack and Run resources that we are going to create afterward.

```shell
kubectl create secret generic spacelift-credentials\
  --from-literal=SPACELIFT_API_KEY_ENDPOINT='https://CHANGEME.app.spacelift.io'\
  --from-literal=SPACELIFT_API_KEY_ID='CHANGEME'\
  --from-literal=SPACELIFT_API_KEY_SECRET='CHANGEME'
```

Then install the operator in the cluster with

```shell
kubectl apply -f controller
```

## Create a stack

```shell
kubectl apply -f examples/stack.yaml
```

See if the stack has been created on spacelift, if not you can inspect the controller logs to see what happened.

```shell
kubectl logs -f -n spacelift-operator-system deploy/spacelift-operator-controller-manager
```

## Trigger a run

```shell
kubectl create -f examples/run.yaml
```

# Setup the kubecon demo

## (optional) Configure org

You can skip this step (recommended) if you want to simply run the demo against the preconfigured preprod env.

- Create a space to run this demo
- Make sure the VCS integration is configured so this repo could be reached
- Create a push policy to ignore all VCS events with label `autoattach:argo`

```rego
package spacelift

ignore {
    true
}
```

- Create an api key in your org, and allow access to the space

```rego
package spacelift

key := "api::CHANGEME"
space_id = "spacelift-operator-CHANGEME"

space_admin[space_id] {
	  input.session.login == key
}
allow {input.session.login == key}
write {input.session.login == key}
sample { true }
```

## Install the operator

You can jump directly to this step and ask for a valid token.

Create a secret in your cluster to allow the controller to perform changes on spacelift backend.
It's simpler to configure a token the dedicated space on spacelift preprod because everything is configured.
The secret should be created in the same namespace as Stack and Run resources that we are going to create afterward.

```shell
kubectl create secret generic spacelift-credentials\
  --from-literal=SPACELIFT_API_KEY_ENDPOINT='https://spacelift-io.app.spacelift.dev'\
  --from-literal=SPACELIFT_API_KEY_ID='CHANGEME'\
  --from-literal=SPACELIFT_API_KEY_SECRET='CHANGEME'
```

Install the operator with the following command

```shell
kubectl apply -f controller
```

# Deployment

## Helm

```shell
# Create a stack
kubectl apply -f infra/spacelift/stack.yaml &&\
kubectl wait --for=jsonpath='{.status.ready}'=true stack/demo-stack --timeout 1h

# Trigger a run
kubectl delete --ignore-not-found=true -f infra/spacelift/run.yaml &&\
kubectl apply -f infra/spacelift/run.yaml &&\
kubectl wait --for=jsonpath='{.status.ready}'=true run/spacelift-operator-demo --timeout 1h

# Deploy the app
helm upgrade --install operator-demo ./infra/helm/ --set 'image.tag=ed35e9a152f61eb79d369cd9b2dd7e4f65629026'
```

Now you can connect to the app and see that the bucket URL has been injected as env var

```shell
kubectl port-forward service/operator-demo 8888
```

And then open http://localhost:8888

### Cleanup

In order to clean up your install, run:

```shell
helm uninstall operator-demo
kubectl delete --ignore-not-found=true -f infra/spacelift/run.yaml
kubectl delete --ignore-not-found=true -f infra/spacelift/stack.yaml
```

Please also remember to clean up stack on the spacelift side, since we do not perform stack deletion.

## Argo

TODO
