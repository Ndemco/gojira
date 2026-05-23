# PUT /rest/api/3/mypreferences
**operationId:** `setPreference`
**Summary:** Set preference

Creates a preference for the user or updates a preference's value by sending a plain text string. For example, `false`. An arbitrary preference can be created with the value containing up to 255 characters. In addition, the following keys define system preferences that can be set or created:

 *  *user.notifications.mimetype* The mime type used in notifications sent to the user. Defaults to `html`.
 *  *user.default.share.private* Whether new [ filters](https://confluence.atlassian.com/x/eQiiLQ)

## Parameters
- `key` [query] (required) string — The key of the preference. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
string
Content-Type: `text/plain`
string

## Responses
- 204: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the key or value is not provided or invalid.
