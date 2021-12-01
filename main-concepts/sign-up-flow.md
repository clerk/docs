---
description: Detailed guide on how the sign up process works
---

# Sign up flow

## Overview

Clerk provides a flexible way to build sign up flows in your application. You can use a single **Sign Up** object to gather information, verify their email address or phone number, add OAuth accounts, and finally, convert them into a **User.**

Every **Sign Up** has a set of requirements it must meet before it is turned into a **User**.  ****  These requirements are defined by the instance settings you selected in the dashboard.  Once all requirements are met, the **Sign Up** will turn into a new **User**, and an active session for that **User** will be created on the current **Client**.

Don't worry about collecting all the required fields at once and passing them to a single request. The API is designed to accommodate a progressive sign up flow, often corresponding to multi-step sign up forms.

## **Required fields**

The **Sign Up** will show the current state of requirement collection. You can consult the `required_fields`, `optional_fields` and `missing_fields` keys for a hint on where things are and what you need to do next.

**`required_fields`**:  all fields that must be collected before the **Sign Up** converts into a **User**.

**`optional_fields`**:  all fields that can be collected, but are not necessary to convert the **Sign Up**.

**`missing_fields`**: __ A subset of `required_fields`. It contains all fields that still need to be collected before a successful **Sign Up**. Note that the `missing_fields` will be updated dynamically. As you add more fields to the **Sign Up**, they will be removed from `missing_fields.` Once it's empty, your **Sign Up** will automatically convert into a **User**.

The values of the collected fields are all accessible on the root of the **Sign Up**, under their corresponding keys; `email_address` and `phone_number` are examples of such keys.

## Verified fields

Some fields, such as `email_address` and `phone_number` , must be verified before it is fully added to the **Sign Up**.  Similar to what happens with required fields, the **Sign Up** contains the current state of all verified fields.  The keys relative to verification are `unverified_fields` and `verifications`.

**`unverified_fields`**: a list of all **User** attributes that need to be verified and are pending verification. This is a list that gets updated dynamically, eventually becoming an empty array, when verification for all required fields  has been successfully completed.

**`verifications`**: an object that describes the current state of verification for the **Sign Up**. There are currently three different keys nested in there; `email_address`, `phone_number` and `external_account`.&#x20;
