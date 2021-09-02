# ServianTechChallenge App Solution

- To implement this solution I leveraged infrastructure as code using Cloudformation templates.
- To implement the Countinuous integration (CI) and Continuous Deployement (CD) I have used Git integrated AWS Codepipeline (Done Manually).

### Architecture Overview
![architecture diagram](/solution/Architecture.png?raw=true)


### Pre-requisites
- AWS Account for deploying this solution.
- AWS Key-Pair for deploying EC2 instances.
- SSH setup on local machine for Testing/debugging.
- Codepipeline IAM role to access EC2,Secrets Manager,S3,CodeDeploy,CodeBuild,Cloudformation,RDS.
- Secrets Created inside AWS Secrets manager for the Postgres Database.

### AWS Services Used
- **AWS Virtual Private Cloud (VPC)** for hosting the application.
- **AWS Relational Database Service(RDS)** to host as a backend database for the ServianTechChallenge WebApp.
- **AWS Elastic Compute Cloud (EC2)** to host the ServianTechChallenge WebApp
- **AWS Clouformation** to automate & deploy AWS Resources using code/template
- **AWS Codepipeline ** to create and automate the CICD process in AWS environment.
- **AWS CodeDeploy ** to deploy applications in automated fashion.


### Steps to deploy the solution.
- Create a Github account and generate an access key token.
- Create a AWS Codepipeline using the service role created already > Choose Source provider as Github and provide Github Repo name and Branch > Skip Build Phase > In Deploy Stage Select provider as Cloudformation > Select Action Mode (Create/Update Stack) > Provide path to cloudformation template in the Repo(~solution/infra.yaml).
- Follow same steps as above except During selection of Action mode select (Delete Stack).
- The CodePipeline should look like below 
  ![CodePipeline](/solution/codepipeline.png?raw=true)
 
- Successful Creation of the pipeline will trigger the first pipeline execution > Monitor Execution details in Cloudformation.
- Upon Successful completion of the Pipeline, you should be able to see the cloudformation stack creation is complete and all resources are created.
- To deploy new changes make changes to the files in the git repo then commit and push to master branch. This action will trigger the CodePipeline Execution again and App will be updated and deployed within few minutes.


### Scope for improvement
- Automate Creation of Codepipeline using cloudformation template and avoid manually creation the pipeline.
- Better Conf.toml file management to retreive secrets from environment variables and avoid storing credentials on server.
- Use of Nested Cloudformation templates to modularise the Resources, this will help in the current deploy phase of Delete then Create stack substantially faster. 

### Application Demo 


[![Application Demo](/solution//mov.gif)]
