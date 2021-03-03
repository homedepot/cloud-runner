<img src="https://github.homedepot.com/cd/cloud-runner/blob/media/cloud-runner.png" width="175" align="left">

# cloud-runner

Cloud Runner is a simple microservice that builds and runs a `gcloud run deploy` command against a given project. For flag support please [visit the wiki](https://github.homedepot.com/cd/cloud-runner/wiki).

[![Jenkins](https://cd-shields.apps-np.homedepot.com/jenkins/build?jobUrl=https%3A%2F%2Fcd-jenkins.apps-np.homedepot.com%2Fjob%2Fcd%2Fjob%2Fcloud-runner%2Fjob%2Fmaster%2F)](http://cd-jenkins.apps-np.homedepot.com/job/cd/job/cloud-runner/job/master/)
[![Sonar Coverage](https://cd-shields.apps-np.homedepot.com/sonar/coverage/cloud-runner?server=https%3A%2F%2Fsonar.homedepot.com&sonarVersion=7.8)](https://sonar.homedepot.com/component_measures?id=cloud-runner&metric=coverage)
![GitHub go.mod Go version](https://cd-shields.apps-np.homedepot.com/github/go-mod/go-version/cd/cloud-runner)

### Development

#### Build
```bash
go build -o cloud-runner cmd/cloud-runner/main.go
```
