# Introduction

If you'd like early access to one of our upcoming features, you can use this API to toggle them on.  But please note, that these may change and are not guaranteed to have a static API.  We may also make minor breaking changes when we publish the beta feature live.

## [Instance Settings](instance-settings.md)

Modify some of your instance's settings.

### Available requests

* **`PATCH`**`/v1/beta_features/instance_settings`

## ****[**Allowlist identifiers**](../allowlist-identifiers.md)****

Setting `restricted_to_allowlist`to `true` in the instance\_settings will block users from signing up, unless you have explicitly added them to the allow list.  You can also notify users if they're allowed to sign up.  This feature is part one of our upcoming "Invitations" feature.

### Available requests

* **`GET`**` ``/v1/allowlist_identifiers`
* **`POST`**`/v1/allowlist_identifiers`
* **`DEL`**` ``/v1/allowlist_identifiers/:id`

## [Tokens](introduction.md#tokens)

Create JSON Web Tokens (JWTs) based on a `jwt_template` defined in your instance.

### Available requests

* **`POST`**` ``/v1/tokens`
