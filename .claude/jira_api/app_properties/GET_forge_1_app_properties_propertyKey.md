# GET /rest/forge/1/app/properties/{propertyKey}
**operationId:** `getForgeAppProperty`
**Summary:** Get app property (Forge)

Returns the value of a Forge app's property.

**[Permissions](#permissions) required:** Only Forge apps can make this request. This API can only be accessed using **[asApp()](https://developer.atlassian.com/platform/forge/apis-reference/fetch-api-product.requestjira/#method-signature)** requests from Forge.

## Parameters
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 200: object:
  - `key`: string
  - `value`: any
- 400: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer
