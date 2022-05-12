# Introduction

If you'd like early access to one of our upcoming features, you can use these API endpoints to enable and try them. But please note, that these may change and are not guaranteed to have a static API. We may also make minor breaking changes when we publish the beta feature live.

## [Domain](domain.md)

Update the domain of a production instance

### Available requests

* **`PUT`**` ``/v1/beta_features/domain`

## [Instance Settings](instance-settings.md)

Modify some of your instance's settings.

### Available requests

* **`PATCH`**`/v1/beta_features/instance_settings`

## [Tokens](introduction.md#tokens)

Create JSON Web Tokens (JWTs) based on a `jwt_template` defined in your instance.

### Available requests

* **`POST`**` ``/v1/tokens`
