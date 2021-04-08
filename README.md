<img src="https://github.com/homedepot/cloud-runner/blob/media/cloud-runner.png" width="175" align="left">

# cloud-runner

Cloud Runner is a simple microservice that builds and runs a `gcloud run deploy` command against a given project. For flag support please [visit the wiki](https://github.com/homedepot/cloud-runner/wiki).

### Development

#### Build
```bash
make build
```

#### Required Environment Variables
| Name | Description |
|-|-:|
| `API_KEY` | Validated for Create/Delete operations |
| `SQL_HOST` | SQL host |
| `SQL_NAME` | SQL database name |
| `SQL_PASS` | SQL password |
| `SQL_USER` | SQL username |
