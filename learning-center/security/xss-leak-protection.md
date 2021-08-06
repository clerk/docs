# XSS leak protection

> Cross-Site Scripting \(XSS\) attacks are a type of injection, in which malicious scripts are injected into otherwise benign and trusted websites.
>
> [The OWASP® Foundation, Cross Site Scripting \(XSS\)](https://owasp.org/www-community/attacks/xss/)

XSS vulnerabilities are incredibly serious and we recommend you reference the [OWASP Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Cross_Site_Scripting_Prevention_Cheat_Sheet.html) to learn how you can prevent an attack.

Prevention is just one aspect of an effective security model, though. To prepare for the worst case scenario where an attack is successful, you should ensure your application is configured to minimize the attack's surface area.

Clerk works to minimize the surface area by using **HttpOnly** cookies to store users' session tokens and prevent them from being leaked during XSS attacks.

## What is a HttpOnly cookie?

**HttpOnly** is a flag on the [**Set-Cookie** header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) that is issued by a server to set a cookie in the browser.

When this flag is set, the cookie can only be read from a server responding to an HTTP request. It is not accessible through Javascript with the **document.cookie** property.

Without the flag, a cookie can also be read by Javascript in the browser.

## How do HttpOnly cookies minimize XSS attacks?

XSS attacks allow the attacker to run arbitrary Javascript in the browser while users are navigating your application.

If Clerk did not use an HttpOnly cookie, but instead used a storage solution that is accessible through Javascript, it would mean the malicious Javascript could access users' session tokens.

Unauthorized access to session tokens is especially problematic because the tokens can be used to take actions on behalf of a user, even after a XSS vulnerability has been resolved. As a result, when session tokens leak, it is recommended you invalidate existing tokens and force users to sign in again.

HttpOnly cookies prevent session tokens from leaking through XSS attacks. Unless your application leaks session tokens another way, there will not be a need to sign your users out after a successful XSS attack.

Other forms of session token storage are susceptible to this issue, so even if you don't use Clerk, we highly recommend avoiding these forms of storage for session tokens:

* Cookies without the HttpOnly flag
* [Window.localStorage](https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage)
* [Window.sessionStorage](https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage)
* [IndexedDB](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API)

## How does Clerk set HttpOnly cookies on my behalf?

The setup process for Clerk requires you to set CNAME records in your DNS. These records enable us to set cookies for your domain on your behalf.

