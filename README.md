[![GoDoc](https://godoc.org/github.com/bulletind/khabar?status.svg)](https://godoc.org/github.com/bulletind/khabar)
[![Build Status](https://travis-ci.org/bulletind/khabar.svg?branch=develop)](https://travis-ci.org/bulletind/khabar)
[![Coverage Status](https://coveralls.io/repos/bulletind/khabar/badge.svg?branch=develop&service=github)](https://coveralls.io/github/bulletind/khabar?branch=develop)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/bulletind/khabar?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

## Khabar

Notifications engine.

It means

> *the latest information; news.*

[google it](https://www.google.com/search?q=define+khabar)

## Table of contents

- [Concept and idea](#concept-and-idea)
- [How does it work?](#how-does-it-work)
- [Development](#development)
  - [Sanity checks](#sanity-checks)
- [Usage](#usage)
- [API](#api)
  - **[Preferences](#preferences)**
    1. [List all preferences of user or org](#list-all-preferences-of-user-or-org) `GET /topics?org=&user=`
    2. [Set an user preference](#set-an-user-preference) `POST /topics/:ident/channels/:channel?org=&user=`
    3. [Unset an user preference](#unset-an-user-preference) `DELETE /topics/:ident/channels/:channel?org=&user=`
    4. [Set an org preference](#set-an-org-preference) `POST /topics/:type/:ident/channels/:channel?org=`
    5. [Unset an org preference](#unset-an-org-preference) `DELETE /topics/:type/:ident/channels/:channel?org=`
  - **[Notifications](#notifications)**
    1. [List all notifications](#list-all-notifications) `GET /notifications`
    2. [Mark a single notification as read](#mark-a-single-notification-as-read) `PUT /notification/:_id`
    3. [Mark all unread notifications as read](#mark-all-unread-notifications-as-read) `PUT /notifications?org=&user=`
    4. [Get notification stats](#get-notification-stats) `GET /notifications/stats?org=&user=`
    5. [Update last seen time stamp](#update-last-seen-time-stamp) `PUT /notifications/stats?org=&user=`
    6. [Send notification](#send-notification) `POST /notifications?topic=`
  - [Some general conventions](#some-general-conventions)
- [Envrionment variables](#environment-variables)
  - [Push notifications](#push-notifications)
  - [Email notifications](#email-notifications)

## Concept and idea

## How does it work?

## Development

```sh
$ go get github.com/codegangsta/gin
$ go get github.com/tools/godep
```

- [gin](http://github.com/codegangsta/gin) is used to to automatically compile files while you are developing
- [godep](http://github.com/tools/godep) is used to manage dependencies

Then run

```sh
$ mkdir -p $GOPATH/src/github.com/bulletind
$ cd $GOPATH/src/github.com/bulletind
$ git clone https://github.com/bulletind/khabar.git # or your fork
$ cd khabar
$ DEBUG=* go get && go install && gin -p 8911 -i # or make dev
```

Now you should be able to access the below API's on port `8911`.

MongoDB config is stored in `config/conf.go`.

After you make the changes (if you import any new deps), don't forget to run

```sh
$ godep save ./... # or make godep
```

#### Sanity checks

```sh
$ make vet    # https://godoc.org/golang.org/x/tools/cmd/vet
$ make lint   # https://github.com/golang/lint
$ make test
```

## Usage

```sh
$ go get github.com/bulletind/khabar
$ khabar
```

## API

1. #### Preferences

  1. ##### List all preferences of user or org

    ```
    GET /topics
    ```

    Query params:

    - `org`: organization id
    - `user`: user id

    If org and user are not sent in query params, then it will send global preferences.

    If org is sent and user is not, then it will send org preferences.

    If org and user are both sent, then it will send user preferences.

    Response:
    ```js
    [
      {
        "_id": "",
        "created_on": 1425547531188,
        "updated_on": 1425879125700,
        "user": "",
        "org": "",
        "app_name": "",
        "channels": [
          "",
          ""
        ],
        "ident": ""
      }
    ]
    ```

  2. ##### Set an user preference

    ```
    POST /topics/:ident/channels/:channel
    ```

    Query params:

    - `org`: (required) organization id
    - `user`: (required) user id

  3. ##### Unset an user preference

    ```
    DELETE /topics/:ident/channels/:channel
    ```

    Query params:

    - `org`: (required) organization id
    - `user`: (required) user id

  4. ##### Set an org preference

    ```
    POST /topics/:type/:ident/channels/:channel
    ```

    `:type` here is either `defaults` or `locked`

    Query params:

    - `org`: (required) organization id

  5. ##### Unset an org preference

    ```
    DELETE /topics/:type/:ident/channels/:channel
    ```

    `:type` here is either `defaults` or `locked`

    Query params:

    - `org`: (required) organization id

2. #### Notifications

  1. ##### List all notifications

    ```
    GET /notifications
    ```

    Query params:

    - `user`: user id
    - `org`: organization id

    Response:
    ```js
    [
      {
        org: "",
        user: "",
        destination_uri: "",
        text: "",
        topic: "",
        destination_uri: "",
        is_read:false,
        created_on: <milliseconds_since_epoch>
       },
       // and so on...
    ]
    ```

    This can be polled periodically

  2. ##### Mark a single notification as read

    ```
    PUT /notification/:_id
    ```

    Request Body:
    ```json
    {
      "destination_uri": "http://link-to-entity",
      "text": "Notification text",
      "topic": "Notification topic"
    }
    ```

    - `destination_uri`: (required) Link to relevant entity. (i.e action, incident)
    - `text`: (required) Notification text (long text)
    - `topic`: (required) Notification topic (short text)

  3. ##### Mark all unread notifications as read

    ```
    PUT /notifications
    ```

    Request Body:
    ```json
    {
      "org": "123",
      "user": "456"
    }
    ```

    - `org`: (required) org id
    - `user`: (required) user id
    - `app_name`: (optional) name of the app or category

  4. ##### Get notification stats

    ```
    GET /notifications/stats
    ```

    Query params:

    - `user`: (required) user id
    - `org`: (required) organization id

    Response:
    ```
    {
      "last_seen": "2015-08-03T14:26:05.860Z",
      "total_count": 37,
      "unread_count": 0,
      "total_unread": 32
    }
    ```

  5. ##### Update last seen time stamp

    ```
    PUT /notifications/stats
    ```

    Query params:

    - `user`: (required) user id
    - `org`: (required) organization id

  6. ##### Send notification

    ```
    POST /notifications?topic=text
    ```

    Request Body:
    ```js
    {
      "created_by" : "5486e02870a0d30200bdcfd3",
      "org" : "5486d3d986ba633a207682b6",
      "app_name" : "myapp", // <- this is the category
      "topic" : "log_incoming",
      "user" : "5486e02870a0d30200bdcfe0",
      "destination_uri" : "http://...",
      "context" : {
        "Organization" : "5486d3d986ba633a207682b6",
        "sender" : "org name",
        "fullname" : "John hopkins",
        "logger" : "Elvis",
        "refnumber" : "IL2958",
        "Collection" : "collection_name",
        "Id" : "554caa744aca430a00de5324",
        "User" : "5486e02870a0d30200bdcfe0",
        "destination_uri" : "http://...",
        "email" : "john.hopkins@email",
        "severity" : "low",
        "subject" : "New log"
      },
      "entity" : "554caa744aca430a00de5324"
    }
    ```

    Query params:

    - `topic`: Text that is used for sending notification (short text)

#### Some general conventions:

- For all of the above request you must pass atleast one of the `org` or `user` or both
- For all the listings, you get a status code of `200`
- When you create a resource you get a status code of `201`
- When you modify/delete a resource you get a status code of `204`
- Response for creation of an entity

  ```js
  {
    "body": "[ID of entity created]",
    "message": "Created",
    "status": 201
  }
  ```
- Response for modifying/deleting an entity

  ```js
  {
    "body": "",
    "message": "NoContent",
    "status": 204
  }
  ```

## Environment variables

We use environment variables to fetch the keys and secrets.

#### Push notifications

We use [parse](https://www.parse.com/) to send push notifications. When you are sending out a notification using the [`POST /notifications`](#send-notification) api call, it looks for certain environment variables. These env variables are based on the categories (`app_name`s) you are using in the `topics_available` collection.

For example: You have an event (ident) `log_incoming` configured for the category (app_name) `myapp` in the `topics_available` collection. Now, when you make a call to `POST /notifications`, it looks for `PARSE_myapp_API_KEY` and `PARSE_myapp_APP_ID` enviroment variables and uses them to send the push notifications.

You can set them by doing

```sh
$ export PARSE_myapp_API_KEY=***
$ export PARSE_myapp_APP_ID=***
```

#### Email notifications

You can configure the email notifications by setting the following env variables

```sh
$ export SMTP_HOSTNAME=***
$ export SMTP_USERNAME=***
$ export SMTP_PASSWORD=***
$ export SMTP_PORT=***
$ export SMTP_FROM=***
```
