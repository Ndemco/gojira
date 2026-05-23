# DELETE /rest/api/3/issuetypescheme/{issueTypeSchemeId}
**operationId:** `deleteIssueTypeScheme`
**Summary:** Delete issue type scheme

Deletes an issue type scheme.

Only issue type schemes used in classic projects can be deleted. Only issue type schemes not associated with a project can be deleted

A validation error will be returned if the specified scheme is associated with one or more projects. Use [Get issue type scheme API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-get) (with the projects expand, and id query parameter) to get a list of project

## Parameters
- `issueTypeSchemeId` [path] (required) integer(int64) — The ID of the issue type scheme.

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
