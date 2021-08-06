# Overview

User requests are responsible for handling any actions on the current user, and you must be signed in as that **User**.  A reference to a **Client** object needs to be sent with __each __request, __through a Clerk-issued JWT_._  This step is handled automatically if you're using one Clerk's Session Management.

## Requests summary

### Current user

* **`GET`**  `/v1/me`
* **`PATCH`**`/v1/me`
* **`DEL`**  `/v1/me`

### Email addresses

* **`GET`** `/v1/me/email_addresses`
* **`POST`**`/v1/me/email_addresses`
* **`GET`** `/v1/me/email_addresses/:id`
* **`POST`**`/v1/me/email_addresses/:id/send_verification_email`
* **`POST`**`/v1/me/email_addresses/:id/verify`
* **`DEL`** `/v1/me/email_addresses/:id`

### Phone numbers

* **`GET`**  `/v1/me/phone_numbers`
* **`POST`** `/v1/me/phone_numbers`
* **`GET`**  `/v1/me/phone_numbers/:id`
* **`PATCH`**`/v1/me/phone_numbers/:id`
* **`POST`** `/v1/me/phone_numbers/:id/send_verification_sms`
* **`POST`** `/v1/me/phone_numbers/:id/verify`
* **`DEL`**  `/v1/me/sessions/:id`

### Profile image

* **`POST`**`/v1/me/profile_image`

### Sessions

* **`GET`** `/v1/me/sessions`
* **`GET`** `/v1/me/sessions/active`
* **`POST`**`/v1/me/sessions/:id/revoke`

### Tokens

* **`POST`**`/v1/me/tokens`

