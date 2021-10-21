# AWS SSM

* [AWS SSM Docs](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html)


* SSM is the location where environmental variable names are stored. Since virtual machines are ephemeral and envioronmental variables are meant to work across the entire production or staging deployment, SSM is used rather than placing different sets of environmental variables spread across different compute units such as EC2.

![](/img/parameternotfound_.png)

The SSM Dashboard can be found within AWS as shown:

![](/img/awsssm-dashboard.png)