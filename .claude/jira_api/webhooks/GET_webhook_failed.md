# GET /rest/api/3/webhook/failed
**operationId:** `getFailedWebhooks`
**Summary:** Get failed webhooks

Returns webhooks that have recently failed to be delivered to the requesting app after the maximum number of retries.

After 72 hours the failure may no longer be returned by this operation.

The oldest failure is returned first.

This method uses a cursor-based pagination. To request the next page use the failure time of the last webhook on the list as the `failedAfter` value or use the URL provided in `next`.

**[Permissions](#permissions) required:** Only [Connect apps](https://developer.atla

## Parameters
- `maxResults` [query] integer(int32) — The maximum number of webhooks to return per page. If obeying the maxResults directive would result in records with the 
- `after` [query] integer(int64) — The time after which any webhook failure must have occurred for the record to be returned, expressed as milliseconds sin

## Responses
- 200: object:
  - `maxResults` (required): integer(int32)
  - `next`: string(uri)
  - `values` (required): []FailedWebhook
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
