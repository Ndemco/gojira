# GET /rest/api/3/events
**operationId:** `getEvents`
**Summary:** Get events

Returns all issue events.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: []object:
  - `id`: integer(int64)
  - `name`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to complete this request.
