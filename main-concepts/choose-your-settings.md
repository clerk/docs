# Instance settings

Before getting started, you need to choose how you want your users to sign in and sign up.  There are two parts to settings, the first part affects how the Frontend API works, and what your User Model should look like.  The second part affects how the UI and UX of the Clerk Modals and Hosted UIs. 

After you create your application, navigate to your development instance then go to **Settings &gt; User Settings**

These settings affect how your users can sign up, sign in, and what attributes are editable in their user profile via the API.

## Contact Information

These requirements will affect what fields are required during sign up, and must always be on their account.

#### Email address

During sign up, a user must verify their email address, and they must always have at least one email address on their account at all times.

#### Phone number

During sign up, a user must verify their phone number, and they must always have at least one phone number on their account at all times.

#### Email address OR phone number

During sign up, a user must verify a phone number OR an email address, and they must always have at least one phone number OR email address on their account at all times.

#### None

Neither an email address or phone number are required at sign up,  in order to select this option, you need to require either a username or a social provider.

## Social Providers

Clerk offers a short list of social providers, but we can add most new ones in about a week.  If you don't see what you need please contact us and we'll most likely be able to build what you're looking for

* **Facebook**
* **Github**
* **Google**
* **Hubspot**
* **TikTok**

#### Don't see what you need? __[_Request others here_](https://www.clerk.dev/support) __

## Authentication strategy

This setting does not affect your Frontend API, it only affects which authentication strategy is displayed to your Users first.

#### Passwordless

Send an email or an SMS with a code to the User, so that they can verify they own their account.

#### Password

Passwords are the standard authentication strategy - everyone knows how they work, and they're reliable.

> Clerk protects your users against weak passwords and leaked passwords by default.

## 2-step verification

For even more protection, and what should be a standard option.  Give your Users the option to further protect their account by forcing an additional level of verification.

* **SMS Code**

**More 2SV options coming soon**

## Personal information

Collect additional information at sign up, required or optional

* **Name**

**More options coming soon**

## How BigCos do it

### Google

* Username is required \(which turns into your `@gmail.com` account\)
* Password is required
* Birthday is required
* Gender is required
* Phone number is optional \(used for 2-step verification\)
* "recovery email address" is optional at sign up
* Multi-Session 

Clerk settings for this setup:

**Contact Information =** Email Address

### Facebook

* Email address OR phone number is required
* Password is required
* Birthday is required
* Gender is required
* Phone number is optional \(used for 2-step verification\)

### Stripe

* Email address
  * unverified - \(forced verification later\)
* Full name
* Country
* Password
* Phone number is optional \(used for 2-step verification\)

### Reddit

* Sign in with Google
* Sign in with Apple
* Email address
  * unverified 
* Username \(randomized at sign up\)
* Password
* Authenticator App is optional \(used for 2-step verification\)
* Multi-Session

