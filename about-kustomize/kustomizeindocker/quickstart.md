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

## Applying the Settings with Kubectl

* Settings get pulled down from a repo within our dockerfile, [k3d-demo](https://github.com/pwdelbloomboard/k3d-demo).
* Once this is cloned into our cluster via the ./composeexec.sh script, the settings are automatically updated with the docker-entrypoint.sh script.


## Applying Kustomization.yaml




## Visiting the Site

Once the pod has finished with its setup, you can visit the site at:

http://localhost:8081/