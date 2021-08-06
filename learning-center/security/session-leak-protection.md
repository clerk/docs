# Session leak protection

When an application relies on a third party to operate a subdomain of their primary domain, it is possible for session tokens to leak to the third party operator.

Although the third party is trusted, leaking session tokens is not advised, since an attack on the third party could be chained into an attack on the application.

## Why would session tokens leak to the third party operator?

Consider an application where:

* _**www**.example.com_ is operated by the application owner
* _**dashboard**.example.com_ is operated by the application owner
* _**support**.example.com_ is operated by a third-party support desk vendor

In this example, the application owner may want to share the session token across **www** and **dashboard**.

The easiest way to accomplish this is by storing the session token in a cookie with the [**Domain** flag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#Attributes) set to **example.com** so it is shared across both subdomains.

Unfortunately, this also shares the cookie with **support**, which is operated by a third party.

## How does Clerk protect against session tokens leaking to third party subdomains

Instead of setting a shared cookie on _example.com_, Clerk sets two cookies: one on _**www**.example.com_ and one on _**dashboard**.example.com_.

### **How does Clerk set these cookies without operating www or dashboard?**

During the setup process, Clerk will ask the developer to configure two CNAME records in their DNS: one for **clerk.www.example.com** and one for **clerk.dashboard.example.com**.

These records allow us to set cookies scoped to **www** and **dashboard** respectively.

_Note: Clerk currently only allows setting a cookie on one subdomain, but this strategy is still used to ensure the cookie is scoped properly and does not leak to other subdomains._

