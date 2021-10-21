# Handling Errors

Frontend APIs are unique in that their request parameters are often tied directly to untrusted user input, while Backend APIs can normally expect to receive validated input.

Because of this, you may want to structure your Frontend API errors differently from your Backend API errors. We recommend:

* Use different error codes for user input errors versus developer implementation errors.
* If more than one user field has an error, return an error for each field. Don't stop validation after finding just one error.
* Return user-readable error messages instead of (or in addition to) developer-readable error messages.

The goal with these recommendations is to allow developers to lean on you for validation, in lieu of building their own client-side validation.

For accessibility reasons, developers should be presenting user input errors alongside the field with a mistake. [Here's an example from Bootstrap](https://getbootstrap.com/docs/5.0/forms/validation/#server-side):

![Bootstrap validation presentation](https://gblobscdn.gitbook.com/assets%2F-MNv9DpyBaNL\_\_j7ZW\_m%2F-MOFYQzJee1WQvjsOxEy%2F-MOFb\_YrfZ7mQlXVfLq0%2FScreen%20Shot%202020-12-11%20at%202.05.08%20AM.png?alt=media\&token=1f16bddd-97df-4608-80f2-c5575227360f)

Your Frontend API errors should be structured in a way that makes it easy for developers to present accessible errors to their end users.\
