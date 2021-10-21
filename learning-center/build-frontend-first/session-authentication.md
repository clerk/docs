# Session authentication

Your Frontend API should authenticate with the current user's session.  To accomplish this safely, the developer must pass you a **session token** that is guaranteed to have been generated on their server.

**A key characteristic of the session token is that it should authenticate every endpoint of your Frontend API.** This makes it so a team's backend developer only has to create one server-side endpoint for generating the token. From there, the team's frontend developer can request a token as-needed to access any endpoint on your Frontend API.

Session tokens can be generated with any number of strategies. Here, we document two strategies that we've seen in the wild.

### Have the developer generate a signed JSON Web Token (JWT)

This strategy is recommended since it is not network-bound. JWTs can be used to sign arbitrary JSON.

JWTs call the JSON keys "claims."  We recommend you ask the developer to put the User ID in the "sub" (subject) claim and the token experiation time in the "exp" claim, which are the appropriate registered claim names under [RFC 7519](https://tools.ietf.org/html/rfc7519#section-4.1).  If you prefer, you can also ask the developer to put the claims an arbitrary field.

You should ensure the expiration is not too far in the future, no greater than a minute.  The developer should be encouraged to generate many, short-lived tokens instead of few, long-lived tokens.

A critical part of this strategy is asking the developer sign the JWT in a manner that allows you to verify it was generated on their server.  You can choose between supporting a symmetric or asymmetric signing algorithm:

* If you choose a symmetric signing algorithm (HMAC), you and the developer will share a secret key. Then, the developer will sign the JWT with that secret key, and you will verify they created it using the same key.
* If you choose an asymmetric signing algorithm (RSA, ECDSA), the developer will generate a private/public key pair and share the public key with you.  Then, the developer will sign the JWT with their private key, and you will verify they created it using the public key.

In either case, the shared secret or public key should be exchanged out-of-band, usually within your service's dashboard.

To see an example of developer-generated JWTs serving as session tokens, please see Hasura's documentation for [Authenticating with JWT](https://hasura.io/docs/1.0/graphql/core/auth/authentication/jwt.html).

### Have the developer request a session token from your Backend API

Although this strategy is network-bound, you may prefer the developer experience since it does not require you to educate developers about JWTs.

Instead of having the developer generate a JWT, you can create an endpoint on your Backend API that generates and responds with a session token.  Then, your Frontend API will be responsible ensuring the session token is one you created and is still valid.

To communicate the current user, the developer should pass you the current user's ID when making the request to your Backend API.

The closest example we've seen to this strategy is Stripe Billing's [Customer Portal](https://stripe.com/docs/billing/subscriptions/integrating-customer-portal#redirect).  It is not a perfect example since the generated session token cannot be used to generically authenticate with a Frontend API, but instead can only be used to redirect the user to a portal.  The mechanism, though, is quite similar: the developer passes Stripe's Backend API a User ID to generate a short-lived session token.
