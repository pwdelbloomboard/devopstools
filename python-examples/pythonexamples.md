## Python Production Tools

### kubernetes

kubernetes is a prerequisite for understanding this example.

* [python-kubernetes official documentation](https://kubernetes.readthedocs.io/en/latest/)

#### General Notes




### logging

* [logging](https://docs.python.org/3/library/logging.html) is a tool that allows a python interpreter to spit out messages into the log.
* All python modules use logging, this is a base library, so you can create your own custom logging messages or utilize logs from other libraries.

Logging has different layers which it allows, 

#### Logging Example

Here is one way that logging is potentially used:

```
logging.info(f'status.succeeded: {job_status.succeeded}. status.failed: {job_status.failed}. status.conditions:\n')
```

* The above simply creates an [f-string literal](https://www.python.org/dev/peps/pep-0498/) (within f'') with the possible messages, given the conditions:

* status.succeeded: {job_status.succeeded}.
* status.failed: {job_status.failed}.
* status.conditions:\n

From the above we infer the following:

* the, "status" class may have several attributes for succeeded, failed and conditions, which appears to indicate other conditions.
* elsewhere in the file, job_status is established through a function: job_status = get_job_status(namespace, job_name, debug=debug)
* so when the above, "info" happens.


Below is an alternate proposed addition:

```
logging.info(f'object.status.phase: {event["object"].status.phase}')
logging.info(f'status.succeeded: {job_status.succeeded}. status.failed: {job_status.failed}. status.conditions:')
```

Which does the following:

* f-string literall within f''
* object.status.phase: {event["object"].status.phase}

the 