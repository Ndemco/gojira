# GET /rest/api/3/user/email
**operationId:** `getUserEmail`
**Summary:** Get user email

Returns a user's email address regardless of the user's profile visibility settings. For Connect apps, this API is only available to apps approved by Atlassian, according to these [guidelines](https://community.developer.atlassian.com/t/guidelines-for-requesting-access-to-email-address/27603). For Forge apps, this API only supports access via asApp() requests.

## Parameters
- `accountId` [query] (required) string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, `5b10ac8d82e0

## Responses
- 200: object:
  - `accountId`: string
  - `email`: string
- 400: Returned if the calling app is not approved to use this API.
- 401: Returned if the authentication credentials are incorrect or missing from the request (for example if a user is trying to access this API).
- 404: Returned if a user with the given `accountId` doesn't exist
- 503: Indicates the API is not currently enabled
