# GET /rest/forge/1/app/properties
**operationId:** `getForgeAppPropertyKeys`
**Summary:** Get app property keys (Forge)

Returns all property keys for the Forge app.

**[Permissions](#permissions) required:** Only Forge apps can make this request. This API can only be accessed using **[asApp()](https://developer.atlassian.com/platform/forge/apis-reference/fetch-api-product.requestjira/#method-signature)** requests from Forge.

## Responses
- 200: object:
  - `keys`: []object
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
