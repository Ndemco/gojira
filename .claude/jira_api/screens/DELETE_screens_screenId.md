# DELETE /rest/api/3/screens/{screenId}
**operationId:** `deleteScreen`
**Summary:** Delete screen

Deletes a screen. A screen cannot be deleted if it is used in a screen scheme, workflow, or workflow draft.

Only screens used in classic projects can be deleted.

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.

## Responses
- 204: Returned if the request is successful.
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
