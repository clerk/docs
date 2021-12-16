# Sign in flow

**Sign in**'s are initiated by creating a **Sign in** object on the current **Client**.  The **Sign in** handles all the state and logic associated with a **Sign in**.  If the **Sign in** is successfully authenticated, it will transform into an active session, on the current **Client**.

## Completing a sign in

There are 3 main steps a user must perform in order to complete a sign in:

### Step 1: Identification

The first step a user needs to make is to identify what account they'd like to sign in to.  This is done with an identifier, which can either be an email address, a phone number, or a username.&#x20;

### Step 2: Factor one verification

Once a user is identified, they need to prove that "they are who they say they are".  This is the process of "authenticating" the user.  There's a number of strategies a user can use to perform authentication, the most basic being the humble **password.** Other authentication strategies that Clerk provides, are sending an **email\_code** or **phone\_code.**  This will send an email or an SMS message with a code, that the user then must provide in order to sign in.&#x20;

### Step 3: Factor two verification (optional)

This step only applies to users that have turned on two factor authentication for their user.  When one form of verification isn't enough, trust two! But seriously.  Forcing two different verification steps _vastly_ increases the security of your account. The most common setup a user will have to protect their account is using a **password** as their first kind of verification, and a **phone\_code** as their second kind of verification.

## Verification Strategies

Now that we've gone over the basic steps of completing a sign in, we'll go over the different types of verification strategies provided, and how to use Clerk's API in each scenario.

A verification can loosely be separated into two main steps _preparing_ and _attempting_ the verification.  Here's a list of each verification strategy, and what steps must be performed:

* **password**
  * Prepare: _no prepare step_.
  * Attempt: user supplies the password.
* **email\_code**
  * Prepare: sends an email with a code.
  * Attempt: user supplies the code.
* **phone\_code**
  * Prepare: sends an SMS with a code.
  * Attempt: user supplies the code.
* **oauth\_google, oauth\_facebook,oauth\_twitter, etc...**
  * Prepare: generates a link the user must visit.
  * Attempt: _no attempt step, this happens externally._

## Sign in with OAuth

There's one more exception we need to talk about.  **oauth\_strategies** _do not need to be identified_ since this happens externally by the oauth provider.  A sign in can be identified before performing an oauth verification.  If this is the case, the identified account must match the one authenticated by the oauth verification.  If it doesn't the oauth verification will fail.

## User management settings

Described above are all possible verification and identification strategies.  Your API may not have all of these, you can change which strategies are activated through the Clerk Dashboard, in the "Settings > User management " page of your instance.&#x20;
