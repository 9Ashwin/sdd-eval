## ADDED Requirements

### Requirement: Get single task

The system SHALL allow clients to retrieve a single task by its ID.

#### Scenario: Task exists

- **WHEN** client sends `GET /tasks/{id}` with an existing task ID
- **THEN** system returns HTTP 200 with the task JSON (`{"id": <int>, "title": "<string>", "done": <bool>}`)

#### Scenario: Task not found

- **WHEN** client sends `GET /tasks/{id}` with a non-existent task ID
- **THEN** system returns HTTP 404 with `{"error":"not found"}`

#### Scenario: Invalid ID format

- **WHEN** client sends `GET /tasks/{id}` with a non-integer `{id}`
- **THEN** system returns HTTP 400 with `{"error":"invalid id"}`

### Requirement: Partial update task

The system SHALL allow clients to partially update a task's `title` and/or `done` fields.

#### Scenario: Update title only

- **WHEN** client sends `PATCH /tasks/{id}` with body `{"title": "new title"}`
- **THEN** system updates the task's title, returns HTTP 200 with the updated task JSON

#### Scenario: Update done only

- **WHEN** client sends `PATCH /tasks/{id}` with body `{"done": true}`
- **THEN** system updates the task's done status, returns HTTP 200 with the updated task JSON

#### Scenario: Update both fields

- **WHEN** client sends `PATCH /tasks/{id}` with body `{"title": "new title", "done": true}`
- **THEN** system updates both fields, returns HTTP 200 with the updated task JSON

#### Scenario: Task not found

- **WHEN** client sends `PATCH /tasks/{id}` with a non-existent task ID
- **THEN** system returns HTTP 404 with `{"error":"not found"}`

#### Scenario: No valid fields provided

- **WHEN** client sends `PATCH /tasks/{id}` with body `{}` or body without `title` or `done`
- **THEN** system returns HTTP 400 with `{"error":"no fields to update"}`

#### Scenario: Invalid ID format

- **WHEN** client sends `PATCH /tasks/{id}` with a non-integer `{id}`
- **THEN** system returns HTTP 400 with `{"error":"invalid id"}`
