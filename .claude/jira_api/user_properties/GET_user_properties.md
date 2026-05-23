# GET /rest/api/3/user/properties
**operationId:** `getUserPropertyKeys`
**Summary:** Get user property keys

Returns the keys of all properties for a user.

Note: This operation does not access the [user properties](https://confluence.atlassian.com/x/8YxjL) created and maintained in Jira.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg), to access the property keys on any user.
 *  Access to Jira, to access the calling user's property keys.

## Parameters
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0
- `userKey` [query] string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http
- `username` [query] string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 400: Returned if `accountId` is missing.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission or is not accessing their user record.
- 404: Returned if the user is not found.
