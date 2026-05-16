title: My first cloud deployment
date: 2026-05-13

Despite me having worked with AWS services for a while, I had never deployed a project
from scratch by myself. After studying and passing the AWS Solutions Architect Associate exam
I figured it was time that I put all of my knowledge together a deploy my own project from scratch.
I had previously created a website, [fightdb.org](https://fightdb.org) and have it hosted on a VPN
(Architecture explanation [here](https://trevornance.dev/architecture/)). Rather than migrating that over, I decided
to host my portfolio site within the AWS ecosystem.

## AWS Architecture

Within AWS, there are multiple ways to deploy a static website.
- [AWS Amplify](https://aws.amazon.com/amplify/)
    - This method is used for deploying web or mobile apps in a server-less ecosystem.
    - Little to no deployment skills required
- [S3 static hosting](https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html)
    - A "object" storage system that has the ability to host your static files
    - This method requires setting up everything yourself

Just to be able to see the pros and cons of both, I deployed my app on each one.

There technically are many other ways to host static sites on AWS. But are overkill (and much more expensive) compared to the former options:
- [EC2 Instance](https://aws.amazon.com/ec2/)
    - This is overkill for most static sites as it gives you the ability to access a server.
    - Is more for dynamic sites and full web apps.
- [Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/)
    - Even though you *technically* could host a static site, this would be drastic overkill.
    - Elastic Beanstalk is meant for hosting full web apps easily. Giving you an easy setup without having to manage each service individually.
- [Elastic Container Service](https://aws.amazon.com/ecs/)
    - This is essentially the AWS version of docker containers.
    - They provide a fully managed container service with little management needed
    - In the same vain as EC2 Instance, this would require a large amount of setup required for a simple static site.

AWS offers many different options for site hosting, as this feature is the most common need out of a cloud service. 
I will talk a little of the pros and cons of Amplify and S3 hosting, as these are the two I used for this project.
I may make a blog post later on regarding the other services. But for now I will not focus on those.

## S3 Static Hosting

S3 itself is just a object storage service that gives you the ablility to store files as "objects" as opposed in file systems.



