# About KServe

## Running an InferenceService

https://kserve.github.io/website/0.9/get_started/first_isvc/#2-create-an-inferenceservice


### Ensuring InferenceService is Running

* Once an InferenceService is deployed, the status of the InferneceService can be tested with:

```
k get InferenceService <name> -n <namespace>
```

* If you're already on the proper namespace profile, you can just use ```k get InferenceService <name>```

* After determining that said InferenceService works, you can find the ingress IP and ports with:

```
k get svc istio-ingressgateway -n istio-system
```

* Which may show, if a load balancer is not set up:

```
$ k get svc istio-ingressgateway -n istio-system
NAME                   TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)
istio-ingressgateway   NodePort   XXX.XX.XXXX.X   <none>
```

* So assuming we're using a NodePort, if the EXTERNAL-IP is none (or perpetually pending), the environment does not provide an external load balancer for the ingress gateway. In this case, you can access the gateway using the service's node port.

* Note - ClusterIP is the default Kubernetes service. Your service will be exposed on a ClusterIP unless you manually define another type.
* NodePort publically exposes a service on a fixed port number. You’ll need to use the cluster’s IP address and the NodePort number—e.g. 123.123.123.123:30000.
* Assuming, "other environments," besides GKE or Minikube, we need to run one of the following commands to get and set the INGRESS_HOST.

```
# Other environment(On Prem)
export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
```
### Performing an Inference

* We can put the following values into a json file in order to put in an inference request on sklearn-iris.

```
cat <<EOF > "./iris-input.json"
{
  "instances": [
    [6.8,  2.8,  4.8,  1.4],
    [6.0,  3.4,  4.5,  1.6]
  ]
}
EOF
```
* One could also create an empty data file with:

```
cat <<EOF > "./empty-input.json"
{"data": []}
EOF
```

* You can call your model using a the curl function, which inputs a json file to the actual endpoint that you have gathered, e.g., using ```curl -d``` which is for data.:

```
curl -v http://sklearn-iris.kserve-test/v1/models/sklearn-iris:predict -d @./iris-input.json
```

* The above domain can be found within the Kubeflow Control Panel under, "models," for the model in question as an, "internal URL."
* In my case, I created a model which is on my namespace, "patrick," with the endpoint shown and fed in empty data:

```
curl -v http://patrick-function-test-via-potato.patrick.svc.cluster.local/v1/models/patrick-function-test-via-potato:predict  -d @./empty-input.json
```
* The result was:

```
* Could not resolve host: patrick-function-test-via-potato.patrick.svc.cluster.local
* Closing connection 0
curl: (6) Could not resolve host: patrick-function-test-via-potato.patrick.svc.cluster.local
```
* This was 



## Deploying a Custom Transformer

https://kserve.github.io/website/modelserving/v1beta1/transformer/torchserve_image_transformer/

