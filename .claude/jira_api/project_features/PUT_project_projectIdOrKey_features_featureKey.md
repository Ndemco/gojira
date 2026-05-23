# PUT /rest/api/3/project/{projectIdOrKey}/features/{featureKey}
**operationId:** `toggleFeatureForProject`
**Summary:** Set project feature state

Sets the state of a project feature.

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or (case-sensitive) key of the project.
- `featureKey` [path] (required) string — The key of the feature.

## Request Body
Content-Type: `application/json`
object:
  - `state`: string

## Responses
- 200: object:
  - `features`: []ProjectFeature
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
- 404: Returned if the project or project feature is not found.
