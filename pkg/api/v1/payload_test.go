package v1_test

const payloadBadRequest = `{
            "error": "invalid character 'd' looking for beginning of value"
          }`

const payloadRequestCredentials = `{
						"projectID": "test-project-id",
						"readGroups": [
							"gg_test"
						],
						"writeGroups": [
							"gg_test"
						]
          }`

const payloadConflictRequest = `{
            "error": "credentials already exists"
          }`

const payloadErrorCreatingCredentials = `{
            "error": "error creating credentials"
          }`

const payloadErrorCreatingReadPermission = `{
            "error": "error creating read permission"
          }`

const payloadErrorCreatingWritePermission = `{
            "error": "error creating write permission"
          }`

const payloadErrorGettingToken = `{
            "error": "error getting token"
          }`

const payloadErrorGettingLCP = `{
            "error": "error getting LCP"
          }`

const payloadErrorGettingRoles = `{
						"error": "error getting roles"
					}`

const payloadUnauthorized = `{
            "error": "X-Spinnaker-User header not set"
          }`

const payloadCredentialsCreated = `{
            "account": "cr-test-project-id-pr",
            "lifecycle": "pr",
            "projectID": "test-project-id",
            "readGroups": [
              "gg_test"
            ],
            "writeGroups": [
              "gg_test"
            ]
          }`

const payloadCredentialsNotFound = `{
						"error": "credentials not found"
					}`

const payloadCredentialsGetGenericError = `{
						"error": "error getting credentials"
					}`

const payloadCredentialsUnauthorized = `{
              "error": "requested credentials only for user, but X-Spinnaker-User header was not set"
            }`

const payloadCredentialsDeleteGenericError = `{
						"error": "error deleting credentials"
					}`

const payloadCredentialsDeleteForbiddentError = `{
            "error": "user test-account does not have access to delete account test-name"
          }`

const payloadCredentialsListEmpty = `{
            "credentials": []
          }`

const payloadCredentialsListFiltered = `{
              "credentials": [
                {
                  "account": "cr-test-project-id2-pr",
                  "lifecycle": "pr",
                  "projectID": "test-project-id2",
                  "readGroups": [
                    "gg_test2"
                  ],
                  "writeGroups": [
                    "gg_test2"
                  ]
                }
              ]
            }`

const payloadCredentialsList = `{
            "credentials": [
              {
                "account": "cr-test-project-id-pr",
                "lifecycle": "pr",
                "projectID": "test-project-id",
                "readGroups": [
                  "gg_test"
                ],
                "writeGroups": [
                  "gg_test"
                ]
              },
              {
                "account": "cr-test-project-id2-pr",
                "lifecycle": "pr",
                "projectID": "test-project-id2",
                "readGroups": [
                  "gg_test2"
                ],
                "writeGroups": [
                  "gg_test2"
                ]
              }
            ]
          }`
