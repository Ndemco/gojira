# GET /rest/api/3/announcementBanner
**operationId:** `getBanner`
**Summary:** Get announcement banner configuration

Returns the current announcement banner configuration.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: object:
  - `hashId`: string
  - `isDismissible`: boolean
  - `isEnabled`: boolean
  - `message`: string
  - `visibility`: string
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
