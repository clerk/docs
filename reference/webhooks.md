---
description: Receive event notifications with Clerk webhooks
---

# Webhooks

If you need to keep your local store in sync with up to date user info,  or you want to send a welcome email, you'll need to setup a webhook, so Clerk can tell you about events that happen on your instance.

For more information on how to setup webhooks in your instance and how to architecture your application in order to be updated with Clerk events the right way, check our [Sync data to your back-end](../popular-guides/sync-data-to-your-back-end.md) guide.

## Supported events

Here is a list of all the events you can choose from:

* `user.created`&#x20;
* `user.updated`&#x20;
* `user.deleted`
* `session.created`
* `session.ended`
* `session.removed`
* `session.revoked`

### Payload structure

The payload of the message includes the type of the event in the `type` property.&#x20;

The `data` property contains the actual payload sent by Clerk. This can either be a user or a session object, depending on the event type.

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

In your backend, when ingesting events, you should follow the [Svix documentation](https://docs.svix.com/receiving/verifying-payloads) to verify that the webhook request is valid before processing it.

{% hint style="danger" %}
If you don't verify the request, your app will be susceptible to a number of attacks since your webhook endpoint is open to the public.
{% endhint %}
