# GET /rest/api/3/application-properties/advanced-settings
**operationId:** `getAdvancedSettings`
**Summary:** Get advanced settings

Returns the application properties that are accessible on the *Advanced Settings* page. To navigate to the *Advanced Settings* page in Jira, choose the Jira icon > **Jira settings** > **System**, **General Configuration** and then click **Advanced Settings** (in the upper right).

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: []object:
  - `allowedValues`: []string
  - `defaultValue`: string
  - `desc`: string
  - `example`: string
  - `id`: string
  - `key`: string
  - `name`: string
  - `type`: string
  - `value`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user is not an administrator.
