# GET /rest/api/3/data-policy/project
**operationId:** `getPolicies`
**Summary:** Get data policy for projects

Returns data policies for the projects specified in the request.

## Parameters
- `ids` [query] string — A list of project identifiers. This parameter accepts a comma-separated list.

## Responses
- 200: object:
  - `projectDataPolicies`: []ProjectWithDataPolicy
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
