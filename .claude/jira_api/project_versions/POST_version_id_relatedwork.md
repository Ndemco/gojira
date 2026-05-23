# POST /rest/api/3/version/{id}/relatedwork
**operationId:** `createRelatedWork`
**Summary:** Create related work

Creates a related work for the given version. You can only create a generic link type of related works via this API. relatedWorkId will be auto-generated UUID, that does not need to be provided.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Resolve issues:* and *Edit issues* [Managing project permissions](https://confluence.atlassian.com/adminjiraserver/managing-project-permissions-938847145.html) for the project that contains the version.

## Parameters
- `id` [path] (required) string — 

## Request Body
Content-Type: `application/json`
object:
  - `category` (required): string
  - `issueId`: integer(int64)
  - `relatedWorkId`: string
  - `title`: string
  - `url`: string(uri)

## Responses
- 201: object:
  - `category` (required): string
  - `issueId`: integer(int64)
  - `relatedWorkId`: string
  - `title`: string
  - `url`: string(uri)
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
- 404: Returned if the version is not found.
