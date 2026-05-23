# POST /rest/api/3/bulk/issues/move
**operationId:** `submitBulkMove`
**Summary:** Bulk move issues

Use this API to submit a bulk issue move request. You can move multiple issues from multiple projects in a single request, but they must all be moved to a single project, issue type, and parent. You can't move more than 1000 issues (including subtasks) at once.

#### Scenarios: ####

This is an early version of the API and it doesn't have full feature parity with the Bulk Move UI experience.

 *  Moving issue of type A to issue of type B in the same project or a different project: `SUPPORTED`
 *

## Request Body
Content-Type: `application/json`
object:
  - `sendBulkNotification`: boolean
  - `targetToSourcesMapping`: object

## Responses
- 201: object:
  - `taskId`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
