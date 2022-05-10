# Kustomize in Docker Quickstart

* Docker must be installed on your local machine.
* k3d must be installed on your local machine.

## Launching the Container

To run, within kustomizeindocker/ do:

```
chmod +x ./composeexec.sh
```

Then do:

```
./composeexec.sh
```

## Applying the Settings

* Settings get pulled down from a repo within our dockerfile, [k3d-demo](https://github.com/pwdelbloomboard/k3d-demo).
* Once this is cloned into our cluster via the ./composeexec.sh script, then we can manually apply the settings to launch the app with:

```
kubectl apply --filename k8s/
```


## Visiting the Site


http://localhost:8081/