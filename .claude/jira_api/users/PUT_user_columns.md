# PUT /rest/api/3/user/columns
**operationId:** `setUserColumns`
**Summary:** Set user default columns

Sets the default [ issue table columns](https://confluence.atlassian.com/x/XYdKLg) for the user. If an account ID is not passed, the calling user's default columns are set. If no column details are sent, then all default columns are removed.

The parameters for this resource are expressed as HTML form data. For example, in curl:

`curl -X PUT -d columns=summary -d columns=description https://your-domain.atlassian.net/rest/api/3/user/columns?accountId=5b10ac8d82e05b22cc7d4ef5'`

**[Permissions](#

## Parameters
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0

## Request Body
Content-Type: `*/*`
object:
  - `columns`: []string
Content-Type: `multipart/form-data`
object:
  - `columns`: []string

## Responses
- 200: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission or is not accessing their user record.
- 404: Returned if the requested user is not found.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
- 500: Returned if an invalid issue table column ID is sent.
