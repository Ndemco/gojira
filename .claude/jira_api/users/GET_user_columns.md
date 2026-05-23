# GET /rest/api/3/user/columns
**operationId:** `getUserDefaultColumns`
**Summary:** Get user default columns

Returns the default [issue table columns](https://confluence.atlassian.com/x/XYdKLg) for the user. If `accountId` is not passed in the request, the calling user's details are returned.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLgl), to get the column details for any user.
 *  Permission to access Jira, to get the calling user's column details.

## Parameters
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0
- `username` [query] string — This parameter is no longer available See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/d

## Responses
- 200: []object:
  - `label`: string
  - `value`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission or is not accessing their user record.
- 404: Returned if the requested user is not found.
