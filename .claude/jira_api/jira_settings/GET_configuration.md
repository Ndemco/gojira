# GET /rest/api/3/configuration
**operationId:** `getConfiguration`
**Summary:** Get global settings

Returns the [global settings](https://confluence.atlassian.com/x/qYXKM) in Jira. These settings determine whether optional features (for example, subtasks, time tracking, and others) are enabled. If time tracking is enabled, this operation also returns the time tracking configuration.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: object:
  - `attachmentsEnabled`: boolean
  - `issueLinkingEnabled`: boolean
  - `subTasksEnabled`: boolean
  - `timeTrackingConfiguration`: allOf(TimeTrackingConfiguration)
  - `timeTrackingEnabled`: boolean
  - `unassignedIssuesAllowed`: boolean
  - `votingEnabled`: boolean
  - `watchingEnabled`: boolean
- 401: Returned if the authentication credentials are incorrect or missing.
