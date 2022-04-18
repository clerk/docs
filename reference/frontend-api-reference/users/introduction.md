# Overview

User requests are responsible for handling any actions on the current user, and you must be signed in as that **User**.  A reference to a **Client** object needs to be sent with __ each __ request, __ through a Clerk-issued JWT_._  This step is handled automatically if you're using one Clerk's Session Management.

## Requests summary

### [Current user requests](user.md)

* **`GET`**`  ``/v1/me`
* **`PATCH`**`/v1/me`

### [Email addresses requests](email-addresses.md)

* **`GET`**` ``/v1/me/email_addresses`
* **`POST`**`/v1/me/email_addresses`
* **`GET`**` ``/v1/me/email_addresses/:id`
* **`POST`**`/v1/me/email_addresses/:id/send_verification_email`
* **`POST`**`/v1/me/email_addresses/:id/verify`
* **`DEL`**` ``/v1/me/email_addresses/:id`

### [Phone numbers requests](phone-numbers.md)

* **`GET`**`  ``/v1/me/phone_numbers`
* **`POST`**` ``/v1/me/phone_numbers`
* **`GET`**`  ``/v1/me/phone_numbers/:id`
* **`PATCH`**`/v1/me/phone_numbers/:id`
* **`POST`**` ``/v1/me/phone_numbers/:id/send_verification_sms`
* **`POST`**` ``/v1/me/phone_numbers/:id/verify`
* **`DEL`**`  ``/v1/me/sessions/:id`

### [Profile image requests](profile-image.md)

* **`POST`**`/v1/me/profile_image`

### [Sessions requests](sessions.md)

* **`GET`**` ``/v1/me/sessions`
* **`GET`**` ``/v1/me/sessions/active`
* **`POST`**`/v1/me/sessions/:id/revoke`

### [Tokens requests](tokens.md)

* **`POST`**`/v1/me/tokens`

### [Organization requests](organizations.md)

* **`GET`**`/v1/me/organizations`
