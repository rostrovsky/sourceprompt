`LICENSE`

```
MIT License

Copyright (c) 2022 Mike Stefanello

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

`Makefile`

```
# Install Ent code-generation module
.PHONY: ent-install
ent-install:
	go get -d entgo.io/ent/cmd/ent

# Generate Ent code
.PHONY: ent-gen
ent-gen:
	go generate ./ent

# Create a new Ent entity
.PHONY: ent-new
ent-new:
	go run entgo.io/ent/cmd/ent new $(name)

# Run the application
.PHONY: run
run:
	clear
	go run cmd/web/main.go

# Run all tests
.PHONY: test
test:
	go test -count=1 -p 1 ./...

# Check for direct dependency updates
.PHONY: check-updates
check-updates:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep "\["
```

`README.md`

````markdown
## Pagoda: Rapid, easy full-stack web development starter kit in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/mikestefanello/pagoda)](https://goreportcard.com/report/github.com/mikestefanello/pagoda)
[![Test](https://github.com/mikestefanello/pagoda/actions/workflows/test.yml/badge.svg)](https://github.com/mikestefanello/pagoda/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikestefanello/pagoda.svg)](https://pkg.go.dev/github.com/mikestefanello/pagoda)
[![GoT](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/pagoda)

<p align="center"><img alt="Logo" src="https://user-images.githubusercontent.com/552328/147838644-0efac538-a97e-4a46-86a0-41e3abdf9f20.png" height="200px"/></p>

## Table of Contents
* [Introduction](#introduction)
    * [Overview](#overview)
    * [Foundation](#foundation)
      * [Backend](#backend)
      * [Frontend](#frontend)
      * [Storage](#storage)
    * [Screenshots](#screenshots)
* [Getting started](#getting-started)
  * [Dependencies](#dependencies)
  * [Start the application](#start-the-application)
  * [Running tests](#running-tests)
* [Service container](#service-container)
  * [Dependency injection](#dependency-injection)
  * [Test dependencies](#test-dependencies)
* [Configuration](#configuration)
    * [Environment overrides](#environment-overrides)
    * [Environments](#environments)
* [Database](#database)
    * [Auto-migrations](#auto-migrations)
    * [Separate test database](#separate-test-database)
* [ORM](#orm)
  * [Entity types](#entity-types)
  * [New entity type](#new-entity-type)
* [Sessions](#sessions)
  * [Encryption](#encryption)
* [Authentication](#authentication)
  * [Login / Logout](#login--logout)
  * [Forgot password](#forgot-password)
  * [Registration](#registration)
  * [Authenticated user](#authenticated-user)
    * [Middleware](#middleware)
  * [Email verification](#email-verification)
* [Routes](#routes)
  * [Custom middleware](#custom-middleware)
  * [Handlers](#handlers)
  * [Errors](#errors)
  * [Testing](#testing)
    * [HTTP server](#http-server)
    * [Request / Request helpers](#request--response-helpers)
    * [Goquery](#goquery)
* [Pages](#pages)
  * [Flash messaging](#flash-messaging)
  * [Pager](#pager)
  * [CSRF](#csrf)
  * [Automatic template parsing](#automatic-template-parsing)
  * [Cached responses](#cached-responses)
    * [Cache tags](#cache-tags)
    * [Cache middleware](#cache-middleware)
  * [Data](#data)
  * [Forms](#forms)
    * [Submission processing](#submission-processing)
    * [Inline validation](#inline-validation)
  * [Headers](#headers)
  * [Status code](#status-code)
  * [Metatags](#metatags)
  * [URL and link generation](#url-and-link-generation)
  * [HTMX support](#htmx-support)
  * [Rendering the page](#rendering-the-page)
* [Template renderer](#template-renderer)
  * [Custom functions](#custom-functions)
  * [Caching](#caching)
  * [Hot-reload for development](#hot-reload-for-development)
  * [File configuration](#file-configuration)
* [Funcmap](#funcmap)
* [Cache](#cache)
  * [Set data](#set-data)
  * [Get data](#get-data)
  * [Flush data](#flush-data)
  * [Flush tags](#flush-tags)
* [Tasks](#tasks)
  * [Queues](#queues)
  * [Dispatcher](#dispatcher)
* [Cron](#cron)
* [Static files](#static-files)
  * [Cache control headers](#cache-control-headers)
  * [Cache-buster](#cache-buster)
* [Email](#email)
* [HTTPS](#https)
* [Logging](#logging)
* [Roadmap](#roadmap)
* [Credits](#credits)

## Introduction

### Overview

_Pagoda_ is not a framework but rather a base starter-kit for rapid, easy full-stack web development in Go, aiming to provide much of the functionality you would expect from a complete web framework as well as establishing patterns, procedures and structure for your web application.

Built on a solid [foundation](#foundation) of well-established frameworks and modules, _Pagoda_ aims to be a starting point for any web application with the benefit over a mega-framework in that you have full control over all of the code, the ability to easily swap any frameworks or modules in or out, no strict patterns or interfaces to follow, and no fear of lock-in.

While separate JavaScript frontends have surged in popularity, many prefer the reliability, simplicity and speed of a full-stack approach with server-side rendered HTML. Even the popular JS frameworks all have SSR options. This project aims to highlight that _Go_ templates can be powerful and easy to work with, and interesting [frontend](#frontend) libraries can provide the same modern functionality and behavior without having to write any JS at all.

### Foundation

While many great projects were used to build this, all of which are listed in the [credits](#credits) section, the following provide the foundation of the back and frontend. It's important to note that you are **not required to use any of these**. Swapping any of them out will be relatively easy.

#### Backend

- [Echo](https://echo.labstack.com/): High performance, extensible, minimalist Go web framework.
- [Ent](https://entgo.io/): Simple, yet powerful ORM for modeling and querying data.

#### Frontend

Go server-side rendered HTML combined with the projects below enable you to create slick, modern UIs without writing any JavaScript or CSS.

- [HTMX](https://htmx.org/): Access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext.
- [Alpine.js](https://alpinejs.dev/): Rugged, minimal tool for composing behavior directly in your markup. Think of it like jQuery for the modern web. Plop in a script tag and get going.
- [Bulma](https://bulma.io/): Provides ready-to-use frontend components that you can easily combine to build responsive web interfaces. No JavaScript dependencies.

#### Storage

- [SQLite](https://sqlite.org/): A small, fast, self-contained, high-reliability, full-featured, SQL database engine and the most used database engine in the world.

Originally, Postgres and Redis were chosen as defaults but since the aim of this project is rapid, simple development, it was changed to SQLite which now provides the primary data storage as well as persistent, background [task queues](#tasks). For [caching](#cache), a simple in-memory solution is provided. If you need to use something like Postgres or Redis, swapping those in can be done quickly and easily. For reference, [this branch](https://github.com/mikestefanello/pagoda/tree/postgres-redis) contains the code that included those (but is no longer maintained).

### Screenshots

#### Inline form validation

<img src="https://user-images.githubusercontent.com/552328/147838632-570a3116-1e74-428f-8bfc-523ed309ef06.png" alt="Inline validation"/>

#### Switch layout templates, user registration

<img src="https://user-images.githubusercontent.com/552328/147838633-c1b3e4f6-bbfd-44e1-b0be-884d1a83f8f4.png" alt="Registration"/>

#### Alpine.js modal, HTMX AJAX request

<img src="https://user-images.githubusercontent.com/552328/147838634-4b84c5d5-dc3b-4280-ac12-247ab22184a3.png" alt="Alpine and HTMX"/>

## Getting started

### Dependencies

Ensure that [Go](https://go.dev/) is installed on your system.

### Start the application

After checking out the repository, from within the root, simply run `make run`:

```
git clone git@github.com:mikestefanello/pagoda.git
cd pagoda
make run
```

Since this repository is a _template_ and not a Go _library_, you **do not** use `go get`.

By default, you should be able to access the application in your browser at `localhost:8000`. This can be changed via the [configuration](#configuration).

By default, your data will be stored within the `dbs` directory. If you ever want to quickly delete all data just remove this directory.

### Running tests

To run all tests in the application, execute `make test`. This ensures that the tests from each package are not run in parallel. This is required since many packages contain tests that connect to the test database which is stored in memory and reset automatically for each package.

## Service container

The container is located at `pkg/services/container.go` and is meant to house all of your application's services and/or dependencies. It is easily extensible and can be created and initialized in a single call. The services currently included in the container are:

- Configuration
- Cache
- Database
- ORM
- Web
- Validator
- Authentication
- Mail
- Template renderer
- Tasks

A new container can be created and initialized via `services.NewContainer()`. It can be later shutdown via `Shutdown()`.

### Dependency injection

The container exists to faciliate easy dependency-injection both for services within the container as well as areas of your application that require any of these dependencies. For example, the container is automatically passed to the `Init()` method of your route [handlers](#handlers) so that the handlers have full, easy access to all services.

### Test dependencies

It is common that your tests will require access to dependencies, like the database, or any of the other services available within the container. Keeping all services in a container makes it especially easy to initialize everything within your tests. You can see an example pattern for doing this [here](#environments).

## Configuration

The `config` package provides a flexible, extensible way to store all configuration for the application. Configuration is added to the `Container` as a _Service_, making it accessible across most of the application.

Be sure to review and adjust all of the default configuration values provided in `config/config.yaml`.

### Environment overrides

Leveraging the functionality of [viper](https://github.com/spf13/viper) to manage configuration, all configuration values can be overridden by environment variables. The name of the variable is determined by the set prefix and the name of the configuration field in `config/config.yaml`.

In `config/config.go`, the prefix is set as `pagoda` via `viper.SetEnvPrefix("pagoda")`. Nested fields require an underscore between levels. For example:

```yaml
http:
  port: 1234
```

can be overridden by setting an environment variable with the name `PAGODA_HTTP_PORT`.

### Environments

The configuration value for the current _environment_ (`Config.App.Environment`) is an important one as it can influence some behavior significantly (will be explained in later sections).

A helper function (`config.SwitchEnvironment`) is available to make switching the environment easy, but this must be executed prior to loading the configuration. The common use-case for this is to switch the environment to `Test` before tests are executed:

```go
func TestMain(m *testing.M) {
    // Set the environment to test
    config.SwitchEnvironment(config.EnvTest)

    // Start a new container
    c = services.NewContainer()

    // Run tests
    exitVal := m.Run()

    // Shutdown the container
    if err := c.Shutdown(); err != nil {
        panic(err)
    }

    os.Exit(exitVal)
}
```

## Database

The database currently used is [SQLite](https://sqlite.org/) but you are free to use whatever you prefer. If you plan to continue using [Ent](https://entgo.io/), the incredible ORM, you can check their supported databases [here](https://entgo.io/docs/dialects). The database driver is provided by [go-sqlite3](https://github.com/mattn/go-sqlite3). A reference to the database is included in the `Container` if direct access is required.

Database configuration can be found and managed within the `config` package.

### Auto-migrations

[Ent](https://entgo.io/) provides automatic migrations which are executed on the database whenever the `Container` is created, which means they will run when the application starts.

### Separate test database

Since many tests can require a database, this application supports a separate database specifically for tests. Within the `config`, the test database can be specified at `Config.Database.TestConnection`, which is the database connection string that will be used. By default, this will be an in-memory SQLite database.

When a `Container` is created, if the [environment](#environments) is set to `config.EnvTest`, the database client will connect to the test database instead and run migrations so your tests start with a clean, ready-to-go database.

When this project was using Postgres, it would automatically drop and recreate the test database. Since the current default is in-memory, that is no longer needed. If you decide to use a test database not in-memory, you can alter the `Container` initialization code to do this for you.

## ORM

As previously mentioned, [Ent](https://entgo.io/) is the supplied ORM. It can swapped out, but I highly recommend it. I don't think there is anything comparable for Go, at the current time. If you're not familiar with Ent, take a look through their top-notch [documentation](https://entgo.io/docs/getting-started).

An Ent client is included in the `Container` to provide easy access to the ORM throughout the application.

Ent relies on code-generation for the entities you create to provide robust, type-safe data operations. Everything within the `ent` package in this repository is generated code for the two entity types listed below with the exception of the schema declaration.

### Entity types

The two included entity types are:
- User
- PasswordToken

### New entity type

While you should refer to their [documentation](https://entgo.io/docs/getting-started) for detailed usage, it's helpful to understand how to create an entity type and generate code. To make this easier, the `Makefile` contains some helpers.

1. Ensure all Ent code is downloaded by executing `make ent-install`.
2. Create the new entity type by executing `make ent-new name=User` where `User` is the name of the entity type. This will generate a file like you can see in `ent/schema/user.go` though the `Fields()` and `Edges()` will be left empty.
3. Populate the `Fields()` and optionally the `Edges()` (which are the relationships to other entity types).
4. When done, generate all code by executing `make ent-gen`.

The generated code is extremely flexible and impressive. An example to highlight this is one used within this application:

```go
entity, err := c.ORM.PasswordToken.
    Query().
    Where(passwordtoken.ID(tokenID)).
    Where(passwordtoken.HasUserWith(user.ID(userID))).
    Where(passwordtoken.CreatedAtGTE(expiration)).
    Only(ctx.Request().Context())
```

This executes a database query to return the _password token_ entity with a given ID that belong to a user with a given ID and has a _created at_ timestamp field that is greater than or equal to a given time.

## Sessions

Sessions are provided and handled via [Gorilla sessions](https://github.com/gorilla/sessions) and configured as middleware in the router located at `pkg/handlers/router.go`. Session data is currently stored in cookies but there are many [options](https://github.com/gorilla/sessions#store-implementations) available if you wish to use something else.

Here's a simple example of loading data from a session and saving new values:

```go
func SomeFunction(ctx echo.Context) error {
    sess, err := session.Get(ctx, "some-session-key")
    if err != nil {
        return err
    }
    sess.Values["hello"] = "world"
    sess.Values["isSomething"] = true
    return sess.Save(ctx.Request(), ctx.Response())
}
```

### Encryption

Session data is encrypted for security purposes. The encryption key is stored in [configuration](#configuration) at `Config.App.EncryptionKey`. While the default is fine for local development, it is **imperative** that you change this value for any live environment otherwise session data can be compromised.

## Authentication

Included are standard authentication features you expect in any web application. Authentication functionality is bundled as a _Service_ within `services/AuthClient` and added to the `Container`. If you wish to handle authentication in a different manner, you could swap this client out or modify it as needed.

Authentication currently requires [sessions](#sessions) and the session middleware.

### Login / Logout

The `AuthClient` has methods `Login()` and `Logout()` to log a user in or out. To track a user's authentication state, data is stored in the session including the user ID and authentication status.

Prior to logging a user in, the method `CheckPassword()` can be used to determine if a user's password matches the hash stored in the database and on the `User` entity.

Routes are provided for the user to login and logout at `user/login` and `user/logout`.

### Forgot password

Users can reset their password in a secure manner by issuing a new password token via the method `GeneratePasswordResetToken()`. This creates a new `PasswordToken` entity in the database belonging to the user. The actual token itself, however, is not stored in the database for security purposes. It is only returned via the method so it can be used to build the reset URL for the email. Rather, a hash of the token is stored, using `bcrypt` the same package used to hash user passwords. The reason for doing this is the same as passwords. You do not want to store a plain-text value in the database that can be used to access an account.

Tokens have a configurable expiration. By default, they expire within 1 hour. This can be controlled in the `config` package. The expiration of the token is not stored in the database, but rather is used only when tokens are loaded for potential usage. This allows you to change the expiration duration and affect existing tokens.

Since the actual tokens are not stored in the database, the reset URL must contain the user and password token ID. Using that, `GetValidPasswordToken()` will load a matching, non-expired _password token_ entity belonging to the user, and use `bcrypt` to determine if the token in the URL matches stored hash of the password token entity.

Once a user claims a valid password token, all tokens for that user should be deleted using `DeletePasswordTokens()`.

Routes are provided to request a password reset email at `user/password` and to reset your password at `user/password/reset/token/:user/:password_token/:token`.

### Registration

The actual registration of a user is not handled within the `AuthClient` but rather just by creating a `User` entity. When creating a user, use `HashPassword()` to create a hash of the user's password, which is what will be stored in the database.

A route is provided for the user to register at `user/register`.

### Authenticated user

The `AuthClient` has two methods available to get either the `User` entity or the ID of the user currently logged in for a given request. Those methods are `GetAuthenticatedUser()` and `GetAuthenticatedUserID()`.

#### Middleware

Registered for all routes is middleware that will load the currently logged in user entity and store it within the request context. The middleware is located at `middleware.LoadAuthenticatedUser()` and, if authenticated, the `User` entity is stored within the context using the key `context.AuthenticatedUserKey`.

If you wish to require either authentication or non-authentication for a given route, you can use either `middleware.RequireAuthentication()` or `middleware.RequireNoAuthentication()`.

### Email verification

Most web applications require the user to verify their email address (or other form of contact information). The `User` entity has a field `Verified` to indicate if they have verified themself. When a user successfully registers, an email is sent to them containing a link with a token that will verify their account when visited. This route is currently accessible at `/email/verify/:token` and handled by `pkg/handlers/auth.go`.

There is currently no enforcement that a `User` must be verified in order to access the application. If that is something you desire, it will have to be added in yourself. It was not included because you may want partial access of certain features until the user verifies; or no access at all.

Verification tokens are [JSON Web Tokens](https://jwt.io/) generated and processed by the [jwt](https://github.com/golang-jwt/jwt) module. The tokens are _signed_ using the encryption key stored in [configuration](#configuration) (`Config.App.EncryptionKey`). **It is imperative** that you override this value from the default in any live environments otherwise the data can be comprimised. JWT was chosen because they are secure tokens that do not have to be stored in the database, since the tokens contain all of the data required, including built-in expirations. These were not chosen for password reset tokens because JWT cannot be withdrawn once they are issued which poses a security risk. Since these tokens do not grant access to an account, the ability to withdraw the tokens is not needed.

By default, verification tokens expire 12 hours after they are issued. This can be changed in configuration at `Config.App.EmailVerificationTokenExpiration`. There is currently not a route or form provided to request a new link.

Be sure to review the [email](#email) section since actual email sending is not fully implemented.

To generate a new verification token, the `AuthClient` has a method `GenerateEmailVerificationToken()` which creates a token for a given email address. To verify the token, pass it in to `ValidateEmailVerificationToken()` which will return the email address associated with the token and an error if the token is invalid.

## Routes

The router functionality is provided by [Echo](https://echo.labstack.com/guide/routing/) and constructed within via the `BuildRouter()` function inside `pkg/handlers/router.go`. Since the _Echo_ instance is a _Service_ on the `Container` which is passed in to `BuildRouter()`, middleware and routes can be added directly to it.

### Custom middleware

By default, a middleware stack is included in the router that makes sense for most web applications. Be sure to review what has been included and what else is available within _Echo_ and the other projects mentioned.

A `middleware` package is included which you can easily add to along with the custom middleware provided.

### Handlers

A `Handler` is a simple type that handles one or more of your routes and allows you to group related routes together (ie, authentication). All provided handlers are located in `pkg/handlers`. _Handlers_ also handle self-registering their routes with the router.

#### Example

The provided patterns are not required, but were designed to make development as easy as possible.

For this example, we'll create a new handler which includes a GET and POST route and uses the ORM. Start by creating a file at `pkg/handlers/example.go`.

1) Define the handler type:

```go
type Example struct {
    orm *ent.Client
    *services.TemplateRenderer
}
```

2) Register the handler so the router automatically includes it

```go
func init() {
    Register(new(Example))
}
```

3) Initialize the handler (and inject any required dependencies from the _Container_). This will be called automatically.

```go
func (e *Example) Init(c *services.Container) error {
    e.TemplateRenderer = c.TemplateRenderer
    e.orm = c.ORM
    return nil
}
```

4) Declare the routes

**It is highly recommended** that you provide a `Name` for your routes. Most methods on the back and frontend leverage the route name and parameters in order to generate URLs.

```go
func (e *Example) Routes(g *echo.Group) {
    g.GET("/example", e.Page).Name = "example"
    g.POST("/example", c.PageSubmit).Name = "example.submit"
}
```

5) Implement your routes

```go
func (e *Example) Page(ctx echo.Context) error {
    // add your code here
}

func (e *Example) PageSubmit(ctx echo.Context) error {
    // add your code here
}
```

### Errors

Routes can return errors to indicate that something wrong happened. Ideally, the error is of type `*echo.HTTPError` to indicate the intended HTTP response code. You can use `return echo.NewHTTPError(http.StatusInternalServerError)`, for example. If an error of a different type is returned, an _Internal Server Error_ is assumed.

The [error handler](https://echo.labstack.com/guide/error-handling/) is set to the provided `Handler` in `pkg/handlers/error.go` in the `BuildRouter()` function. That means that if any middleware or route return an error, the request gets routed there. This route conveniently constructs and renders a `Page` which uses the template `templates/pages/error.gohtml`. The status code is passed to the template so you can easily alter the markup depending on the error type.

### Redirects

The `pkg/redirect` package makes it easy to perform redirects, especially if you provide names for your routes. The `Redirect` type provides the ability to chain redirect options and also supports automatically handling HTMX redirects for boosted requests.

For example, if your route name is `user_profile` with a URL pattern of `/user/profile/:id`, you can perform a redirect by doing:

```go
return redirect.New(ctx).
    Route("user_profile").
    Params(userID).
    Query(queryParams).
    Go()
```

### Testing

Since most of your web application logic will live in your routes, being able to easily test them is important. The following aims to help facilitate that.

The test setup and helpers reside in `pkg/handlers/router_test.go`.

Only a brief example of route tests were provided in order to highlight what is available. Adding full tests did not seem logical since these routes will most likely be changed or removed in your project.

#### HTTP server

When the route tests initialize, a new `Container` is created which provides full access to all of the _Services_ that will be available during normal application execution. Also provided is a test HTTP server with the router added. This means your tests can make requests and expect responses exactly as the application would behave outside of tests. You do not need to mock the requests and responses.

#### Request / Response helpers

With the test HTTP server setup, test helpers for making HTTP requests and evaluating responses are made available to reduce the amount of code you need to write. See `httpRequest` and `httpResponse` within `pkg/handlers/router_test.go`.

Here is an example how to easily make a request and evaluate the response:

```go
func TestAbout_Get(t *testing.T) {
    doc := request(t).
        setRoute("about").
        get().
        assertStatusCode(http.StatusOK).
        toDoc()
}
```

#### Goquery

A helpful, included package to test HTML markup from HTTP responses is [goquery](https://github.com/PuerkitoBio/goquery). This allows you to use jQuery-style selectors to parse and extract HTML values, attributes, and so on.

In the example above, `toDoc()` will return a `*goquery.Document` created from the HTML response of the test HTTP server.

Here is a simple example of how to use it, along with [testify](https://github.com/stretchr/testify) for making assertions:

```go
h1 := doc.Find("h1.title")
assert.Len(t, h1.Nodes, 1)
assert.Equal(t, "About", h1.Text())
```

## Pages

The `Page` is the major building block of your `Handler` responses. It is a _struct_ type located at `pkg/page/page.go`. The concept of the `Page` is that it provides a consistent structure for building responses and transmitting data and functionality to the templates. Pages are rendered with the `TemplateRenderer`.

All example routes provided construct and _render_ a `Page`. It's recommended that you review both the `Page` and the example routes as they try to illustrate all included functionality.

As you develop your application, the `Page` can be easily extended to include whatever data or functions you want to provide to your templates.

Initializing a new page is simple:

```go
func (c *home) Get(ctx echo.Context) error {
    p := page.New(ctx)
}
```

Using the `echo.Context`, the `Page` will be initialized with the following fields populated:

- `Context`: The passed in _context_
- `Path`: The requested URL path
- `URL`: The requested URL
- `StatusCode`: Defaults to 200
- `Pager`: Initialized `Pager` (see below)
- `RequestID`: The request ID, if the middleware is being used
- `IsHome`: If the request was for the homepage
- `IsAuth`: If the user is authenticated
- `AuthUser`: The logged in user entity, if one
- `CSRF`: The CSRF token, if the middleware is being used
- `HTMX.Request`: Data from the HTMX headers, if HTMX made the request (see below)

### Flash messaging

Flash messaging functionality is provided within the `msg` package. It is used to provide one-time status messages to users.

Flash messaging requires that [sessions](#sessions) and the session middleware are in place since that is where the messages are stored.

#### Creating messages

There are four types of messages, and each can be created as follows:
- Success: `msg.Success(ctx echo.Context, message string)`
- Info: `msg.Info(ctx echo.Context, message string)`
- Warning: `msg.Warning(ctx echo.Context, message string)`
- Danger: `msg.Danger(ctx echo.Context, message string)`

The _message_ string can contain HTML.

#### Rendering messages

When a flash message is retrieved from storage in order to be rendered, it is deleted from storage so that it cannot be rendered again.

The `Page` has a method that can be used to fetch messages for a given type from within the template: `Page.GetMessages(typ msg.Type)`. This is used rather than the _funcmap_ because the `Page` contains the request context which is required in order to access the session data. Since the `Page` is the data destined for the templates, you can use: `{{.GetMessages "success"}}` for example.

To make things easier, a template _component_ is already provided, located at `templates/components/messages.gohtml`. This will render all messages of all types simply by using `{{template "messages" .}}` either within your page or layout template.

### Pager

A very basic mechanism is provided to handle and facilitate paging located in `pkg/page/pager.go`. When a `Page` is initialized, so is a `Pager` at `Page.Pager`. If the requested URL contains a `page` query parameter with a numeric value, that will be set as the page number in the pager.

During initialization, the _items per page_ amount will be set to the default, controlled via constant, which has a value of 20. It can be overridden by changing `Pager.ItemsPerPage` but should be done before other values are set in order to not provide incorrect calculations.

Methods include:

- `SetItems(items int)`: Set the total amount of items in the entire result-set
- `IsBeginning()`: Determine if the pager is at the beginning of the pages
- `IsEnd()`: Determine if the pager is at the end of the pages
- `GetOffset()`: Get the offset which can be useful is constructing a paged database query

There is currently no template (yet) to easily render a pager.

### CSRF

By default, all non GET requests will require a CSRF token be provided as a form value. This is provided by middleware and can be adjusted or removed in the router.

The `Page` will contain the CSRF token for the given request. There is a CSRF helper component template which can be used to easily render a hidden form element in your form which will contain the CSRF token and the proper element name. Simply include `{{template "csrf" .}}` within your form.

### Automatic template parsing

Dealing with templates can be quite tedious and annoying so the `Page` aims to make it as simple as possible with the help of the [template renderer](#template-renderer). To start, templates for _pages_ are grouped in the following directories within the `templates` directory:

- `layouts`: Base templates that provide the entire HTML wrapper/layout. This template should include a call to `{{template "content" .}}` to render the content of the `Page`.
- `pages`: Templates that are specific for a given route/page. These must contain `{{define "content"}}{{end}}` which will be injected in to the _layout_ template.
- `components`: A shared library of common components that the layout and base template can leverage.

Specifying which templates to render for a given `Page` is as easy as:

```go
page.Name = "home"
page.Layout = "main"
```

That alone will result in the following templates being parsed and executed when the `Page` is rendered:

1) `layouts/main.gohtml` as the base template
2) `pages/home.gohtml` to provide the `content` template for the layout
3) All template files located within the `components` directory
4) The entire [funcmap](#funcmap)

The [template renderer](#template-renderer) also provides caching and local hot-reloading.

### Cached responses

A `Page` can have cached enabled just by setting `Page.Cache.Enabled` to `true`. The `TemplateRenderer` will automatically handle caching the HTML output, headers and status code. Cached pages are stored using a key that matches the full request URL and [middleware](#cache-middleware) is used to serve it on matching requests.

By default, the cache expiration time will be set according to the configuration value located at `Config.Cache.Expiration.Page` but it can be set per-page at `Page.Cache.Expiration`.

#### Cache tags

You can optionally specify cache tags for the `Page` by setting a slice of strings on `Page.Cache.Tags`. This provides the ability to build in cache invalidation logic in your application driven by events such as entity operations, for example.

You can use the [cache client](#cache) on the `Container` to easily [flush cache tags](#flush-tags), if needed.

#### Cache middleware

Cached pages are served via the middleware `ServeCachedPage()` in the `middleware` package.

The cache is bypassed if the requests meet any of the following criteria:
1) Is not a GET request
2) Is made by an authenticated user

Cached pages are looked up for a key that matches the exact, full URL of the given request.

### Data

The `Data` field on the `Page` is of type `any` and is what allows your route to pass whatever it requires to the templates, alongside the `Page` itself.

### Forms

The `Form` field on the `Page` is similar to the `Data` field, but it's meant to store a struct that represents a form being rendered on the page.

An example of this pattern is:

```go
type ContactForm struct {
    Email      string `form:"email" validate:"required,email"`
    Message    string `form:"message" validate:"required"`
    form.Submission
}
```

Embedding `form.Submission` satisfies the `form.Form` interface and makes dealing with submissions and validation extremely easy.

Then in your page:

```go
p := page.New(ctx)
p.Form = form.Get[ContactForm](ctx)
```

This will either initialize a new form to be rendered, or load one previously stored in the context (ie, if it was already submitted). How the _form_ gets populated with values so that your template can render them is covered in the next section.

#### Submission processing

Form submission processing is made extremely simple by leveraging functionality provided by [Echo binding](https://echo.labstack.com/guide/binding/), [validator](https://github.com/go-playground/validator) and the `Submission` struct located in `pkg/form/submission.go`.

Using the example form above, this is all you would have to do within the _POST_ callback for your route:

Start by submitting the form along with the request context. This will:
1. Store a pointer to the form so that your _GET_ callback can access the form values (shown previously). That allows the form to easily be re-rendered with any validation errors it may have as well as the values that were provided.
2. Parse the input in the _POST_ data to map to the struct so the fields becomes populated. This uses the `form` struct tags to map form input values to the struct fields.
3. Validate the values in the struct fields according to the rules provided in the optional `validate` struct tags.

```go
var input ContactForm

err := form.Submit(ctx, &input)
```

Check the error returned, and act accordingly. For example:
```go
switch err.(type) {
case nil:
    // All good!
case validator.ValidationErrors:
    // The form input was not valid, so re-render the form
    return c.Page(ctx)
default:
    // Request failed, show the error page
    return err
}
```

And finally, your template:
```html
<form id="contact" method="post" hx-post="{{url "contact.post"}}">
    <input id="email" name="email" type="email" class="input" value="{{.Form.Email}}">
    <input id="message" name="message" type="text" class="input" value="{{.Form.Message}}">
</form
```

#### Inline validation

The `Submission` makes inline validation easier because it will store all validation errors in a map, keyed by the form struct field name. It also contains helper methods that your templates can use to provide classes and extract the error messages.

While [validator](https://github.com/go-playground/validator) is a great package that is used to validate based on struct tags, the downside is that the messaging, by default, is not very human-readable or easy to override. Within `Submission.setErrorMessages()` the validation errors are converted to more readable messages based on the tag that failed validation. Only a few tags are provided as an example, so be sure to expand on that as needed.

To provide the inline validation in your template, there are two things that need to be done.

First, include a status class on the element so it will highlight green or red based on the validation:
```html
<input id="email" name="email" type="email" class="input {{.Form.GetFieldStatusClass "Email"}}" value="{{.Form.Email}}">
```

Second, render the error messages, if there are any for a given field:
```go
{{template "field-errors" (.Form.GetFieldErrors "Email")}}
```

### Headers

HTTP headers can be set either via the `Page` or the _context_:

```go
p := page.New(ctx)
p.Headers["HeaderName"] = "header-value"
```

```go
ctx.Response().Header().Set("HeaderName", "header-value")
```

### Status code

The HTTP response status code can be set either via the `Page` or the _context_:

```go
p := page.New(ctx)
p.StatusCode = http.StatusTooManyRequests
```

```go
ctx.Response().Status = http.StatusTooManyRequests
```

### Metatags

The `Page` provides the ability to set basic HTML metatags which can be especially useful if your web application is publicly accessible. Only fields for the _description_ and _keywords_ are provided but adding additional fields is very easy.

```go
p := page.New(ctx)
p.Metatags.Description = "The page description."
p.Metatags.Keywords = []string{"Go", "Software"}
```

A _component_ template is included to render metatags in `core.gohtml` which can be used by adding `{{template "metatags" .}}` to your _layout_.

### URL and link generation

Generating URLs in the templates is made easy if you follow the [routing patterns](#patterns) and provide names for your routes. Echo provides a `Reverse` function to generate a route URL with a given route name and optional parameters. This function is made accessible to the templates via _funcmap_ function `url`.

As an example, if you have route such as:
```go
e.GET("/user/profile/:user", handler.Get).Name = "user_profile"
```

And you want to generate a URL in the template, you can:
```go
{{url "user_profile" 1}
```

Which will generate: `/user/profile/1`

There is also a helper function provided in the [funcmap](#funcmap) to generate links which has the benefit of adding an _active_ class if the link URL matches the current path. This is especially useful for navigation menus.

```go
{{link (url "user_profile" .AuthUser.ID) "Profile" .Path "extra-class"}}
```

Will generate:
```html
<a href="/user/profile/1" class="is-active extra-class">Profile</a>
```
Assuming the current _path_ is `/user/profile/1`; otherwise the `is-active` class will be excluded.

### HTMX support

[HTMX](https://htmx.org/) is an awesome JavaScript library allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext.

Many examples of its usage are available in the included examples:
- All navigation links use [boost](https://htmx.org/docs/#boosting) which dynamically replaces the page content with an AJAX request, providing a SPA-like experience.
- All forms use either [boost](https://htmx.org/docs/#boosting) or [hx-post](https://htmx.org/docs/#triggers) to submit via AJAX.
- The mock search autocomplete modal uses [hx-get](https://htmx.org/docs/#targets) to fetch search results from the server via AJAX and update the UI.
- The mock posts on the homepage/dashboard use [hx-get](https://htmx.org/docs/#targets) to fetch and page posts via AJAX.

All of this can be easily accomplished without writing any JavaScript at all.

Another benefit of [HTMX](https://htmx.org/) is that it's completely backend-agnostic and does not require any special tools or integrations on the backend. But to make things easier, included is a small package to read and write [HTTP headers](https://htmx.org/docs/#requests) that HTMX uses to communicate additional information and commands.

The `htmx` package contains the headers for the _request_ and _response_. When a `Page` is initialized, `Page.HTMX.Request` will also be initialized and populated with the headers that HTMX provides, if HTMX made the request. This allows you to determine if HTMX is making the given request and what exactly it is doing, which could be useful both in your _route_ as well as your _templates_.

If you need to set any HTMX headers in your `Page` response, this can be done by altering `Page.HTMX.Response`.

#### Layout template override

To facilitate easy partial rendering for HTMX requests, the `Page` will automatically change your _Layout_ template to use `htmx.gohtml`, which currently only renders `{{template "content" .}}`. This allows you to use an HTMX request to only update the content portion of the page, rather than the entire HTML.

This override only happens if the HTMX request being made is **not a boost** request because **boost** requests replace the entire `body` element so there is no need to do a partial render.

#### Conditional processing / rendering

Since HTMX communicates what it is doing with the server, you can use the request headers to conditionally process in your _route_ or render in your _template_, if needed. If your routes aren't doing multiple things, you may not need this, but it's worth knowing how flexible you can be.

A simple example of this:

```go
if page.HTMX.Request.Target == "search" {
    // You know this request HTMX is fetching content just for the #search element
}
```

```go
{{if eq .HTMX.Request.Target "search"}}
    // Render content for the #search element
{{end}}
```

#### CSRF token

If [CSRF](#csrf) protection is enabled, the token value will automatically be passed to HTMX to be included in all non-GET requests. This is done in the `footer` template by leveraging HTMX [events](https://htmx.org/reference/#events).

### Rendering the page

Once your `Page` is fully built, rendering it via the embedded `TemplateRenderer` in your _handler_ can be done simply by calling `RenderPage()`:

```go
func (c *home) Get(ctx echo.Context) error {
    p := page.New(ctx)
    p.Layout = templates.LayoutMain
    p.Name = templates.PageHome
    return c.RenderPage(ctx, p)
}
```

## Template renderer

The _template renderer_ is a _Service_ on the `Container` that aims to make template parsing and rendering easy and flexible. It is the mechanism that allows the `Page` to do [automatic template parsing](#automatic-template-parsing). The standard `html/template` is still the engine used behind the scenes. The code can be found in `pkg/services/template_renderer.go`.

Here is an example of a complex rendering that uses multiple template files as well as an entire directory of template files:

```go
buf, err = c.TemplateRenderer.
    Parse().
    Group("page").
    Key("home").
    Base("main").
    Files("layouts/main", "pages/home").
    Directories("components").
    Execute(data)
```

This will do the following:
- [Cache](#caching) the parsed template with a _group_ of `page` and _key_ of `home` so this parse only happens once
- Set the _base template file_ as `main`
- Include the templates `templates/layout/main.gohtml` and `templates/pages/home.gohtml`
- Include all templates located within the directory `templates/components`
- Include the [funcmap](#funcmap)
- Execute the parsed template with `data` being passed in to the templates

Using the example from the [page rendering](#rendering-the-page), this is will execute:

```go
buf, err = c.TemplateRenderer.
    Parse().
    Group("page").
    Key(page.Name).
    Base(page.Layout).
    Files(
        fmt.Sprintf("layouts/%s", page.Layout),
        fmt.Sprintf("pages/%s", page.Name),
    ).
    Directories("components").
    Execute(page)
```

If you have a need to _separately_ parse and cache the templates then later execute, you can separate the operations:

```go
_, err := c.TemplateRenderer.
    Parse().
    Group("my-group").
    Key("my-key").
    Base("auth").
    Files("layouts/auth", "pages/login").
    Directories("components").
    Store()
```

```go
tpl, err := c.TemplateRenderer.Load("my-group", "my-key")
buf, err := tpl.Execute(data)
```

### Custom functions

All templates will be parsed with the [funcmap](#funcmap) so all of your custom functions as well as the functions provided by [sprig](https://github.com/Masterminds/sprig) will be available.

### Caching

Parsed templates will be cached within a `sync.Map` so the operation will only happen once per cache _group_ and _ID_. Be careful with your cache _group_ and _ID_ parameters to avoid collisions.

### Hot-reload for development

If the current [environment](#environments) is set to `config.EnvLocal`, which is the default, the cache will be bypassed and templates will be parsed every time they are requested. This allows you to have hot-reloading without having to restart the application so you can see your HTML changes in the browser immediately.

### File configuration

To make things easier and less repetitive, parameters given to the _template renderer_ must not include the `templates` directory or the template file extensions. The file extension is stored as a constant (`TemplateExt`) within the `config` package.

## Funcmap

The `funcmap` package provides a _function map_ (`template.FuncMap`) which will be included for all templates rendered with the [template renderer](#template-renderer). Aside from a few custom functions, [sprig](https://github.com/Masterminds/sprig) is included which provides over 100 commonly used template functions. The full list is available [here](http://masterminds.github.io/sprig/).

To include additional custom functions, add to the map in `NewFuncMap()` and define the function in the package. It will then become automatically available in all templates.

## Cache

As previously mentioned, the default cache implementation is a simple in-memory store, backed by [otter](https://github.com/maypok86/otter), a lockless cache that uses [S3-FIFO](https://s3fifo.com/) eviction. The `Container` houses a `CacheClient` which is a useful, wrapper to interact with the cache (see examples below). Within the `CacheClient` is the underlying store interface `CacheStore`. If you wish to use a different store, such as Redis, and want to keep using the `CacheClient`, simply implement the `CacheStore` interface with a Redis library and adjust the `Container` initialization to use that.

The built-in usage of the cache is currently only for optional [page caching](#cached-responses) and a simple example route located at `/cache` where you can set and view the value of a given cache entry.

Since the current cache is in-memory, there's no need to adjust the `Container` during tests. When this project used Redis, the configuration had a separate database that would be used strictly for tests to avoid writing to your primary database. If you need that functionality, it is easy to add back in.

### Set data

**Set data with just a key:**

```go
err := c.Cache.
    Set().
    Key("my-key").
    Data(myData).
    Expiration(time.Hour * 2).
    Save(ctx)
```

**Set data within a group:**

```go
err := c.Cache.
    Set().
    Group("my-group").
    Key("my-key").
    Expiration(time.Hour * 2).
    Data(myData).
    Save(ctx)
```

**Include cache tags:**

```go
err := c.Cache.
    Set().
    Key("my-key").
    Tags("tag1", "tag2").
    Expiration(time.Hour * 2).
    Data(myData).
    Save(ctx)
```

### Get data

```go
data, err := c.Cache.
    Get().
    Group("my-group").
    Key("my-key").
    Fetch(ctx)
```

### Flush data

```go
err := c.Cache.
    Flush().
    Group("my-group").
    Key("my-key").
    Execute(ctx)
```

### Flush tags

This will flush all cache entries that were tagged with the given tags.

```go
err := c.Cache.
    Flush().
    Tags("tag1", "tag2").
    Execute(ctx)
```

### Tagging

As shown in the previous examples, cache tags were provided because they can be convenient. However, maintaining them comes at a cost and it may not be a good fit for your application depending on your needs. When including tags, the `CacheClient` must lock in order to keep the tag index in sync. And since the tag index cannot support eviction, since that could result in a flush call not actually flushing the tag's keys, the maps that provide the index do not have a size limit. See the code for more details.

## Tasks

Tasks are queued operations to be executed in the background, either immediately, at a specfic time, or after a given amount of time has passed. Some examples of tasks could be long-running operations, bulk processing, cleanup, notifications, etc.

Since we're already using [SQLite](https://sqlite.org/) for our database, it's available to act as a persistent store for queued tasks so that tasks are never lost, can be retried until successful, and their concurrent execution can be managed. [Backlite](https://github.com/mikestefanello/backlite) is the library chosen to interface with [SQLite](https://sqlite.org/) and handle queueing tasks and processing them asynchronously. I wrote that specifically to address the requirements I wanted to satisfy for this project.

To make things easy, the _Backlite_ client is provided as a _Service_ on the `Container` which allows you to register queues and add tasks.

Configuration for the _Backlite_ client is exposed through the app's yaml configuration. The required database schema will be automatically installed when the app starts.

### Queues

A full example of a queue implementation can be found in `pkg/tasks` with an interactive form to create a task and add to the queue at `/task` (see `pkg/handlers/task.go`). Also refer to the [Backlite](https://github.com/mikestefanello/backlite) documentation for reference and examples.

See `pkg/tasks/register.go` for a simple way to register all of your queues and to easily pass the `Container` to them so the queue processor callbacks have access to all of your app's dependencies.

### Dispatcher

The _task dispatcher_ is what manages the worker pool used for executing tasks in the background. It monitors incoming and scheduled tasks and handles sending them to the pool for execution by the queue's processor callback. This must be started in order for this to happen. In `cmd/web/main.go`, the _task dispatcher_ is automatically started when the app starts via:

```go
c.Tasks.Start(ctx)
```

The app [configuration](#configuration) contains values to configure the client and dispatcher including how many goroutines to use, when to release stuck tasks back into the queue, and how often to cleanup retained tasks in the database.

When the app is shutdown, the dispatcher is given 10 seconds to wait for any in-progress tasks to finish execution. This can be changed in `cmd/web/main.go`.

## Cron

By default, no cron solution is provided because it's very easy to add yourself if you need this. You can either use a [ticker](https://pkg.go.dev/time#Ticker) or a [library](https://github.com/robfig/cron).

## Static files

Static files are currently configured in the router (`pkg/handler/router.go`) to be served from the `static` directory. If you wish to change the directory, alter the constant `config.StaticDir`. The URL prefix for static files is `/files` which is controlled via the `config.StaticPrefix` constant.

### Cache control headers

Static files are grouped separately so you can apply middleware only to them. Included is a custom middleware to set cache control headers (`middleware.CacheControl`) which has been added to the static files router group.

The cache max-life is controlled by the configuration at `Config.Cache.Expiration.StaticFile` and defaults to 6 months.

### Cache-buster

While it's ideal to use cache control headers on your static files so browsers cache the files, you need a way to bust the cache in case the files are changed. In order to do this, a function is provided in the [funcmap](#funcmap) to generate a static file URL for a given file that appends a cache-buster query. This query string is randomly generated and persisted until the application restarts.

For example, to render a file located in `static/picture.png`, you would use:
```html
<img src="{{file "picture.png"}}"/>
```

Which would result in:
```html
<img src="/files/picture.png?v=9fhe73kaf3"/>
```

Where `9fhe73kaf3` is the randomly-generated cache-buster.

## Email

An email client was added as a _Service_ to the `Container` but it is just a skeleton without any actual email-sending functionality. The reason is because there are a lot of ways to send email and most prefer using a SaaS solution for that. That makes it difficult to provide a generic solution that will work for most applications.

The structure in the client (`MailClient`) makes composing emails very easy and you have the option to construct the body using either a simple string or with a template by leveraging the [template renderer](#template-renderer). The standard library can be used if you wish to send email via SMTP and most SaaS providers have a Go package that can be used if you choose to go that direction. **You must** finish the implementation of `MailClient.send`.

The _from_ address will default to the configuration value at `Config.Mail.FromAddress`. This can be overridden per-email by calling `From()` on the email and passing in the desired address.

See below for examples on how to use the client to compose emails.

**Sending with a string body**:

```go
err = c.Mail.
    Compose().
    To("hello@example.com").
    Subject("Welcome!").
    Body("Thank you for registering.").
    Send(ctx)
```

**Sending with a template body**:

```go
err = c.Mail.
    Compose().
    To("hello@example.com").
    Subject("Welcome!").
    Template("welcome").
    TemplateData(templateData).
    Send(ctx)
```

This will use the template located at `templates/emails/welcome.gohtml` and pass `templateData` to it.

## HTTPS

By default, the application will not use HTTPS but it can be enabled easily. Just alter the following configuration:

- `Config.HTTP.TLS.Enabled`: `true`
- `Config.HTTP.TLS.Certificate`: Full path to the certificate file
- `Config.HTTP.TLS.Key`: Full path to the key file

To use _Let's Encrypt_ follow [this guide](https://echo.labstack.com/cookbook/auto-tls/#server).

## Logging

By default, the [Echo logger](https://echo.labstack.com/guide/customization/#logging) is not used for the following reasons:

1) It does not support structured logging, which makes it difficult to deal with variables, especially if you intend to store a logger in context with pre-populated attributes.
2) The upcoming v5 release of Echo will not contain a logger.
3) It should be easy to use whatever logger you prefer (even if that is Echo's logger).

The provided implementation uses the relatively new [log/slog](https://go.dev/blog/slog) library which was added to Go in version 1.21 but swapping that out for whichever you prefer is very easy.

### Context

The simple `pkg/log` package provides the ability to set and get a logger from the Echo context. This is especially useful when you use the provided logger middleware (see below). If you intend to use a different logger, modify these methods to receive and return the logger of your choice. 

### Usage

Adding a logger to the context:
```go
logger := slog.New(logHandler).With("id", requestId)
log.Set(ctx, logger)
```

Access and use the logger:
```go
func (h *handler) Page(ctx echo.Context) error {
    log.Ctx(ctx).Info("send a message to the log",
      "value_one", valueOne,
      "value_two", valueTwo,
    )
}
```

### Log level

When the _Container_ configuration is initialized (`initConfig()`), the `slog` default log level is set based on the environment. `INFO` is used for production and `DEBUG` for everything else.

### Middleware

The `SetLogger()` middleware has been added to the router which sets an initialized logger on the request context. It's recommended that this remains after Echo's `RequestID()` middleware because it will add the request ID to the logger which means that all logs produced for that given request will contain the same ID, so they can be linked together. If you want to include more attributes on all request logs, set those fields here.

The `LogRequest()` middleware is a replacement for Echo's `Logger()` middleware which produces a log of every request made, but uses our logger rather than Echo's.

```
2024/06/15 09:07:11 INFO GET /contact request_id=gNblvugTKcyLnBYPMPTwMPEqDOioVLKp ip=::1 host=localhost:8000 referer="" status=200 bytes_in=0 bytes_out=5925 latency=107.527803ms
```

## Roadmap

Future work includes but is not limited to:

- Flexible pager templates
- Expanded HTMX examples and integration
- Admin section

## Credits

Thank you to all of the following amazing projects for making this possible.

- [alpinejs](https://github.com/alpinejs/alpine)
- [backlite](https://github.com/mikestefanello/backlite)
- [bulma](https://github.com/jgthms/bulma)
- [echo](https://github.com/labstack/echo)
- [ent](https://github.com/ent/ent)
- [go](https://go.dev/)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [goquery](https://github.com/PuerkitoBio/goquery)
- [htmx](https://github.com/bigskysoftware/htmx)
- [jwt](https://github.com/golang-jwt/jwt)
- [otter](https://github.com/maypok86/otter)
- [sessions](https://github.com/gorilla/sessions)
- [sprig](https://github.com/Masterminds/sprig)
- [sqlite](https://sqlite.org/)
- [testify](https://github.com/stretchr/testify)
- [validator](https://github.com/go-playground/validator)
- [viper](https://github.com/spf13/viper)
````markdown

`cmd/web/main.go`

```go
package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/mikestefanello/pagoda/pkg/handlers"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/tasks"
)

func main() {
	// Start a new container
	c := services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	// Build the router
	if err := handlers.BuildRouter(c); err != nil {
		log.Fatalf("failed to build the router: %v", err)
	}

	// Start the server
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WriteTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}

		if c.Config.HTTP.TLS.Enabled {
			certs, err := tls.LoadX509KeyPair(c.Config.HTTP.TLS.Certificate, c.Config.HTTP.TLS.Key)
			if err != nil {
				log.Fatalf("cannot load TLS certificate: %v", err)
			}

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := c.Web.StartServer(&srv); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Register all task queues
	tasks.Register(c)

	// Start the task runner to execute queued tasks
	c.Tasks.Start(context.Background())

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit

	// Shutdown both the task runner and web server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.Tasks.Stop(ctx)
	}()

	go func() {
		defer wg.Done()
		if err := c.Web.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
```go

`config/config.go`

```go
package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	// TemplateExt stores the extension used for the template files
	TemplateExt = ".gohtml"

	// StaticDir stores the name of the directory that will serve static files
	StaticDir = "static"

	// StaticPrefix stores the URL prefix used when serving static files
	StaticPrefix = "files"
)

type environment string

const (
	// EnvLocal represents the local environment
	EnvLocal environment = "local"

	// EnvTest represents the test environment
	EnvTest environment = "test"

	// EnvDevelop represents the development environment
	EnvDevelop environment = "dev"

	// EnvStaging represents the staging environment
	EnvStaging environment = "staging"

	// EnvQA represents the qa environment
	EnvQA environment = "qa"

	// EnvProduction represents the production environment
	EnvProduction environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env environment) {
	if err := os.Setenv("PAGODA_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration
	Config struct {
		HTTP     HTTPConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
		Tasks    TasksConfig
		Mail     MailConfig
		OAuth    OAuthConfig
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Hostname     string
		Port         uint16
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
		TLS          struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	// AppConfig stores application configuration
	AppConfig struct {
		Name          string
		Environment   environment
		EncryptionKey string
		Timeout       time.Duration
		PasswordToken struct {
			Expiration time.Duration
			Length     int
		}
		EmailVerificationTokenExpiration time.Duration
	}

	// CacheConfig stores the cache configuration
	CacheConfig struct {
		Capacity   int
		Expiration struct {
			StaticFile time.Duration
			Page       time.Duration
		}
	}

	// DatabaseConfig stores the database configuration
	DatabaseConfig struct {
		Driver         string
		Connection     string
		TestConnection string
	}

	// TasksConfig stores the tasks configuration
	TasksConfig struct {
		Goroutines      int
		ReleaseAfter    time.Duration
		CleanupInterval time.Duration
	}

	// MailConfig stores the mail configuration
	MailConfig struct {
		Hostname    string
		Port        uint16
		User        string
		Password    string
		FromAddress string
	}

	// OAuth stores OAuth2 configuration
	OAuthConfig struct {
		Google struct {
			ClientID     string
			ClientSecret string
			RedirectURL  string
		}
		Facebook struct {
			ClientID     string
			ClientSecret string
			RedirectURL  string
		}
	}
)

// GetConfig loads and returns configuration
func GetConfig() (Config, error) {
	var c Config

	// Load the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// Load env variables
	viper.SetEnvPrefix("pagoda")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}
```go

`config/config.yaml`

```yaml
http:
  hostname: ""
  port: 8080
  readTimeout: "5s"
  writeTimeout: "10s"
  idleTimeout: "2m"
  tls:
    enabled: false
    certificate: ""
    key: ""

app:
  name: "Pagoda ToDo App"
  environment: "local"
  # Change this on any live environments
  encryptionKey: "?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf"
  timeout: "20s"
  passwordToken:
      expiration: "60m"
      length: 64
  emailVerificationTokenExpiration: "12h"

cache:
  capacity: 100000
  expiration:
    staticFile: "4380h"
    page: "24h"

database:
  driver: "sqlite3"
  connection: "dbs/main.db?_journal=WAL&_timeout=5000&_fk=true"
  testConnection: ":memory:?_journal=WAL&_timeout=5000&_fk=true"

tasks:
  goroutines: 1
  releaseAfter: "15m"
  cleanupInterval: "1h"

mail:
  hostname: "localhost"
  port: 25
  user: "admin"
  password: "admin"
  fromAddress: "admin@localhost"

oauth:
  google:
    clientID: "your-google-client-id"
    clientSecret: "your-google-client-secret"
    redirectURL: "http://localhost:8080/auth/google/callback"
  facebook:
    clientID: "your-facebook-client-id"
    clientSecret: "your-facebook-client-secret"
    redirectURL: "http://localhost:8080/auth/facebook/callback"
```yaml

`config/config_test.go`

```go
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig()
	require.NoError(t, err)

	var env environment
	env = "abc"
	SwitchEnvironment(env)
	cfg, err := GetConfig()
	require.NoError(t, err)
	assert.Equal(t, env, cfg.App.Environment)
}
```go

`ent/client.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/mikestefanello/pagoda/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// PasswordToken is the client for interacting with the PasswordToken builders.
	PasswordToken *PasswordTokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.PasswordToken = NewPasswordTokenClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		PasswordToken: NewPasswordTokenClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		PasswordToken: NewPasswordTokenClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		PasswordToken.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.PasswordToken.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.PasswordToken.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PasswordTokenMutation:
		return c.PasswordToken.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PasswordTokenClient is a client for the PasswordToken schema.
type PasswordTokenClient struct {
	config
}

// NewPasswordTokenClient returns a client for the PasswordToken from the given config.
func NewPasswordTokenClient(c config) *PasswordTokenClient {
	return &PasswordTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `passwordtoken.Hooks(f(g(h())))`.
func (c *PasswordTokenClient) Use(hooks ...Hook) {
	c.hooks.PasswordToken = append(c.hooks.PasswordToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `passwordtoken.Intercept(f(g(h())))`.
func (c *PasswordTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.PasswordToken = append(c.inters.PasswordToken, interceptors...)
}

// Create returns a builder for creating a PasswordToken entity.
func (c *PasswordTokenClient) Create() *PasswordTokenCreate {
	mutation := newPasswordTokenMutation(c.config, OpCreate)
	return &PasswordTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PasswordToken entities.
func (c *PasswordTokenClient) CreateBulk(builders ...*PasswordTokenCreate) *PasswordTokenCreateBulk {
	return &PasswordTokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PasswordTokenClient) MapCreateBulk(slice any, setFunc func(*PasswordTokenCreate, int)) *PasswordTokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PasswordTokenCreateBulk{err: fmt.Errorf("calling to PasswordTokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PasswordTokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PasswordTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PasswordToken.
func (c *PasswordTokenClient) Update() *PasswordTokenUpdate {
	mutation := newPasswordTokenMutation(c.config, OpUpdate)
	return &PasswordTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PasswordTokenClient) UpdateOne(pt *PasswordToken) *PasswordTokenUpdateOne {
	mutation := newPasswordTokenMutation(c.config, OpUpdateOne, withPasswordToken(pt))
	return &PasswordTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PasswordTokenClient) UpdateOneID(id int) *PasswordTokenUpdateOne {
	mutation := newPasswordTokenMutation(c.config, OpUpdateOne, withPasswordTokenID(id))
	return &PasswordTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PasswordToken.
func (c *PasswordTokenClient) Delete() *PasswordTokenDelete {
	mutation := newPasswordTokenMutation(c.config, OpDelete)
	return &PasswordTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PasswordTokenClient) DeleteOne(pt *PasswordToken) *PasswordTokenDeleteOne {
	return c.DeleteOneID(pt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PasswordTokenClient) DeleteOneID(id int) *PasswordTokenDeleteOne {
	builder := c.Delete().Where(passwordtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PasswordTokenDeleteOne{builder}
}

// Query returns a query builder for PasswordToken.
func (c *PasswordTokenClient) Query() *PasswordTokenQuery {
	return &PasswordTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePasswordToken},
		inters: c.Interceptors(),
	}
}

// Get returns a PasswordToken entity by its id.
func (c *PasswordTokenClient) Get(ctx context.Context, id int) (*PasswordToken, error) {
	return c.Query().Where(passwordtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PasswordTokenClient) GetX(ctx context.Context, id int) *PasswordToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a PasswordToken.
func (c *PasswordTokenClient) QueryUser(pt *PasswordToken) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(passwordtoken.Table, passwordtoken.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, passwordtoken.UserTable, passwordtoken.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PasswordTokenClient) Hooks() []Hook {
	return c.hooks.PasswordToken
}

// Interceptors returns the client interceptors.
func (c *PasswordTokenClient) Interceptors() []Interceptor {
	return c.inters.PasswordToken
}

func (c *PasswordTokenClient) mutate(ctx context.Context, m *PasswordTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PasswordTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PasswordTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PasswordTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PasswordTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PasswordToken mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a User.
func (c *UserClient) QueryOwner(u *User) *PasswordTokenQuery {
	query := (&PasswordTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(passwordtoken.Table, passwordtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.OwnerTable, user.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		PasswordToken, User []ent.Hook
	}
	inters struct {
		PasswordToken, User []ent.Interceptor
	}
)
```go

`ent/ent.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// ent aliases to avoid import conflicts in user's code.
type (
	Op            = ent.Op
	Hook          = ent.Hook
	Value         = ent.Value
	Query         = ent.Query
	QueryContext  = ent.QueryContext
	Querier       = ent.Querier
	QuerierFunc   = ent.QuerierFunc
	Interceptor   = ent.Interceptor
	InterceptFunc = ent.InterceptFunc
	Traverser     = ent.Traverser
	TraverseFunc  = ent.TraverseFunc
	Policy        = ent.Policy
	Mutator       = ent.Mutator
	Mutation      = ent.Mutation
	MutateFunc    = ent.MutateFunc
)

type clientCtxKey struct{}

// FromContext returns a Client stored inside a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Client {
	c, _ := ctx.Value(clientCtxKey{}).(*Client)
	return c
}

// NewContext returns a new context with the given Client attached.
func NewContext(parent context.Context, c *Client) context.Context {
	return context.WithValue(parent, clientCtxKey{}, c)
}

type txCtxKey struct{}

// TxFromContext returns a Tx stored inside a context, or nil if there isn't one.
func TxFromContext(ctx context.Context) *Tx {
	tx, _ := ctx.Value(txCtxKey{}).(*Tx)
	return tx
}

// NewTxContext returns a new context with the given Tx attached.
func NewTxContext(parent context.Context, tx *Tx) context.Context {
	return context.WithValue(parent, txCtxKey{}, tx)
}

// OrderFunc applies an ordering on the sql selector.
// Deprecated: Use Asc/Desc functions or the package builders instead.
type OrderFunc func(*sql.Selector)

var (
	initCheck   sync.Once
	columnCheck sql.ColumnCheck
)

// columnChecker checks if the column exists in the given table.
func checkColumn(table, column string) error {
	initCheck.Do(func() {
		columnCheck = sql.NewColumnCheck(map[string]func(string) bool{
			passwordtoken.Table: passwordtoken.ValidColumn,
			user.Table:          user.ValidColumn,
		})
	})
	return columnCheck(table, column)
}

// Asc applies the given fields in ASC order.
func Asc(fields ...string) func(*sql.Selector) {
	return func(s *sql.Selector) {
		for _, f := range fields {
			if err := checkColumn(s.TableName(), f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Asc(s.C(f)))
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) func(*sql.Selector) {
	return func(s *sql.Selector) {
		for _, f := range fields {
			if err := checkColumn(s.TableName(), f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Desc(s.C(f)))
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
type AggregateFunc func(*sql.Selector) string

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.As(fn(s), end)
	}
}

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.Count("*")
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		if err := checkColumn(s.TableName(), field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Max(s.C(field))
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		if err := checkColumn(s.TableName(), field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Avg(s.C(field))
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		if err := checkColumn(s.TableName(), field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Min(s.C(field))
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		if err := checkColumn(s.TableName(), field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Sum(s.C(field))
	}
}

// ValidationError returns when validating a field or edge fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "ent: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "ent: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "ent: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "ent: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}

// selector embedded by the different Select/GroupBy builders.
type selector struct {
	label string
	flds  *[]string
	fns   []AggregateFunc
	scan  func(context.Context, any) error
}

// ScanX is like Scan, but panics if an error occurs.
func (s *selector) ScanX(ctx context.Context, v any) {
	if err := s.scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (s *selector) Strings(ctx context.Context) ([]string, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (s *selector) StringsX(ctx context.Context) []string {
	v, err := s.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (s *selector) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = s.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (s *selector) StringX(ctx context.Context) string {
	v, err := s.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (s *selector) Ints(ctx context.Context) ([]int, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (s *selector) IntsX(ctx context.Context) []int {
	v, err := s.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (s *selector) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = s.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (s *selector) IntX(ctx context.Context) int {
	v, err := s.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (s *selector) Float64s(ctx context.Context) ([]float64, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (s *selector) Float64sX(ctx context.Context) []float64 {
	v, err := s.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (s *selector) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = s.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (s *selector) Float64X(ctx context.Context) float64 {
	v, err := s.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (s *selector) Bools(ctx context.Context) ([]bool, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (s *selector) BoolsX(ctx context.Context) []bool {
	v, err := s.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (s *selector) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = s.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (s *selector) BoolX(ctx context.Context) bool {
	v, err := s.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// withHooks invokes the builder operation with the given hooks, if any.
func withHooks[V Value, M any, PM interface {
	*M
	Mutation
}](ctx context.Context, exec func(context.Context) (V, error), mutation PM, hooks []Hook) (value V, err error) {
	if len(hooks) == 0 {
		return exec(ctx)
	}
	var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
		mutationT, ok := any(m).(PM)
		if !ok {
			return nil, fmt.Errorf("unexpected mutation type %T", m)
		}
		// Set the mutation to the builder.
		*mutation = *mutationT
		return exec(ctx)
	})
	for i := len(hooks) - 1; i >= 0; i-- {
		if hooks[i] == nil {
			return value, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
		}
		mut = hooks[i](mut)
	}
	v, err := mut.Mutate(ctx, mutation)
	if err != nil {
		return value, err
	}
	nv, ok := v.(V)
	if !ok {
		return value, fmt.Errorf("unexpected node type %T returned from %T", v, mutation)
	}
	return nv, nil
}

// setContextOp returns a new context with the given QueryContext attached (including its op) in case it does not exist.
func setContextOp(ctx context.Context, qc *QueryContext, op string) context.Context {
	if ent.QueryFromContext(ctx) == nil {
		qc.Op = op
		ctx = ent.NewQueryContext(ctx, qc)
	}
	return ctx
}

func querierAll[V Value, Q interface {
	sqlAll(context.Context, ...queryHook) (V, error)
}]() Querier {
	return QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		query, ok := q.(Q)
		if !ok {
			return nil, fmt.Errorf("unexpected query type %T", q)
		}
		return query.sqlAll(ctx)
	})
}

func querierCount[Q interface {
	sqlCount(context.Context) (int, error)
}]() Querier {
	return QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		query, ok := q.(Q)
		if !ok {
			return nil, fmt.Errorf("unexpected query type %T", q)
		}
		return query.sqlCount(ctx)
	})
}

func withInterceptors[V Value](ctx context.Context, q Query, qr Querier, inters []Interceptor) (v V, err error) {
	for i := len(inters) - 1; i >= 0; i-- {
		qr = inters[i].Intercept(qr)
	}
	rv, err := qr.Query(ctx, q)
	if err != nil {
		return v, err
	}
	vt, ok := rv.(V)
	if !ok {
		return v, fmt.Errorf("unexpected type %T returned from %T. expected type: %T", vt, q, v)
	}
	return vt, nil
}

func scanWithInterceptors[Q1 ent.Query, Q2 interface {
	sqlScan(context.Context, Q1, any) error
}](ctx context.Context, rootQuery Q1, selectOrGroup Q2, inters []Interceptor, v any) error {
	rv := reflect.ValueOf(v)
	var qr Querier = QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		query, ok := q.(Q1)
		if !ok {
			return nil, fmt.Errorf("unexpected query type %T", q)
		}
		if err := selectOrGroup.sqlScan(ctx, query, v); err != nil {
			return nil, err
		}
		if k := rv.Kind(); k == reflect.Pointer && rv.Elem().CanInterface() {
			return rv.Elem().Interface(), nil
		}
		return v, nil
	})
	for i := len(inters) - 1; i >= 0; i-- {
		qr = inters[i].Intercept(qr)
	}
	vv, err := qr.Query(ctx, rootQuery)
	if err != nil {
		return err
	}
	switch rv2 := reflect.ValueOf(vv); {
	case rv.IsNil(), rv2.IsNil(), rv.Kind() != reflect.Pointer:
	case rv.Type() == rv2.Type():
		rv.Elem().Set(rv2.Elem())
	case rv.Elem().Type() == rv2.Type():
		rv.Elem().Set(rv2)
	}
	return nil
}

// queryHook describes an internal hook for the different sqlAll methods.
type queryHook func(context.Context, *sqlgraph.QuerySpec)
```go

`ent/enttest/enttest.go`

```go
// Code generated by ent, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/mikestefanello/pagoda/ent"
	// required by schema hooks.
	_ "github.com/mikestefanello/pagoda/ent/runtime"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/mikestefanello/pagoda/ent/migrate"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...any)
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []ent.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...ent.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls ent.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *ent.Client {
	o := newOptions(opts)
	c, err := ent.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	migrateSchema(t, c, o)
	return c
}

// NewClient calls ent.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *ent.Client {
	o := newOptions(opts)
	c := ent.NewClient(o.opts...)
	migrateSchema(t, c, o)
	return c
}
func migrateSchema(t TestingT, c *ent.Client, o *options) {
	tables, err := schema.CopyTables(migrate.Tables)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := migrate.Create(context.Background(), c.Schema, tables, o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
```go

`ent/generate.go`

```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```go

`ent/hook/hook.go`

```go
// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/mikestefanello/pagoda/ent"
)

// The PasswordTokenFunc type is an adapter to allow the use of ordinary
// function as PasswordToken mutator.
type PasswordTokenFunc func(context.Context, *ent.PasswordTokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PasswordTokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PasswordTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PasswordTokenMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
```go

`ent/migrate/migrate.go`

```go
// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"context"
	"fmt"
	"io"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

var (
	// WithGlobalUniqueID sets the universal ids options to the migration.
	// If this option is enabled, ent migration will allocate a 1<<32 range
	// for the ids of each entity (table).
	// Note that this option cannot be applied on tables that already exist.
	WithGlobalUniqueID = schema.WithGlobalUniqueID
	// WithDropColumn sets the drop column option to the migration.
	// If this option is enabled, ent migration will drop old columns
	// that were used for both fields and edges. This defaults to false.
	WithDropColumn = schema.WithDropColumn
	// WithDropIndex sets the drop index option to the migration.
	// If this option is enabled, ent migration will drop old indexes
	// that were defined in the schema. This defaults to false.
	// Note that unique constraints are defined using `UNIQUE INDEX`,
	// and therefore, it's recommended to enable this option to get more
	// flexibility in the schema changes.
	WithDropIndex = schema.WithDropIndex
	// WithForeignKeys enables creating foreign-key in schema DDL. This defaults to true.
	WithForeignKeys = schema.WithForeignKeys
)

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv dialect.Driver
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema { return &Schema{drv: drv} }

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context, opts ...schema.MigrateOption) error {
	return Create(ctx, s, Tables, opts...)
}

// Create creates all table resources using the given schema driver.
func Create(ctx context.Context, s *Schema, tables []*schema.Table, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %w", err)
	}
	return migrate.Create(ctx, tables...)
}

// WriteTo writes the schema changes to w instead of running them against the database.
//
//	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
//		log.Fatal(err)
//	}
func (s *Schema) WriteTo(ctx context.Context, w io.Writer, opts ...schema.MigrateOption) error {
	return Create(ctx, &Schema{drv: &schema.WriteDriver{Writer: w, Driver: s.drv}}, Tables, opts...)
}
```go

`ent/migrate/schema.go`

```go
// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PasswordTokensColumns holds the columns for the "password_tokens" table.
	PasswordTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "password_token_user", Type: field.TypeInt},
	}
	// PasswordTokensTable holds the schema information for the "password_tokens" table.
	PasswordTokensTable = &schema.Table{
		Name:       "password_tokens",
		Columns:    PasswordTokensColumns,
		PrimaryKey: []*schema.Column{PasswordTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_tokens_users_user",
				Columns:    []*schema.Column{PasswordTokensColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PasswordTokensTable,
		UsersTable,
	}
)

func init() {
	PasswordTokensTable.ForeignKeys[0].RefTable = UsersTable
}
```go

`ent/mutation.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePasswordToken = "PasswordToken"
	TypeUser          = "User"
)

// PasswordTokenMutation represents an operation that mutates the PasswordToken nodes in the graph.
type PasswordTokenMutation struct {
	config
	op            Op
	typ           string
	id            *int
	hash          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*PasswordToken, error)
	predicates    []predicate.PasswordToken
}

var _ ent.Mutation = (*PasswordTokenMutation)(nil)

// passwordtokenOption allows management of the mutation configuration using functional options.
type passwordtokenOption func(*PasswordTokenMutation)

// newPasswordTokenMutation creates new mutation for the PasswordToken entity.
func newPasswordTokenMutation(c config, op Op, opts ...passwordtokenOption) *PasswordTokenMutation {
	m := &PasswordTokenMutation{
		config:        c,
		op:            op,
		typ:           TypePasswordToken,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPasswordTokenID sets the ID field of the mutation.
func withPasswordTokenID(id int) passwordtokenOption {
	return func(m *PasswordTokenMutation) {
		var (
			err   error
			once  sync.Once
			value *PasswordToken
		)
		m.oldValue = func(ctx context.Context) (*PasswordToken, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().PasswordToken.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPasswordToken sets the old PasswordToken of the mutation.
func withPasswordToken(node *PasswordToken) passwordtokenOption {
	return func(m *PasswordTokenMutation) {
		m.oldValue = func(context.Context) (*PasswordToken, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PasswordTokenMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PasswordTokenMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PasswordTokenMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PasswordTokenMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().PasswordToken.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetHash sets the "hash" field.
func (m *PasswordTokenMutation) SetHash(s string) {
	m.hash = &s
}

// Hash returns the value of the "hash" field in the mutation.
func (m *PasswordTokenMutation) Hash() (r string, exists bool) {
	v := m.hash
	if v == nil {
		return
	}
	return *v, true
}

// OldHash returns the old "hash" field's value of the PasswordToken entity.
// If the PasswordToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordTokenMutation) OldHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHash: %w", err)
	}
	return oldValue.Hash, nil
}

// ResetHash resets all changes to the "hash" field.
func (m *PasswordTokenMutation) ResetHash() {
	m.hash = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *PasswordTokenMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PasswordTokenMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the PasswordToken entity.
// If the PasswordToken object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PasswordTokenMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PasswordTokenMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *PasswordTokenMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *PasswordTokenMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *PasswordTokenMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *PasswordTokenMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *PasswordTokenMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *PasswordTokenMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the PasswordTokenMutation builder.
func (m *PasswordTokenMutation) Where(ps ...predicate.PasswordToken) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PasswordTokenMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PasswordTokenMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.PasswordToken, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PasswordTokenMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PasswordTokenMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (PasswordToken).
func (m *PasswordTokenMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PasswordTokenMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.hash != nil {
		fields = append(fields, passwordtoken.FieldHash)
	}
	if m.created_at != nil {
		fields = append(fields, passwordtoken.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PasswordTokenMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case passwordtoken.FieldHash:
		return m.Hash()
	case passwordtoken.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PasswordTokenMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case passwordtoken.FieldHash:
		return m.OldHash(ctx)
	case passwordtoken.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown PasswordToken field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PasswordTokenMutation) SetField(name string, value ent.Value) error {
	switch name {
	case passwordtoken.FieldHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHash(v)
		return nil
	case passwordtoken.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown PasswordToken field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PasswordTokenMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PasswordTokenMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PasswordTokenMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown PasswordToken numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PasswordTokenMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PasswordTokenMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PasswordTokenMutation) ClearField(name string) error {
	return fmt.Errorf("unknown PasswordToken nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PasswordTokenMutation) ResetField(name string) error {
	switch name {
	case passwordtoken.FieldHash:
		m.ResetHash()
		return nil
	case passwordtoken.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown PasswordToken field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PasswordTokenMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, passwordtoken.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PasswordTokenMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case passwordtoken.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PasswordTokenMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PasswordTokenMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PasswordTokenMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, passwordtoken.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PasswordTokenMutation) EdgeCleared(name string) bool {
	switch name {
	case passwordtoken.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PasswordTokenMutation) ClearEdge(name string) error {
	switch name {
	case passwordtoken.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown PasswordToken unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PasswordTokenMutation) ResetEdge(name string) error {
	switch name {
	case passwordtoken.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown PasswordToken edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	email         *string
	password      *string
	verified      *bool
	created_at    *time.Time
	clearedFields map[string]struct{}
	owner         map[int]struct{}
	removedowner  map[int]struct{}
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetVerified sets the "verified" field.
func (m *UserMutation) SetVerified(b bool) {
	m.verified = &b
}

// Verified returns the value of the "verified" field in the mutation.
func (m *UserMutation) Verified() (r bool, exists bool) {
	v := m.verified
	if v == nil {
		return
	}
	return *v, true
}

// OldVerified returns the old "verified" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldVerified(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVerified is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVerified requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVerified: %w", err)
	}
	return oldValue.Verified, nil
}

// ResetVerified resets all changes to the "verified" field.
func (m *UserMutation) ResetVerified() {
	m.verified = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddOwnerIDs adds the "owner" edge to the PasswordToken entity by ids.
func (m *UserMutation) AddOwnerIDs(ids ...int) {
	if m.owner == nil {
		m.owner = make(map[int]struct{})
	}
	for i := range ids {
		m.owner[ids[i]] = struct{}{}
	}
}

// ClearOwner clears the "owner" edge to the PasswordToken entity.
func (m *UserMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the PasswordToken entity was cleared.
func (m *UserMutation) OwnerCleared() bool {
	return m.clearedowner
}

// RemoveOwnerIDs removes the "owner" edge to the PasswordToken entity by IDs.
func (m *UserMutation) RemoveOwnerIDs(ids ...int) {
	if m.removedowner == nil {
		m.removedowner = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.owner, ids[i])
		m.removedowner[ids[i]] = struct{}{}
	}
}

// RemovedOwner returns the removed IDs of the "owner" edge to the PasswordToken entity.
func (m *UserMutation) RemovedOwnerIDs() (ids []int) {
	for id := range m.removedowner {
		ids = append(ids, id)
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
func (m *UserMutation) OwnerIDs() (ids []int) {
	for id := range m.owner {
		ids = append(ids, id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *UserMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
	m.removedowner = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.verified != nil {
		fields = append(fields, user.FieldVerified)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldVerified:
		return m.Verified()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldVerified:
		return m.OldVerified(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldVerified:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVerified(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldVerified:
		m.ResetVerified()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, user.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOwner:
		ids := make([]ent.Value, 0, len(m.owner))
		for id := range m.owner {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedowner != nil {
		edges = append(edges, user.EdgeOwner)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOwner:
		ids := make([]ent.Value, 0, len(m.removedowner))
		for id := range m.removedowner {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, user.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
```go

`ent/passwordtoken/passwordtoken.go`

```go
// Code generated by ent, DO NOT EDIT.

package passwordtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the passwordtoken type in the database.
	Label = "password_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the passwordtoken in the database.
	Table = "password_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "password_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "password_token_user"
)

// Columns holds all SQL columns for passwordtoken fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "password_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"password_token_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the PasswordToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
```go

`ent/passwordtoken/where.go`

```go
// Code generated by ent, DO NOT EDIT.

package passwordtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldID, id))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldHash, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldCreatedAt, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldContainsFold(FieldHash, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PasswordToken {
	return predicate.PasswordToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PasswordToken {
	return predicate.PasswordToken(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.NotPredicates(p))
}
```go

`ent/passwordtoken.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// PasswordToken is the model entity for the PasswordToken schema.
type PasswordToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PasswordTokenQuery when eager-loading is set.
	Edges               PasswordTokenEdges `json:"edges"`
	password_token_user *int
	selectValues        sql.SelectValues
}

// PasswordTokenEdges holds the relations/edges for other nodes in the graph.
type PasswordTokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PasswordTokenEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PasswordToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case passwordtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case passwordtoken.FieldHash:
			values[i] = new(sql.NullString)
		case passwordtoken.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case passwordtoken.ForeignKeys[0]: // password_token_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PasswordToken fields.
func (pt *PasswordToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case passwordtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case passwordtoken.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				pt.Hash = value.String
			}
		case passwordtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case passwordtoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field password_token_user", value)
			} else if value.Valid {
				pt.password_token_user = new(int)
				*pt.password_token_user = int(value.Int64)
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PasswordToken.
// This includes values selected through modifiers, order, etc.
func (pt *PasswordToken) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the PasswordToken entity.
func (pt *PasswordToken) QueryUser() *UserQuery {
	return NewPasswordTokenClient(pt.config).QueryUser(pt)
}

// Update returns a builder for updating this PasswordToken.
// Note that you need to call PasswordToken.Unwrap() before calling this method if this PasswordToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PasswordToken) Update() *PasswordTokenUpdateOne {
	return NewPasswordTokenClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the PasswordToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PasswordToken) Unwrap() *PasswordToken {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PasswordToken is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PasswordToken) String() string {
	var builder strings.Builder
	builder.WriteString("PasswordToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("hash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PasswordTokens is a parsable slice of PasswordToken.
type PasswordTokens []*PasswordToken
```go

`ent/passwordtoken_create.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// PasswordTokenCreate is the builder for creating a PasswordToken entity.
type PasswordTokenCreate struct {
	config
	mutation *PasswordTokenMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (ptc *PasswordTokenCreate) SetHash(s string) *PasswordTokenCreate {
	ptc.mutation.SetHash(s)
	return ptc
}

// SetCreatedAt sets the "created_at" field.
func (ptc *PasswordTokenCreate) SetCreatedAt(t time.Time) *PasswordTokenCreate {
	ptc.mutation.SetCreatedAt(t)
	return ptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptc *PasswordTokenCreate) SetNillableCreatedAt(t *time.Time) *PasswordTokenCreate {
	if t != nil {
		ptc.SetCreatedAt(*t)
	}
	return ptc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptc *PasswordTokenCreate) SetUserID(id int) *PasswordTokenCreate {
	ptc.mutation.SetUserID(id)
	return ptc
}

// SetUser sets the "user" edge to the User entity.
func (ptc *PasswordTokenCreate) SetUser(u *User) *PasswordTokenCreate {
	return ptc.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptc *PasswordTokenCreate) Mutation() *PasswordTokenMutation {
	return ptc.mutation
}

// Save creates the PasswordToken in the database.
func (ptc *PasswordTokenCreate) Save(ctx context.Context) (*PasswordToken, error) {
	ptc.defaults()
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PasswordTokenCreate) SaveX(ctx context.Context) *PasswordToken {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *PasswordTokenCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *PasswordTokenCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptc *PasswordTokenCreate) defaults() {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		v := passwordtoken.DefaultCreatedAt()
		ptc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *PasswordTokenCreate) check() error {
	if _, ok := ptc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "PasswordToken.hash"`)}
	}
	if v, ok := ptc.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "PasswordToken.hash": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PasswordToken.created_at"`)}
	}
	if _, ok := ptc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "PasswordToken.user"`)}
	}
	return nil
}

func (ptc *PasswordTokenCreate) sqlSave(ctx context.Context) (*PasswordToken, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ptc.mutation.id = &_node.ID
	ptc.mutation.done = true
	return _node, nil
}

func (ptc *PasswordTokenCreate) createSpec() (*PasswordToken, *sqlgraph.CreateSpec) {
	var (
		_node = &PasswordToken{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(passwordtoken.Table, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	)
	if value, ok := ptc.mutation.Hash(); ok {
		_spec.SetField(passwordtoken.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := ptc.mutation.CreatedAt(); ok {
		_spec.SetField(passwordtoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ptc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.password_token_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PasswordTokenCreateBulk is the builder for creating many PasswordToken entities in bulk.
type PasswordTokenCreateBulk struct {
	config
	err      error
	builders []*PasswordTokenCreate
}

// Save creates the PasswordToken entities in the database.
func (ptcb *PasswordTokenCreateBulk) Save(ctx context.Context) ([]*PasswordToken, error) {
	if ptcb.err != nil {
		return nil, ptcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*PasswordToken, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasswordTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *PasswordTokenCreateBulk) SaveX(ctx context.Context) []*PasswordToken {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *PasswordTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *PasswordTokenCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
```go

`ent/passwordtoken_delete.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// PasswordTokenDelete is the builder for deleting a PasswordToken entity.
type PasswordTokenDelete struct {
	config
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// Where appends a list predicates to the PasswordTokenDelete builder.
func (ptd *PasswordTokenDelete) Where(ps ...predicate.PasswordToken) *PasswordTokenDelete {
	ptd.mutation.Where(ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *PasswordTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ptd.sqlExec, ptd.mutation, ptd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *PasswordTokenDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *PasswordTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(passwordtoken.Table, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ptd.mutation.done = true
	return affected, err
}

// PasswordTokenDeleteOne is the builder for deleting a single PasswordToken entity.
type PasswordTokenDeleteOne struct {
	ptd *PasswordTokenDelete
}

// Where appends a list predicates to the PasswordTokenDelete builder.
func (ptdo *PasswordTokenDeleteOne) Where(ps ...predicate.PasswordToken) *PasswordTokenDeleteOne {
	ptdo.ptd.mutation.Where(ps...)
	return ptdo
}

// Exec executes the deletion query.
func (ptdo *PasswordTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{passwordtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *PasswordTokenDeleteOne) ExecX(ctx context.Context) {
	if err := ptdo.Exec(ctx); err != nil {
		panic(err)
	}
}
```go

`ent/passwordtoken_query.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

// PasswordTokenQuery is the builder for querying PasswordToken entities.
type PasswordTokenQuery struct {
	config
	ctx        *QueryContext
	order      []passwordtoken.OrderOption
	inters     []Interceptor
	predicates []predicate.PasswordToken
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PasswordTokenQuery builder.
func (ptq *PasswordTokenQuery) Where(ps ...predicate.PasswordToken) *PasswordTokenQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PasswordTokenQuery) Limit(limit int) *PasswordTokenQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PasswordTokenQuery) Offset(offset int) *PasswordTokenQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PasswordTokenQuery) Unique(unique bool) *PasswordTokenQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PasswordTokenQuery) Order(o ...passwordtoken.OrderOption) *PasswordTokenQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryUser chains the current query on the "user" edge.
func (ptq *PasswordTokenQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(passwordtoken.Table, passwordtoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, passwordtoken.UserTable, passwordtoken.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PasswordToken entity from the query.
// Returns a *NotFoundError when no PasswordToken was found.
func (ptq *PasswordTokenQuery) First(ctx context.Context) (*PasswordToken, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{passwordtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PasswordTokenQuery) FirstX(ctx context.Context) *PasswordToken {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PasswordToken ID from the query.
// Returns a *NotFoundError when no PasswordToken ID was found.
func (ptq *PasswordTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{passwordtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PasswordTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PasswordToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PasswordToken entity is found.
// Returns a *NotFoundError when no PasswordToken entities are found.
func (ptq *PasswordTokenQuery) Only(ctx context.Context) (*PasswordToken, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{passwordtoken.Label}
	default:
		return nil, &NotSingularError{passwordtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PasswordTokenQuery) OnlyX(ctx context.Context) *PasswordToken {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PasswordToken ID in the query.
// Returns a *NotSingularError when more than one PasswordToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PasswordTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{passwordtoken.Label}
	default:
		err = &NotSingularError{passwordtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PasswordTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PasswordTokens.
func (ptq *PasswordTokenQuery) All(ctx context.Context) ([]*PasswordToken, error) {
	ctx = setContextOp(ctx, ptq.ctx, "All")
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PasswordToken, *PasswordTokenQuery]()
	return withInterceptors[[]*PasswordToken](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PasswordTokenQuery) AllX(ctx context.Context) []*PasswordToken {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PasswordToken IDs.
func (ptq *PasswordTokenQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, "IDs")
	if err = ptq.Select(passwordtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PasswordTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PasswordTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Count")
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PasswordTokenQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PasswordTokenQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PasswordTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Exist")
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PasswordTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PasswordTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PasswordTokenQuery) Clone() *PasswordTokenQuery {
	if ptq == nil {
		return nil
	}
	return &PasswordTokenQuery{
		config:     ptq.config,
		ctx:        ptq.ctx.Clone(),
		order:      append([]passwordtoken.OrderOption{}, ptq.order...),
		inters:     append([]Interceptor{}, ptq.inters...),
		predicates: append([]predicate.PasswordToken{}, ptq.predicates...),
		withUser:   ptq.withUser.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PasswordTokenQuery) WithUser(opts ...func(*UserQuery)) *PasswordTokenQuery {
	query := (&UserClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withUser = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PasswordToken.Query().
//		GroupBy(passwordtoken.FieldHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PasswordTokenQuery) GroupBy(field string, fields ...string) *PasswordTokenGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PasswordTokenGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = passwordtoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hash string `json:"hash,omitempty"`
//	}
//
//	client.PasswordToken.Query().
//		Select(passwordtoken.FieldHash).
//		Scan(ctx, &v)
func (ptq *PasswordTokenQuery) Select(fields ...string) *PasswordTokenSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PasswordTokenSelect{PasswordTokenQuery: ptq}
	sbuild.label = passwordtoken.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PasswordTokenSelect configured with the given aggregations.
func (ptq *PasswordTokenQuery) Aggregate(fns ...AggregateFunc) *PasswordTokenSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PasswordTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !passwordtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PasswordTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PasswordToken, error) {
	var (
		nodes       = []*PasswordToken{}
		withFKs     = ptq.withFKs
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withUser != nil,
		}
	)
	if ptq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PasswordToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PasswordToken{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withUser; query != nil {
		if err := ptq.loadUser(ctx, query, nodes, nil,
			func(n *PasswordToken, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PasswordTokenQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*PasswordToken, init func(*PasswordToken), assign func(*PasswordToken, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PasswordToken)
	for i := range nodes {
		if nodes[i].password_token_user == nil {
			continue
		}
		fk := *nodes[i].password_token_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "password_token_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptq *PasswordTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PasswordTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(passwordtoken.Table, passwordtoken.Columns, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.FieldID)
		for i := range fields {
			if fields[i] != passwordtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PasswordTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(passwordtoken.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = passwordtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PasswordTokenGroupBy is the group-by builder for PasswordToken entities.
type PasswordTokenGroupBy struct {
	selector
	build *PasswordTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PasswordTokenGroupBy) Aggregate(fns ...AggregateFunc) *PasswordTokenGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PasswordTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, "GroupBy")
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordTokenQuery, *PasswordTokenGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PasswordTokenGroupBy) sqlScan(ctx context.Context, root *PasswordTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PasswordTokenSelect is the builder for selecting fields of PasswordToken entities.
type PasswordTokenSelect struct {
	*PasswordTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PasswordTokenSelect) Aggregate(fns ...AggregateFunc) *PasswordTokenSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PasswordTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, "Select")
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordTokenQuery, *PasswordTokenSelect](ctx, pts.PasswordTokenQuery, pts, pts.inters, v)
}

func (pts *PasswordTokenSelect) sqlScan(ctx context.Context, root *PasswordTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
```go

`ent/passwordtoken_update.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

// PasswordTokenUpdate is the builder for updating PasswordToken entities.
type PasswordTokenUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// Where appends a list predicates to the PasswordTokenUpdate builder.
func (ptu *PasswordTokenUpdate) Where(ps ...predicate.PasswordToken) *PasswordTokenUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetHash sets the "hash" field.
func (ptu *PasswordTokenUpdate) SetHash(s string) *PasswordTokenUpdate {
	ptu.mutation.SetHash(s)
	return ptu
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (ptu *PasswordTokenUpdate) SetNillableHash(s *string) *PasswordTokenUpdate {
	if s != nil {
		ptu.SetHash(*s)
	}
	return ptu
}

// SetCreatedAt sets the "created_at" field.
func (ptu *PasswordTokenUpdate) SetCreatedAt(t time.Time) *PasswordTokenUpdate {
	ptu.mutation.SetCreatedAt(t)
	return ptu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptu *PasswordTokenUpdate) SetNillableCreatedAt(t *time.Time) *PasswordTokenUpdate {
	if t != nil {
		ptu.SetCreatedAt(*t)
	}
	return ptu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptu *PasswordTokenUpdate) SetUserID(id int) *PasswordTokenUpdate {
	ptu.mutation.SetUserID(id)
	return ptu
}

// SetUser sets the "user" edge to the User entity.
func (ptu *PasswordTokenUpdate) SetUser(u *User) *PasswordTokenUpdate {
	return ptu.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptu *PasswordTokenUpdate) Mutation() *PasswordTokenMutation {
	return ptu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ptu *PasswordTokenUpdate) ClearUser() *PasswordTokenUpdate {
	ptu.mutation.ClearUser()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PasswordTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ptu.sqlSave, ptu.mutation, ptu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PasswordTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PasswordTokenUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PasswordTokenUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *PasswordTokenUpdate) check() error {
	if v, ok := ptu.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "PasswordToken.hash": %w`, err)}
		}
	}
	if _, ok := ptu.mutation.UserID(); ptu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PasswordToken.user"`)
	}
	return nil
}

func (ptu *PasswordTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ptu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(passwordtoken.Table, passwordtoken.Columns, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.Hash(); ok {
		_spec.SetField(passwordtoken.FieldHash, field.TypeString, value)
	}
	if value, ok := ptu.mutation.CreatedAt(); ok {
		_spec.SetField(passwordtoken.FieldCreatedAt, field.TypeTime, value)
	}
	if ptu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptu.mutation.done = true
	return n, nil
}

// PasswordTokenUpdateOne is the builder for updating a single PasswordToken entity.
type PasswordTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// SetHash sets the "hash" field.
func (ptuo *PasswordTokenUpdateOne) SetHash(s string) *PasswordTokenUpdateOne {
	ptuo.mutation.SetHash(s)
	return ptuo
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (ptuo *PasswordTokenUpdateOne) SetNillableHash(s *string) *PasswordTokenUpdateOne {
	if s != nil {
		ptuo.SetHash(*s)
	}
	return ptuo
}

// SetCreatedAt sets the "created_at" field.
func (ptuo *PasswordTokenUpdateOne) SetCreatedAt(t time.Time) *PasswordTokenUpdateOne {
	ptuo.mutation.SetCreatedAt(t)
	return ptuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptuo *PasswordTokenUpdateOne) SetNillableCreatedAt(t *time.Time) *PasswordTokenUpdateOne {
	if t != nil {
		ptuo.SetCreatedAt(*t)
	}
	return ptuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptuo *PasswordTokenUpdateOne) SetUserID(id int) *PasswordTokenUpdateOne {
	ptuo.mutation.SetUserID(id)
	return ptuo
}

// SetUser sets the "user" edge to the User entity.
func (ptuo *PasswordTokenUpdateOne) SetUser(u *User) *PasswordTokenUpdateOne {
	return ptuo.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptuo *PasswordTokenUpdateOne) Mutation() *PasswordTokenMutation {
	return ptuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ptuo *PasswordTokenUpdateOne) ClearUser() *PasswordTokenUpdateOne {
	ptuo.mutation.ClearUser()
	return ptuo
}

// Where appends a list predicates to the PasswordTokenUpdate builder.
func (ptuo *PasswordTokenUpdateOne) Where(ps ...predicate.PasswordToken) *PasswordTokenUpdateOne {
	ptuo.mutation.Where(ps...)
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PasswordTokenUpdateOne) Select(field string, fields ...string) *PasswordTokenUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PasswordToken entity.
func (ptuo *PasswordTokenUpdateOne) Save(ctx context.Context) (*PasswordToken, error) {
	return withHooks(ctx, ptuo.sqlSave, ptuo.mutation, ptuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PasswordTokenUpdateOne) SaveX(ctx context.Context) *PasswordToken {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PasswordTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PasswordTokenUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *PasswordTokenUpdateOne) check() error {
	if v, ok := ptuo.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "PasswordToken.hash": %w`, err)}
		}
	}
	if _, ok := ptuo.mutation.UserID(); ptuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PasswordToken.user"`)
	}
	return nil
}

func (ptuo *PasswordTokenUpdateOne) sqlSave(ctx context.Context) (_node *PasswordToken, err error) {
	if err := ptuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(passwordtoken.Table, passwordtoken.Columns, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PasswordToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.FieldID)
		for _, f := range fields {
			if !passwordtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passwordtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.Hash(); ok {
		_spec.SetField(passwordtoken.FieldHash, field.TypeString, value)
	}
	if value, ok := ptuo.mutation.CreatedAt(); ok {
		_spec.SetField(passwordtoken.FieldCreatedAt, field.TypeTime, value)
	}
	if ptuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PasswordToken{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptuo.mutation.done = true
	return _node, nil
}
```go

`ent/predicate/predicate.go`

```go
// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// PasswordToken is the predicate function for passwordtoken builders.
type PasswordToken func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
```go

`ent/runtime/runtime.go`

```go
// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/schema"
	"github.com/mikestefanello/pagoda/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[3].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)
```go

`ent/runtime.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

// The schema-stitching logic is generated in github.com/mikestefanello/pagoda/ent/runtime/runtime.go
```go

`ent/schema/passwordtoken.go`

```go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PasswordToken holds the schema definition for the PasswordToken entity.
type PasswordToken struct {
	ent.Schema
}

// Fields of the PasswordToken.
func (PasswordToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").
			Sensitive().
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the PasswordToken.
func (PasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique(),
	}
}
```go

`ent/schema/user.go`

```go
package schema

import (
	"context"
	"strings"
	"time"

	ge "github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/hook"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.Bool("verified").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", PasswordToken.Type).
			Ref("user"),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ge.UserMutation) (ent.Value, error) {
					if v, exists := m.Email(); exists {
						m.SetEmail(strings.ToLower(v))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
```go

`ent/tx.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// PasswordToken is the client for interacting with the PasswordToken builders.
	PasswordToken *PasswordTokenClient
	// User is the client for interacting with the User builders.
	User *UserClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.PasswordToken = NewPasswordTokenClient(tx.config)
	tx.User = NewUserClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: PasswordToken.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
```go

`ent/user/user.go`

```go
// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the user in the database.
	Table = "users"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "password_tokens"
	// OwnerInverseTable is the table name for the PasswordToken entity.
	// It exists in this package in order to avoid circular dependency with the "passwordtoken" package.
	OwnerInverseTable = "password_tokens"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "password_token_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldVerified,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/mikestefanello/pagoda/ent/runtime"
var (
	Hooks [1]ent.Hook
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByVerified orders the results by the verified field.
func ByVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerified, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOwnerCount orders the results by owner count.
func ByOwnerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnerStep(), opts...)
	}
}

// ByOwner orders the results by owner terms.
func ByOwner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OwnerTable, OwnerColumn),
	)
}
```go

`ent/user/where.go`

```go
// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Verified applies equality check predicate on the "verified" field. It's identical to VerifiedEQ.
func Verified(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVerified, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// VerifiedEQ applies the EQ predicate on the "verified" field.
func VerifiedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVerified, v))
}

// VerifiedNEQ applies the NEQ predicate on the "verified" field.
func VerifiedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldVerified, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.PasswordToken) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
```go

`ent/user.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Verified holds the value of the "verified" field.
	Verified bool `json:"verified,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Owner holds the value of the owner edge.
	Owner []*PasswordToken `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnerOrErr() ([]*PasswordToken, error) {
	if e.loadedTypes[0] {
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldVerified:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldEmail, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified", values[i])
			} else if value.Valid {
				u.Verified = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the User entity.
func (u *User) QueryOwner() *PasswordTokenQuery {
	return NewUserClient(u.config).QueryOwner(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("verified=")
	builder.WriteString(fmt.Sprintf("%v", u.Verified))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
```go

`ent/user_create.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetVerified sets the "verified" field.
func (uc *UserCreate) SetVerified(b bool) *UserCreate {
	uc.mutation.SetVerified(b)
	return uc
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uc *UserCreate) SetNillableVerified(b *bool) *UserCreate {
	if b != nil {
		uc.SetVerified(*b)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// AddOwnerIDs adds the "owner" edge to the PasswordToken entity by IDs.
func (uc *UserCreate) AddOwnerIDs(ids ...int) *UserCreate {
	uc.mutation.AddOwnerIDs(ids...)
	return uc
}

// AddOwner adds the "owner" edges to the PasswordToken entity.
func (uc *UserCreate) AddOwner(p ...*PasswordToken) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddOwnerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.Verified(); !ok {
		v := user.DefaultVerified
		uc.mutation.SetVerified(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		if user.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Verified(); !ok {
		return &ValidationError{Name: "verified", err: errors.New(`ent: missing required field "User.verified"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
		_node.Verified = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
```go

`ent/user_delete.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

// UserDelete is the builder for deleting a User entity.
type UserDelete struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserDelete builder.
func (ud *UserDelete) Where(ps ...predicate.User) *UserDelete {
	ud.mutation.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ud.sqlExec, ud.mutation, ud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ud.mutation.done = true
	return affected, err
}

// UserDeleteOne is the builder for deleting a single User entity.
type UserDeleteOne struct {
	ud *UserDelete
}

// Where appends a list predicates to the UserDelete builder.
func (udo *UserDeleteOne) Where(ps ...predicate.User) *UserDeleteOne {
	udo.ud.mutation.Where(ps...)
	return udo
}

// Exec executes the deletion query.
func (udo *UserDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{user.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserDeleteOne) ExecX(ctx context.Context) {
	if err := udo.Exec(ctx); err != nil {
		panic(err)
	}
}
```go

`ent/user_query.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx        *QueryContext
	order      []user.OrderOption
	inters     []Interceptor
	predicates []predicate.User
	withOwner  *PasswordTokenQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryOwner chains the current query on the "owner" edge.
func (uq *UserQuery) QueryOwner() *PasswordTokenQuery {
	query := (&PasswordTokenClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(passwordtoken.Table, passwordtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.OwnerTable, user.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:     uq.config,
		ctx:        uq.ctx.Clone(),
		order:      append([]user.OrderOption{}, uq.order...),
		inters:     append([]Interceptor{}, uq.inters...),
		predicates: append([]predicate.User{}, uq.predicates...),
		withOwner:  uq.withOwner.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithOwner(opts ...func(*PasswordTokenQuery)) *UserQuery {
	query := (&PasswordTokenClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withOwner = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldName).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withOwner; query != nil {
		if err := uq.loadOwner(ctx, query, nodes,
			func(n *User) { n.Edges.Owner = []*PasswordToken{} },
			func(n *User, e *PasswordToken) { n.Edges.Owner = append(n.Edges.Owner, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadOwner(ctx context.Context, query *PasswordTokenQuery, nodes []*User, init func(*User), assign func(*User, *PasswordToken)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PasswordToken(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.OwnerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.password_token_user
		if fk == nil {
			return fmt.Errorf(`foreign-key "password_token_user" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "password_token_user" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
```go

`ent/user_update.go`

```go
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetVerified sets the "verified" field.
func (uu *UserUpdate) SetVerified(b bool) *UserUpdate {
	uu.mutation.SetVerified(b)
	return uu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetVerified(*b)
	}
	return uu
}

// AddOwnerIDs adds the "owner" edge to the PasswordToken entity by IDs.
func (uu *UserUpdate) AddOwnerIDs(ids ...int) *UserUpdate {
	uu.mutation.AddOwnerIDs(ids...)
	return uu
}

// AddOwner adds the "owner" edges to the PasswordToken entity.
func (uu *UserUpdate) AddOwner(p ...*PasswordToken) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddOwnerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearOwner clears all "owner" edges to the PasswordToken entity.
func (uu *UserUpdate) ClearOwner() *UserUpdate {
	uu.mutation.ClearOwner()
	return uu
}

// RemoveOwnerIDs removes the "owner" edge to PasswordToken entities by IDs.
func (uu *UserUpdate) RemoveOwnerIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveOwnerIDs(ids...)
	return uu
}

// RemoveOwner removes "owner" edges to PasswordToken entities.
func (uu *UserUpdate) RemoveOwner(p ...*PasswordToken) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if uu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !uu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetVerified sets the "verified" field.
func (uuo *UserUpdateOne) SetVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetVerified(b)
	return uuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetVerified(*b)
	}
	return uuo
}

// AddOwnerIDs adds the "owner" edge to the PasswordToken entity by IDs.
func (uuo *UserUpdateOne) AddOwnerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddOwnerIDs(ids...)
	return uuo
}

// AddOwner adds the "owner" edges to the PasswordToken entity.
func (uuo *UserUpdateOne) AddOwner(p ...*PasswordToken) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddOwnerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearOwner clears all "owner" edges to the PasswordToken entity.
func (uuo *UserUpdateOne) ClearOwner() *UserUpdateOne {
	uuo.mutation.ClearOwner()
	return uuo
}

// RemoveOwnerIDs removes the "owner" edge to PasswordToken entities by IDs.
func (uuo *UserUpdateOne) RemoveOwnerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveOwnerIDs(ids...)
	return uuo
}

// RemoveOwner removes "owner" edges to PasswordToken entities.
func (uuo *UserUpdateOne) RemoveOwner(p ...*PasswordToken) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveOwnerIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if uuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !uuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnerTable,
			Columns: []string{user.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
```go

`go.mod`

```
module github.com/mikestefanello/pagoda

go 1.22.4

require (
	entgo.io/ent v0.13.1
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/PuerkitoBio/goquery v1.9.1
	github.com/go-playground/validator/v10 v10.19.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/context v1.1.2
	github.com/gorilla/sessions v1.2.2
	github.com/labstack/echo/v4 v4.12.0
	github.com/labstack/gommon v0.4.2
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/maypok86/otter v1.2.1
	github.com/mikestefanello/backlite v0.1.0
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.22.0
	golang.org/x/oauth2 v0.21.0
)

require (
	ariga.io/atlas v0.21.1 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dolthub/maphash v0.1.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gammazero/deque v0.2.1 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.20.1 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

`go.sum`

```
ariga.io/atlas v0.21.1 h1:Eg9XYhKTH3UHoqP7tKMWFV+Z5JnpVOJCgO3MHrUtKmk=
ariga.io/atlas v0.21.1/go.mod h1:VPlcXdd4w2KqKnH54yEZcry79UAhpaWaxEsmn5JRNoE=
cloud.google.com/go/compute/metadata v0.3.0 h1:Tz+eQXMEqDIKRsmY3cHTL6FVaynIjX2QxYC4trgAKZc=
cloud.google.com/go/compute/metadata v0.3.0/go.mod h1:zFmK7XCadkQkj6TtorcaGlCW1hT1fIilQDwofLpJ20k=
entgo.io/ent v0.13.1 h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=
entgo.io/ent v0.13.1/go.mod h1:qCEmo+biw3ccBn9OyL4ZK5dfpwg++l1Gxwac5B1206A=
github.com/DATA-DOG/go-sqlmock v1.5.0 h1:Shsta01QNfFxHCfpW6YH2STWB0MudeXXEWMr20OEh60=
github.com/DATA-DOG/go-sqlmock v1.5.0/go.mod h1:f/Ixk793poVmq4qj/V1dPUg2JEAKC73Q5eFN3EC/SaM=
github.com/Masterminds/goutils v1.1.1 h1:5nUrii3FMTL5diU80unEVvNevw1nH4+ZV4DSLVJLSYI=
github.com/Masterminds/goutils v1.1.1/go.mod h1:8cTjp+g8YejhMuvIA5y2vz3BpJxksy863GQaJW2MFNU=
github.com/Masterminds/semver v1.5.0 h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=
github.com/Masterminds/semver v1.5.0/go.mod h1:MB6lktGJrhw8PrUyiEoblNEGEQ+RzHPF078ddwwvV3Y=
github.com/Masterminds/sprig v2.22.0+incompatible h1:z4yfnGrZ7netVz+0EDJ0Wi+5VZCSYp4Z0m2dk6cEM60=
github.com/Masterminds/sprig v2.22.0+incompatible/go.mod h1:y6hNFY5UBTIWBxnzTeuNhlNS5hqE0NB0E6fgfo2Br3o=
github.com/PuerkitoBio/goquery v1.9.1 h1:mTL6XjbJTZdpfL+Gwl5U2h1l9yEkJjhmlTeV9VPW7UI=
github.com/PuerkitoBio/goquery v1.9.1/go.mod h1:cW1n6TmIMDoORQU5IU/P1T3tGFunOeXEpGP2WHRwkbY=
github.com/agext/levenshtein v1.2.3 h1:YB2fHEn0UJagG8T1rrWknE3ZQzWM06O8AMAatNn7lmo=
github.com/agext/levenshtein v1.2.3/go.mod h1:JEDfjyjHDjOF/1e4FlBE/PkbqA9OfWu2ki2W0IB5558=
github.com/andybalholm/cascadia v1.3.2 h1:3Xi6Dw5lHF15JtdcmAHD3i1+T8plmv7BQ/nsViSLyss=
github.com/andybalholm/cascadia v1.3.2/go.mod h1:7gtRlve5FxPPgIgX36uWBX58OdBsSS6lUvCFb+h7KvU=
github.com/apparentlymart/go-textseg/v15 v15.0.0 h1:uYvfpb3DyLSCGWnctWKGj857c6ew1u1fNQOlOtuGxQY=
github.com/apparentlymart/go-textseg/v15 v15.0.0/go.mod h1:K8XmNZdhEBkdlyDdvbmmsvpAG721bKi0joRfFdHIWJ4=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc h1:U9qPSI2PIWSS1VwoXQT9A3Wy9MM3WgvqSxFWenqJduM=
github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/dolthub/maphash v0.1.0 h1:bsQ7JsF4FkkWyrP3oCnFJgrCUAFbFf3kOl4L/QxPDyQ=
github.com/dolthub/maphash v0.1.0/go.mod h1:gkg4Ch4CdCDu5h6PMriVLawB7koZ+5ijb9puGMV50a4=
github.com/frankban/quicktest v1.14.6 h1:7Xjx+VpznH+oBnejlPUj8oUpdxnVs4f8XU8WnHkI4W8=
github.com/frankban/quicktest v1.14.6/go.mod h1:4ptaffx2x8+WTWXmUCuVU6aPUX1/Mz7zb5vbUoiM6w0=
github.com/fsnotify/fsnotify v1.7.0 h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=
github.com/fsnotify/fsnotify v1.7.0/go.mod h1:40Bi/Hjc2AVfZrqy+aj+yEI+/bRxZnMJyTJwOpGvigM=
github.com/gabriel-vasile/mimetype v1.4.3 h1:in2uUcidCuFcDKtdcBxlR0rJ1+fsokWf+uqxgUFjbI0=
github.com/gabriel-vasile/mimetype v1.4.3/go.mod h1:d8uq/6HKRL6CGdk+aubisF/M5GcPfT7nKyLpA0lbSSk=
github.com/gammazero/deque v0.2.1 h1:qSdsbG6pgp6nL7A0+K/B7s12mcCY/5l5SIUpMOl+dC0=
github.com/gammazero/deque v0.2.1/go.mod h1:LFroj8x4cMYCukHJDbxFCkT+r9AndaJnFMuZDV34tuU=
github.com/go-openapi/inflect v0.21.0 h1:FoBjBTQEcbg2cJUWX6uwL9OyIW8eqc9k4KhN4lfbeYk=
github.com/go-openapi/inflect v0.21.0/go.mod h1:INezMuUu7SJQc2AyR3WO0DqqYUJSj8Kb4hBd7WtjlAw=
github.com/go-playground/assert/v2 v2.2.0 h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=
github.com/go-playground/assert/v2 v2.2.0/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
github.com/go-playground/locales v0.14.1 h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=
github.com/go-playground/locales v0.14.1/go.mod h1:hxrqLVvrK65+Rwrd5Fc6F2O76J/NuW9t0sjnWqG1slY=
github.com/go-playground/universal-translator v0.18.1 h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=
github.com/go-playground/universal-translator v0.18.1/go.mod h1:xekY+UJKNuX9WP91TpwSH2VMlDf28Uj24BCp08ZFTUY=
github.com/go-playground/validator/v10 v10.19.0 h1:ol+5Fu+cSq9JD7SoSqe04GMI92cbn0+wvQ3bZ8b/AU4=
github.com/go-playground/validator/v10 v10.19.0/go.mod h1:dbuPbCMFw/DrkbEynArYaCwl3amGuJotoKCe95atGMM=
github.com/go-test/deep v1.0.3 h1:ZrJSEWsXzPOxaZnFteGEfooLba+ju3FYIbOrS+rQd68=
github.com/go-test/deep v1.0.3/go.mod h1:wGDj63lr65AM2AQyKZd/NYHGb0R+1RLqB8NKt3aSFNA=
github.com/golang-jwt/jwt v3.2.2+incompatible h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=
github.com/golang-jwt/jwt v3.2.2+incompatible/go.mod h1:8pz2t5EyA70fFQQSrl6XZXzqecmYZeUEB8OUGHkxJ+I=
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
github.com/google/gofuzz v1.2.0 h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=
github.com/google/gofuzz v1.2.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/gorilla/context v1.1.2 h1:WRkNAv2uoa03QNIc1A6u4O7DAGMUVoopZhkiXWA2V1o=
github.com/gorilla/context v1.1.2/go.mod h1:KDPwT9i/MeWHiLl90fuTgrt4/wPcv75vFAZLaOOcbxM=
github.com/gorilla/securecookie v1.1.2 h1:YCIWL56dvtr73r6715mJs5ZvhtnY73hBvEF8kXD8ePA=
github.com/gorilla/securecookie v1.1.2/go.mod h1:NfCASbcHqRSY+3a8tlWJwsQap2VX5pwzwo4h3eOamfo=
github.com/gorilla/sessions v1.2.2 h1:lqzMYz6bOfvn2WriPUjNByzeXIlVzURcPmgMczkmTjY=
github.com/gorilla/sessions v1.2.2/go.mod h1:ePLdVu+jbEgHH+KWw8I1z2wqd0BAdAQh/8LRvBeoNcQ=
github.com/hashicorp/hcl v1.0.0 h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=
github.com/hashicorp/hcl v1.0.0/go.mod h1:E5yfLk+7swimpb2L/Alb/PJmXilQ/rhwaUYs4T20WEQ=
github.com/hashicorp/hcl/v2 v2.20.1 h1:M6hgdyz7HYt1UN9e61j+qKJBqR3orTWbI1HKBJEdxtc=
github.com/hashicorp/hcl/v2 v2.20.1/go.mod h1:TZDqQ4kNKCbh1iJp99FdPiUaVDDUPivbqxZulxDYqL4=
github.com/huandu/xstrings v1.4.0 h1:D17IlohoQq4UcpqD7fDk80P7l+lwAmlFaBHgOipl2FU=
github.com/huandu/xstrings v1.4.0/go.mod h1:y5/lhBue+AyNmUVz9RLU9xbLR0o4KIIExikq4ovT0aE=
github.com/imdario/mergo v0.3.16 h1:wwQJbIsHYGMUyLSPrEq1CT16AhnhNJQ51+4fdHUnCl4=
github.com/imdario/mergo v0.3.16/go.mod h1:WBLT9ZmE3lPoWsEzCh9LPo3TiwVN+ZKEjmz+hD27ysY=
github.com/kr/pretty v0.2.1/go.mod h1:ipq/a2n7PKx3OHsz4KJII5eveXtPO4qwEXGdVfWzfnI=
github.com/kr/pretty v0.3.1 h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=
github.com/kr/pretty v0.3.1/go.mod h1:hoEshYVHaxMs3cyo3Yncou5ZscifuDolrwPKZanG3xk=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/labstack/echo/v4 v4.12.0 h1:IKpw49IMryVB2p1a4dzwlhP1O2Tf2E0Ir/450lH+kI0=
github.com/labstack/echo/v4 v4.12.0/go.mod h1:UP9Cr2DJXbOK3Kr9ONYzNowSh7HP0aG0ShAyycHSJvM=
github.com/labstack/gommon v0.4.2 h1:F8qTUNXgG1+6WQmqoUWnz8WiEU60mXVVw0P4ht1WRA0=
github.com/labstack/gommon v0.4.2/go.mod h1:QlUFxVM+SNXhDL/Z7YhocGIBYOiwB0mXm1+1bAPHPyU=
github.com/leodido/go-urn v1.4.0 h1:WT9HwE9SGECu3lg4d/dIA+jxlljEa1/ffXKmRjqdmIQ=
github.com/leodido/go-urn v1.4.0/go.mod h1:bvxc+MVxLKB4z00jd1z+Dvzr47oO32F/QSNjSBOlFxI=
github.com/magiconair/properties v1.8.7 h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=
github.com/magiconair/properties v1.8.7/go.mod h1:Dhd985XPs7jluiymwWYZ0G4Z61jb3vdS329zhj2hYo0=
github.com/mattn/go-colorable v0.1.13 h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=
github.com/mattn/go-colorable v0.1.13/go.mod h1:7S9/ev0klgBDR4GtXTXX8a3vIGJpMovkB8vQcUbaXHg=
github.com/mattn/go-isatty v0.0.16/go.mod h1:kYGgaQfpe5nmfYZH+SKPsOc2e4SrIfOl2e/yFXSvRLM=
github.com/mattn/go-isatty v0.0.20 h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=
github.com/mattn/go-isatty v0.0.20/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=
github.com/mattn/go-sqlite3 v1.14.22 h1:2gZY6PC6kBnID23Tichd1K+Z0oS6nE/XwU+Vz/5o4kU=
github.com/mattn/go-sqlite3 v1.14.22/go.mod h1:Uh1q+B4BYcTPb+yiD3kU8Ct7aC0hY9fxUwlHK0RXw+Y=
github.com/maypok86/otter v1.2.1 h1:xyvMW+t0vE1sKt/++GTkznLitEl7D/msqXkAbLwiC1M=
github.com/maypok86/otter v1.2.1/go.mod h1:mKLfoI7v1HOmQMwFgX4QkRk23mX6ge3RDvjdHOWG4R4=
github.com/mikestefanello/backlite v0.1.0 h1:bIiZJXPZB8V5PXWvDmkTepY015w3gJdeRrP3QrEV4Ls=
github.com/mikestefanello/backlite v0.1.0/go.mod h1:/vj8LPZWG/xqK/3uHaqOtu5JRLDEWqeyJKWTAlADTV0=
github.com/mitchellh/copystructure v1.2.0 h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=
github.com/mitchellh/copystructure v1.2.0/go.mod h1:qLl+cE2AmVv+CoeAwDPye/v+N2HKCj9FbZEVFJRxO9s=
github.com/mitchellh/go-wordwrap v1.0.1 h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=
github.com/mitchellh/go-wordwrap v1.0.1/go.mod h1:R62XHJLzvMFRBbcrT7m7WgmE1eOyTSsCt+hzestvNj0=
github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/reflectwalk v1.0.2 h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=
github.com/mitchellh/reflectwalk v1.0.2/go.mod h1:mSTlrgnPZtwu0c4WaC2kGObEpuNDbx0jmZXqmk4esnw=
github.com/pelletier/go-toml/v2 v2.2.1 h1:9TA9+T8+8CUCO2+WYnDLCgrYi9+omqKXyjDtosvtEhg=
github.com/pelletier/go-toml/v2 v2.2.1/go.mod h1:1t835xjRzz80PqgE6HHgN2JOsmgYu/h4qDAS4n929Rs=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 h1:Jamvg5psRIccs7FGNTlIRMkT8wgtp5eCXdBlqhYGL6U=
github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/rogpeppe/go-internal v1.10.0 h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=
github.com/rogpeppe/go-internal v1.10.0/go.mod h1:UQnix2H7Ngw/k4C5ijL5+65zddjncjaFoBhdsK/akog=
github.com/sagikazarmark/locafero v0.4.0 h1:HApY1R9zGo4DBgr7dqsTH/JJxLTTsOt7u6keLGt6kNQ=
github.com/sagikazarmark/locafero v0.4.0/go.mod h1:Pe1W6UlPYUk/+wc/6KFhbORCfqzgYEpgQ3O5fPuL3H4=
github.com/sagikazarmark/slog-shim v0.1.0 h1:diDBnUNK9N/354PgrxMywXnAwEr1QZcOr6gto+ugjYE=
github.com/sagikazarmark/slog-shim v0.1.0/go.mod h1:SrcSrq8aKtyuqEI1uvTDTK1arOWRIczQRv+GVI1AkeQ=
github.com/sourcegraph/conc v0.3.0 h1:OQTbbt6P72L20UqAkXXuLOj79LfEanQ+YQFNpLA9ySo=
github.com/sourcegraph/conc v0.3.0/go.mod h1:Sdozi7LEKbFPqYX2/J+iBAM6HpqSLTASQIKqDmF7Mt0=
github.com/spf13/afero v1.11.0 h1:WJQKhtpdm3v2IzqG8VMqrr6Rf3UYpEF239Jy9wNepM8=
github.com/spf13/afero v1.11.0/go.mod h1:GH9Y3pIexgf1MTIWtNGyogA5MwRIDXGUr+hbWNoBjkY=
github.com/spf13/cast v1.6.0 h1:GEiTHELF+vaR5dhz3VqZfFSzZjYbgeKDpBxQVS4GYJ0=
github.com/spf13/cast v1.6.0/go.mod h1:ancEpBxwJDODSW/UG4rDrAqiKolqNNh2DX3mk86cAdo=
github.com/spf13/pflag v1.0.5 h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=
github.com/spf13/pflag v1.0.5/go.mod h1:McXfInJRrz4CZXVZOBLb0bTZqETkiAhM9Iw0y3An2Bg=
github.com/spf13/viper v1.18.2 h1:LUXCnvUvSM6FXAsj6nnfc8Q2tp1dIgUfY9Kc8GsSOiQ=
github.com/spf13/viper v1.18.2/go.mod h1:EKmWIqdnk5lOcmR72yw6hS+8OPYcwD0jteitLMVB+yk=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.4.0/go.mod h1:YvHI0jy2hoMjB+UWwv71VJQ9isScKT/TqJzVSSt89Yw=
github.com/stretchr/objx v0.5.0/go.mod h1:Yh+to48EsGEfYuaHDzXPcE3xhTkx73EhmCGUpEOglKo=
github.com/stretchr/objx v0.5.2/go.mod h1:FRsXN1f5AsAjCGJKqEizvkpNtU+EGNCLh3NxZ/8L+MA=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.0/go.mod h1:yNjHg4UonilssWZ8iaSj1OCr/vHnekPRkoO+kdMU+MU=
github.com/stretchr/testify v1.8.4/go.mod h1:sz/lmYIOXD/1dqDmKjjqLyZ2RngseejIcXlSw2iwfAo=
github.com/stretchr/testify v1.9.0 h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=
github.com/stretchr/testify v1.9.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/subosito/gotenv v1.6.0 h1:9NlTDc1FTs4qu0DDq7AEtTPNw6SVm7uBMsUCUjABIf8=
github.com/subosito/gotenv v1.6.0/go.mod h1:Dk4QP5c2W3ibzajGcXpNraDfq2IrhjMIvMSWPKKo0FU=
github.com/valyala/bytebufferpool v1.0.0 h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=
github.com/valyala/bytebufferpool v1.0.0/go.mod h1:6bBcMArwyJ5K/AmCkWv1jt77kVWyCJ6HpOuEn7z0Csc=
github.com/valyala/fasttemplate v1.2.2 h1:lxLXG0uE3Qnshl9QyaK6XJxMXlQZELvChBOCmQD0Loo=
github.com/valyala/fasttemplate v1.2.2/go.mod h1:KHLXt3tVN2HBp8eijSv/kGJopbvo7S+qRAEEKiv+SiQ=
github.com/yuin/goldmark v1.4.13/go.mod h1:6yULJ656Px+3vBD8DxQVa3kxgyrAnzto9xy5taEt/CY=
github.com/zclconf/go-cty v1.14.4 h1:uXXczd9QDGsgu0i/QFR/hzI5NYCHLf6NQw/atrbnhq8=
github.com/zclconf/go-cty v1.14.4/go.mod h1:VvMs5i0vgZdhYawQNq5kePSpLAoz8u1xvZgrPIxfnZE=
github.com/zclconf/go-cty-debug v0.0.0-20191215020915-b22d67c1ba0b h1:FosyBZYxY34Wul7O/MSKey3txpPYyCqVO5ZyceuQJEI=
github.com/zclconf/go-cty-debug v0.0.0-20191215020915-b22d67c1ba0b/go.mod h1:ZRKQfBXbGkpdV6QMzT3rU1kSTAnfu1dO8dPKjYprgj8=
go.uber.org/multierr v1.11.0 h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=
go.uber.org/multierr v1.11.0/go.mod h1:20+QtiLqy0Nd6FdQB9TLXag12DsQkrbs3htMFfDN80Y=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20210921155107-089bfa567519/go.mod h1:GvvjBRRGRdwPK5ydBHafDWAxML/pGHZbMvKqRZ5+Abc=
golang.org/x/crypto v0.22.0 h1:g1v0xeRhjcugydODzvb3mEM9SQ0HGp9s/nh3COQ/C30=
golang.org/x/crypto v0.22.0/go.mod h1:vr6Su+7cTlO45qkww3VDJlzDn0ctJvRgYbC2NvXHt+M=
golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f h1:99ci1mjWVBWwJiEKYY6jWa4d2nTQVIEhZIptnrVb1XY=
golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f/go.mod h1:/lliqkxwWAhPjf5oSOIJup2XcqJaw8RGS6k3TGEc7GI=
golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4/go.mod h1:jJ57K6gSWd91VN4djpZkiMVwK6gcyfeH4XE8wZrZaV4=
golang.org/x/mod v0.8.0/go.mod h1:iBbtSCu2XBx23ZKBPSOrRkjjQPZFPuis4dIYUhu/chs=
golang.org/x/mod v0.17.0 h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=
golang.org/x/mod v0.17.0/go.mod h1:hTbmBsO62+eylJbnUtE2MGJUyE7QWk4xUqPFrRgJ+7c=
golang.org/x/net v0.0.0-20190620200207-3b0461eec859/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
golang.org/x/net v0.0.0-20210226172049-e18ecbb05110/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
golang.org/x/net v0.0.0-20220722155237-a158d28d115b/go.mod h1:XRhObCWvk6IyKnWLug+ECip1KBveYUHfp+8e9klMJ9c=
golang.org/x/net v0.6.0/go.mod h1:2Tu9+aMcznHK/AK1HMvgo6xiTLG5rD5rZLDS+rp2Bjs=
golang.org/x/net v0.9.0/go.mod h1:d48xBJpPfHeWQsugry2m+kC02ZBRGRgulfHnEXEuWns=
golang.org/x/net v0.24.0 h1:1PcaxkF854Fu3+lvBIx5SYn9wRlBzzcnHZSiaFFAb0w=
golang.org/x/net v0.24.0/go.mod h1:2Q7sJY5mzlzWjKtYUEXSlBWCdyaioyXzRB2RtU8KVE8=
golang.org/x/oauth2 v0.21.0 h1:tsimM75w1tF/uws5rbeHzIWxEqElMehnc+iW793zsZs=
golang.org/x/oauth2 v0.21.0/go.mod h1:XYTD2NtWslqkgxebSiOHnXEap4TF09sJSc7H1sXbhtI=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.1.0/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.7.0 h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=
golang.org/x/sync v0.7.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.5.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.7.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.19.0 h1:q5f1RH2jigJ1MoAWp2KTp3gm5zAGFUTarQZ5U386+4o=
golang.org/x/sys v0.19.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/term v0.0.0-20210927222741-03fcf44c2211/go.mod h1:jbD1KX2456YbFQfuXm/mYQcufACuNUgVhRMnK/tPxf8=
golang.org/x/term v0.5.0/go.mod h1:jMB1sMXY+tzblOD4FWmEbocvup2/aLOaQEp7JmGp78k=
golang.org/x/term v0.7.0/go.mod h1:P32HKFT3hSsZrRxla30E9HqToFYAQPCMs/zFMBUFqPY=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/text v0.7.0/go.mod h1:mrYo+phRRbMaCq/xk9113O4dZlRixOauAjOtrjsXDZ8=
golang.org/x/text v0.9.0/go.mod h1:e1OnstbJyHTd6l/uOt8jFFHp6TRDWZR/bV3emEE/zU8=
golang.org/x/text v0.14.0 h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=
golang.org/x/text v0.14.0/go.mod h1:18ZOQIKpY8NJVqYksKHtTdi31H5itFRjB5/qKTNYzSU=
golang.org/x/time v0.5.0 h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=
golang.org/x/time v0.5.0/go.mod h1:3BpzKBy/shNhVucY/MWOyx10tF3SFh9QdLuxbVysPQM=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
golang.org/x/tools v0.1.12/go.mod h1:hNGJHUnrk76NpqgfD5Aqm5Crs+Hm0VOH/i9J2+nxYbc=
golang.org/x/tools v0.6.0/go.mod h1:Xwgl3UAJ/d3gWutnCtw505GrjyAbvKui8lOU390QaIU=
golang.org/x/tools v0.20.0 h1:hz/CVckiOxybQvFw6h7b/q80NTr9IUQb4s1IIzW7KNY=
golang.org/x/tools v0.20.0/go.mod h1:WvitBU7JJf6A4jOdg4S1tviW9bhUxkgeCui/0JHctQg=
golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/ini.v1 v1.67.0 h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=
gopkg.in/ini.v1 v1.67.0/go.mod h1:pNLf8WUiyNEtQjuu5G5vTm06TEv9tsIgeAvK8hOrP4k=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

`pkg/context/context.go`

```go
package context

import (
	"context"
	"errors"
)

const (
	// AuthenticatedUserKey is the key value used to store the authenticated user in context
	AuthenticatedUserKey = "auth_user"

	// UserKey is the key value used to store a user in context
	UserKey = "user"

	// FormKey is the key value used to store a form in context
	FormKey = "form"

	// PasswordTokenKey is the key value used to store a password token in context
	PasswordTokenKey = "password_token"

	// LoggerKey is the key value used to store a structured logger in context
	LoggerKey = "logger"

	// SessionKey is the key value used to store the session data in context
	SessionKey = "session"
)

// IsCanceledError determines if an error is due to a context cancelation
func IsCanceledError(err error) bool {
	return errors.Is(err, context.Canceled)
}
```go

`pkg/context/context_test.go`

```go
package context

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	assert.False(t, IsCanceledError(ctx.Err()))
	cancel()
	assert.True(t, IsCanceledError(ctx.Err()))

	ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond*5)
	<-ctx.Done()
	cancel()
	assert.False(t, IsCanceledError(ctx.Err()))

	assert.False(t, IsCanceledError(errors.New("test error")))
}
```go

`pkg/form/form.go`

```go
package form

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// Form represents a form that can be submitted and validated
type Form interface {
	// Submit marks the form as submitted, stores a pointer to it in the context, binds the request
	// values to the struct fields, and validates the input based on the struct tags.
	// Returns a validator.ValidationErrors if the form values were not valid.
	// Returns an echo.HTTPError if the request failed to process.
	Submit(c echo.Context, form any) error

	// IsSubmitted returns true if the form was submitted
	IsSubmitted() bool

	// IsValid returns true if the form has no validation errors
	IsValid() bool

	// IsDone returns true if the form was submitted and has no validation errors
	IsDone() bool

	// FieldHasErrors returns true if a given struct field has validation errors
	FieldHasErrors(fieldName string) bool

	// SetFieldError sets a validation error message for a given struct field
	SetFieldError(fieldName string, message string)

	// GetFieldErrors returns the validation errors for a given struct field
	GetFieldErrors(fieldName string) []string

	// GetFieldStatusClass returns a CSS class to be used for a given struct field
	GetFieldStatusClass(fieldName string) string
}

// Get gets a form from the context or initializes a new copy if one is not set
func Get[T any](ctx echo.Context) *T {
	if v := ctx.Get(context.FormKey); v != nil {
		return v.(*T)
	}
	var v T
	return &v
}

// Clear removes the form set in the context
func Clear(ctx echo.Context) {
	ctx.Set(context.FormKey, nil)
}

// Submit submits a form
// See Form.Submit()
func Submit(ctx echo.Context, form Form) error {
	return form.Submit(ctx, form)
}
```go

`pkg/form/form_test.go`

```go
package form

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockForm struct {
	called bool
	Submission
}

func (m *mockForm) Submit(_ echo.Context, _ any) error {
	m.called = true
	return nil
}

func TestSubmit(t *testing.T) {
	m := mockForm{}
	ctx, _ := tests.NewContext(echo.New(), "/")
	err := Submit(ctx, &m)
	require.NoError(t, err)
	assert.True(t, m.called)
}

func TestGetClear(t *testing.T) {
	e := echo.New()

	type example struct {
		Name string `form:"name"`
	}

	t.Run("get empty context", func(t *testing.T) {
		// Empty context, still return a form
		ctx, _ := tests.NewContext(e, "/")
		form := Get[example](ctx)
		assert.NotNil(t, form)
	})

	t.Run("get non-empty context", func(t *testing.T) {
		form := example{
			Name: "test",
		}
		ctx, _ := tests.NewContext(e, "/")
		ctx.Set(context.FormKey, &form)

		// Get again and expect the values were stored
		got := Get[example](ctx)
		require.NotNil(t, got)
		assert.Equal(t, "test", form.Name)

		// Clear
		Clear(ctx)
		got = Get[example](ctx)
		require.NotNil(t, got)
		assert.Empty(t, got.Name)
	})
}
```go

`pkg/form/submission.go`

```go
package form

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mikestefanello/pagoda/pkg/context"

	"github.com/labstack/echo/v4"
)

// Submission represents the state of the submission of a form, not including the form itself.
// This satisfies the Form interface.
type Submission struct {
	// isSubmitted indicates if the form has been submitted
	isSubmitted bool

	// errors stores a slice of error message strings keyed by form struct field name
	errors map[string][]string
}

func (f *Submission) Submit(ctx echo.Context, form any) error {
	f.isSubmitted = true

	// Set in context so the form can later be retrieved
	ctx.Set(context.FormKey, form)

	// Bind the values from the incoming request to the form struct
	if err := ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to bind form: %v", err))
	}

	// Validate the form
	if err := ctx.Validate(form); err != nil {
		f.setErrorMessages(err)
		return err
	}

	return nil
}

func (f *Submission) IsSubmitted() bool {
	return f.isSubmitted
}

func (f *Submission) IsValid() bool {
	if f.errors == nil {
		return true
	}
	return len(f.errors) == 0
}

func (f *Submission) IsDone() bool {
	return f.IsSubmitted() && f.IsValid()
}

func (f *Submission) FieldHasErrors(fieldName string) bool {
	return len(f.GetFieldErrors(fieldName)) > 0
}

func (f *Submission) SetFieldError(fieldName string, message string) {
	if f.errors == nil {
		f.errors = make(map[string][]string)
	}
	f.errors[fieldName] = append(f.errors[fieldName], message)
}

func (f *Submission) GetFieldErrors(fieldName string) []string {
	if f.errors == nil {
		return []string{}
	}
	return f.errors[fieldName]
}

func (f *Submission) GetFieldStatusClass(fieldName string) string {
	if f.isSubmitted {
		if f.FieldHasErrors(fieldName) {
			return "is-danger"
		}
		return "is-success"
	}
	return ""
}

// setErrorMessages sets errors messages on the submission for all fields that failed validation
func (f *Submission) setErrorMessages(err error) {
	// Only this is supported right now
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return
	}

	for _, ve := range ves {
		var message string

		// Provide better error messages depending on the failed validation tag
		// This should be expanded as you use additional tags in your validation
		switch ve.Tag() {
		case "required":
			message = "This field is required."
		case "email":
			message = "Enter a valid email address."
		case "eqfield":
			message = "Does not match."
		case "gte":
			message = fmt.Sprintf("Must be greater than or equal to %v.", ve.Param())
		default:
			message = "Invalid value."
		}

		// Add the error
		f.SetFieldError(ve.Field(), message)
	}
}
```go

`pkg/form/submission_test.go`

```go
package form

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormSubmission(t *testing.T) {
	type formTest struct {
		Name  string `form:"name" validate:"required"`
		Email string `form:"email" validate:"required,email"`
		Submission
	}

	e := echo.New()
	e.Validator = services.NewValidator()

	t.Run("valid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("email=a@a.com"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx := e.NewContext(req, httptest.NewRecorder())

		var form formTest
		err := form.Submit(ctx, &form)
		assert.IsType(t, validator.ValidationErrors{}, err)

		assert.Empty(t, form.Name)
		assert.Equal(t, "a@a.com", form.Email)
		assert.False(t, form.IsValid())
		assert.True(t, form.FieldHasErrors("Name"))
		assert.False(t, form.FieldHasErrors("Email"))
		require.Len(t, form.GetFieldErrors("Name"), 1)
		assert.Len(t, form.GetFieldErrors("Email"), 0)
		assert.Equal(t, "This field is required.", form.GetFieldErrors("Name")[0])
		assert.Equal(t, "is-danger", form.GetFieldStatusClass("Name"))
		assert.Equal(t, "is-success", form.GetFieldStatusClass("Email"))
		assert.False(t, form.IsDone())

		formInCtx := Get[formTest](ctx)
		require.NotNil(t, formInCtx)
		assert.Equal(t, form.Email, formInCtx.Email)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("abc=abc"))
		ctx := e.NewContext(req, httptest.NewRecorder())
		var form formTest
		err := form.Submit(ctx, &form)
		assert.IsType(t, new(echo.HTTPError), err)
	})
}
```go

`pkg/funcmap/funcmap.go`

```go
package funcmap

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/mikestefanello/pagoda/config"
)

var (
	// CacheBuster stores a random string used as a cache buster for static files.
	CacheBuster = random.String(10)
)

type funcMap struct {
	web *echo.Echo
}

// NewFuncMap provides a template function map
func NewFuncMap(web *echo.Echo) template.FuncMap {
	fm := &funcMap{web: web}

	// See http://masterminds.github.io/sprig/ for all provided funcs
	funcs := sprig.FuncMap()

	// Include all the custom functions
	funcs["hasField"] = fm.hasField
	funcs["file"] = fm.file
	funcs["link"] = fm.link
	funcs["url"] = fm.url

	return funcs
}

// hasField checks if an interface contains a given field
func (fm *funcMap) hasField(v any, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// file appends a cache buster to a given filepath so it can remain cached until the app is restarted
func (fm *funcMap) file(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, CacheBuster)
}

// link outputs HTML for a link element, providing the ability to dynamically set the active class
func (fm *funcMap) link(url, text, currentPath string, classes ...string) template.HTML {
	if currentPath == url {
		classes = append(classes, "is-active")
	}

	html := fmt.Sprintf(`<a class="%s" href="%s">%s</a>`, strings.Join(classes, " "), url, text)
	return template.HTML(html)
}

// url generates a URL from a given route name and optional parameters
func (fm *funcMap) url(routeName string, params ...any) string {
	return fm.web.Reverse(routeName, params...)
}
```go

`pkg/funcmap/funcmap_test.go`

```go
package funcmap

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"

	"github.com/stretchr/testify/assert"
)

func TestNewFuncMap(t *testing.T) {
	f := NewFuncMap(echo.New())
	assert.NotNil(t, f["hasField"])
	assert.NotNil(t, f["link"])
	assert.NotNil(t, f["file"])
	assert.NotNil(t, f["url"])
}

func TestHasField(t *testing.T) {
	type example struct {
		name string
	}
	var e example
	f := new(funcMap)
	assert.True(t, f.hasField(e, "name"))
	assert.False(t, f.hasField(e, "abcd"))
}

func TestLink(t *testing.T) {
	f := new(funcMap)

	link := string(f.link("/abc", "Text", "/abc"))
	expected := `<a class="is-active" href="/abc">Text</a>`
	assert.Equal(t, expected, link)

	link = string(f.link("/abc", "Text", "/abc", "first", "second"))
	expected = `<a class="first second is-active" href="/abc">Text</a>`
	assert.Equal(t, expected, link)

	link = string(f.link("/abc", "Text", "/def"))
	expected = `<a class="" href="/abc">Text</a>`
	assert.Equal(t, expected, link)
}

func TestFile(t *testing.T) {
	f := new(funcMap)

	file := f.file("test.png")
	expected := fmt.Sprintf("/%s/test.png?v=%s", config.StaticPrefix, CacheBuster)
	assert.Equal(t, expected, file)
}

func TestUrl(t *testing.T) {
	f := new(funcMap)
	f.web = echo.New()
	f.web.GET("/mypath/:id", func(c echo.Context) error {
		return nil
	}).Name = "test"
	out := f.url("test", 5)
	assert.Equal(t, "/mypath/5", out)
}
```go

`pkg/handlers/auth.go`

```go
package handlers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/redirect"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

const (
	routeNameForgotPassword       = "forgot_password"
	routeNameForgotPasswordSubmit = "forgot_password.submit"
	routeNameLogin                = "login"
	routeNameLoginSubmit          = "login.submit"
	routeNameLogout               = "logout"
	routeNameRegister             = "register"
	routeNameRegisterSubmit       = "register.submit"
	routeNameResetPassword        = "reset_password"
	routeNameResetPasswordSubmit  = "reset_password.submit"
	routeNameVerifyEmail          = "verify_email"
)

type (
	Auth struct {
		auth *services.AuthClient
		mail *services.MailClient
		orm  *ent.Client
		*services.TemplateRenderer
	}

	forgotPasswordForm struct {
		Email string `form:"email" validate:"required,email"`
		form.Submission
	}

	loginForm struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
		form.Submission
	}

	registerForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}

	resetPasswordForm struct {
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}
)

func init() {
	Register(new(Auth))
}

func (h *Auth) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.orm = c.ORM
	h.auth = c.Auth
	h.mail = c.Mail
	return nil
}

func (h *Auth) Routes(g *echo.Group) {
	g.GET("/logout", h.Logout, middleware.RequireAuthentication()).Name = routeNameLogout
	g.GET("/email/verify/:token", h.VerifyEmail).Name = routeNameVerifyEmail

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	noAuth.GET("/login", h.LoginPage).Name = routeNameLogin
	noAuth.POST("/login", h.LoginSubmit).Name = routeNameLoginSubmit
	noAuth.GET("/register", h.RegisterPage).Name = routeNameRegister
	noAuth.POST("/register", h.RegisterSubmit).Name = routeNameRegisterSubmit
	noAuth.GET("/password", h.ForgotPasswordPage).Name = routeNameForgotPassword
	noAuth.POST("/password", h.ForgotPasswordSubmit).Name = routeNameForgotPasswordSubmit

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(h.orm),
		middleware.LoadValidPasswordToken(h.auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", h.ResetPasswordPage).Name = routeNameResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", h.ResetPasswordSubmit).Name = routeNameResetPasswordSubmit
}

func (h *Auth) ForgotPasswordPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageForgotPassword
	p.Title = "Forgot password"
	p.Form = form.Get[forgotPasswordForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var input forgotPasswordForm

	succeed := func() error {
		form.Clear(ctx)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return h.ForgotPasswordPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ForgotPasswordPage(ctx)
	default:
		return err
	}

	// Attempt to load the user
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return fail(err, "error querying user during forgot password")
	}

	// Generate the token
	token, pt, err := h.auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return fail(err, "error generating password reset token")
	}

	log.Ctx(ctx).Info("generated password reset token",
		"user_id", u.ID,
	)

	// Email the user
	url := ctx.Echo().Reverse(routeNameResetPassword, u.ID, pt.ID, token)
	err = h.mail.
		Compose().
		To(u.Email).
		Subject("Reset your password").
		Body(fmt.Sprintf("Go here to reset your password: %s", url)).
		Send(ctx)

	if err != nil {
		return fail(err, "error sending password reset email")
	}

	return succeed()
}

func (h *Auth) LoginPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageLogin
	p.Title = "Log in"
	p.Form = form.Get[loginForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) LoginSubmit(ctx echo.Context) error {
	var input loginForm

	authFailed := func() error {
		input.SetFieldError("Email", "")
		input.SetFieldError("Password", "")
		msg.Danger(ctx, "Invalid credentials. Please try again.")
		return h.LoginPage(ctx)
	}

	oauth2Required := func() error {
		input.SetFieldError("Email", "")
		input.SetFieldError("Password", "")
		msg.Danger(ctx, "Account requires OAuth2 login. Please try again.")
		return h.LoginPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.LoginPage(ctx)
	default:
		return err
	}

	// Attempt to load the user
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return authFailed()
	case nil:
	default:
		return fail(err, "error querying user during login")
	}

	if u.Password == "__oauth2__" {
		return oauth2Required()
	}

	// Check if the password is correct
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, <strong>%s</strong>. You are now logged in.", u.Name))

	return redirect.New(ctx).
		Route(routeNameHome).
		Go()
}

func (h *Auth) Logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err == nil {
		msg.Success(ctx, "You have been logged out successfully.")
	} else {
		msg.Danger(ctx, "An error occurred. Please try again.")
	}
	return redirect.New(ctx).
		Route(routeNameHome).
		Go()
}

func (h *Auth) RegisterPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageRegister
	p.Title = "Register"
	p.Form = form.Get[registerForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) RegisterSubmit(ctx echo.Context) error {
	var input registerForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.RegisterPage(ctx)
	default:
		return err
	}

	// Hash the password
	pwHash, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return fail(err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := h.orm.User.
		Create().
		SetName(input.Name).
		SetEmail(input.Email).
		SetPassword(pwHash).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		log.Ctx(ctx).Info("user created",
			"user_name", u.Name,
			"user_id", u.ID,
		)
	case *ent.ConstraintError:
		msg.Warning(ctx, "A user with this email address already exists. Please log in.")
		return redirect.New(ctx).
			Route(routeNameLogin).
			Go()
	default:
		return fail(err, "unable to create user")
	}

	// Log the user in
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		log.Ctx(ctx).Error("unable to log user in",
			"error", err,
			"user_id", u.ID,
		)
		msg.Info(ctx, "Your account has been created.")
		return redirect.New(ctx).
			Route(routeNameLogin).
			Go()
	}

	msg.Success(ctx, "Your account has been created. You are now logged in.")

	// Send the verification email
	h.sendVerificationEmail(ctx, u)

	return redirect.New(ctx).
		Route(routeNameHome).
		Go()
}

func (h *Auth) sendVerificationEmail(ctx echo.Context, usr *ent.User) {
	// Generate a token
	token, err := h.auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		log.Ctx(ctx).Error("unable to generate email verification token",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	// Send the email
	url := ctx.Echo().Reverse(routeNameVerifyEmail, token)
	err = h.mail.
		Compose().
		To(usr.Email).
		Subject("Confirm your email address").
		Body(fmt.Sprintf("Click here to confirm your email address: %s", url)).
		Send(ctx)

	if err != nil {
		log.Ctx(ctx).Error("unable to send email verification link",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	msg.Info(ctx, "An email was sent to you to verify your email address.")
}

func (h *Auth) ResetPasswordPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageResetPassword
	p.Title = "Reset password"
	p.Form = form.Get[resetPasswordForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var input resetPasswordForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ResetPasswordPage(ctx)
	default:
		return err
	}

	// Hash the new password
	hash, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return fail(err, "unable to hash password")
	}

	// Get the requesting user
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user
	_, err = usr.
		Update().
		SetPassword(hash).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to update password")
	}

	// Delete all password tokens for this user
	err = h.auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return fail(err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return redirect.New(ctx).
		Route(routeNameLogin).
		Go()
}

func (h *Auth) VerifyEmail(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := h.auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return redirect.New(ctx).
			Route(routeNameHome).
			Go()
	}

	// Check if it matches the authenticated user
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed
	if usr == nil {
		usr, err = h.orm.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed
	if !usr.Verified {
		usr, err = usr.
			Update().
			SetVerified(true).
			Save(ctx.Request().Context())

		if err != nil {
			return fail(err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")
	return redirect.New(ctx).
		Route(routeNameHome).
		Go()
}
```go

`pkg/handlers/error.go`

```go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

type Error struct {
	*services.TemplateRenderer
}

func (e *Error) Page(err error, ctx echo.Context) {
	if ctx.Response().Committed || context.IsCanceledError(err) {
		return
	}

	// Determine the error status code
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// Log the error
	logger := log.Ctx(ctx)
	switch {
	case code >= 500:
		logger.Error(err.Error())
	case code >= 400:
		logger.Warn(err.Error())
	}

	// Render the error page
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageError
	p.Title = http.StatusText(code)
	p.StatusCode = code
	p.HTMX.Request.Enabled = false

	if err = e.RenderPage(ctx, p); err != nil {
		log.Ctx(ctx).Error("failed to render error page",
			"error", err,
		)
	}
}
```go

`pkg/handlers/handlers.go`

```go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
)

var handlers []Handler

// Handler handles one or more HTTP routes
type Handler interface {
	// Routes allows for self-registration of HTTP routes on the router
	Routes(g *echo.Group)

	// Init provides the service container to initialize
	Init(*services.Container) error
}

// Register registers a handler
func Register(h Handler) {
	handlers = append(handlers, h)
}

// GetHandlers returns all handlers
func GetHandlers() []Handler {
	return handlers
}

// fail is a helper to fail a request by returning a 500 error and logging the error
func fail(err error, log string) error {
	// The error handler will handle logging
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", log, err))
}
```go

`pkg/handlers/handlers_test.go`

```go
package handlers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSetHandlers(t *testing.T) {
	handlers = []Handler{}
	assert.Empty(t, GetHandlers())
	h := new(Pages)
	Register(h)
	got := GetHandlers()
	require.Len(t, got, 1)
	assert.Equal(t, h, got[0])
}

func TestFail(t *testing.T) {
	err := fail(errors.New("err message"), "log message")
	require.IsType(t, new(echo.HTTPError), err)
	he := err.(*echo.HTTPError)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "log message: err message", he.Message)
}
```go

`pkg/handlers/oauth.go`

```go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type (
	OAuth struct {
		auth  *services.AuthClient
		oauth *services.OAuthClient
		*services.TemplateRenderer
	}
)

func init() {
	Register(new(OAuth))
}

func (h *OAuth) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.auth = c.Auth
	h.oauth = c.OAuth
	return nil
}

func (h *OAuth) Routes(g *echo.Group) {
	g.GET("/auth/google", h.GoogleLogin)
	g.GET("/auth/google/callback", h.GoogleCallback)
	g.GET("/auth/facebook", h.FacebookLogin)
	g.GET("/auth/facebook/callback", h.FacebookCallback)
}

func (h *OAuth) GoogleLogin(c echo.Context) error {
	return h.oauthLogin(c, "google")
}

func (h *OAuth) GoogleCallback(c echo.Context) error {
	return h.oauthCallback(c, "google")
}

func (h *OAuth) FacebookLogin(c echo.Context) error {
	return h.oauthLogin(c, "facebook")
}

func (h *OAuth) FacebookCallback(c echo.Context) error {
	return h.oauthCallback(c, "facebook")
}

func (h *OAuth) oauthLogin(c echo.Context, provider string) error {
	url := h.oauth.GetAuthCodeURL(provider, c)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OAuth) oauthCallback(c echo.Context, provider string) error {
	user, err := h.oauth.HandleCallback(provider, c)
	if err != nil {
		msg.Danger(c, fmt.Sprintf("Failed to authenticate with %s: %v", provider, err))
		return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("login"))
	}

	err = h.auth.Login(c, user.ID)
	if err != nil {
		return err
	}

	msg.Success(c, fmt.Sprintf("Successfully logged in with %s", provider))
	return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("home"))
}
```go

`pkg/handlers/pages.go`

```go
package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
	"gopkg.in/yaml.v3"
)

const (
	routeNameAbout  = "about"
	routeNameHome   = "home"
	routeNameConfig = "config"
)

type (
	Pages struct {
		*services.TemplateRenderer
		Config *config.Config
	}

	post struct {
		Title string
		Body  string
	}
)

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.Config = c.Config
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home).Name = routeNameHome
	g.GET("/config", h.GetConfig, middleware.RequireLocalEnv(h.Config)).Name = routeNameConfig
}

func (h *Pages) Home(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageHome
	p.Metatags.Description = "Welcome to the homepage."
	p.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	p.Pager = page.NewPager(ctx, 4)
	p.Data = h.fetchPosts(&p.Pager)

	return h.RenderPage(ctx, p)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (h *Pages) fetchPosts(pager *page.Pager) []post {
	pager.SetItems(20)
	posts := make([]post, 20)

	for k := range posts {
		posts[k] = post{
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}

func (h *Pages) GetConfig(ctx echo.Context) error {
	if h.Config.App.Environment != config.EnvLocal {
		return echo.ErrNotFound
	}

	yamlData, err := yaml.Marshal(h.Config)
	if err != nil {
		return echo.ErrInternalServerError
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageConfig
	p.Title = "Application Configuration"
	p.Data = string(yamlData)

	return h.RenderPage(ctx, p)
}
```go

`pkg/handlers/pages_test.go`

```go
package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple example of how to test routes and their markup using the test HTTP server spun up within
// this test package
func TestPages__About(t *testing.T) {
	doc := request(t).
		setRoute(routeNameAbout).
		get().
		assertStatusCode(http.StatusOK).
		toDoc()

	// Goquery is an excellent package to use for testing HTML markup
	h1 := doc.Find("h1.title")
	assert.Len(t, h1.Nodes, 1)
	assert.Equal(t, "About", h1.Text())
}
```go

`pkg/handlers/router.go`

```go
package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) error {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non-static file route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		middleware.Session(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.TemplateRenderer),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	// Error handler
	err := Error{c.TemplateRenderer}
	c.Web.HTTPErrorHandler = err.Page

	// Initialize and register all handlers
	for _, h := range GetHandlers() {
		if err := h.Init(c); err != nil {
			return err
		}

		h.Routes(g)
	}

	return nil
}
```go

`pkg/handlers/router_test.go`

```go
package handlers

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

var (
	srv *httptest.Server
	c   *services.Container
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Start a new container
	c = services.NewContainer()

	// Start a test HTTP server
	if err := BuildRouter(c); err != nil {
		panic(err)
	}
	srv = httptest.NewServer(c.Web)

	// Run tests
	exitVal := m.Run()

	// Shutdown the container and test server
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
	srv.Close()

	os.Exit(exitVal)
}

type httpRequest struct {
	route  string
	client http.Client
	body   url.Values
	t      *testing.T
}

func request(t *testing.T) *httpRequest {
	jar, err := cookiejar.New(nil)
	require.NoError(t, err)
	r := httpRequest{
		t:    t,
		body: url.Values{},
		client: http.Client{
			Jar: jar,
		},
	}
	return &r
}

func (h *httpRequest) setClient(client http.Client) *httpRequest {
	h.client = client
	return h
}

func (h *httpRequest) setRoute(route string, params ...any) *httpRequest {
	h.route = srv.URL + c.Web.Reverse(route, params)
	return h
}

func (h *httpRequest) setBody(body url.Values) *httpRequest {
	h.body = body
	return h
}

func (h *httpRequest) get() *httpResponse {
	resp, err := h.client.Get(h.route)
	require.NoError(h.t, err)
	r := httpResponse{
		t:        h.t,
		Response: resp,
	}
	return &r
}

func (h *httpRequest) post() *httpResponse {
	// Make a get request to get the CSRF token
	doc := h.get().
		assertStatusCode(http.StatusOK).
		toDoc()

	// Extract the CSRF and include it in the POST request body
	csrf := doc.Find(`input[name="csrf"]`).First()
	token, exists := csrf.Attr("value")
	assert.True(h.t, exists)
	h.body["csrf"] = []string{token}

	// Make the POST requests
	resp, err := h.client.PostForm(h.route, h.body)
	require.NoError(h.t, err)
	r := httpResponse{
		t:        h.t,
		Response: resp,
	}
	return &r
}

type httpResponse struct {
	*http.Response
	t *testing.T
}

func (h *httpResponse) assertStatusCode(code int) *httpResponse {
	assert.Equal(h.t, code, h.Response.StatusCode)
	return h
}

func (h *httpResponse) assertRedirect(t *testing.T, route string, params ...any) *httpResponse {
	assert.Equal(t, c.Web.Reverse(route, params), h.Header.Get("Location"))
	return h
}

func (h *httpResponse) toDoc() *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(h.Body)
	require.NoError(h.t, err)
	err = h.Body.Close()
	assert.NoError(h.t, err)
	return doc
}
```go

`pkg/htmx/htmx.go`

```go
package htmx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request headers: https://htmx.org/docs/#request-headers
const (
	HeaderBoosted               = "HX-Boosted"
	HeaderHistoryRestoreRequest = "HX-History-Restore-Request"
	HeaderPrompt                = "HX-Prompt"
	HeaderRequest               = "HX-Request"
	HeaderTarget                = "HX-Target"
	HeaderTrigger               = "HX-Trigger"
	HeaderTriggerName           = "HX-Trigger-Name"
)

// Response headers: https://htmx.org/docs/#response-headers
const (
	HeaderPushURL            = "HX-Push-Url"
	HeaderRedirect           = "HX-Redirect"
	HeaderReplaceURL         = "HX-Replace-Url"
	HeaderRefresh            = "HX-Refresh"
	HeaderTriggerAfterSettle = "HX-Trigger-After-Settle"
	HeaderTriggerAfterSwap   = "HX-Trigger-After-Swap"
)

type (
	// Request contains data that HTMX provides during requests
	Request struct {
		Enabled        bool
		Boosted        bool
		HistoryRestore bool
		Trigger        string
		TriggerName    string
		Target         string
		Prompt         string
	}

	// Response contain data that the server can communicate back to HTMX
	Response struct {
		PushURL            string
		Redirect           string
		Refresh            bool
		ReplaceURL         string
		Trigger            string
		TriggerAfterSwap   string
		TriggerAfterSettle string
		NoContent          bool
	}
)

// GetRequest extracts HTMX data from the request
func GetRequest(ctx echo.Context) Request {
	return Request{
		Enabled:        ctx.Request().Header.Get(HeaderRequest) == "true",
		Boosted:        ctx.Request().Header.Get(HeaderBoosted) == "true",
		Trigger:        ctx.Request().Header.Get(HeaderTrigger),
		TriggerName:    ctx.Request().Header.Get(HeaderTriggerName),
		Target:         ctx.Request().Header.Get(HeaderTarget),
		Prompt:         ctx.Request().Header.Get(HeaderPrompt),
		HistoryRestore: ctx.Request().Header.Get(HeaderHistoryRestoreRequest) == "true",
	}
}

// Apply applies data from a Response to a server response
func (r Response) Apply(ctx echo.Context) {
	if r.PushURL != "" {
		ctx.Response().Header().Set(HeaderPushURL, r.PushURL)
	}
	if r.Redirect != "" {
		ctx.Response().Header().Set(HeaderRedirect, r.Redirect)
	}
	if r.Refresh {
		ctx.Response().Header().Set(HeaderRefresh, "true")
	}
	if r.Trigger != "" {
		ctx.Response().Header().Set(HeaderTrigger, r.Trigger)
	}
	if r.TriggerAfterSwap != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSwap, r.TriggerAfterSwap)
	}
	if r.TriggerAfterSettle != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSettle, r.TriggerAfterSettle)
	}
	if r.ReplaceURL != "" {
		ctx.Response().Header().Set(HeaderReplaceURL, r.ReplaceURL)
	}
	if r.NoContent {
		ctx.Response().Status = http.StatusNoContent
	}
}
```go

`pkg/htmx/htmx_test.go`

```go
package htmx

import (
	"net/http"
	"testing"

	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestSetRequest(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	ctx.Request().Header.Set(HeaderRequest, "true")
	ctx.Request().Header.Set(HeaderBoosted, "true")
	ctx.Request().Header.Set(HeaderTrigger, "a")
	ctx.Request().Header.Set(HeaderTriggerName, "b")
	ctx.Request().Header.Set(HeaderTarget, "c")
	ctx.Request().Header.Set(HeaderPrompt, "d")
	ctx.Request().Header.Set(HeaderHistoryRestoreRequest, "true")

	r := GetRequest(ctx)
	assert.Equal(t, true, r.Enabled)
	assert.Equal(t, true, r.Boosted)
	assert.Equal(t, true, r.HistoryRestore)
	assert.Equal(t, "a", r.Trigger)
	assert.Equal(t, "b", r.TriggerName)
	assert.Equal(t, "c", r.Target)
	assert.Equal(t, "d", r.Prompt)
}

func TestResponse_Apply(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	r := Response{
		PushURL:            "a",
		Redirect:           "b",
		ReplaceURL:         "f",
		Refresh:            true,
		Trigger:            "c",
		TriggerAfterSwap:   "d",
		TriggerAfterSettle: "e",
		NoContent:          true,
	}
	r.Apply(ctx)

	assert.Equal(t, "a", ctx.Response().Header().Get(HeaderPushURL))
	assert.Equal(t, "b", ctx.Response().Header().Get(HeaderRedirect))
	assert.Equal(t, "true", ctx.Response().Header().Get(HeaderRefresh))
	assert.Equal(t, "c", ctx.Response().Header().Get(HeaderTrigger))
	assert.Equal(t, "d", ctx.Response().Header().Get(HeaderTriggerAfterSwap))
	assert.Equal(t, "e", ctx.Response().Header().Get(HeaderTriggerAfterSettle))
	assert.Equal(t, "f", ctx.Response().Header().Get(HeaderReplaceURL))
	assert.Equal(t, http.StatusNoContent, ctx.Response().Status)
}
```go

`pkg/log/log.go`

```go
package log

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// Set sets a logger in the context
func Set(ctx echo.Context, logger *slog.Logger) {
	ctx.Set(context.LoggerKey, logger)
}

// Ctx returns the logger stored in context, or provides the default logger if one is not present
func Ctx(ctx echo.Context) *slog.Logger {
	if l, ok := ctx.Get(context.LoggerKey).(*slog.Logger); ok {
		return l
	}

	return Default()
}

// Default returns the default logger
func Default() *slog.Logger {
	return slog.Default()
}
```go

`pkg/log/log_test.go`

```go
package log

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestCtxSet(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	logger := Ctx(ctx)
	assert.NotNil(t, logger)

	logger = logger.With("a", "b")
	Set(ctx, logger)

	got := Ctx(ctx)
	assert.Equal(t, got, logger)
}
```go

`pkg/middleware/auth.go`

```go
package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/labstack/echo/v4"
)

// LoadAuthenticatedUser loads the authenticated user, if one, and stores in context
func LoadAuthenticatedUser(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u, err := authClient.GetAuthenticatedUser(c)
			switch err.(type) {
			case *ent.NotFoundError:
				log.Ctx(c).Warn("auth user not found")
			case services.NotAuthenticatedError:
			case nil:
				c.Set(context.AuthenticatedUserKey, u)
			default:
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("error querying for authenticated user: %v", err),
				)
			}

			return next(c)
		}
	}
}

// LoadValidPasswordToken loads a valid password token entity that matches the user and token
// provided in path parameters
// If the token is invalid, the user will be redirected to the forgot password route
// This requires that the user owning the token is loaded in to context
func LoadValidPasswordToken(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the user parameter
			if c.Get(context.UserKey) == nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			usr := c.Get(context.UserKey).(*ent.User)

			// Extract the token ID
			tokenID, err := strconv.Atoi(c.Param("password_token"))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			// Attempt to load a valid password token
			token, err := authClient.GetValidPasswordToken(
				c,
				usr.ID,
				tokenID,
				c.Param("token"),
			)

			switch err.(type) {
			case nil:
				c.Set(context.PasswordTokenKey, token)
				return next(c)
			case services.InvalidPasswordTokenError:
				msg.Warning(c, "The link is either invalid or has expired. Please request a new one.")
				// TODO use the const for route name
				return c.Redirect(http.StatusFound, c.Echo().Reverse("forgot_password"))
			default:
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("error loading password token: %v", err),
				)
			}
		}
	}
}

// RequireAuthentication requires that the user be authenticated in order to proceed
func RequireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u == nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

// RequireNoAuthentication requires that the user not be authenticated in order to proceed
func RequireNoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u != nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
```go

`pkg/middleware/auth_test.go`

```go
package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestLoadAuthenticatedUser(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)
	mw := LoadAuthenticatedUser(c.Auth)

	// Not authenticated
	_ = tests.ExecuteMiddleware(ctx, mw)
	assert.Nil(t, ctx.Get(context.AuthenticatedUserKey))

	// Login
	err := c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)

	// Verify the midldeware returns the authenticated user
	_ = tests.ExecuteMiddleware(ctx, mw)
	require.NotNil(t, ctx.Get(context.AuthenticatedUserKey))
	ctxUsr, ok := ctx.Get(context.AuthenticatedUserKey).(*ent.User)
	require.True(t, ok)
	assert.Equal(t, usr.ID, ctxUsr.ID)
}

func TestRequireAuthentication(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Not logged in
	err := tests.ExecuteMiddleware(ctx, RequireAuthentication())
	tests.AssertHTTPErrorCode(t, err, http.StatusUnauthorized)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireAuthentication())
	assert.Nil(t, err)
}

func TestRequireNoAuthentication(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Not logged in
	err := tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	assert.Nil(t, err)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	tests.AssertHTTPErrorCode(t, err, http.StatusForbidden)
}

func TestLoadValidPasswordToken(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Missing user context
	err := tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	tests.AssertHTTPErrorCode(t, err, http.StatusInternalServerError)

	// Add user and password token context but no token and expect a redirect
	ctx.SetParamNames("user", "password_token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), "1")
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, ctx.Response().Status)

	// Add user context and invalid password token and expect a redirect
	ctx.SetParamNames("user", "password_token", "token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), "1", "faketoken")
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, ctx.Response().Status)

	// Create a valid token
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)

	// Add user and valid password token
	ctx.SetParamNames("user", "password_token", "token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), fmt.Sprintf("%d", pt.ID), token)
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.Nil(t, err)
	ctxPt, ok := ctx.Get(context.PasswordTokenKey).(*ent.PasswordToken)
	require.True(t, ok)
	assert.Equal(t, pt.ID, ctxPt.ID)
}
```go

`pkg/middleware/cache.go`

```go
package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/labstack/echo/v4"
)

// ServeCachedPage attempts to load a page from the cache by matching on the complete request URL
// If a page is cached for the requested URL, it will be served here and the request terminated.
// Any request made by an authenticated user or that is not a GET will be skipped.
func ServeCachedPage(t *services.TemplateRenderer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Skip non GET requests
			if ctx.Request().Method != http.MethodGet {
				return next(ctx)
			}

			// Skip if the user is authenticated
			if ctx.Get(context.AuthenticatedUserKey) != nil {
				return next(ctx)
			}

			// Attempt to load from cache
			page, err := t.GetCachedPage(ctx, ctx.Request().URL.String())

			if err != nil {
				switch {
				case errors.Is(err, services.ErrCacheMiss):
				case context.IsCanceledError(err):
					return nil
				default:
					log.Ctx(ctx).Error("failed getting cached page",
						"error", err,
					)
				}

				return next(ctx)
			}

			// Set any headers
			if page.Headers != nil {
				for k, v := range page.Headers {
					ctx.Response().Header().Set(k, v)
				}
			}

			log.Ctx(ctx).Debug("serving cached page")

			return ctx.HTMLBlob(page.StatusCode, page.HTML)
		}
	}
}

// CacheControl sets a Cache-Control header with a given max age
func CacheControl(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			v := "no-cache, no-store"
			if maxAge > 0 {
				v = fmt.Sprintf("public, max-age=%.0f", maxAge.Seconds())
			}
			ctx.Response().Header().Set("Cache-Control", v)
			return next(ctx)
		}
	}
}
```go

`pkg/middleware/cache_test.go`

```go
package middleware

import (
	"net/http"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/mikestefanello/pagoda/templates"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestServeCachedPage(t *testing.T) {
	// Cache a page
	ctx, rec := tests.NewContext(c.Web, "/cache")
	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageHome
	p.Cache.Enabled = true
	p.Cache.Expiration = time.Minute
	p.StatusCode = http.StatusCreated
	p.Headers["a"] = "b"
	p.Headers["c"] = "d"
	err := c.TemplateRenderer.RenderPage(ctx, p)
	output := rec.Body.Bytes()
	require.NoError(t, err)

	// Request the URL of the cached page
	ctx, rec = tests.NewContext(c.Web, "/cache")
	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.TemplateRenderer))
	assert.NoError(t, err)
	assert.Equal(t, p.StatusCode, ctx.Response().Status)
	assert.Equal(t, p.Headers["a"], ctx.Response().Header().Get("a"))
	assert.Equal(t, p.Headers["c"], ctx.Response().Header().Get("c"))
	assert.Equal(t, output, rec.Body.Bytes())

	// Login and try again
	tests.InitSession(ctx)
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))
	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.TemplateRenderer))
	assert.Nil(t, err)
}

func TestCacheControl(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_ = tests.ExecuteMiddleware(ctx, CacheControl(time.Second*5))
	assert.Equal(t, "public, max-age=5", ctx.Response().Header().Get("Cache-Control"))
	_ = tests.ExecuteMiddleware(ctx, CacheControl(0))
	assert.Equal(t, "no-cache, no-store", ctx.Response().Header().Get("Cache-Control"))
}
```go

`pkg/middleware/entity.go`

```go
package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"

	"github.com/labstack/echo/v4"
)

// LoadUser loads the user based on the ID provided as a path parameter
func LoadUser(orm *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := strconv.Atoi(c.Param("user"))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			u, err := orm.User.
				Query().
				Where(user.ID(userID)).
				Only(c.Request().Context())

			switch err.(type) {
			case nil:
				c.Set(context.UserKey, u)
				return next(c)
			case *ent.NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound)
			default:
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("error querying user: %v", err),
				)
			}
		}
	}
}
```go

`pkg/middleware/entity_test.go`

```go
package middleware

import (
	"fmt"
	"testing"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadUser(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	ctx.SetParamNames("user")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID))
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	ctxUsr, ok := ctx.Get(context.UserKey).(*ent.User)
	require.True(t, ok)
	assert.Equal(t, usr.ID, ctxUsr.ID)
}
```go

`pkg/middleware/local_env.go`

```go
package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
)

// RequireLocalEnv is a middleware that ensures the application is running in the local environment
func RequireLocalEnv(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cfg.App.Environment != config.EnvLocal {
				return echo.NewHTTPError(http.StatusNotFound)
			}
			return next(c)
		}
	}
}
```go

`pkg/middleware/log.go`

```go
package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/log"
)

// SetLogger initializes a logger for the current request and stores it in the context.
// It's recommended to have this executed after Echo's RequestID() middleware because it will add
// the request ID to the logger so that all log messages produced from this request have the
// request ID in it. You can modify this code to include any other fields that you want to always
// appear.
func SetLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Include the request ID in the logger
			rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
			logger := log.Ctx(ctx).With("request_id", rID)

			// TODO include other fields you may want in all logs for this request
			log.Set(ctx, logger)
			return next(ctx)
		}
	}
}

// LogRequest logs the current request
// Echo provides middleware similar to this, but we want to use our own logger
func LogRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			req := ctx.Request()
			res := ctx.Response()

			// Track how long the request takes to complete
			start := time.Now()
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}
			stop := time.Now()

			sub := log.Ctx(ctx).With(
				"ip", ctx.RealIP(),
				"host", req.Host,
				"referer", req.Referer(),
				"status", res.Status,
				"bytes_in", func() string {
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					return cl
				}(),
				"bytes_out", strconv.FormatInt(res.Size, 10),
				"latency", stop.Sub(start).String(),
			)

			msg := fmt.Sprintf("%s %s", req.Method, req.URL.RequestURI())

			if res.Status >= 500 {
				sub.Error(msg)
			} else {
				sub.Info(msg)
			}

			return nil
		}
	}
}
```go

`pkg/middleware/log_test.go`

```go
package middleware

import (
	"context"
	"log/slog"
	"testing"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type mockLogHandler struct {
	msg   string
	level string
	group string
	attr  []slog.Attr
}

func (m *mockLogHandler) Enabled(_ context.Context, l slog.Level) bool {
	return true
}

func (m *mockLogHandler) Handle(_ context.Context, r slog.Record) error {
	m.level = r.Level.String()
	m.msg = r.Message
	return nil
}

func (m *mockLogHandler) WithAttrs(as []slog.Attr) slog.Handler {
	if m.attr == nil {
		m.attr = make([]slog.Attr, 0)
	}
	m.attr = append(m.attr, as...)
	return m
}

func (m *mockLogHandler) WithGroup(name string) slog.Handler {
	m.group = name
	return m
}

func (m *mockLogHandler) GetAttr(key string) string {
	if m.attr == nil {
		return ""
	}
	for _, attr := range m.attr {
		if attr.Key == key {
			return attr.Value.String()
		}
	}

	return ""
}

func TestLogRequestID(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")

	h := new(mockLogHandler)
	logger := slog.New(h)
	log.Set(ctx, logger)

	require.NoError(t, tests.ExecuteMiddleware(ctx, echomw.RequestID()))
	require.NoError(t, tests.ExecuteMiddleware(ctx, SetLogger()))

	log.Ctx(ctx).Info("test")
	rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	assert.Equal(t, rID, h.GetAttr("request_id"))
}

func TestLogRequest(t *testing.T) {
	statusCode := 200
	h := new(mockLogHandler)

	exec := func() {
		ctx, _ := tests.NewContext(c.Web, "http://test.localhost/abc?d=1&e=2")
		logger := slog.New(h).With("previous", "param")
		log.Set(ctx, logger)
		ctx.Request().Header.Set("Referer", "ref.com")
		ctx.Request().Header.Set(echo.HeaderXRealIP, "21.12.12.21")

		require.NoError(t, tests.ExecuteHandler(ctx, func(ctx echo.Context) error {
			return ctx.String(statusCode, "hello")
		},
			SetLogger(),
			LogRequest(),
		))
	}

	exec()
	assert.Equal(t, "param", h.GetAttr("previous"))
	assert.Equal(t, "21.12.12.21", h.GetAttr("ip"))
	assert.Equal(t, "test.localhost", h.GetAttr("host"))
	assert.Equal(t, "ref.com", h.GetAttr("referer"))
	assert.Equal(t, "200", h.GetAttr("status"))
	assert.Equal(t, "0", h.GetAttr("bytes_in"))
	assert.Equal(t, "5", h.GetAttr("bytes_out"))
	assert.NotEmpty(t, h.GetAttr("latency"))
	assert.Equal(t, "INFO", h.level)
	assert.Equal(t, "GET /abc?d=1&e=2", h.msg)

	statusCode = 500
	exec()
	assert.Equal(t, "ERROR", h.level)
}
```go

`pkg/middleware/middleware_test.go`

```go
package middleware

import (
	"os"
	"testing"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/tests"
)

var (
	c   *services.Container
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Create a user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
```go

`pkg/middleware/session.go`

```go
package middleware

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/session"
)

// Session sets the session storage in the request context
func Session(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer context.Clear(ctx.Request())
			session.Store(ctx, store)
			return next(ctx)
		}
	}
}
```go

`pkg/middleware/session_test.go`

```go
package middleware

import (
	"testing"

	"github.com/gorilla/sessions"
	"github.com/mikestefanello/pagoda/pkg/session"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSession(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_, err := session.Get(ctx, "test")
	assert.Equal(t, session.ErrStoreNotFound, err)

	store := sessions.NewCookieStore([]byte("secret"))
	err = tests.ExecuteMiddleware(ctx, Session(store))
	require.NoError(t, err)

	_, err = session.Get(ctx, "test")
	assert.NotEqual(t, session.ErrStoreNotFound, err)
}
```go

`pkg/msg/msg.go`

```go
package msg

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/session"
)

// Type is a message type
type Type string

const (
	// TypeSuccess represents a success message type
	TypeSuccess Type = "success"

	// TypeInfo represents a info message type
	TypeInfo Type = "info"

	// TypeWarning represents a warning message type
	TypeWarning Type = "warning"

	// TypeDanger represents a danger message type
	TypeDanger Type = "danger"
)

const (
	// sessionName stores the name of the session which contains flash messages
	sessionName = "msg"
)

// Success sets a success flash message
func Success(ctx echo.Context, message string) {
	Set(ctx, TypeSuccess, message)
}

// Info sets an info flash message
func Info(ctx echo.Context, message string) {
	Set(ctx, TypeInfo, message)
}

// Warning sets a warning flash message
func Warning(ctx echo.Context, message string) {
	Set(ctx, TypeWarning, message)
}

// Danger sets a danger flash message
func Danger(ctx echo.Context, message string) {
	Set(ctx, TypeDanger, message)
}

// Set adds a new flash message of a given type into the session storage.
// Errors will be logged and not returned.
func Set(ctx echo.Context, typ Type, message string) {
	if sess, err := getSession(ctx); err == nil {
		sess.AddFlash(message, string(typ))
		save(ctx, sess)
	}
}

// Get gets flash messages of a given type from the session storage.
// Errors will be logged and not returned.
func Get(ctx echo.Context, typ Type) []string {
	var msgs []string

	if sess, err := getSession(ctx); err == nil {
		if flash := sess.Flashes(string(typ)); len(flash) > 0 {
			save(ctx, sess)

			for _, m := range flash {
				msgs = append(msgs, m.(string))
			}
		}
	}

	return msgs
}

// getSession gets the flash message session
func getSession(ctx echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(ctx, sessionName)
	if err != nil {
		log.Ctx(ctx).Error("cannot load flash message session",
			"error", err,
		)
	}
	return sess, err
}

// save saves the flash message session
func save(ctx echo.Context, sess *sessions.Session) {
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		log.Ctx(ctx).Error("failed to set flash message",
			"error", err,
		)
	}
}
```go

`pkg/msg/msg_test.go`

```go
package msg

import (
	"testing"

	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/labstack/echo/v4"
)

func TestMsg(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	tests.InitSession(ctx)

	assertMsg := func(typ Type, message string) {
		ret := Get(ctx, typ)
		require.Len(t, ret, 1)
		assert.Equal(t, message, ret[0])
		ret = Get(ctx, typ)
		require.Len(t, ret, 0)
	}

	text := "aaa"
	Success(ctx, text)
	assertMsg(TypeSuccess, text)

	text = "bbb"
	Info(ctx, text)
	assertMsg(TypeInfo, text)

	text = "ccc"
	Danger(ctx, text)
	assertMsg(TypeDanger, text)

	text = "ddd"
	Warning(ctx, text)
	assertMsg(TypeWarning, text)

	text = "eee"
	Set(ctx, TypeSuccess, text)
	assertMsg(TypeSuccess, text)
}
```go

`pkg/page/page.go`

```go
package page

import (
	"html/template"
	"net/http"
	"time"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/templates"

	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// Page consists of all data that will be used to render a page response for a given route.
// While it's not required for a handler to render a Page on a route, this is the common data
// object that will be passed to the templates, making it easy for all handlers to share
// functionality both on the back and frontend. The Page can be expanded to include anything else
// your app wants to support.
// Methods on this page also then become available in the templates, which can be more useful than
// the funcmap if your methods require data stored in the page, such as the context.
type Page struct {
	// AppName stores the name of the application.
	// If omitted, the configuration value will be used.
	AppName string

	// Title stores the title of the page
	Title string

	// Context stores the request context
	Context echo.Context

	// Path stores the path of the current request
	Path string

	// URL stores the URL of the current request
	URL string

	// Data stores whatever additional data that needs to be passed to the templates.
	// This is what the handler uses to pass the content of the page.
	Data any

	// Form stores a struct that represents a form on the page.
	// This should be a struct with fields for each form field, using both "form" and "validate" tags
	// It should also contain form.FormSubmission if you wish to have validation
	// messages and markup presented to the user
	Form any

	// Layout stores the name of the layout base template file which will be used when the page is rendered.
	// This should match a template file located within the layouts directory inside the templates directory.
	// The template extension should not be included in this value.
	Layout templates.Layout

	// Name stores the name of the page as well as the name of the template file which will be used to render
	// the content portion of the layout template.
	// This should match a template file located within the pages directory inside the templates directory.
	// The template extension should not be included in this value.
	Name templates.Page

	// IsHome stores whether the requested page is the home page or not
	IsHome bool

	// IsAuth stores whether the user is authenticated
	IsAuth bool

	// AuthUser stores the authenticated user
	AuthUser *ent.User

	// StatusCode stores the HTTP status code that will be returned
	StatusCode int

	// Metatags stores metatag values
	Metatags struct {
		// Description stores the description metatag value
		Description string

		// Keywords stores the keywords metatag values
		Keywords []string
	}

	// Pager stores a pager which can be used to page lists of results
	Pager Pager

	// CSRF stores the CSRF token for the given request.
	// This will only be populated if the CSRF middleware is in effect for the given request.
	// If this is populated, all forms must include this value otherwise the requests will be rejected.
	CSRF string

	// Headers stores a list of HTTP headers and values to be set on the response
	Headers map[string]string

	// RequestID stores the ID of the given request.
	// This will only be populated if the request ID middleware is in effect for the given request.
	RequestID string

	// HTMX provides the ability to interact with the HTMX library
	HTMX struct {
		// Request contains the information provided by HTMX about the current request
		Request htmx.Request

		// Response contains values to pass back to HTMX
		Response *htmx.Response
	}

	// Cache stores values for caching the response of this page
	Cache struct {
		// Enabled dictates if the response of this page should be cached.
		// Cached responses are served via middleware.
		Enabled bool

		// Expiration stores the amount of time that the cache entry should live for before expiring.
		// If omitted, the configuration value will be used.
		Expiration time.Duration

		// Tags stores a list of tags to apply to the cache entry.
		// These are useful when invalidating cache for dynamic events such as entity operations.
		Tags []string
	}
}

// New creates and initiatizes a new Page for a given request context
func New(ctx echo.Context) Page {
	p := Page{
		Context:    ctx,
		Path:       ctx.Request().URL.Path,
		URL:        ctx.Request().URL.String(),
		StatusCode: http.StatusOK,
		Pager:      NewPager(ctx, DefaultItemsPerPage),
		Headers:    make(map[string]string),
		RequestID:  ctx.Response().Header().Get(echo.HeaderXRequestID),
	}

	p.IsHome = p.Path == "/"

	if csrf := ctx.Get(echomw.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		p.IsAuth = true
		p.AuthUser = u.(*ent.User)
	}

	p.HTMX.Request = htmx.GetRequest(ctx)

	return p
}

// GetMessages gets all flash messages for a given type.
// This allows for easy access to flash messages from the templates.
func (p Page) GetMessages(typ msg.Type) []template.HTML {
	strs := msg.Get(p.Context, typ)
	ret := make([]template.HTML, len(strs))
	for k, v := range strs {
		ret[k] = template.HTML(v)
	}
	return ret
}
```go

`pkg/page/page_test.go`

```go
package page

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/tests"

	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	p := New(ctx)
	assert.Same(t, ctx, p.Context)
	assert.Equal(t, "/", p.Path)
	assert.Equal(t, "/", p.URL)
	assert.Equal(t, http.StatusOK, p.StatusCode)
	assert.Equal(t, NewPager(ctx, DefaultItemsPerPage), p.Pager)
	assert.Empty(t, p.Headers)
	assert.True(t, p.IsHome)
	assert.False(t, p.IsAuth)
	assert.Empty(t, p.CSRF)
	assert.Empty(t, p.RequestID)
	assert.False(t, p.Cache.Enabled)

	ctx, _ = tests.NewContext(e, "/abc?def=123")
	usr := &ent.User{
		ID: 1,
	}
	ctx.Set(context.AuthenticatedUserKey, usr)
	ctx.Set(echomw.DefaultCSRFConfig.ContextKey, "csrf")
	p = New(ctx)
	assert.Equal(t, "/abc", p.Path)
	assert.Equal(t, "/abc?def=123", p.URL)
	assert.False(t, p.IsHome)
	assert.True(t, p.IsAuth)
	assert.Equal(t, usr, p.AuthUser)
	assert.Equal(t, "csrf", p.CSRF)
}

func TestPage_GetMessages(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	tests.InitSession(ctx)
	p := New(ctx)

	// Set messages
	msgTests := make(map[msg.Type][]string)
	msgTests[msg.TypeWarning] = []string{
		"abc",
		"def",
	}
	msgTests[msg.TypeInfo] = []string{
		"123",
		"456",
	}
	for typ, values := range msgTests {
		for _, value := range values {
			msg.Set(ctx, typ, value)
		}
	}

	// Get the messages
	for typ, values := range msgTests {
		msgs := p.GetMessages(typ)

		for i, message := range msgs {
			assert.Equal(t, values[i], string(message))
		}
	}
}
```go

`pkg/page/pager.go`

```go
package page

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	// DefaultItemsPerPage stores the default amount of items per page
	DefaultItemsPerPage = 20

	// PageQueryKey stores the query key used to indicate the current page
	PageQueryKey = "page"
)

// Pager provides a mechanism to allow a user to page results via a query parameter
type Pager struct {
	// Items stores the total amount of items in the result set
	Items int

	// Page stores the current page number
	Page int

	// ItemsPerPage stores the amount of items to display per page
	ItemsPerPage int

	// Pages stores the total amount of pages in the result set
	Pages int
}

// NewPager creates a new Pager
func NewPager(ctx echo.Context, itemsPerPage int) Pager {
	p := Pager{
		ItemsPerPage: itemsPerPage,
		Page:         1,
	}

	if page := ctx.QueryParam(PageQueryKey); page != "" {
		if pageInt, err := strconv.Atoi(page); err == nil {
			if pageInt > 0 {
				p.Page = pageInt
			}
		}
	}

	return p
}

// SetItems sets the amount of items in total for the pager and calculate the amount
// of total pages based off on the item per page.
// This should be used rather than setting either items or pages directly.
func (p *Pager) SetItems(items int) {
	p.Items = items
	p.Pages = int(math.Ceil(float64(items) / float64(p.ItemsPerPage)))

	if p.Page > p.Pages {
		p.Page = p.Pages
	}
}

// IsBeginning determines if the pager is at the beginning of the pages
func (p Pager) IsBeginning() bool {
	return p.Page == 1
}

// IsEnd determines if the pager is at the end of the pages
func (p Pager) IsEnd() bool {
	return p.Page >= p.Pages
}

// GetOffset determines the offset of the results in order to get the items for
// the current page
func (p Pager) GetOffset() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.ItemsPerPage
}
```go

`pkg/page/pager_test.go`

```go
package page

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
)

func TestNewPager(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	pgr := NewPager(ctx, 10)
	assert.Equal(t, 10, pgr.ItemsPerPage)
	assert.Equal(t, 1, pgr.Page)
	assert.Equal(t, 0, pgr.Items)
	assert.Equal(t, 0, pgr.Pages)

	ctx, _ = tests.NewContext(e, fmt.Sprintf("/abc?%s=%d", PageQueryKey, 2))
	pgr = NewPager(ctx, 10)
	assert.Equal(t, 2, pgr.Page)

	ctx, _ = tests.NewContext(e, fmt.Sprintf("/abc?%s=%d", PageQueryKey, -2))
	pgr = NewPager(ctx, 10)
	assert.Equal(t, 1, pgr.Page)
}

func TestPager_SetItems(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.SetItems(100)
	assert.Equal(t, 100, pgr.Items)
	assert.Equal(t, 5, pgr.Pages)
}

func TestPager_IsBeginning(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.Pages = 10
	assert.True(t, pgr.IsBeginning())
	pgr.Page = 2
	assert.False(t, pgr.IsBeginning())
	pgr.Page = 1
	assert.True(t, pgr.IsBeginning())
}

func TestPager_IsEnd(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.Pages = 10
	assert.False(t, pgr.IsEnd())
	pgr.Page = 10
	assert.True(t, pgr.IsEnd())
	pgr.Page = 1
	assert.False(t, pgr.IsEnd())
}

func TestPager_GetOffset(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	assert.Equal(t, 0, pgr.GetOffset())
	pgr.Page = 2
	assert.Equal(t, 20, pgr.GetOffset())
	pgr.Page = 3
	assert.Equal(t, 40, pgr.GetOffset())
}
```go

`pkg/redirect/redirect.go`

```go
package redirect

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
)

// Redirect is a helper to perform HTTP redirects.
type Redirect struct {
	ctx         echo.Context
	url         string
	routeName   string
	routeParams []any
	status      int
	query       url.Values
}

// New initializes a new Redirect
func New(ctx echo.Context) *Redirect {
	return &Redirect{
		ctx:    ctx,
		status: http.StatusFound,
	}
}

// Route sets the route name to redirect to.
// Use either this or URL()
func (r *Redirect) Route(name string) *Redirect {
	r.routeName = name
	return r
}

// Params sets the route params
func (r *Redirect) Params(params ...any) *Redirect {
	r.routeParams = params
	return r
}

// StatusCode sets the HTTP status code which defaults to http.StatusFound.
// Does not apply to HTMX redirects.
func (r *Redirect) StatusCode(code int) *Redirect {
	r.status = code
	return r
}

// Query sets a URL query
func (r *Redirect) Query(query url.Values) *Redirect {
	r.query = query
	return r
}

// URL sets the URL to redirect to
// Use either this or Route()
func (r *Redirect) URL(url string) *Redirect {
	r.url = url
	return r
}

// Go performs the redirect
// If the request is HTMX boosted, an HTMX redirect will be performed instead of an HTTP redirect
func (r *Redirect) Go() error {
	if r.routeName == "" && r.url == "" {
		return errors.New("no redirect provided")
	}

	var dest string
	if r.url != "" {
		dest = r.url
	} else {
		dest = r.ctx.Echo().Reverse(r.routeName, r.routeParams...)
	}

	if len(r.query) > 0 {
		dest = fmt.Sprintf("%s?%s", dest, r.query.Encode())
	}

	if htmx.GetRequest(r.ctx).Boosted {
		htmx.Response{
			Redirect: dest,
		}.Apply(r.ctx)

		return nil
	} else {
		return r.ctx.Redirect(r.status, dest)
	}
}
```go

`pkg/redirect/redirect_test.go`

```go
package redirect

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedirect(t *testing.T) {
	e := echo.New()
	e.GET("/path/:first/and/:second", func(c echo.Context) error {
		return nil
	}).Name = "test"

	redirect := func() (*Redirect, echo.Context) {
		ctx, _ := tests.NewContext(e, "/")
		return New(ctx), ctx
	}

	t.Run("route", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		r.Route("test")
		r.Params("one", "two")
		r.Query(q)
		r.StatusCode(http.StatusTemporaryRedirect)
		require.NoError(t, r.Go())
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusTemporaryRedirect, ctx.Response().Status)
	})

	t.Run("route htmx", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		r.Route("test")
		r.Params("one", "two")
		r.Query(q)
		require.NoError(t, r.Go())
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})

	t.Run("url", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		r.URL("https://localhost.dev")
		r.Query(q)
		r.StatusCode(http.StatusTemporaryRedirect)
		require.NoError(t, r.Go())
		assert.Equal(t, "https://localhost.dev?a=1&b=2", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusTemporaryRedirect, ctx.Response().Status)
	})

	t.Run("url htmx", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		r.URL("https://localhost.dev")
		r.Query(q)
		require.NoError(t, r.Go())
		assert.Equal(t, "https://localhost.dev?a=1&b=2", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})
}
```go

`pkg/services/auth.go`

```go
package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/session"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	// authSessionName stores the name of the session which contains authentication data
	authSessionName = "ua"

	// authSessionKeyUserID stores the key used to store the user ID in the session
	authSessionKeyUserID = "user_id"

	// authSessionKeyAuthenticated stores the key used to store the authentication status in the session
	authSessionKeyAuthenticated = "authenticated"
)

// NotAuthenticatedError is an error returned when a user is not authenticated
type NotAuthenticatedError struct{}

// Error implements the error interface.
func (e NotAuthenticatedError) Error() string {
	return "user not authenticated"
}

// InvalidPasswordTokenError is an error returned when an invalid token is provided
type InvalidPasswordTokenError struct{}

// Error implements the error interface.
func (e InvalidPasswordTokenError) Error() string {
	return "invalid password token"
}

// AuthClient is the client that handles authentication requests
type AuthClient struct {
	config *config.Config
	orm    *ent.Client
}

// NewAuthClient creates a new authentication client
func NewAuthClient(cfg *config.Config, orm *ent.Client) *AuthClient {
	return &AuthClient{
		config: cfg,
		orm:    orm,
	}
}

// Login logs in a user of a given ID
func (c *AuthClient) Login(ctx echo.Context, userID int) error {
	sess, err := session.Get(ctx, authSessionName)
	if err != nil {
		return err
	}
	sess.Values[authSessionKeyUserID] = userID
	sess.Values[authSessionKeyAuthenticated] = true
	return sess.Save(ctx.Request(), ctx.Response())
}

// Logout logs the requesting user out
func (c *AuthClient) Logout(ctx echo.Context) error {
	sess, err := session.Get(ctx, authSessionName)
	if err != nil {
		return err
	}
	sess.Values[authSessionKeyAuthenticated] = false
	return sess.Save(ctx.Request(), ctx.Response())
}

// GetAuthenticatedUserID returns the authenticated user's ID, if the user is logged in
func (c *AuthClient) GetAuthenticatedUserID(ctx echo.Context) (int, error) {
	sess, err := session.Get(ctx, authSessionName)
	if err != nil {
		return 0, err
	}

	if sess.Values[authSessionKeyAuthenticated] == true {
		return sess.Values[authSessionKeyUserID].(int), nil
	}

	return 0, NotAuthenticatedError{}
}

// GetAuthenticatedUser returns the authenticated user if the user is logged in
func (c *AuthClient) GetAuthenticatedUser(ctx echo.Context) (*ent.User, error) {
	if userID, err := c.GetAuthenticatedUserID(ctx); err == nil {
		return c.orm.User.Query().
			Where(user.ID(userID)).
			Only(ctx.Request().Context())
	}

	return nil, NotAuthenticatedError{}
}

// HashPassword returns a hash of a given password
func (c *AuthClient) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword check if a given password matches a given hash
func (c *AuthClient) CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GeneratePasswordResetToken generates a password reset token for a given user.
// For security purposes, the token itself is not stored in the database but rather
// a hash of the token, exactly how passwords are handled. This method returns both
// the generated token as well as the token entity which only contains the hash.
func (c *AuthClient) GeneratePasswordResetToken(ctx echo.Context, userID int) (string, *ent.PasswordToken, error) {
	// Generate the token, which is what will go in the URL, but not the database
	token, err := c.RandomToken(c.config.App.PasswordToken.Length)
	if err != nil {
		return "", nil, err
	}

	// Hash the token, which is what will be stored in the database
	hash, err := c.HashPassword(token)
	if err != nil {
		return "", nil, err
	}

	// Create and save the password reset token
	pt, err := c.orm.PasswordToken.
		Create().
		SetHash(hash).
		SetUserID(userID).
		Save(ctx.Request().Context())

	return token, pt, err
}

// GetValidPasswordToken returns a valid, non-expired password token entity for a given user, token ID and token.
// Since the actual token is not stored in the database for security purposes, if a matching password token entity is
// found a hash of the provided token is compared with the hash stored in the database in order to validate.
func (c *AuthClient) GetValidPasswordToken(ctx echo.Context, userID, tokenID int, token string) (*ent.PasswordToken, error) {
	// Ensure expired tokens are never returned
	expiration := time.Now().Add(-c.config.App.PasswordToken.Expiration)

	// Query to find a password token entity that matches the given user and token ID
	pt, err := c.orm.PasswordToken.
		Query().
		Where(passwordtoken.ID(tokenID)).
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Where(passwordtoken.CreatedAtGTE(expiration)).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
	case nil:
		// Check the token for a hash match
		if err := c.CheckPassword(token, pt.Hash); err == nil {
			return pt, nil
		}
	default:
		if !context.IsCanceledError(err) {
			return nil, err
		}
	}

	return nil, InvalidPasswordTokenError{}
}

// DeletePasswordTokens deletes all password tokens in the database for a belonging to a given user.
// This should be called after a successful password reset.
func (c *AuthClient) DeletePasswordTokens(ctx echo.Context, userID int) error {
	_, err := c.orm.PasswordToken.
		Delete().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Exec(ctx.Request().Context())

	return err
}

// RandomToken generates a random token string of a given length
func (c *AuthClient) RandomToken(length int) (string, error) {
	b := make([]byte, (length/2)+1)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	return token[:length], nil
}

// GenerateEmailVerificationToken generates an email verification token for a given email address using JWT which
// is set to expire based on the duration stored in configuration
func (c *AuthClient) GenerateEmailVerificationToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(c.config.App.EmailVerificationTokenExpiration).Unix(),
	})

	return token.SignedString([]byte(c.config.App.EncryptionKey))
}

// ValidateEmailVerificationToken validates an email verification token and returns the associated email address if
// the token is valid and has not expired
func (c *AuthClient) ValidateEmailVerificationToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(c.config.App.EncryptionKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims["email"].(string), nil
	}

	return "", errors.New("invalid or expired token")
}
```go

`pkg/services/auth_test.go`

```go
package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestAuthClient_Auth(t *testing.T) {
	assertNoAuth := func() {
		_, err := c.Auth.GetAuthenticatedUserID(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
		_, err = c.Auth.GetAuthenticatedUser(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
	}

	assertNoAuth()

	err := c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)

	uid, err := c.Auth.GetAuthenticatedUserID(ctx)
	require.NoError(t, err)
	assert.Equal(t, usr.ID, uid)

	u, err := c.Auth.GetAuthenticatedUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, u.ID, usr.ID)

	err = c.Auth.Logout(ctx)
	require.NoError(t, err)

	assertNoAuth()
}

func TestAuthClient_PasswordHashing(t *testing.T) {
	pw := "testcheckpassword"
	hash, err := c.Auth.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEqual(t, hash, pw)
	err = c.Auth.CheckPassword(pw, hash)
	assert.NoError(t, err)
}

func TestAuthClient_GeneratePasswordResetToken(t *testing.T) {
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	assert.Len(t, token, c.Config.App.PasswordToken.Length)
	assert.NoError(t, c.Auth.CheckPassword(token, pt.Hash))
}

func TestAuthClient_GetValidPasswordToken(t *testing.T) {
	// Check that a fake token is not valid
	_, err := c.Auth.GetValidPasswordToken(ctx, usr.ID, 1, "faketoken")
	assert.Error(t, err)

	// Generate a valid token and check that it is returned
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	pt2, err := c.Auth.GetValidPasswordToken(ctx, usr.ID, pt.ID, token)
	require.NoError(t, err)
	assert.Equal(t, pt.ID, pt2.ID)

	// Expire the token by pushing the date far enough back
	count, err := c.ORM.PasswordToken.
		Update().
		SetCreatedAt(time.Now().Add(-(c.Config.App.PasswordToken.Expiration + time.Hour))).
		Where(passwordtoken.ID(pt.ID)).
		Save(context.Background())
	require.NoError(t, err)
	require.Equal(t, 1, count)

	// Expired tokens should not be valid
	_, err = c.Auth.GetValidPasswordToken(ctx, usr.ID, pt.ID, token)
	assert.Error(t, err)
}

func TestAuthClient_DeletePasswordTokens(t *testing.T) {
	// Create three tokens for the user
	for i := 0; i < 3; i++ {
		_, _, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
		require.NoError(t, err)
	}

	// Delete all tokens for the user
	err := c.Auth.DeletePasswordTokens(ctx, usr.ID)
	require.NoError(t, err)

	// Check that no tokens remain
	count, err := c.ORM.PasswordToken.
		Query().
		Where(passwordtoken.HasUserWith(user.ID(usr.ID))).
		Count(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestAuthClient_RandomToken(t *testing.T) {
	length := c.Config.App.PasswordToken.Length
	a, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	b, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	assert.Len(t, a, length)
	assert.Len(t, b, length)
	assert.NotEqual(t, a, b)
}

func TestAuthClient_EmailVerificationToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		tokenEmail, err := c.Auth.ValidateEmailVerificationToken(token)
		require.NoError(t, err)
		assert.Equal(t, email, tokenEmail)
	})

	t.Run("invalid token", func(t *testing.T) {
		badToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWxob3N0LmNvbSIsImV4cCI6MTkxNzg2NDAwMH0.ScJCpfEEzlilKfRs_aVouzwPNKI28M3AIm-hyImQHUQ"
		_, err := c.Auth.ValidateEmailVerificationToken(badToken)
		assert.Error(t, err)
	})

	t.Run("expired token", func(t *testing.T) {
		c.Config.App.EmailVerificationTokenExpiration = -time.Hour
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		_, err = c.Auth.ValidateEmailVerificationToken(token)
		assert.Error(t, err)

		c.Config.App.EmailVerificationTokenExpiration = time.Hour * 12
	})
}
```go

`pkg/services/cache.go`

```go
package services

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/maypok86/otter"
)

// ErrCacheMiss indicates that the requested key does not exist in the cache
var ErrCacheMiss = errors.New("cache miss")

type (
	// CacheStore provides an interface for cache storage
	CacheStore interface {
		// get attempts to get a cached value
		get(context.Context, *CacheGetOp) (any, error)

		// set attempts to set an entry in the cache
		set(context.Context, *CacheSetOp) error

		// flush removes a given key and/or tags from the cache
		flush(context.Context, *CacheFlushOp) error

		// close shuts down the cache storage
		close()
	}

	// CacheClient is the client that allows you to interact with the cache
	CacheClient struct {
		// store holds the Cache storage
		store CacheStore
	}

	// CacheSetOp handles chaining a set operation
	CacheSetOp struct {
		client     *CacheClient
		key        string
		group      string
		data       any
		expiration time.Duration
		tags       []string
	}

	// CacheGetOp handles chaining a get operation
	CacheGetOp struct {
		client *CacheClient
		key    string
		group  string
	}

	// CacheFlushOp handles chaining a flush operation
	CacheFlushOp struct {
		client *CacheClient
		key    string
		group  string
		tags   []string
	}

	// inMemoryCacheStore is a cache store implementation in memory
	inMemoryCacheStore struct {
		store    *otter.CacheWithVariableTTL[string, any]
		tagIndex *tagIndex
	}

	// tagIndex maintains an index to support cache tags for in-memory cache stores.
	// There is a performance and memory impact to using cache tags since set and get operations using tags will require
	// locking, and we need to keep track of this index in order to keep everything in sync.
	// If using something like Redis for caching, you can leverage sets to store the index.
	// Cache tags can be useful and convenient, so you should decide if your app benefits enough from this.
	// As it stands here, there is no limiting how much memory this will consume and it will track all keys
	// and tags added and removed from the cache. You could store these in the cache itself but allowing these to
	// be evicted poses challenges.
	tagIndex struct {
		sync.Mutex
		tags map[string]map[string]struct{} // tag->keys
		keys map[string]map[string]struct{} // key->tags
	}
)

// NewCacheClient creates a new cache client
func NewCacheClient(store CacheStore) *CacheClient {
	return &CacheClient{store: store}
}

// Close closes the connection to the cache
func (c *CacheClient) Close() {
	c.store.close()
}

// Set creates a cache set operation
func (c *CacheClient) Set() *CacheSetOp {
	return &CacheSetOp{
		client: c,
	}
}

// Get creates a cache get operation
func (c *CacheClient) Get() *CacheGetOp {
	return &CacheGetOp{
		client: c,
	}
}

// Flush creates a cache flush operation
func (c *CacheClient) Flush() *CacheFlushOp {
	return &CacheFlushOp{
		client: c,
	}
}

// cacheKey formats a cache key with an optional group
func (c *CacheClient) cacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s::%s", group, key)
	}
	return key
}

// Key sets the cache key
func (c *CacheSetOp) Key(key string) *CacheSetOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheSetOp) Group(group string) *CacheSetOp {
	c.group = group
	return c
}

// Data sets the data to cache
func (c *CacheSetOp) Data(data any) *CacheSetOp {
	c.data = data
	return c
}

// Expiration sets the expiration duration of the cached data
func (c *CacheSetOp) Expiration(expiration time.Duration) *CacheSetOp {
	c.expiration = expiration
	return c
}

// Tags sets the cache tags
func (c *CacheSetOp) Tags(tags ...string) *CacheSetOp {
	c.tags = tags
	return c
}

// Save saves the data in the cache
func (c *CacheSetOp) Save(ctx context.Context) error {
	switch {
	case c.key == "":
		return errors.New("no cache key specified")
	case c.data == nil:
		return errors.New("no cache data specified")
	case c.expiration == 0:
		return errors.New("no cache expiration specified")
	}

	return c.client.store.set(ctx, c)
}

// Key sets the cache key
func (c *CacheGetOp) Key(key string) *CacheGetOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheGetOp) Group(group string) *CacheGetOp {
	c.group = group
	return c
}

// Fetch fetches the data from the cache
func (c *CacheGetOp) Fetch(ctx context.Context) (any, error) {
	if c.key == "" {
		return nil, errors.New("no cache key specified")
	}

	return c.client.store.get(ctx, c)
}

// Key sets the cache key
func (c *CacheFlushOp) Key(key string) *CacheFlushOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheFlushOp) Group(group string) *CacheFlushOp {
	c.group = group
	return c
}

// Tags sets the cache tags
func (c *CacheFlushOp) Tags(tags ...string) *CacheFlushOp {
	c.tags = tags
	return c
}

// Execute flushes the data from the cache
func (c *CacheFlushOp) Execute(ctx context.Context) error {
	return c.client.store.flush(ctx, c)
}

// newInMemoryCache creates a new in-memory CacheStore
func newInMemoryCache(capacity int) (CacheStore, error) {
	s := &inMemoryCacheStore{
		tagIndex: newTagIndex(),
	}

	store, err := otter.MustBuilder[string, any](capacity).
		WithVariableTTL().
		DeletionListener(func(key string, value any, cause otter.DeletionCause) {
			s.tagIndex.purgeKeys(key)
		}).
		Build()

	if err != nil {
		return nil, err
	}

	s.store = &store

	return s, nil
}

func (s *inMemoryCacheStore) get(_ context.Context, op *CacheGetOp) (any, error) {
	v, exists := s.store.Get(op.client.cacheKey(op.group, op.key))

	if !exists {
		return nil, ErrCacheMiss
	}

	return v, nil
}

func (s *inMemoryCacheStore) set(_ context.Context, op *CacheSetOp) error {
	key := op.client.cacheKey(op.group, op.key)

	added := s.store.Set(
		key,
		op.data,
		op.expiration,
	)

	if len(op.tags) > 0 {
		s.tagIndex.setTags(key, op.tags...)
	}

	if !added {
		return errors.New("cache set failed")
	}

	return nil
}

func (s *inMemoryCacheStore) flush(_ context.Context, op *CacheFlushOp) error {
	keys := make([]string, 0)

	if key := op.client.cacheKey(op.group, op.key); key != "" {
		keys = append(keys, key)
	}

	if len(op.tags) > 0 {
		keys = append(keys, s.tagIndex.purgeTags(op.tags...)...)
	}

	for _, key := range keys {
		s.store.Delete(key)
	}

	s.tagIndex.purgeKeys(keys...)

	return nil
}

func (s *inMemoryCacheStore) close() {
	s.store.Close()
}

func newTagIndex() *tagIndex {
	return &tagIndex{
		tags: make(map[string]map[string]struct{}),
		keys: make(map[string]map[string]struct{}),
	}
}

func (i *tagIndex) setTags(key string, tags ...string) {
	i.Lock()
	defer i.Unlock()

	if _, exists := i.keys[key]; !exists {
		i.keys[key] = make(map[string]struct{})
	}

	for _, tag := range tags {
		if _, exists := i.tags[tag]; !exists {
			i.tags[tag] = make(map[string]struct{})
		}
		i.tags[tag][key] = struct{}{}
		i.keys[key][tag] = struct{}{}
	}
}

func (i *tagIndex) purgeTags(tags ...string) []string {
	i.Lock()
	defer i.Unlock()

	keys := make([]string, 0)

	for _, tag := range tags {
		if tagKeys, exists := i.tags[tag]; exists {
			delete(i.tags, tag)

			for key := range tagKeys {
				delete(i.keys[key], tag)
				if len(i.keys[key]) == 0 {
					delete(i.keys, key)
				}

				keys = append(keys, key)
			}
		}
	}

	return keys
}

func (i *tagIndex) purgeKeys(keys ...string) {
	i.Lock()
	defer i.Unlock()

	for _, key := range keys {
		if keyTags, exists := i.keys[key]; exists {
			delete(i.keys, key)

			for tag := range keyTags {
				delete(i.tags[tag], key)
				if len(i.tags[tag]) == 0 {
					delete(i.tags, tag)
				}
			}
		}
	}
}
```go

`pkg/services/cache_test.go`

```go
package services

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheClient(t *testing.T) {
	type cacheTest struct {
		Value string
	}

	// Cache some data
	data := cacheTest{Value: "abcdef"}
	group := "testgroup"
	key := "testkey"
	err := c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Expiration(500 * time.Millisecond).
		Save(context.Background())
	require.NoError(t, err)

	// Get the data
	fromCache, err := c.Cache.
		Get().
		Group(group).
		Key(key).
		Fetch(context.Background())
	require.NoError(t, err)
	cast, ok := fromCache.(cacheTest)
	require.True(t, ok)
	assert.Equal(t, data, cast)

	// The same key with the wrong group should fail
	_, err = c.Cache.
		Get().
		Key(key).
		Fetch(context.Background())
	assert.Equal(t, ErrCacheMiss, err)

	// Flush the data
	err = c.Cache.
		Flush().
		Group(group).
		Key(key).
		Execute(context.Background())
	require.NoError(t, err)

	// The data should be gone
	assertFlushed := func(key string) {
		// The data should be gone
		_, err = c.Cache.
			Get().
			Group(group).
			Key(key).
			Fetch(context.Background())
		assert.Equal(t, ErrCacheMiss, err)
	}
	assertFlushed(key)

	// Set with tags
	key = "testkey2"
	err = c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Tags("tag1", "tag2").
		Expiration(time.Hour).
		Save(context.Background())
	require.NoError(t, err)

	// Check the tag index
	index := c.Cache.store.(*inMemoryCacheStore).tagIndex
	gk := c.Cache.cacheKey(group, key)
	_, exists := index.tags["tag1"][gk]
	assert.True(t, exists)
	_, exists = index.tags["tag2"][gk]
	assert.True(t, exists)
	_, exists = index.keys[gk]["tag1"]
	assert.True(t, exists)
	_, exists = index.keys[gk]["tag2"]
	assert.True(t, exists)

	// Flush one of tags
	err = c.Cache.
		Flush().
		Tags("tag1").
		Execute(context.Background())
	require.NoError(t, err)

	// The data should be gone
	assertFlushed(key)

	// The index should be empty
	assert.Empty(t, index.tags)
	assert.Empty(t, index.keys)
}
```go

`pkg/services/container.go`

```go
package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mikestefanello/backlite"
	"log/slog"
	"os"
	"strings"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/funcmap"
	"github.com/mikestefanello/pagoda/pkg/log"

	// Require by ent
	_ "github.com/mikestefanello/pagoda/ent/runtime"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests
type Container struct {
	// Validator stores a validator
	Validator *Validator

	// Web stores the web framework
	Web *echo.Echo

	// Config stores the application configuration
	Config *config.Config

	// Cache contains the cache client
	Cache *CacheClient

	// Database stores the connection to the database
	Database *sql.DB

	// ORM stores a client to the ORM
	ORM *ent.Client

	// Mail stores an email sending client
	Mail *MailClient

	// Auth stores an authentication client
	Auth *AuthClient

	// TemplateRenderer stores a service to easily render and cache templates
	TemplateRenderer *TemplateRenderer

	// Tasks stores the task client
	Tasks *backlite.Client

	// OAuth stores OAuth client
	OAuth *OAuthClient
}

// NewContainer creates and initializes a new Container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initORM()
	c.initAuth()
	c.initTemplateRenderer()
	c.initMail()
	c.initTasks()
	c.initOAuth()
	return c
}

// Shutdown shuts the Container down and disconnects all connections.
// If the task runner was started, cancel the context to shut it down prior to calling this.
func (c *Container) Shutdown() error {
	if err := c.ORM.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}
	c.Cache.Close()

	return nil
}

// initConfig initializes configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg

	// Configure logging
	switch cfg.App.Environment {
	case config.EnvProduction:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web framework
func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true
	c.Web.Validator = c.Validator
}

// initCache initializes the cache
func (c *Container) initCache() {
	store, err := newInMemoryCache(c.Config.Cache.Capacity)
	if err != nil {
		panic(err)
	}

	c.Cache = NewCacheClient(store)
}

// initDatabase initializes the database
func (c *Container) initDatabase() {
	var err error
	var connection string

	switch c.Config.App.Environment {
	case config.EnvTest:
		// TODO: Drop/recreate the DB, if this isn't in memory?
		connection = c.Config.Database.TestConnection
	default:
		connection = c.Config.Database.Connection
	}

	c.Database, err = openDB(c.Config.Database.Driver, connection)
	if err != nil {
		panic(err)
	}
}

// initORM initializes the ORM
func (c *Container) initORM() {
	drv := entsql.OpenDB(c.Config.Database.Driver, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}

// initAuth initializes the authentication client
func (c *Container) initAuth() {
	c.Auth = NewAuthClient(c.Config, c.ORM)
}

// initTemplateRenderer initializes the template renderer
func (c *Container) initTemplateRenderer() {
	c.TemplateRenderer = NewTemplateRenderer(c.Config, c.Cache, funcmap.NewFuncMap(c.Web))
}

// initMail initialize the mail client
func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config, c.TemplateRenderer)
	if err != nil {
		panic(fmt.Sprintf("failed to create mail client: %v", err))
	}
}

// initTasks initializes the task client
func (c *Container) initTasks() {
	var err error
	// You could use a separate database for tasks, if you'd like. but using one
	// makes transaction support easier
	c.Tasks, err = backlite.NewClient(backlite.ClientConfig{
		DB:              c.Database,
		Logger:          log.Default(),
		NumWorkers:      c.Config.Tasks.Goroutines,
		ReleaseAfter:    c.Config.Tasks.ReleaseAfter,
		CleanupInterval: c.Config.Tasks.CleanupInterval,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to create task client: %v", err))
	}

	if err = c.Tasks.Install(); err != nil {
		panic(fmt.Sprintf("failed to install task schema: %v", err))
	}
}

// initOAuth initializes OAuth client
func (c *Container) initOAuth() {
	c.OAuth = NewOAuthClient(c.Config, c.ORM)
}

// openDB opens a database connection
func openDB(driver, connection string) (*sql.DB, error) {
	// Helper to automatically create the directories that the specified sqlite file
	// should reside in, if one
	if driver == "sqlite3" {
		d := strings.Split(connection, "/")

		if len(d) > 1 {
			path := strings.Join(d[:len(d)-1], "/")

			if err := os.MkdirAll(path, 0755); err != nil {
				return nil, err
			}
		}
	}

	return sql.Open(driver, connection)
}
```go

`pkg/services/container_test.go`

```go
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.NotNil(t, c.Web)
	assert.NotNil(t, c.Config)
	assert.NotNil(t, c.Validator)
	assert.NotNil(t, c.Cache)
	assert.NotNil(t, c.Database)
	assert.NotNil(t, c.ORM)
	assert.NotNil(t, c.Mail)
	assert.NotNil(t, c.Auth)
	assert.NotNil(t, c.TemplateRenderer)
	assert.NotNil(t, c.Tasks)
}
```go

`pkg/services/mail.go`

```go
package services

import (
	"errors"
	"fmt"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/log"

	"github.com/labstack/echo/v4"
)

type (
	// MailClient provides a client for sending email
	// This is purposely not completed because there are many different methods and services
	// for sending email, many of which are very different. Choose what works best for you
	// and populate the methods below
	MailClient struct {
		// config stores application configuration
		config *config.Config

		// templates stores the template renderer
		templates *TemplateRenderer
	}

	// mail represents an email to be sent
	mail struct {
		client       *MailClient
		from         string
		to           string
		subject      string
		body         string
		template     string
		templateData any
	}
)

// NewMailClient creates a new MailClient
func NewMailClient(cfg *config.Config, templates *TemplateRenderer) (*MailClient, error) {
	return &MailClient{
		config:    cfg,
		templates: templates,
	}, nil
}

// Compose creates a new email
func (m *MailClient) Compose() *mail {
	return &mail{
		client: m,
		from:   m.config.Mail.FromAddress,
	}
}

// skipSend determines if mail sending should be skipped
func (m *MailClient) skipSend() bool {
	return m.config.App.Environment != config.EnvProduction
}

// send attempts to send the email
func (m *MailClient) send(email *mail, ctx echo.Context) error {
	switch {
	case email.to == "":
		return errors.New("email cannot be sent without a to address")
	case email.body == "" && email.template == "":
		return errors.New("email cannot be sent without a body or template")
	}

	// Check if a template was supplied
	if email.template != "" {
		// Parse and execute template
		buf, err := m.templates.
			Parse().
			Group("mail").
			Key(email.template).
			Base(email.template).
			Files(fmt.Sprintf("emails/%s", email.template)).
			Execute(email.templateData)

		if err != nil {
			return err
		}

		email.body = buf.String()
	}

	// Check if mail sending should be skipped
	if m.skipSend() {
		log.Ctx(ctx).Debug("skipping email delivery",
			"to", email.to,
		)
		return nil
	}

	// TODO: Finish based on your mail sender of choice!
	return nil
}

// From sets the email from address
func (m *mail) From(from string) *mail {
	m.from = from
	return m
}

// To sets the email address this email will be sent to
func (m *mail) To(to string) *mail {
	m.to = to
	return m
}

// Subject sets the subject line of the email
func (m *mail) Subject(subject string) *mail {
	m.subject = subject
	return m
}

// Body sets the body of the email
// This is not required and will be ignored if a template via Template()
func (m *mail) Body(body string) *mail {
	m.body = body
	return m
}

// Template sets the template to be used to produce the body of the email
// The template name should only include the filename without the extension or directory.
// The template must reside within the emails sub-directory.
// The funcmap will be automatically added to the template.
// Use TemplateData() to supply the data that will be passed in to the template.
func (m *mail) Template(template string) *mail {
	m.template = template
	return m
}

// TemplateData sets the data that will be passed to the template specified when calling Template()
func (m *mail) TemplateData(data any) *mail {
	m.templateData = data
	return m
}

// Send attempts to send the email
func (m *mail) Send(ctx echo.Context) error {
	return m.client.send(m, ctx)
}
```go

`pkg/services/mail_test.go`

```go
package services

// Fill this in once you implement your mail client
```go

`pkg/services/oauth.go`

```go
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type OAuthClient struct {
	config  *config.Config
	orm     *ent.Client
	configs map[string]*oauth2.Config
}

type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

type FacebookUserInfo struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ProfilePicURL string `json:"picture.data.url"`
}

func NewOAuthClient(cfg *config.Config, orm *ent.Client) *OAuthClient {
	client := &OAuthClient{
		config:  cfg,
		orm:     orm,
		configs: make(map[string]*oauth2.Config),
	}

	client.configs["google"] = &oauth2.Config{
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		RedirectURL:  cfg.OAuth.Google.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	client.configs["facebook"] = &oauth2.Config{
		ClientID:     cfg.OAuth.Facebook.ClientID,
		ClientSecret: cfg.OAuth.Facebook.ClientSecret,
		RedirectURL:  cfg.OAuth.Facebook.RedirectURL,
		Scopes:       []string{"email", "public_profile"},
		Endpoint:     facebook.Endpoint,
	}

	return client
}

func (c *OAuthClient) GetAuthCodeURL(provider string, ctx echo.Context) string {
	return c.configs[provider].AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (c *OAuthClient) HandleCallback(provider string, ctx echo.Context) (*ent.User, error) {
	code := ctx.QueryParam("code")
	token, err := c.configs[provider].Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := c.configs[provider].Client(context.Background(), token)
	userInfo, err := c.getUserInfo(provider, client)
	if err != nil {
		return nil, err
	}

	user, err := c.findOrCreateUser(provider, userInfo)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *OAuthClient) getUserInfo(provider string, client *http.Client) (map[string]interface{}, error) {
	var userInfoURL string
	switch provider {
	case "google":
		userInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"
	case "facebook":
		userInfoURL = "https://graph.facebook.com/me?fields=id,name,email"
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	resp, err := client.Get(userInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (c *OAuthClient) findOrCreateUser(provider string, userInfo map[string]interface{}) (*ent.User, error) {
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)

	user, err := c.orm.User.Query().Where(user.Email(email)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			user, err = c.orm.User.Create().
				SetEmail(email).
				SetName(name).
				SetVerified(true).
				SetPassword("__oauth2__"). // Set a placeholder password
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return user, nil
}
```go

`pkg/services/services_test.go`

```go
package services

import (
	"os"
	"testing"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/labstack/echo/v4"
)

var (
	c   *Container
	ctx echo.Context
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = NewContainer()

	// Create a web context
	ctx, _ = tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Create a test user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
```go

`pkg/services/template_renderer.go`

```go
package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/templates"
)

// cachedPageGroup stores the cache group for cached pages
const cachedPageGroup = "page"

type (
	// TemplateRenderer provides a flexible and easy to use method of rendering simple templates or complex sets of
	// templates while also providing caching and/or hot-reloading depending on your current environment
	TemplateRenderer struct {
		// templateCache stores a cache of parsed page templates
		templateCache sync.Map

		// funcMap stores the template function map
		funcMap template.FuncMap

		// config stores application configuration
		config *config.Config

		// cache stores the cache client
		cache *CacheClient
	}

	// TemplateParsed is a wrapper around parsed templates which are stored in the TemplateRenderer cache
	TemplateParsed struct {
		// Template is the parsed template
		Template *template.Template

		// build stores the build data used to parse the template
		build *templateBuild
	}

	// templateBuild stores the build data used to parse a template
	templateBuild struct {
		group       string
		key         string
		base        string
		files       []string
		directories []string
	}

	// templateBuilder handles chaining a template parse operation
	templateBuilder struct {
		build    *templateBuild
		renderer *TemplateRenderer
	}

	// CachedPage is what is used to store a rendered Page in the cache
	CachedPage struct {
		// URL stores the URL of the requested page
		URL string

		// HTML stores the complete HTML of the rendered Page
		HTML []byte

		// StatusCode stores the HTTP status code
		StatusCode int

		// Headers stores the HTTP headers
		Headers map[string]string
	}
)

// NewTemplateRenderer creates a new TemplateRenderer
func NewTemplateRenderer(cfg *config.Config, cache *CacheClient, fm template.FuncMap) *TemplateRenderer {
	return &TemplateRenderer{
		templateCache: sync.Map{},
		funcMap:       fm,
		config:        cfg,
		cache:         cache,
	}
}

// Parse creates a template build operation
func (t *TemplateRenderer) Parse() *templateBuilder {
	return &templateBuilder{
		renderer: t,
		build:    &templateBuild{},
	}
}

// RenderPage renders a Page as an HTTP response
func (t *TemplateRenderer) RenderPage(ctx echo.Context, page page.Page) error {
	var buf *bytes.Buffer
	var err error
	templateGroup := "page"

	// Page name is required
	if page.Name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "page render failed due to missing name")
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = t.config.App.Name
	}

	// Check if this is an HTMX non-boosted request which indicates that only partial
	// content should be rendered
	if page.HTMX.Request.Enabled && !page.HTMX.Request.Boosted {
		// Switch the layout which will only render the page content
		page.Layout = templates.LayoutHTMX

		// Alter the template group so this is cached separately
		templateGroup = "page:htmx"
	}

	// Parse and execute the templates for the Page
	// As mentioned in the documentation for the Page struct, the templates used for the page will be:
	// 1. The layout/base template specified in Page.Layout
	// 2. The content template specified in Page.Name
	// 3. All templates within the components directory
	// Also included is the function map provided by the funcmap package
	buf, err = t.
		Parse().
		Group(templateGroup).
		Key(string(page.Name)).
		Base(string(page.Layout)).
		Files(
			fmt.Sprintf("layouts/%s", page.Layout),
			fmt.Sprintf("pages/%s", page.Name),
		).
		Directories("components").
		Execute(page)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to parse and execute templates: %s", err),
		)
	}

	// Set the status code
	ctx.Response().Status = page.StatusCode

	// Set any headers
	for k, v := range page.Headers {
		ctx.Response().Header().Set(k, v)
	}

	// Apply the HTMX response, if one
	if page.HTMX.Response != nil {
		page.HTMX.Response.Apply(ctx)
	}

	// Cache this page, if caching was enabled
	t.cachePage(ctx, page, buf)

	return ctx.HTMLBlob(ctx.Response().Status, buf.Bytes())
}

// cachePage caches the HTML for a given Page if the Page has caching enabled
func (t *TemplateRenderer) cachePage(ctx echo.Context, page page.Page, html *bytes.Buffer) {
	if !page.Cache.Enabled || page.IsAuth {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = t.config.Cache.Expiration.Page
	}

	// Extract the headers
	headers := make(map[string]string)
	for k, v := range ctx.Response().Header() {
		headers[k] = v[0]
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests
	key := ctx.Request().URL.String()
	cp := &CachedPage{
		URL:        key,
		HTML:       html.Bytes(),
		Headers:    headers,
		StatusCode: ctx.Response().Status,
	}

	err := t.cache.
		Set().
		Group(cachedPageGroup).
		Key(key).
		Tags(page.Cache.Tags...).
		Expiration(page.Cache.Expiration).
		Data(cp).
		Save(ctx.Request().Context())

	switch {
	case err == nil:
		log.Ctx(ctx).Debug("cached page")
	case !context.IsCanceledError(err):
		log.Ctx(ctx).Error("failed to cache page",
			"error", err,
		)
	}
}

// GetCachedPage attempts to fetch a cached page for a given URL
func (t *TemplateRenderer) GetCachedPage(ctx echo.Context, url string) (*CachedPage, error) {
	p, err := t.cache.
		Get().
		Group(cachedPageGroup).
		Key(url).
		Fetch(ctx.Request().Context())

	if err != nil {
		return nil, err
	}

	return p.(*CachedPage), nil
}

// getCacheKey gets a cache key for a given group and ID
func (t *TemplateRenderer) getCacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s:%s", group, key)
	}
	return key
}

// parse parses a set of templates and caches them for quick execution
// If the application environment is set to local, the cache will be bypassed and templates will be
// parsed upon each request so hot-reloading is possible without restarts.
// Also included will be the function map provided by the funcmap package.
func (t *TemplateRenderer) parse(build *templateBuild) (*TemplateParsed, error) {
	var tp *TemplateParsed
	var err error

	switch {
	case build.key == "":
		return nil, errors.New("cannot parse template without key")
	case len(build.files) == 0 && len(build.directories) == 0:
		return nil, errors.New("cannot parse template without files or directories")
	case build.base == "":
		return nil, errors.New("cannot parse template without base")
	}

	// Generate the cache key
	cacheKey := t.getCacheKey(build.group, build.key)

	// Check if the template has not yet been parsed or if the app environment is local, so that
	// templates reflect changes without having the restart the server
	if tp, err = t.Load(build.group, build.key); err != nil || t.config.App.Environment == config.EnvLocal {
		// Initialize the parsed template with the function map
		parsed := template.New(build.base + config.TemplateExt).
			Funcs(t.funcMap)

		// Format the requested files
		for k, v := range build.files {
			build.files[k] = fmt.Sprintf("%s%s", v, config.TemplateExt)
		}

		// Include all files within the requested directories
		for k, v := range build.directories {
			build.directories[k] = fmt.Sprintf("%s/*%s", v, config.TemplateExt)
		}

		// Get the templates
		var tpl fs.FS
		if t.config.App.Environment == config.EnvLocal {
			tpl = templates.GetOS()
		} else {
			tpl = templates.Get()
		}

		// Parse the templates
		parsed, err = parsed.ParseFS(tpl, append(build.files, build.directories...)...)
		if err != nil {
			return nil, err
		}

		// Store the template so this process only happens once
		tp = &TemplateParsed{
			Template: parsed,
			build:    build,
		}
		t.templateCache.Store(cacheKey, tp)
	}

	return tp, nil
}

// Load loads a template from the cache
func (t *TemplateRenderer) Load(group, key string) (*TemplateParsed, error) {
	load, ok := t.templateCache.Load(t.getCacheKey(group, key))
	if !ok {
		return nil, errors.New("uncached page template requested")
	}

	tmpl, ok := load.(*TemplateParsed)
	if !ok {
		return nil, errors.New("unable to cast cached template")
	}

	return tmpl, nil
}

// Execute executes a template with the given data and provides the output
func (t *TemplateParsed) Execute(data any) (*bytes.Buffer, error) {
	if t.Template == nil {
		return nil, errors.New("cannot execute template: template not initialized")
	}

	buf := new(bytes.Buffer)
	err := t.Template.ExecuteTemplate(buf, t.build.base+config.TemplateExt, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Group sets the cache group for the template being built
func (t *templateBuilder) Group(group string) *templateBuilder {
	t.build.group = group
	return t
}

// Key sets the cache key for the template being built
func (t *templateBuilder) Key(key string) *templateBuilder {
	t.build.key = key
	return t
}

// Base sets the name of the base template to be used during template parsing and execution.
// This should be only the file name without a directory or extension.
func (t *templateBuilder) Base(base string) *templateBuilder {
	t.build.base = base
	return t
}

// Files sets a list of template files to include in the parse.
// This should not include the file extension and the paths should be relative to the templates directory.
func (t *templateBuilder) Files(files ...string) *templateBuilder {
	t.build.files = files
	return t
}

// Directories sets a list of directories that all template files within will be parsed.
// The paths should be relative to the templates directory.
func (t *templateBuilder) Directories(directories ...string) *templateBuilder {
	t.build.directories = directories
	return t
}

// Store parsed the templates and stores them in the cache
func (t *templateBuilder) Store() (*TemplateParsed, error) {
	return t.renderer.parse(t.build)
}

// Execute executes the template with the given data.
// If the template has not already been cached, this will parse and cache the template
func (t *templateBuilder) Execute(data any) (*bytes.Buffer, error) {
	tp, err := t.Store()
	if err != nil {
		return nil, err
	}

	return tp.Execute(data)
}
```go

`pkg/services/template_renderer_test.go`

```go
package services

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/mikestefanello/pagoda/templates"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateRenderer(t *testing.T) {
	group := "test"
	id := "parse"

	// Should not exist yet
	_, err := c.TemplateRenderer.Load(group, id)
	assert.Error(t, err)

	// Parse in to the cache
	tpl, err := c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("htmx").
		Files("layouts/htmx", "pages/error").
		Directories("components").
		Store()
	require.NoError(t, err)

	// Should exist now
	parsed, err := c.TemplateRenderer.Load(group, id)
	require.NoError(t, err)

	// Check that all expected templates are included
	expectedTemplates := make(map[string]bool)
	expectedTemplates["htmx"+config.TemplateExt] = true
	expectedTemplates["error"+config.TemplateExt] = true
	components, err := templates.Get().ReadDir("components")
	require.NoError(t, err)
	for _, f := range components {
		expectedTemplates[f.Name()] = true
	}
	for _, v := range parsed.Template.Templates() {
		delete(expectedTemplates, v.Name())
	}
	assert.Empty(t, expectedTemplates)

	data := struct {
		StatusCode int
	}{
		StatusCode: 500,
	}
	buf, err := tpl.Execute(data)
	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Please try again")

	buf, err = c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("htmx").
		Files("htmx", "pages/error").
		Directories("components").
		Execute(data)

	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Please try again")
}

func TestTemplateRenderer_RenderPage(t *testing.T) {
	setup := func() (echo.Context, *httptest.ResponseRecorder, page.Page) {
		ctx, rec := tests.NewContext(c.Web, "/test/TestTemplateRenderer_RenderPage")
		tests.InitSession(ctx)

		p := page.New(ctx)
		p.Name = "home"
		p.Layout = "main"
		p.Cache.Enabled = false
		p.Headers["A"] = "b"
		p.Headers["C"] = "d"
		p.StatusCode = http.StatusCreated
		return ctx, rec, p
	}

	t.Run("missing name", func(t *testing.T) {
		// Rendering should fail if the Page has no name
		ctx, _, p := setup()
		p.Name = ""
		err := c.TemplateRenderer.RenderPage(ctx, p)
		assert.Error(t, err)
	})

	t.Run("no page cache", func(t *testing.T) {
		ctx, _, p := setup()
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Check status code and headers
		assert.Equal(t, http.StatusCreated, ctx.Response().Status)
		for k, v := range p.Headers {
			assert.Equal(t, v, ctx.Response().Header().Get(k))
		}

		// Check the template cache
		parsed, err := c.TemplateRenderer.Load("page", string(p.Name))
		require.NoError(t, err)

		// Check that all expected templates were parsed.
		// This includes the name, layout and all components
		expectedTemplates := make(map[string]bool)
		expectedTemplates[fmt.Sprintf("%s%s", p.Name, config.TemplateExt)] = true
		expectedTemplates[fmt.Sprintf("%s%s", p.Layout, config.TemplateExt)] = true
		components, err := templates.Get().ReadDir("components")
		require.NoError(t, err)
		for _, f := range components {
			expectedTemplates[f.Name()] = true
		}

		for _, v := range parsed.Template.Templates() {
			delete(expectedTemplates, v.Name())
		}
		assert.Empty(t, expectedTemplates)
	})

	t.Run("htmx rendering", func(t *testing.T) {
		ctx, _, p := setup()
		p.HTMX.Request.Enabled = true
		p.HTMX.Response = &htmx.Response{
			Trigger: "trigger",
		}
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Check HTMX header
		assert.Equal(t, "trigger", ctx.Response().Header().Get(htmx.HeaderTrigger))

		// Check the template cache
		parsed, err := c.TemplateRenderer.Load("page:htmx", string(p.Name))
		require.NoError(t, err)

		// Check that all expected templates were parsed.
		// This includes the name, htmx and all components
		expectedTemplates := make(map[string]bool)
		expectedTemplates[fmt.Sprintf("%s%s", p.Name, config.TemplateExt)] = true
		expectedTemplates["htmx"+config.TemplateExt] = true
		components, err := templates.Get().ReadDir("components")
		require.NoError(t, err)
		for _, f := range components {
			expectedTemplates[f.Name()] = true
		}

		for _, v := range parsed.Template.Templates() {
			delete(expectedTemplates, v.Name())
		}
		assert.Empty(t, expectedTemplates)
	})

	t.Run("page cache", func(t *testing.T) {
		ctx, rec, p := setup()
		p.Cache.Enabled = true
		p.Cache.Tags = []string{"tag1"}
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Fetch from the cache
		cp, err := c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		require.NoError(t, err)

		// Compare the cached page
		assert.Equal(t, p.URL, cp.URL)
		assert.Equal(t, p.Headers, cp.Headers)
		assert.Equal(t, p.StatusCode, cp.StatusCode)
		assert.Equal(t, rec.Body.Bytes(), cp.HTML)

		// Clear the tag
		err = c.Cache.
			Flush().
			Tags(p.Cache.Tags[0]).
			Execute(context.Background())
		require.NoError(t, err)

		// Refetch from the cache and expect no results
		_, err = c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		assert.Error(t, err)
	})
}
```go

`pkg/services/validator.go`

```go
package services

import (
	"github.com/go-playground/validator/v10"
)

// Validator provides validation mainly validating structs within the web context
type Validator struct {
	// validator stores the underlying validator
	validator *validator.Validate
}

// NewValidator creats a new Validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a struct
func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
```go

`pkg/services/validator_test.go`

```go
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	type example struct {
		Value string `validate:"required"`
	}
	e := example{}
	err := c.Validator.Validate(e)
	assert.Error(t, err)
	e.Value = "a"
	err = c.Validator.Validate(e)
	assert.NoError(t, err)
}
```go

`pkg/session/session.go`

```go
package session

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// ErrStoreNotFound indicates that the session store was not present in the context
var ErrStoreNotFound = errors.New("session store not found")

// Get returns a session
func Get(ctx echo.Context, name string) (*sessions.Session, error) {
	s := ctx.Get(context.SessionKey)
	if s == nil {
		return nil, ErrStoreNotFound
	}
	store := s.(sessions.Store)
	return store.Get(ctx.Request(), name)
}

// Store sets the session storage in the context
func Store(ctx echo.Context, store sessions.Store) {
	ctx.Set(context.SessionKey, store)
}
```go

`pkg/session/session_test.go`

```go
package session

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetStore(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())
	_, err := Get(ctx, "test")
	assert.Equal(t, ErrStoreNotFound, err)

	Store(ctx, sessions.NewCookieStore([]byte("secret")))
	_, err = Get(ctx, "test")
	assert.NoError(t, err)
}
```go

`pkg/tasks/example.go`

```go
package tasks

import (
	"context"
	"github.com/mikestefanello/backlite"
	"time"

	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"
)

// ExampleTask is an example implementation of backlite.Task
// This represents the task that can be queued for execution via the task client and should contain everything
// that your queue processor needs to process the task.
type ExampleTask struct {
	Message string
}

// Config satisfies the backlite.Task interface by providing configuration for the queue that these items will be
// placed into for execution.
func (t ExampleTask) Config() backlite.QueueConfig {
	return backlite.QueueConfig{
		Name:        "ExampleTask",
		MaxAttempts: 3,
		Timeout:     5 * time.Second,
		Backoff:     10 * time.Second,
		Retention: &backlite.Retention{
			Duration:   24 * time.Hour,
			OnlyFailed: false,
			Data: &backlite.RetainData{
				OnlyFailed: false,
			},
		},
	}
}

// NewExampleTaskQueue provides a Queue that can process ExampleTask tasks
// The service container is provided so the subscriber can have access to the app dependencies.
// All queues must be registered in the Register() function.
// Whenever an ExampleTask is added to the task client, it will be queued and eventually sent here for execution.
func NewExampleTaskQueue(c *services.Container) backlite.Queue {
	return backlite.NewQueue[ExampleTask](func(ctx context.Context, task ExampleTask) error {
		log.Default().Info("Example task received",
			"message", task.Message,
		)
		log.Default().Info("This can access the container for dependencies",
			"echo", c.Web.Reverse("home"),
		)
		return nil
	})
}
```go

`pkg/tasks/register.go`

```go
package tasks

import (
	"github.com/mikestefanello/pagoda/pkg/services"
)

// Register registers all task queues with the task client
func Register(c *services.Container) {
	c.Tasks.Register(NewExampleTaskQueue(c))
}
```go

`pkg/tests/tests.go`

```go
package tests

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/session"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// NewContext creates a new Echo context for tests using an HTTP test request and response recorder
func NewContext(e *echo.Echo, url string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

// InitSession initializes a session for a given Echo context
func InitSession(ctx echo.Context) {
	session.Store(ctx, sessions.NewCookieStore([]byte("secret")))
}

// ExecuteMiddleware executes a middleware function on a given Echo context
func ExecuteMiddleware(ctx echo.Context, mw echo.MiddlewareFunc) error {
	handler := mw(func(c echo.Context) error {
		return nil
	})
	return handler(ctx)
}

// ExecuteHandler executes a handler with an optional stack of middleware
func ExecuteHandler(ctx echo.Context, handler echo.HandlerFunc, mw ...echo.MiddlewareFunc) error {
	return ExecuteMiddleware(ctx, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			run := handler

			for _, w := range mw {
				run = w(run)
			}

			return run(ctx)
		}
	})
}

// AssertHTTPErrorCode asserts an HTTP status code on a given Echo HTTP error
func AssertHTTPErrorCode(t *testing.T, err error, code int) {
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, code, httpError.Code)
}

// CreateUser creates a random user entity
func CreateUser(orm *ent.Client) (*ent.User, error) {
	seed := fmt.Sprintf("%d-%d", time.Now().UnixMilli(), rand.Intn(1000000))
	return orm.User.
		Create().
		SetEmail(fmt.Sprintf("testuser-%s@localhost.localhost", seed)).
		SetPassword("password").
		SetName(fmt.Sprintf("Test User %s", seed)).
		Save(context.Background())
}
```go

`templates/components/core.gohtml`

```html
{{define "metatags"}}
    <title>{{ .AppName }}{{ if .Title }} | {{ .Title }}{{ end }}</title>
    <link rel="icon" href="{{file "favicon.png"}}">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    {{- if .Metatags.Description}}
        <meta name="description" content="{{.Metatags.Description}}">
    {{- end}}
    {{- if .Metatags.Keywords}}
        <meta name="keywords" content="{{.Metatags.Keywords | join ", "}}">
    {{- end}}
{{end}}

{{define "css"}}
    {{/* <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"> */}}
{{end}}

{{define "js"}}
    <script src="https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"></script>
    <script defer src="https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"></script>
{{end}}

{{define "footer"}}
    {{- if .CSRF}}
        <script>
            document.body.addEventListener('htmx:configRequest', function(evt)  {
                if (evt.detail.verb !== "get") {
                    evt.detail.parameters['csrf'] = '{{.CSRF}}';
                }
            })
        </script>
    {{end}}
    <script>
        document.body.addEventListener('htmx:beforeSwap', function(evt) {
            if (evt.detail.xhr.status >= 400){
                evt.detail.shouldSwap = true;
                evt.detail.target = htmx.find("body");
            }
        });
    </script>
{{end}}
```html

`templates/components/forms.gohtml`

```html
{{define "csrf"}}
    <input type="hidden" name="csrf" value="{{.CSRF}}"/>
{{end}}

{{define "field-errors"}}
    {{- range .}}
        <p class="help is-danger">{{.}}</p>
    {{- end}}
{{end}}
```html

`templates/components/messages.gohtml`

```html
{{define "messages"}}
    {{- range (.GetMessages "success")}}
        {{template "message" dict "Type" "success" "Text" .}}
    {{- end}}
    {{- range (.GetMessages "info")}}
        {{template "message" dict "Type" "info" "Text" .}}
    {{- end}}
    {{- range (.GetMessages "warning")}}
        {{template "message" dict "Type" "warning" "Text" .}}
    {{- end}}
    {{- range (.GetMessages "danger")}}
        {{template "message" dict "Type" "danger" "Text" .}}
    {{- end}}
{{end}}

{{define "message"}}
    <div class="notification is-light is-{{.Type}}" x-data="{show: true}" x-show="show">
        <button class="delete" @click="show = false"></button>
        MSG: {{.Text}}
    </div>
{{end}}
```html

`templates/emails/test.gohtml`

```html
Test email template. See services/mail.go to provide your implementation.
```html

`templates/layouts/auth.gohtml`

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        {{template "metatags" .}}
        {{template "css" .}}
        {{template "js" .}}
    </head>
    <body>
        <section class="hero is-info is-fullheight">
            <div class="hero-body">
                <div class="container">
                    <div class="columns is-centered">
                        <div class="column is-half">
                            {{- if .Title}}
                                <h1 class="title">{{.Title}}</h1>
                            {{- end}}
                            <div class="box">
                                {{template "messages" .}}
                                {{template "content" .}}

                                <div class="content is-small has-text-centered" hx-boost="true">
                                    <a href="{{url "login"}}">Login</a> &#9676;
                                    <a href="{{url "register"}}">Create an account</a> &#9676;
                                    <a href="{{url "forgot_password"}}">Forgot password?</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        {{template "footer" .}}
    </body>
</html>
```html

`templates/layouts/htmx.gohtml`

```html
{{template "content" .}}
```html

`templates/layouts/main.gohtml`

```html
<!DOCTYPE html>
<html lang="en" style="height:100%;">
    <head>
        {{template "metatags" .}}
        {{template "css" .}}
        {{template "js" .}}
    </head>
    <body class="has-background-light" style="min-height:100%;">
        
        <nav class="navbar is-dark">
            <a href="{{url "home"}}" class="navbar-item">{{.AppName}}</a>
        </nav>

        <p class="menu-label">Account</p>
        <ul class="menu-list">
            {{- if .IsAuth}}
                <li><a href="{{url "logout"}}">Logout</a></li>
            {{- else}}
                <li><a href="{{url "login"}}">Login</a></li>
            {{- end}}
        </ul>


        {{- if .Title}}
            <h1 class="title">{{.Title}}</h1>
        {{- end}}

        {{template "messages" .}}
        
        {{template "content" .}}

        {{template "footer" .}}
    </body>
</html>
```html

`templates/pages/config.gohtml`

```html
{{define "content"}}
    <pre>{{.Data}}</pre>
{{end}}
```html

`templates/pages/error.gohtml`

```html
{{define "content"}}
    {{if ge .StatusCode 500}}
        <p>Please try again.</p>
    {{else if  or (eq .StatusCode 403) (eq .StatusCode 401)}}
        <p>You are not authorized to view the requested page.</p>
    {{else if eq .StatusCode 404}}
        <p>Click {{link (url "home") "here" .Path}} to return home</p>
    {{else}}
        <p>Something went wrong</p>
    {{end}}
{{end}}
```html

`templates/pages/forgot-password.gohtml`

```html
{{define "content"}}
    <form method="post" hx-boost="true" action="{{url "forgot_password.submit"}}">
        <div class="content">
            <p>Enter your email address and we'll email you a link that allows you to reset your password.</p>
        </div>
        <div class="field">
            <label for="email" class="label">Email address</label>
            <div class="control">
                <input id="email" type="email" name="email" class="input {{.Form.Submission.GetFieldStatusClass "Email"}}" value="{{.Form.Email}}">
                {{template "field-errors" (.Form.Submission.GetFieldErrors "Email")}}
            </div>
        </div>
        <div class="field is-grouped">
            <p class="control">
                <button class="button is-primary">Reset password</button>
            </p>
            <p class="control">
                <a href="{{url "home"}}" class="button is-light">Cancel</a>
            </p>
        </div>
        {{template "csrf" .}}
    </form>
{{end}}
```html

`templates/pages/home.gohtml`

```html
{{define "content"}}
    Hello, please log in to view TODOs
{{end}}
```html

`templates/pages/login.gohtml`

```html
{{define "content"}}
    <form method="post" hx-boost="true" action="{{url "login.submit"}}">
        {{template "messages" .}}
        <div class="field">
            <label for="email" class="label">Email address</label>
            <div class="control">
                <input id="email" type="email" name="email" class="input {{.Form.Submission.GetFieldStatusClass "Email"}}" value="{{.Form.Email}}">
                {{template "field-errors" (.Form.Submission.GetFieldErrors "Email")}}
            </div>
        </div>
        <div class="field">
            <label for="password" class="label">Password</label>
            <div class="control">
                <input id="password" type="password" name="password" placeholder="*******" class="input {{.Form.Submission.GetFieldStatusClass "Password"}}">
                {{template "field-errors" (.Form.Submission.GetFieldErrors "Password")}}
            </div>
        </div>
        <div class="field is-grouped">
            <p class="control">
                <button class="button is-primary">Log in</button>
            </p>
            <p class="control">
                <a href="{{url "home"}}" class="button is-light">Cancel</a>
            </p>
        </div>
        {{template "csrf" .}}
    </form>
    
    <div class="is-divider" data-content="OR"></div>
    
    <div class="buttons">
        <a href="/auth/google" class="button is-danger">
            <span class="icon">
                <i class="fab fa-google"></i>
            </span>
            <span>Login with Google</span>
        </a>
        <a href="/auth/facebook" class="button is-info">
            <span class="icon">
                <i class="fab fa-facebook"></i>
            </span>
            <span>Login with Facebook</span>
        </a>
    </div>
{{end}}
```html

`templates/pages/register.gohtml`

```html
{{define "content"}}
    <form method="post" hx-boost="true" action="{{url "register.submit"}}">
        <div class="field">
            <label for="name" class="label">Name</label>
            <div class="control">
                <input type="text" id="name" name="name" class="input {{.Form.GetFieldStatusClass "Name"}}" value="{{.Form.Name}}">
                {{template "field-errors" (.Form.GetFieldErrors "Name")}}
            </div>
        </div>
        <div class="field">
            <label for="email" class="label">Email address</label>
            <div class="control">
                <input type="email" id="email" name="email" class="input {{.Form.GetFieldStatusClass "Email"}}" value="{{.Form.Email}}">
                {{template "field-errors" (.Form.GetFieldErrors "Email")}}
            </div>
        </div>
        <div class="field">
            <label for="password" class="label">Password</label>
            <div class="control">
                <input type="password" id="password" name="password" placeholder="*******" class="input {{.Form.GetFieldStatusClass "Password"}}">
                {{template "field-errors" (.Form.GetFieldErrors "Password")}}
            </div>
        </div>
        <div class="field">
            <label for="password-confirm" class="label">Confirm password</label>
            <div class="control">
                <input type="password" id="password-confirm" name="password-confirm" placeholder="*******" class="input {{.Form.GetFieldStatusClass "ConfirmPassword"}}">
                {{template "field-errors" (.Form.GetFieldErrors "ConfirmPassword")}}
            </div>
        </div>
        <div class="field is-grouped">
            <p class="control">
                <button class="button is-primary">Register</button>
            </p>
            <p class="control">
                <a href="{{url "home"}}" class="button is-light">Cancel</a>
            </p>
        </div>
        {{template "csrf" .}}
    </form>
{{end}}
```html

`templates/pages/reset-password.gohtml`

```html
{{define "content"}}
    <form method="post" hx-boost="true" action="{{.Path}}">
        <div class="field">
            <label for="password" class="label">Password</label>
            <div class="control">
                <input type="password" id="password" name="password" placeholder="*******" class="input {{.Form.GetFieldStatusClass "Password"}}">
                {{template "field-errors" (.Form.GetFieldErrors "Password")}}
            </div>
        </div>
        <div class="field">
            <label for="password-confirm" class="label">Confirm password</label>
            <div class="control">
                <input type="password" id="password-confirm" name="password-confirm" placeholder="*******" class="input {{.Form.GetFieldStatusClass "ConfirmPassword"}}">
                {{template "field-errors" (.Form.GetFieldErrors "ConfirmPassword")}}
            </div>
        </div>
        <div class="field is-grouped">
            <p class="control">
                <button class="button is-primary">Update password</button>
            </p>
        </div>
        {{template "csrf" .}}
    </form>
{{end}}
```html

`templates/templates.go`

```go
package templates

import (
	"embed"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type (
	Layout string
	Page   string
)

const (
	LayoutMain Layout = "main"
	LayoutAuth Layout = "auth"
	LayoutHTMX Layout = "htmx"
)

const (
	PageError          Page = "error"
	PageForgotPassword Page = "forgot-password"
	PageHome           Page = "home"
	PageLogin          Page = "login"
	PageRegister       Page = "register"
	PageResetPassword  Page = "reset-password"
	PageConfig         Page = "config"
)

//go:embed *
var templates embed.FS

// Get returns a file system containing all templates via embed.FS
func Get() embed.FS {
	return templates
}

// GetOS returns a file system containing all templates which will load the files directly from the operating system.
// This should only be used for local development in order to facilitate live reloading.
func GetOS() fs.FS {
	// Gets the complete templates directory path
	// This is needed in case this is called from a package outside of main, such as within tests
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	p := filepath.Join(filepath.Dir(d), "templates")
	return os.DirFS(p)
}
```go

`templates/templates_test.go`

```go
package templates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	_, err := Get().Open(fmt.Sprintf("pages/%s.gohtml", PageHome))
	require.NoError(t, err)
}

func TestGetOS(t *testing.T) {
	_, err := GetOS().Open(fmt.Sprintf("pages/%s.gohtml", PageHome))
	require.NoError(t, err)
}
```go

