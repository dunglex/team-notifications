 curl -X POST -H 'content-type: application-json' -L http://localhost:3978/pr/created -d '{
  "subscriptionId": "70ca59a9-32ab-4b9f-87de-2e46732e17c3",
  "notificationId": 16,
  "id": "0982ed73-d39a-4e22-ad4b-4a8d72c61f8a",
  "eventType": "git.pullrequest.created",
  "publisherId": "tfs",
  "message": null,
  "detailedMessage": null,
  "resource": {
    "repository": {
      "id": "f1ff22bf-6527-4b21-9659-b20a88f9a3b6",
      "name": "cma-backend",
      "url": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6",
      "project": {
        "id": "ba06509b-dd73-433e-98de-2fb7e5b94006",
        "name": "hr-digitalization",
        "url": "https://dev.azure.com/hc-vn/_apis/projects/ba06509b-dd73-433e-98de-2fb7e5b94006",
        "state": "wellFormed",
        "revision": 1521,
        "visibility": "private",
        "lastUpdateTime": "2024-08-02T03:15:13.33Z"
      },
      "size": 321467,
      "remoteUrl": "https://hc-vn@dev.azure.com/hc-vn/hr-digitalization/_git/cma-backend",
      "sshUrl": "git@ssh.dev.azure.com:v3/hc-vn/hr-digitalization/cma-backend",
      "webUrl": "https://dev.azure.com/hc-vn/hr-digitalization/_git/cma-backend",
      "isDisabled": false,
      "isInMaintenance": false
    },
    "pullRequestId": 35125,
    "codeReviewId": 35125,
    "status": "active",
    "createdBy": {
      "displayName": "Dung Le Xuan",
      "url": "https://spsprodsea1.vssps.visualstudio.com/A3fbf69ec-3f40-4ff1-9d23-11fb77c6fa04/_apis/Identities/90835b2c-1f72-61ae-96e7-3dbb608f68c1",
      "_links": {
        "avatar": {
          "href": "https://dev.azure.com/hc-vn/_apis/GraphProfile/MemberAvatars/aad.OTA4MzViMmMtMWY3Mi03MWFlLTk2ZTctM2RiYjYwOGY2OGMx"
        }
      },
      "id": "90835b2c-1f72-61ae-96e7-3dbb608f68c1",
      "uniqueName": "Dung.LeX@homecredit.vn",
      "imageUrl": "https://dev.azure.com/hc-vn/_api/_common/identityImage?id=90835b2c-1f72-61ae-96e7-3dbb608f68c1",
      "descriptor": "aad.OTA4MzViMmMtMWY3Mi03MWFlLTk2ZTctM2RiYjYwOGY2OGMx"
    },
    "creationDate": "2024-10-17T03:59:24.4363543Z",
    "title": "HRDIGI-87 Add support for API accounts authentication",
    "description": "Pull request created in: https://sd.homecredit.vn/browse/HRDIGI-87",
    "sourceRefName": "refs/heads/HRDIGI-87-add-support-for-api-accounts-authentication",
    "targetRefName": "refs/heads/dev",
    "mergeStatus": "succeeded",
    "isDraft": false,
    "mergeId": "7247627b-ffb4-4401-82a8-3872a9128d71",
    "lastMergeSourceCommit": {
      "commitId": "3c7dfad504702cc8b3783febae01c89b1ae5bf29",
      "url": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/commits/3c7dfad504702cc8b3783febae01c89b1ae5bf29"
    },
    "lastMergeTargetCommit": {
      "commitId": "3c7dfad504702cc8b3783febae01c89b1ae5bf29",
      "url": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/commits/3c7dfad504702cc8b3783febae01c89b1ae5bf29"
    },
    "lastMergeCommit": {
      "commitId": "3c7dfad504702cc8b3783febae01c89b1ae5bf29",
      "author": {
        "name": "Dung Le Xuan",
        "email": "Dung.LeX@homecredit.vn",
        "date": "2024-10-16T09:52:04Z"
      },
      "committer": {
        "name": "Dung Le Xuan",
        "email": "Dung.LeX@homecredit.vn",
        "date": "2024-10-16T09:52:04Z"
      },
      "comment": "Merged PR 35105: Update azure-pipelines-dev.yml for Azure Pipelines",
      "url": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/commits/3c7dfad504702cc8b3783febae01c89b1ae5bf29"
    },
    "reviewers": [
      {
        "reviewerUrl": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/pullRequests/35125/reviewers/eb6473e3-326e-44b0-8ebb-621249196fe5",
        "vote": 0,
        "hasDeclined": false,
        "isFlagged": false,
        "displayName": "[hr-digitalization]\\hr-digitalization Team",
        "url": "https://spsprodsea1.vssps.visualstudio.com/A3fbf69ec-3f40-4ff1-9d23-11fb77c6fa04/_apis/Identities/eb6473e3-326e-44b0-8ebb-621249196fe5",
        "_links": {
          "avatar": {
            "href": "https://dev.azure.com/hc-vn/_apis/GraphProfile/MemberAvatars/vssgp.Uy0xLTktMTU1MTM3NDI0NS0yNjA1NzEzMDgyLTE5NDM4NzkyMzUtMjU2NDY5ODAzOS0zODU0MTIzMDE0LTEtMjQyNjM4NTcyLTM3OTU3ODA5MzEtMjUxODM2NTcwMy0xMTM1NjMwMTYw"
          }
        },
        "id": "eb6473e3-326e-44b0-8ebb-621249196fe5",
        "uniqueName": "vstfs:///Classification/TeamProject/ba06509b-dd73-433e-98de-2fb7e5b94006\\hr-digitalization Team",
        "imageUrl": "https://dev.azure.com/hc-vn/_api/_common/identityImage?id=eb6473e3-326e-44b0-8ebb-621249196fe5",
        "isContainer": true
      }
    ],
    "url": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/pullRequests/35125",
    "_links": {
      "web": {
        "href": "https://dev.azure.com/hc-vn/hr-digitalization/_git/cma-backend/pullrequest/35125"
      },
      "statuses": {
        "href": "https://dev.azure.com/hc-vn/ba06509b-dd73-433e-98de-2fb7e5b94006/_apis/git/repositories/f1ff22bf-6527-4b21-9659-b20a88f9a3b6/pullRequests/35125/statuses"
      }
    },
    "supportsIterations": true,
    "artifactId": "vstfs:///Git/PullRequestId/ba06509b-dd73-433e-98de-2fb7e5b94006%2ff1ff22bf-6527-4b21-9659-b20a88f9a3b6%2f35125"
  },
  "resourceVersion": "1.0",
  "resourceContainers": {
    "collection": {
      "id": "9d911673-5925-4ae3-8d80-007352f9f67d",
      "baseUrl": "https://dev.azure.com/hc-vn/"
    },
    "account": {
      "id": "3fbf69ec-3f40-4ff1-9d23-11fb77c6fa04",
      "baseUrl": "https://dev.azure.com/hc-vn/"
    },
    "project": {
      "id": "ba06509b-dd73-433e-98de-2fb7e5b94006",
      "baseUrl": "https://dev.azure.com/hc-vn/"
    }
  },
  "createdDate": "2024-10-17T03:59:30.8449315Z"
}'