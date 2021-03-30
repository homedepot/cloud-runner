<img src="https://github.com/homedepot/cloud-runner/blob/media/cloud-runner.png" width="175" align="left">

# cloud-runner

Cloud Runner is a simple microservice that builds and runs a `gcloud run deploy` command against a given project. For flag support please [visit the wiki](https://github.com/homedepot/cloud-runner/wiki).

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
| `THD_IDENTITY_CLIENT_ID` | THD IDP client ID (e.g. `spiffe://homedepot.dev/om-api-security-client`)  |
| `THD_IDENTITY_CLIENT_SECRET` | THD IDP client secret |
| `THD_IDENTITY_RESOURCE` | THD IDP resource (e.g. `spiffe://homedepot.dev/om-api-security-api`) |
| `THD_IDENTITY_TOKEN_ENDPOINT` | THD IDP token endpoint (e.g. `https://identity-qa.homedepot.com/as/token.oauth2`) |
