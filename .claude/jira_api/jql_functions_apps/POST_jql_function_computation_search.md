# POST /rest/api/3/jql/function/computation/search
**operationId:** `getPrecomputationsByID`
**Summary:** Get precomputations by ID (apps)

Returns function precomputations by IDs, along with information about when they were created, updated, and last used. Each precomputation has a `value` \- the JQL fragment to replace the custom function clause with.

**[Permissions](#permissions) required:** This API is only accessible to apps and apps can only inspect their own functions.

The new `read:app-data:jira` OAuth scope is 100% optional now, and not using it won't break your app. However, we recommend adding it to your app's scope lis

## Parameters
- `orderBy` [query] string — [Order](#ordering) the results by a field:

## Request Body
Content-Type: `application/json`
object:
  - `precomputationIDs`: []string

## Responses
- 200: object:
  - `notFoundPrecomputationIDs`: []string
  - `precomputations`: []JqlFunctionPrecomputationBean
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the request is not authenticated as the app that provided the function.
- 404: Returned if the function is not found.
