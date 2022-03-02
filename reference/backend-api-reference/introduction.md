# Introduction

### API Reference&#x20;

The Clerk Backend API is organized around [REST](http://en.wikipedia.org/wiki/Representational\_State\_Transfer) and is built to be used from your backend code. All requests accept [JSON-encoded](https://www.json.org/json-en.html) (recommended) and [form-encoded](https://en.wikipedia.org/wiki/POST\_\(HTTP\)#Use\_for\_submitting\_web\_forms) request bodies, return [JSON-encoded](https://www.json.org/json-en.html) responses and uses standard HTTP response codes, authentication, and verbs.

{% hint style="info" %}
Base URL: **https://api.clerk.dev/**
{% endhint %}

### Authentication

The Clerk Backend API uses secret API keys. You can find and create your instance API keys in Clerk Dashboard **API Keys -> Backend API Keys** page.&#x20;

{% hint style="warning" %}
These keys should never be shared with anyone, as they allow the holder to access all your Clerk resources.
{% endhint %}

&#x20;An example request is the following:

```bash
curl \
 -H "Authorization: Bearer test_abc...xyz" \
 -H "Content-type: application/json" \ 
 https://api.clerk.dev/v1/users
```

{% hint style="info" %}
Be careful not to expose these keys in your Git repository, frontend code, or anywhere else that is public.&#x20;
{% endhint %}

