# PUT /rest/forge/1/app/properties/{propertyKey}
**operationId:** `putForgeAppProperty`
**Summary:** Set app property (Forge)

Sets the value of a Forge app's property.
These values can be retrieved in [Jira expressions](/cloud/jira/platform/jira-expressions/)
through the `app` [context variable](/cloud/jira/platform/jira-expressions/#context-variables).
They are also available in [entity property display conditions](/platform/forge/manifest-reference/display-conditions/entity-property-conditions/).

For other use cases, use the [Storage API](/platform/forge/runtime-reference/storage-api/).

The value of the request bod

## Parameters
- `propertyKey` [path] (required) string — The key of the property.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 201: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 400: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
