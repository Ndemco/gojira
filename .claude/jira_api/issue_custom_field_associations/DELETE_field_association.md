# DELETE /rest/api/3/field/association
**operationId:** `removeAssociations`
**Summary:** Remove associations

Unassociates a set of fields with a project and issue type context.

Fields will be unassociated with all projects/issue types that share the same field configuration which the provided project and issue types are using. This means that while the field will be unassociated with the provided project and issue types, it will also be unassociated with any other projects and issue types that share the same field configuration.

If a success response is returned it means that the field association ha

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
