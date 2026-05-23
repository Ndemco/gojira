# GET /rest/api/3/redact/status/{jobId}
**operationId:** `getRedactionStatus`
**Summary:** Get redaction status

Retrieves the current status of a redaction job ID.

The jobStatus will be one of the following:

 *  IN\_PROGRESS - The redaction job is currently in progress
 *  COMPLETED - The redaction job has completed successfully.
 *  PENDING - The redaction job has not started yet

## Parameters
- `jobId` [path] (required) string — Redaction job id

## Responses
- 200: object:
  - `bulkRedactionResponse`: BulkRedactionResponse
  - `jobStatus`: string
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
