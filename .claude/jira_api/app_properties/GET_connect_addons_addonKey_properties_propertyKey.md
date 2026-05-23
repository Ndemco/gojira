# GET /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey}
**operationId:** `AddonPropertiesResource.getAddonProperty_get`
**Summary:** Get app property

Returns the key and value of an app's property. The property key `connect_client_key_019cdff3-8bfb-71fe-9628-875b700aebb8`
is reserved. It returns a synthetic, read-only property containing the Connect `clientKey` for the requested tenant.
This is intended for Forge apps with `app.connect.key` to retrieve the Connect client key during migration.

**[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.
Additionally, Forge apps can access Co

## Parameters
- `addonKey` [path] (required) string — The key of the app, as defined in its descriptor.
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 200: object:
  - `key`: string
  - `value`: any
- 400: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer
