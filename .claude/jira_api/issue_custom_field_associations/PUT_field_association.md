# PUT /rest/api/3/field/association
**operationId:** `createAssociations`
**Summary:** Create associations

Associates fields with projects.

Fields will be associated with each issue type on the requested projects.

Fields will be associated with all projects that share the same field configuration which the provided projects are using. This means that while the field will be associated with the requested projects, it will also be associated with any other projects that share the same field configuration.

If a success response is returned it means that the field association has been created in any a

## Request Body
Content-Type: `application/json`
object:
  - `associationContexts` (required): []AssociationContextObject
  - `fields` (required): []FieldIdentifierObject

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the field, project, or issue type is not found.
