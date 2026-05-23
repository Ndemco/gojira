# GET /rest/api/3/mypreferences
**operationId:** `getPreference`
**Summary:** Get preference

Returns the value of a preference of the current user.

Note that these keys are deprecated:

 *  *jira.user.locale* The locale of the user. By default this is not set and the user takes the locale of the instance.
 *  *jira.user.timezone* The time zone of the user. By default this is not set and the user takes the timezone of the instance.

These system preferences keys will be deprecated by 15/07/2024. You can still retrieve these keys, but it will not have any impact on Notification behaviour

## Parameters
- `key` [query] (required) string — The key of the preference.

## Responses
- 200: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the key is not provided or not found.
