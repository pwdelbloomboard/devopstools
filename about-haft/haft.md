# High Availability Fault Tolerance

## Linkedin Article from 2020

* [High Availability vs. Fault Tolerance](https://www.linkedin.com/pulse/high-availability-vs-fault-tolerance-jon-bonso)

Summary:

* The author posits that most professionals don't even really know what HAFT systems are, even if they have been working in I.T. for a long time.
* High Availability differs from Fault Tolerance.
* High Availability has to do with creating a situation where you have a miniscule percentage of downtime through weak redundancy, minimal cost and failovers, minimizing recovery time. The, "availability," in high availability has to do with the system being able to make itself available through recovery, perhaps another term for high availability could be, "fast recovery."
* Fault Tolerant systems have to do with having absolutely 100% uptime, with as large as necessary redundancy, and having higher costs, achieved through synchronization and failover. The, "tolerance" in fault tolerance implies that the system can tolerate any component failure to avoid performance impact. In other words, there is no emphasis on recovery, but total emphasis on ever having a problem.
* High Availability may be considered fire fighting, while Fault Tolerance may be considered fire prevention.
* Systems can be either or, or both, and Fault Tolerance has a higher cost.

## Slack Article from 2022


* [Slack’s Incident on 2-22-22 - Trouble with Datacenters](https://slack.engineering/slacks-incident-on-2-22-22/)

Summary:

* The scenario was a complex systems failure which involved a cascading failure.
* User tickets started coming in pointing to problems connecting to Slack at 6AM PST, the start of the work day for the USA.
* Vitess, a horizontal scaling system for MySQL.
* Keyspaces, logical databases.
* Group Direct Messaging (GDM) conversations by user.
* Memcached is the caching tier
* Mcrouter scales cache fleet horizontally
* Mcrib is the control plane, and generates up-to-date Mcrouter configurations.
* Consul, service discovery system.

* Predictions for cascading failures, bottlenecks and dependency on a warm cache seem obvious in retrospect, but it is easier to make broad brush statements about the sorts of incidents that may occur to pinpoint every specific sequence of events or problem that leads to a cascading failure.
* Complex systems contain changing mixutres of failures latent within them. The complexity of these systems make it impossible for them to run without multiple flaws present. Eradication of all failures is limited by economic costs because it is difficult before the fact to see how such failures might contribute to an incident.


## Uber Michealangelo

* [Meet Michelangelo: Uber’s Machine Learning Platform](https://www.uber.com/en-CA/blog/michelangelo-machine-learning-platform/)

## Blog Post from Jayendra Patil on AWS HAFT

* [Blog Post](https://jayendrapatil.com/aws-high-availability-fault-tolerance-architecture-certification/)

* Further, [AWS Certified Solutions Architect Exam Guide](https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Exam-Guide.pdf)

* https://aws.amazon.com/certification/certified-solutions-architect-associate/

# Job Requirements

* Types of job requirements listed for Machine Learning Engineers:

* Recommendation systems and deep learning
* Applying machine learning methods in settings like recommender systems, search, user modeling, image recognition, graph
* Representation learning, natural language processing.
* Experience building pipelines for large scale data processing
* Comprehensive knowledge of ML concepts, including classification, regression, and unsupervised learning.
* Hands-on experience with traditional machine learning as well as deep learning algorithms.
* Proficiency with Python programming, including numeric, data analysis, and ML libraries such as NumPy, pandas, Scikit-learn, and/or PyTorch.
* Solid understanding of statistical and mathematical methods used in ML, evaluation, and hypothesis testing.
* Proficiency with deploying models via a platform like MLFLow or AWS SageMaker.
* Proficiency with SQL for data analysis.
* Direct experience with feature engineering using structured or unstructured data.
* Large scale data processing technologies (e.g. TensorFlow, SciKit learn, Dataflow, Hadoop, Scalding, Spark, Storm)
* Experience in using ML techniques to optimize customer facing product features
* Strong mathematical background in statistics and machine learning
* Agile software processes, data-driven development, reliability, and responsible experimentation
* Machine learning publications or work on open source