# GET /rest/api/3/permissionscheme
**operationId:** `getAllPermissionSchemes`
**Summary:** Get all permission schemes

Returns all permission schemes.

### About permission schemes and grants ###

A permission scheme is a collection of permission grants. A permission grant consists of a `holder` and a `permission`.

#### Holder object ####

The `holder` object contains information about the user or group being granted the permission. For example, the *Administer projects* permission is granted to a group named *Teams in space administrators*. In this case, the type is `"type": "group"`, and the parameter is the 

## Parameters
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Responses
- 200: object:
  - `permissionSchemes`: []PermissionScheme
- 401: Returned if the authentication credentials are incorrect or missing.
