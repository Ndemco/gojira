# DELETE /rest/api/3/user/properties/{propertyKey}
**operationId:** `deleteUserProperty`
**Summary:** Delete user property

Deletes a property from a user.

Note: This operation does not access the [user properties](https://confluence.atlassian.com/x/8YxjL) created and maintained in Jira.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg), to delete a property from any user.
 *  Access to Jira, to delete a property from the calling user's record.

## Parameters
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0
- `userKey` [query] string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http
- `username` [query] string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http
- `propertyKey` [path] (required) string — The key of the user's property.

## Responses
- 204: Returned if the user property is deleted.
- 400: Returned if `accountId` is missing.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission or is not accessing their user record.
- 404: Returned if the user or the property is not found.
