---
description: Receive event notifications with Clerk webhooks
---

# Webhooks

If you need to keep your local store in sync with up to date user info,  or you want to send a welcome email, you'll need to setup a webhook, so Clerk can tell you about events that happen on your instance.

Webhooks can be configured from [your dashboard](https://dashboard.clerk.dev).  You can turn them on by selecting your Instance, then going to the **Integrations** page, and toggling on "Svix".  ****Clerk uses Svix as a provider to ensure a robust and reliable webhook experience.

Clicking on **Manage Webhooks** will bring you to the webhook management UI which will allow you to add/remove endpoints, and see event logs.

## Supported events

You can select which events you'd like to listen to via the webhook management UI.  Here is a list of all the events you can choose from:

* `user.created` 
* `user.updated` 
* `user.deleted`
* `session.created`
* `session.ended`
* `session.removed`
* `session.revoked`

### Payload structure

```javascript
{
   "object": "event",
   "type": "user.created",
   "data": {
      // user object      
      // or
      // session object
  }
}

```

## Verifying requests

In your backend, when ingesting events, you should follow [this documentation](https://docs.svix.com/receiving/verifying-payloads) to verify that the webhook request is valid before processing it.

**Warning: If you don't verify the request, your app will be susceptible to a number of attacks since your webhook endpoint is open to the public.**

