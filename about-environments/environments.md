# Environment Definition

* An environment is a collection of resources that you can target with deployments from a pipeline.
* Typical examples of environment names are Dev, Test, QA, Staging, and Production.
## Different Levels of Environments

![](/img/environment-stages.png)

* Having too many environment stages becomes too expensive and costly, the cost of maintaining and ensuring that the environments mimic or do what they are supposed to do becomes high.
* Having insufficient environment stages becomes risky, particularly not having a staging environment - you run the risk of pushing errors into production by not mimicking 1:1 between staging and production.
### Staging

> Staging environments are generally meant to be identical or nearly identical to production. This means that they have the same hardware, software, and configuration. The closer you can get to this, the more useful your staging environment will become.

> That level of sameness between staging and production ensures that testing on your staging environment reflects what would happen in production under like circumstances.

> Unlike development or limited integration test environments, staging environments utilize the same back-end and up-and-downstream services. They also have the same architecture, the same scale, and have highly similar or identical configurations to the production environment.

> Depending on any regulatory factors (such as GDPR requirements) and your organization’s level of ability to anonymize data, a staging environment may even have anonymized or complete sets of production data in order to more closely mimic the real world production environment. That means that a staging environment is typically not released or made available to your production user base, but rather it is made available to an internal or pilot user base.
## Compute Environments and Database Environments

### Compute Environments

Compute environments are basically areas that are hosting virtual machines which host applications, namely, they are hosting Docker containers or k3d/k8s clusters and nodes.

Kubernetes (k8s) nodes can have a master node, or in the case of k3d this is known as a server. It is possible to have more than one master node - the purpose of this would be to improve failover rates, in that essentially if one master goes down, an API that may be served out of the cluster is essentially dead and you cannot modify the cluster.

The cluster itself does not go down when the master node goes down, so the master node can be brought back up and the API can resume service as it had previously. However if you had multiple master nodes, the possibility of failure would be lower, because it would be less likely that both master nodes would go down at one time.

The disadvantage to having multiple master nodes is that it costs more, it is just more costly to have more master/servers which could otherwise be used as worker/agent nodes, doing work and running applications, or perhaps just not having an extra master/server node at all. All compute time has a cost.

### Database Environments

There are different uses of the word, "master" - one is the, "master nodes," within k8s clusters, and another is, "master database," or, "database masters."

Abstractly, within an architecture, you could have multiple database masters, which are defined as databases which define all of the system-level information, which includes instance-wide metadata such as login accounts, endpoints, servers and system configuration settings. The master database is also the database that records the existence of all other databases and the location of database files. The master database must be available for other databases to access it.

Basically, the master is the database that recieves the writes and is the official source of truth. Data is replicated to one or more replicas. Read-only operations can be sent to replica to offload traffic from a master.

System performance could be improved by having multiple database masters. Whereas with Compute Environments having multiple master nodes would improve failover rate but would not improve performance, having multiple database masters located in different geographies could reduce latency if a given architecture is running application servers geographically closer to users.

However, the first step in improving database performance might not be distribution of geography, it might be scaling up a cluster size to obtain more resources.

# Resources

* [Azure Devops Pipelines](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/environments?view=azure-devops)
* [Staging vs Production Environments](https://dev.to/heroku/staging-environments-are-overlooked-here-s-why-they-matter-3ghd)
* [Master Database Definition](https://docs.microsoft.com/en-us/sql/relational-databases/databases/master-database?view=sql-server-ver15#:~:text=The%20master%20database%20records%20all,for%20a%20SQL%20Server%20system.&text=Also%2C%20master%20is%20the%20database,initialization%20information%20for%20SQL%20Server.)