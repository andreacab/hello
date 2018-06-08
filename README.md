# Go example app for CI-CD with jenkins
## Objectives:
1. [x] trigger pipeline when changes are pushed onto bitbucket hello master branch
2. [x] Add test stage before building go app
3. [x] Cross compile/build the go app in a docker container (using Dockerfile-build or standard go image)
4. [ ] Build new docker image on every build with new binary (using Dockerfile)
5. [ ] Test image works and is running properly
6. [ ] Publish new image to docker hub or amazon ECR
7. [ ] Deploy new image to AWS infrastructure