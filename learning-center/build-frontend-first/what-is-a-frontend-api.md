---
description: A primer for Clerk's developer tools philosophy
---

# What is a Frontend API?

A Frontend API is designed to be accessed directly from the client.  It usually does not replace a traditional Backend API,  but instead augments it to improve the developer experience.

The primary difference between a Frontend API and a Backend API is the authentication strategy. While a Backend API authenticates with a secret key, a Frontend API authenticates with the current user's session. Because of this, Frontend API endpoints should only read data and execute actions if the current user has permission.

Sometimes, Frontend API endpoints require elevated permissions, which can granted by coordinating with a Backend API.

Because Frontend APIs are accessed from a different environment, additional considerations should be made around error handling and, for browsers, cross-origin resource sharing (CORS).
