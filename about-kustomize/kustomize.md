## About Kustomize

* With Kubernetes, you have tons of yaml files and are applying them.
* However, if you're using nothing but yaml files, you are copying and pasting a lot of the same thing.
* At the same time, different environments need different levels of customizations.
* Similarly, you might have a different number of replicas in staging vs. production and so on and sofourth, the list goes on and on!
* You will almost certainly have variations across different environments, and you have to apply different changes against whatever third-party settings, while at the same time copy and pasting what are essentially, templates.

Helm solved this problem using, "templates."  However, this may not always be the best choice.

Essentially, it allows us to get the same results as what helm provides, but without templates, purely with yaml files.

It allows you to modify yaml files to accomplish all of the above, but in an easier way.

## Prerequisites for Kustomize

* Need a Kubernetes Cluster.
* Need an Nginx Ingress.
* Need the Kustomize CLI.

 


## Resources

[Kustomize - Hwo to Simplify Kubernetes Configuration Management](https://www.youtube.com/watch?v=Twtbg6LFnAg)