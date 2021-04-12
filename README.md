<img src="https://github.com/homedepot/cloud-runner/blob/media/cloud-runner.png" width="175" align="left">

# cloud-runner

Cloud Runner is a simple microservice that builds and runs a `gcloud run deploy` command against a given GCP project ID. For flag support please [visit the wiki](https://github.com/homedepot/cloud-runner/wiki).

### Development

#### Environment Variables
| Name | Description | Required | Notes
|-|-:|-:|-:|
| `API_KEY` | Validated for Create/Delete operations | ✔️  | |
| `SQL_HOST` | SQL host | | If not set will default to local sqlite DB |
| `SQL_NAME` | SQL database name | | If not set will default to local sqlite DB |
| `SQL_PASS` | SQL password | | If not set will default to local sqlite DB |
| `SQL_USER` | SQL username | | If not set will default to local sqlite DB |

#### Build
```bash
make build
```

#### Test
```bash
make test
```

#### Run Locally
The following will show you how to run Cloud Runner locally and onboard your first account! Running locally will by default create a SQLite DB named `cloud-runner.db` in the current directory.

1. Run cloud-runner
```bash
$ export API_KEY=test
$ make run
```
2. Create an account. Cloud Runner connects to Spinnaker's fiat (at http://spin-fiat.spinnaker:7003) when deploying to GCP to verify the current user has read and write access to the account. The user is defined in the `X-Spinnaker-User` request header. When onboarding an account into Cloud Runner make sure to define the read and write groups correctly! If no `account` field is provided one will be generated in the format `cr-<GCP_PROJECT_ID>`.
```bash
$ curl -H "API-Key: test" localhost:80/v1/credentials -d '{
  "account": "test-account-name",
  "projectID": "test-project-id",
  "readGroups": [
    "test-group"
  ],
  "writeGroups": [
    "test-group"
  ]
}' | jq
```
You should see the response
```json
{
  "account": "test-account-name",
  "projectID": "test-project-id",
  "readGroups": [
    "test-group"
  ],
  "writeGroups": [
    "test-group"
  ]
}
```
3. List credentials
```
$ curl localhost:80/v1/credentials | jq
```
You should see the response
```json
{
  "credentials": [
    {
      "account": "test-account-name",
      "projectID": "test-project-id",
      "readGroups": [
        "test-group"
      ],
      "writeGroups": [
        "test-group"
      ]
    }
  ]
}
```
To generate CURL commands to create and monitor a deployment, reference the swagger YAML at `api/swagger.yaml`.
