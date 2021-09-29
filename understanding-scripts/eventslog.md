# eventslog.txt

* eventslog.txt is created by running "kubectl get events > events.txt"

* Information about what "get events" is can be found at the [kubernetes cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

"get events" is essentially a way to find resources relevant to events and display them.

### Events Reasons

#### Normal

* ScalingReplicaSet
* Scheduled
* Sync
* Pulled
* Created
* Started
* Killing
* RELOAD
* NodeNotReady

#### Warnings

* FailedScheduling
* Unhealthy 
* NodeNotReady
* FailedPreStopHook
* FailedMount
* SystemOOM
* InvalidDiskCapacity

##### InvalidDiskCapacity

* Set your Docker VM space higher, to 100GB or more if possible.
* [Docker dockemtnation on VM Size](https://docs.docker.com/desktop/mac/space/)
* Go to Preferences (Gear Icon on Docker Desktop) >> Resources >> Advanced
