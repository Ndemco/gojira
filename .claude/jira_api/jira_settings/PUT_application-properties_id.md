# PUT /rest/api/3/application-properties/{id}
**operationId:** `setApplicationProperty`
**Summary:** Set application property

Changes the value of an application property. For example, you can change the value of the `jira.clone.prefix` from its default value of *CLONE -* to *Clone -* if you prefer sentence case capitalization. Editable properties are described below along with their default values.

#### Advanced settings ####

The advanced settings below are also accessible in [Jira](https://confluence.atlassian.com/x/vYXKM).

| Key | Description | Default value |  
| -- | -- | -- |  
| `jira.clone.prefix` | The stri

## Parameters
- `id` [path] (required) string — The key of the application property to update.

## Request Body
Content-Type: `application/json`
object:
  - `id`: string
  - `value`: string

## Responses
- 200: object:
  - `allowedValues`: []string
  - `defaultValue`: string
  - `desc`: string
  - `example`: string
  - `id`: string
  - `key`: string
  - `name`: string
  - `type`: string
  - `value`: string
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
