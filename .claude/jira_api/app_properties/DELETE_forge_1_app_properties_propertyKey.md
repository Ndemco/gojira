# DELETE /rest/forge/1/app/properties/{propertyKey}
**operationId:** `deleteForgeAppProperty`
**Summary:** Delete app property (Forge)

Deletes a Forge app's property.

**[Permissions](#permissions) required:** Only Forge apps can make this request. This API can only be accessed using **[asApp()](https://developer.atlassian.com/platform/forge/apis-reference/fetch-api-product.requestjira/#method-signature)** requests from Forge.

The new `write:app-data:jira` OAuth scope is 100% optional now, and not using it won't break your app. However, we recommend adding it to your app's scope list because we will eventually make it mandator

## Parameters
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 204: Returned if the request is successful.
- 400: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer
