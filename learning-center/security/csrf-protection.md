# CSRF protection

> CSRF is an attack that tricks the victim into submitting a malicious request. It inherits the identity and privileges of the victim to perform an undesired function on the victim’s behalf. For most sites, browser requests automatically include any credentials associated with the site, such as the user’s session cookie, IP address, Windows domain credentials, and so forth. Therefore, if the user is currently authenticated to the site, the site will have no way to distinguish between the forged request sent by the victim and a legitimate request sent by the victim.
>
> [The OWASP® Foundation, Cross Site Request Forgery (CSRF)](https://owasp.org/www-community/attacks/xss/)

Most Cross Site Request Forgery (CSRF) attacks can be protected against by properly configuring the way session tokens are stored. Clerk handles the necessary configuration on your behalf by configuring cookies with the **SameSite** flag.

## How does SameSite help prevent CSRF attacks?

**SameSite** is a flag on the [**Set-Cookie** header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) that is issued by a server to set a cookie in the browser.

When a cookie's SameSite flag is set to **Strict** or **Lax**, the cookie will not be sent with HTTP requests that originate from a third party site (a "cross-site" request). The cookie will only be sent when the originating site shares the same root domain, or more precisely, the same [public suffix](https://publicsuffix.org).

For example, consider a cookie that is set on _foo.example.com_:

* The cookie **will** be sent when _**bar**.example.com_ initiates a request to _foo.example.com_
* The cookie **will not** be sent when _bar.example.**org**_ initiates a request to _foo.example.com_

**Lax** alleviates this restriction slightly and allows for the cookie to be sent when the user is navigating from a third party site. In the example above, if the user is on _bar.example.**org**_ and clicks a link to _foo.example.com_, then the initial request to load _foo.example.com_ will include the cookie.

Clerk sets the SameSite flag to **Lax**, which is the default in modern browsers.

## Do I need to take additional steps to prevent CSRF attacks?

You do not, but it is still possible to accidentally make your application vulnerable to XSS attacks.

Because Clerk uses the **Lax** setting, it is critical to remember that navigation alone should never trigger a mutation in your backend. Otherwise, the user can be tricked into clicking a link that takes an action they did not intend.
