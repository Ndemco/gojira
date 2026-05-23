# GET /rest/atlassian-connect/1/addons/{addonKey}/properties
**operationId:** `AddonPropertiesResource.getAddonProperties_get`
**Summary:** Get app properties

Gets all the properties of an app. The reserved key `connect_client_key_019cdff3-8bfb-71fe-9628-875b700aebb8` is not returned.

**[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.
Additionally, Forge apps can access Connect app properties (stored against the same `app.connect.key`).

## Parameters
- `addonKey` [path] (required) string — The key of the app, as defined in its descriptor.

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
