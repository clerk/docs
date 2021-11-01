# Overview

Client requests are responsible for handling the **Sign in**, **Sign up**, and **Session** processes.  More specifically, they operate on the **Client **object.  A reference to a **Client** object needs to be sent with_ every **Frontend API **request_,_ _through a Clerk-issued JWT_.  _Usually, a "device", such as a browser, or a mobile app, will store one **Client** object to send up with every request. By default, Clerk handles this process, and is explained in the [Session management](../../../main-concepts/session-management.md) section.&#x20;

## Requests summary

### [Client requests](../../backend-api-reference/clients.md)

* **`GET`**`/v1/client`
* **`PUT`**`/v1/client`
* **`DEL`**`/v1/client`

### [Sessions requests](sessions.md)

* **`GET `**`/v1/client/sessions/:id`
* **`POST`**`/v1/client/sessions/:id/end`
* **`POST`**`/v1/client/sessions/:id/remove`

### [Sign in requests](sign-ins.md)

#### **Create**

* **`PUT `**`/v1/client/sign_in_attempt`

#### Identify

* **`POST`**`/v1/client/sign_in_attempt/identify`

#### Factor one

* **`POST`**`/v1/client/sign_in_attempt/prepare_factor_one`
* **`POST`**`/v1/client/sign_in_attempt/attempt_factor_one`

#### Factor two

* **`POST`**`/v1/client/sign_in_attempt/prepare_factor_two`
* **`POST`**`/v1/client/sign_in_attempt/attempt_factor_two`

### [Sign up requests](sign-ups.md)

#### **Create/Update**

* **`PUT  `**`/v1/client/sessions/sign_up_attempt`
* **`PATCH`**`/v1/client/sessions/sign_up_attempt`

#### **Email Address**

* **`PUT `**`/v1/client/sign_up_attempt/email_address`
* **`POST`**`/v1/client/sign_up_attempt/email_address/send_verification_email`
* **`POST`**`/v1/client/sign_up_attempt/email_address/verify`

**Phone Number**

* **`PUT `**`/v1/client/sign_up_attempt/phone_number`
* **`POST`**`/v1/client/sign_up_attempt/phone_number/send_verification_sms`
* **`POST`**`/v1/client/sign_up_attempt/verify`

**External Account**

* **`PUT`**`/v1/client/sign_up_attempt/external_accounts`
