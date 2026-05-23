# POST /rest/internal/api/latest/worklog/bulk
**operationId:** `getWorklogsByIssueIdAndWorklogId`
**Summary:** Get worklogs by issue id and worklog id

Returns worklog details for a list of issue ID and worklog ID pairs.

This is an internal API for bulk fetching worklogs by their issue and worklog IDs. Worklogs that don't exist will be filtered out from the response.

The returned list of worklogs is limited to 1000 items.

**[Permissions](#permissions) required:** This is an internal service-to-service API that requires ASAP authentication. No user permission checks are performed as this bypasses normal user context.

## Request Body
Content-Type: `application/json`
object:
  - `requests`: []WorklogCompositeKey

## Responses
- 200: object:
  - `worklogs`: []WorklogKeyResult
- 400: Returned if the request contains more than 1000 worklog pairs, is empty, or has invalid format.
- 401: Returned if the authentication credentials are incorrect or missing.
- 500: Returned if there is an internal server error.
