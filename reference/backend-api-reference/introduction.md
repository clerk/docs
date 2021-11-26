# Introduction

The Clerk Backend API is built to be used from your backend code. We make our best attempt to make your Clerk resources available to you in a predictable manner. All requests accept form-encoded request bodies, and respond with a JSON-encoded body.

> Base URL: **https://api.clerk.dev/**

### Authentication

The Clerk Backend API uses secret API keys.  You can find and create your instance API keys in Clerk Dashboard **API Keys **page.&#x20;

{% hint style="warning" %}
These keys should never be shared with anyone, as they allow the holder to access all your Clerk resources.
{% endhint %}

&#x20;An example, request is the following:

```bash
curl \
 -H "Authorization: Bearer test_abc...xyz" \
 -H "Content-type: application/json" \ 
 https://api.clerk.dev/v1/users
```

{% hint style="info" %}
Be careful not to expose these keys in your Git repository, frontend code, or anywhere else that is public.&#x20;
{% endhint %}

