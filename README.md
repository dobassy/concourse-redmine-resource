# Redmine ticket creation resource

[![CircleCI](https://circleci.com/gh/dobassy/concourse-redmine-resource/tree/master.svg?style=svg)](https://circleci.com/gh/dobassy/concourse-redmine-resource/tree/master)

## Getting Started

```
resource_types:
- name: redmine
  type: registry-image
  source:
    repository: dobassy/concourse-redmine-resource
    tag: latest
    username: ((docker_username))
    password: ((docker_password))

resources:
- name: redmine
  type: redmine
  source:
    uri: ((redmine_uri))
    apikey: ((redmine_apikey))

jobs:
- name: post-redmine
  plan:
  - task: build text
    ...some tasks

  - put: redmine
    params:
      project_id: 1
      tracker_id: 1
      content_file: assets/message
```

## Motivation

t.b.d.

## Source Configuration

### Preparation:

Generate a redmine api key on the account page. It probably has a link at the top right of the screen.

### Parameters:

- `uri`: (**Required**)<br>
  `e.g. http(s)://your.redmine.domain/`
- `apikey`: (**Required**)
- `insecure`: (Optional. default: `false`) skip SSL validation.

## Behavior

### `check`:

Not provided.

### `in`:

Not provided.

### `out`:

Create a simple issue (ticket) with the title (subject) and body (description).

#### Parameters:

- `project_id`: (**Required**) Project ID must be a number. You should be able to find it in `/projects.json` url.
- `tracker_id`: (**Required**)
- `subject`: (Optional) Used when creating a simple ticket with only a title specified.
- `content_file`: (Optional, **Recommended**) Use this if you want to specifiy the body. This parameter specifies the file name containing `input` name. `e.g. assets/message`. See also `out/fixtures/ticket_content_file.txt` example.
- `status_id`: (Optional)
