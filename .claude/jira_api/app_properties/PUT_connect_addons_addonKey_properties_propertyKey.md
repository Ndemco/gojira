# PUT /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey}
**operationId:** `AddonPropertiesResource.putAddonProperty_put`
**Summary:** Set app property

Sets the value of an app's property. Use this resource to store custom data for your app.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

**[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.
Additionally, Forge apps can access Connect app properties (stored against the same `app.connect.key`).

## Parameters
- `addonKey` [path] (required) string — The key of the app, as defined in its descriptor.
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
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 403: object:
  - `message` (required): string
  - `statusCode` (required): integer
