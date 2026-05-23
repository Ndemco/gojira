# GET /rest/api/3/issue/createmeta
**operationId:** `getCreateIssueMeta`
**Summary:** Get create issue metadata

Returns details of projects, issue types within projects, and, when requested, the create screen fields for each issue type for the user. Use the information to populate the requests in [ Create issue](#api-rest-api-3-issue-post) and [Create issues](#api-rest-api-3-issue-bulk-post).

Deprecated, see [Create Issue Meta Endpoint Deprecation Notice](https://developer.atlassian.com/cloud/jira/platform/changelog/#CHANGE-1304).

The request can be restricted to specific projects or issue types using t

## Parameters
- `projectIds` [query] []string — List of project IDs. This parameter accepts a comma-separated list. Multiple project IDs can also be provided using an a
- `projectKeys` [query] []string — List of project keys. This parameter accepts a comma-separated list. Multiple project keys can also be provided using an
- `issuetypeIds` [query] []string — List of issue type IDs. This parameter accepts a comma-separated list. Multiple issue type IDs can also be provided usin
- `issuetypeNames` [query] []string — List of issue type names. This parameter accepts a comma-separated list. Multiple issue type names can also be provided 
- `expand` [query] string — Use [expand](#expansion) to include additional information about issue metadata in the response. This parameter accepts 

## Responses
- 200: object:
  - `expand`: string
  - `projects`: []ProjectIssueCreateMetadata
- 401: Returned if the authentication credentials are incorrect or missing.
