# DELETE /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey}
**operationId:** `AddonPropertiesResource.deleteAddonProperty_delete`
**Summary:** Delete app property

Deletes an app's property.

**[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.
Additionally, Forge apps can access Connect app properties (stored against the same `app.connect.key`).

## Parameters
- `addonKey` [path] (required) string — The key of the app, as defined in its descriptor.
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 204: Returned if the request is successful.
- 400: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 403: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer
