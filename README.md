# sary-task

url: a74dfa04af530480d91ba174e610afe9-2139051881.eu-central-1.elb.amazonaws.com:8080

### Infra
Folder is for infrastructure, using pulumi as the Iac.

### kube
Folder kubernetes deployments, manifests and helm charts.

### fake-stackoverflow
Folder where the Golang implementation of the fake stackoverflow.

### .github/workflows
For CD github actions is used. still missing some features.

## Requirments

The project is a very simple Q/A forum with the following user stories:
```bash
curl -X POST -H "Content-Type: application/json"     -d '{"name": "test1"}'    a74dfa04af530480d91ba174e610afe9-2139051881.eu-central-1.elb.amazonaws.com:8080/user
```
- As a user, I should be able to submit questions.
```bash
curl -X POST -H "Content-Type: application/json"     -d '{"userrefer": 1, "tital":"why is tomato?", "body": "red, so much!?"}'    a74dfa04af530480d91ba174e610afe9-2139051881.eu-central-1.elb.amazonaws.com:8080/question
```
- As a user, I should be able to answer submitted questions.
```bash
curl -X POST -H "Content-Type: application/json"     -d '{"userrefer": 1, "questionrefer": 2, "body": "sorry tomato is depreciated, use orange"}'    a74dfa04af530480d91ba174e610afe9-2139051881.eu-central-1.elb.amazonaws.com:8080/question/answer
```
- As a user, I should be able to assign multiple tags to questions.
```bash
curl -X POST -H "Content-Type: application/json"     -d '{"userrefer": 1, "tital":"why", "body": "why potato not feature", "tags":[{"ID":2, "name": "veg"}] }'    a74dfa04af530480d91ba174e610afe9-2139051881.eu-central-1.elb.amazonaws.com:8080/question
```
