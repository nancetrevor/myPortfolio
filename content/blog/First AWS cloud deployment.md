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

Just to be able to see the pros and cons of both, I tried deploying my app using each service listed above.

