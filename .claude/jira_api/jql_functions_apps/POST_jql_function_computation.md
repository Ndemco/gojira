# POST /rest/api/3/jql/function/computation
**operationId:** `updatePrecomputations`
**Summary:** Update precomputations (apps)

Update the precomputation value of a function created by a Forge/Connect app.

**[Permissions](#permissions) required:** An API for apps to update their own precomputations.

The new `write:app-data:jira` OAuth scope is 100% optional now, and not using it won't break your app. However, we recommend adding it to your app's scope list because we will eventually make it mandatory.

## Parameters
- `skipNotFoundPrecomputations` [query] boolean — 

## Request Body
Content-Type: `application/json`
object:
  - `values`: []JqlFunctionPrecomputationUpdateBean

## Responses
- 200: object:
  - `notFoundPrecomputationIDs`: []string
- 204: any
- 400: object:
  - `errorMessages`: []string
  - `notFoundPrecomputationIDs`: []string
- 403: object:
  - `errorMessages`: []string
  - `notFoundPrecomputationIDs`: []string
- 404: object:
  - `errorMessages`: []string
  - `notFoundPrecomputationIDs`: []string
