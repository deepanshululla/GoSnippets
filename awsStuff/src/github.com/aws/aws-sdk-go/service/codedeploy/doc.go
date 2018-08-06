// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codedeploy provides the client and types for making API
// requests to AWS CodeDeploy.
//
// AWS CodeDeploy is a deployment service that automates application deployments
// to Amazon EC2 instances, on-premises instances running in your own facility,
// or serverless AWS Lambda functions.
//
// You can deploy a nearly unlimited variety of application content, such as
// an updated Lambda function, code, web and configuration files, executables,
// packages, main, multimedia files, and so on. AWS CodeDeploy can deploy
// application content stored in Amazon S3 buckets, GitHub repositories, or
// Bitbucket repositories. You do not need to make changes to your existing
// code before you can use AWS CodeDeploy.
//
// AWS CodeDeploy makes it easier for you to rapidly release new features, helps
// you avoid downtime during application deployment, and handles the complexity
// of updating your applications, without many of the risks associated with
// error-prone manual deployments.
//
// AWS CodeDeploy Components
//
// Use the information in this guide to help you work with the following AWS
// CodeDeploy components:
//
//    * Application: A name that uniquely identifies the application you want
//    to deploy. AWS CodeDeploy uses this name, which functions as a container,
//    to ensure the correct combination of revision, deployment configuration,
//    and deployment group are referenced during a deployment.
//
//    * Deployment group: A set of individual instances or CodeDeploy Lambda
//    applications. A Lambda deployment group contains a group of applications.
//    An EC2/On-premises deployment group contains individually tagged instances,
//    Amazon EC2 instances in Auto Scaling groups, or both.
//
//    * Deployment configuration: A set of deployment rules and deployment success
//    and failure conditions used by AWS CodeDeploy during a deployment.
//
//    * Deployment: The process and the components used in the process of updating
//    a Lambda function or of installing content on one or more instances.
//
//    * Application revisions: For an AWS Lambda deployment, this is an AppSpec
//    file that specifies the Lambda function to update and one or more functions
//    to validate deployment lifecycle events. For an EC2/On-premises deployment,
//    this is an archive file containing source content—source code, web pages,
//    executable files, and deployment main—along with an AppSpec file. Revisions
//    are stored in Amazon S3 buckets or GitHub repositories. For Amazon S3,
//    a revision is uniquely identified by its Amazon S3 object key and its
//    ETag, version, or both. For GitHub, a revision is uniquely identified
//    by its commit ID.
//
// This guide also contains information to help you get details about the instances
// in your deployments, to make on-premises instances available for AWS CodeDeploy
// deployments, and to get details about a Lambda function deployment.
//
// AWS CodeDeploy Information Resources
//
//    * AWS CodeDeploy User Guide (http://docs.aws.amazon.com/codedeploy/latest/userguide)
//
//    * AWS CodeDeploy API Reference Guide (http://docs.aws.amazon.com/codedeploy/latest/APIReference/)
//
//    * AWS CLI Reference for AWS CodeDeploy (http://docs.aws.amazon.com/cli/latest/reference/deploy/index.html)
//
//    * AWS CodeDeploy Developer Forum (https://forums.aws.amazon.com/forum.jspa?forumID=179)
//
// See https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06 for more information on this service.
//
// See codedeploy package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/
//
// Using the Client
//
// To contact AWS CodeDeploy with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeDeploy client CodeDeploy for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/#New
package codedeploy
