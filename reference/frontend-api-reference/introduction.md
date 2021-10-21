# Introduction

The **Frontend API** is built to be used from your frontend code. It is the same API that our pre-built UIs use.  If you need to augment your Sign in, Sign up, or User profile processes, this API is for you.   We follow a familiar [REST](https://en.wikipedia.org/wiki/Representational\_state\_transfer) style, so hopefully this API is clear.  You can always [reach out](https://www.clerk.dev/support) for help.

The **Frontend API** is accessed via the Frontend API Domain listed in your [dashboard](https://dashboard.clerk.dev).  Navigate to your instance, then it will be in the **Home>Instance Configuration **section.

In development, the URL will follow the following pattern:

> Development Base URL: **https://clerk.abc1.edf2.lcl.dev/**

In production, it will be accessed via your own domain, and will follow the following pattern:

> Production Base URL: **https://clerk.your-domain.com**

All requests accept form-encoded request bodies, and respond with a JSON-encoded body.

## Frontend API Configuration

All configuration for the frontend happens in the [dashboard](https://dashboard.clerk.dev).  Navigate to your instance, then to **Settings>User management **to see and modify these settings.

The **User management **settings you select _will modify certain parts _of the **Frontend API**.

For example, if you choose to require users to sign in with a phone number, using a passwordless** **authentication mechanism (an SMS sent to the phone), then, that instance's **Frontend API **will not allow sign in's with an email address or username.  This, in turn, modifies what parameters the `/sign_in` request can accept.

## Frontend API structure

The **Frontend API** is split into two main categories.  "Client requests" and "User requests".&#x20;

"Client requests" make up a bulk of the **Frontend API.  **It is what powers the Clerk-provided Sign in and Sign up pages.  "User requests" part of the **Frontend API** is what powers the Clerk-provided User profile page.  Again, the Clerk-provided UIs are built using this exact same API, so you can be in complete control of all the UIs if you like.  You can always default to the Clerk-provided UIs, or start with them, then modify yours later. &#x20;

## Client requests

[Client requests](user-api/) are responsible for handling the **Sign in**, **Sign up**, and **Session** processes.  More specifically, they operate on the **Client **object.  A reference to a **Client** object needs to be sent with_ every **Frontend API **request_,_ _through a Clerk-issued JWT_.  _Usually, a "device", such as a browser, or a mobile app, will store one **Client** object to send up with every request. By default, Clerk handles this process, and is explained in the [Session management](broken-reference) section.&#x20;

### **Client object**

The **Client **object contains the current state of the **Sign in**,** Sign up**,** and Session**(s) objects, as you can see below.

> Note: Clerk allow's a Client object to have multiple active **Session**'s.  This enables someone to be logged into multiple accounts on the same device, and freely switch between them.  You can see this in action on the Clerk dashboard, by clicking on your Profile Image in the top right, and selecting "Add Account". &#x20;
>
> By default, Clerk runs in "Single-session mode", in which case there can be only zero or one sessions. For more information on how to support "Multi-session mode" check out our Multi Session Guide _(coming soon)._ The feature is heavily inspired by GMail's "Account switcher", and a feature we think every web-application should include.

**Client object schema**

```javascript
{
        "object": "client",
        "id": "client_1sV3b42PUQh1kX7wi26opkRXR4x",
        "sessions": [],
        "sign_in_attempt": null,
        "sign_up_attempt": null,
        "last_active_session_id": null,
        "created_at": 1620943997,
        "updated_at": 1620943997
}
```

For more information see the in-depth [Client requests](user-api/) page.

## User requests

[User requests](users/) are mainly responsible for modifying any active **User**'s.  This happens through **Session**'s. Each **Session **contains one **User**. &#x20;

Once someone is signed in there will be an active **Session **on the **Client** object,** **which is automatically authorized to make requests on behalf of the containing **User**.  Some examples of actions a **Session** can take for the **User** are:

* Changing the first or last name
* Changing the password
* Enabling/Disabling 2FA
* Revoking active sessions on other devices for their user.
* ... and much more

### **User object**

The **User **object contains all security/profile information. Here's an example:

```javascript
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user",
    "username": "boss-clerk",
    "first_name": "Boss",
    "last_name": "Clerk",
    "profile_image_url": "https://images.clerk.services/clerk/default-profile.svg",
    "primary_email_address_id": "idn_1oBNgISXFbSf5m0uP2Wl0qWtNGX",
    "primary_phone_number_id": null,
    "password_enabled": true,
    "two_factor_enabled": false,
    "email_addresses": [
        {
            "id": "idn_1oBNgISXFbSf5m0uP2Wl0qWtNGX",
            "object": "email_address",
            "email_address": "boss@clerk.dev",
            "verification": {
                "status": "verified",
                "strategy": "email_code",
                "attempts": 1,
                "expire_at": 1612756733
            },
            "linked_to": []
        }
    ],
    "phone_numbers": [
        {
            "id": "idn_1q8Uq8Mc4t7WWMy9Z6Og0gNJVui",
            "object": "phone_number",
            "phone_number": "+15555555555",
            "reserved_for_second_factor": false,
            "verification": {
                "status": "verified",
                "strategy": "phone_code",
                "attempts": 1,
                "expire_at": 1616461499
            },
            "linked_to": []
        }
    ],
    "external_accounts": [],
    "public_metadata": {},
    "private_metadata": {},
    "created_at": 1612756155,
    "updated_at": 1612756155
}
```

For more information see the in-depth [User requests](users/introduction.md#user-requests) page.

## Session management

The Clerk Frontend API handles session management for you by default, secured using the [best practices](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-63b.pdf) such as HTTP-only cookies.  Clerk will automatically create and update cookies during sign in and sign up's so you don't have to.  These cookies will be sent to the **Frontend API** automatically and power the authentication of all **Frontend API **requests. &#x20;

These cookies will also get sent up to _your API._   This happens automatically for development and staging instances.  For production instances, you specified the subdomain which to send cookies during the instance creation process.  You can see what subdomain you selected by navigating to: **Instance>Settings>DNS>Domains>Backend Host**

> While production instances **always** use HTTP-only cookies, which is the most secure way to manage sessions, development instances do not use http only cookies.  This is due to limitations that occur when developing on your local machine. You should not notice these differences, but because browser security is constantly shifting, something may go wrong from time to time.  If you notice anything unexpected on a development instance, [we are always available to help](https://www.clerk.dev/support).  For more information on how Clerk manages sessions on development instances visit our \[in-depth-session-management-guide _(coming soon!)_]

Non-browser environments, such as iOS and Android apps, do not use cookies.  Instead, you should store the Clerk-issued JWT on the device, and send it up in an `Authorization` header.

You can also opt-out of Clerk's Session management feature in a browser, this will put you in charge of managing the Clerk-issued JWT, sending it up with requests to both the **Frontend API **and to _your API.  _This feature is coming soon, if you would like to turn off session management, please [contact us](https://www.clerk.dev/support). &#x20;
