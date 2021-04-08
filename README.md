<img src="https://github.com/homedepot/cloud-runner/blob/media/cloud-runner.png" width="175" align="left">

# cloud-runner

Cloud Runner is a simple microservice that builds and runs a `gcloud run deploy` command against a given GCP project ID. For flag support please [visit the wiki](https://github.com/homedepot/cloud-runner/wiki).

### Development

#### Build
```bash
make build
```

#### Environment Variables
| Name | Description | Required | Notes
|-|-:|-:|-:|
| `API_KEY` | Validated for Create/Delete operations | ✔️  | |
| `SQL_HOST` | SQL host | | If not set will default to local sqlite DB |
| `SQL_NAME` | SQL database name | | If not set will default to local sqlite DB |
| `SQL_PASS` | SQL password | | If not set will default to local sqlite DB |
| `SQL_USER` | SQL username | | If not set will default to local sqlite DB |

#### Run Locally
The following will show you how to run `cloud-runner` locally and onboard your first account!

1. Run cloud-runner
```bash
$ export API_KEY=test
$ make run
```
2. Create an account. Cloud Runner connects to Spinnaker's fiat (at http://spin-fiat.spinnaker:7003) when deploying to verify the current user has read and write access to the account. When onboarding an account into Spinnaker make sure to define the read and write groups correctly!
```bash
$ curl -H "API-Key: test" localhost:80/v1/credentials -d '{
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
  "account": "cr-test-project-id",
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
      "account": "cr-test-project-id",
      "projectID": "test-project-id",
      "readGroups": [
        "test-read-group"
      ],
      "writeGroups": [
        "test-group"
      ]
    }
  ]
}
```
To generate CURL commands to create and monitor a deployment, reference the swagger YAML at `api/swagger.yaml`.
