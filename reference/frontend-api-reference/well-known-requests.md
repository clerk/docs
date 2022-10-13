---
description: Frontend API .well-known endpoints
---

# Well known requests

## Available requests



* **`GET`**` ``/.well-known/jwks.json`
* **`GET`**` ``/.well-known/openid-configuration`

{% swagger method="get" path="/.well-known/jwks.json" baseUrl="https://clerk.example.com" summary="" %}
{% swagger-description %}
Clerk uses [JSON Web Keys (JWK)](https://tools.ietf.org/html/rfc7517) to share the asymmetric, public cryptographic key used to validate the signature of a signed Clerk session JWT. The JWK is unique for each Clerk instance.

This endpoint is used by the official Clerk Backend SDKs to execute a network-less session JWT verification. The `jwks.json` should be cached in order to validate multiple JWTs without extra network calls.
{% endswagger-description %}

{% swagger-response status="200: OK" description="" %}
```json
{
  "keys": [
    {
      "use": "sig",
      "kty": "RSA",
      "kid": "ins_1lyWDZiobr600AKUeQDoSlrEmoM",
      "alg": "RS256",
      "n": "r8T9D2xqXhQ3zMFx3kGYx5NUJ4z1TynESh2YbLdC9WvtUS58TGMWpaku_VOBN7crxWTKC7KsEBbgI88FK_YmHJLuYXUdEPWj_dMnInAxeOMyQhFwuXkIZbU000N62GbpT5_xWFp5UGV_Vsl2try-1nrrD7xzid7DcxVaglvvRBmd51yJayxjEnsLMcRaGvBDDdjpgIVOq5Fz1RO520PFarcKQvlsltnCQbEUtUwprQMxQdE4F9L0D3mY6qV30Yz4mUprVCvwa2xLyaS9Ht6JMf2Uv0Sxvbhoq8H4SN4LysPMGGdTB6zwxpvRlJreOsWaO9HMr91xkAHeS8VqgfVaTw",
      "e": "AQAB"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/.well-known/openid-configuration" baseUrl="https://clerk.example.com" summary="" %}
{% swagger-description %}
**A minimal Well-known URI Discovery Mechanism for** [OpenID Connect](https://openid.net/specs/openid-connect-discovery-1\_0.html) which provides configuration information about Clerk as an Identity Provider (IDP).&#x20;

This endpoint can be used by API Gateways such as AWS and GCloud API Gateway.
{% endswagger-description %}

{% swagger-response status="200: OK" description="" %}
```javascript
{
  "issuer": "https://clerk.clerk.dev",
  "jwks_uri": "https://clerk.clerk.dev/.well-known/jwks.json"
}
```
{% endswagger-response %}
{% endswagger %}
