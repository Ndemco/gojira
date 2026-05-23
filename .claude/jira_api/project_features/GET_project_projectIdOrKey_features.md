# GET /rest/api/3/project/{projectIdOrKey}/features
**operationId:** `getFeaturesForProject`
**Summary:** Get project features

Returns the list of features for a project.

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or (case-sensitive) key of the project.

## Responses
- 200: object:
  - `features`: []ProjectFeature
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
- 404: Returned if the project is not found.
