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

#### Required Environment Variables
| Name | Description |
|-|-:|
| `API_KEY` | Validated for Create/Delete operations |
| `SNOWQL_URL` | URL for SnowQL |
| `SQL_HOST` | SQL host |
| `SQL_NAME` | SQL database name |
| `SQL_PASS` | SQL password |
| `SQL_USER` | SQL username |
| `THD_IDENTITY_CLIENT_ID` | THD IDP client key (e.g. `spiffe://homedepot.dev/om-api-security-client`)  |
| `THD_IDENTITY_CLIENT_SECRET` | THD IDP client secret |
| `THD_IDENTITY_RESOURCE` | THD IDP resource (e.g. `spiffe://homedepot.dev/om-api-security-api`) |
| `THD_IDENTITY_TOKEN_ENDPOINT` | THD IDP token endpoint (e.g. `https://identity-qa.homedepot.com/as/token.oauth2`) |
