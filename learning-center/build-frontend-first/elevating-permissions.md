# Elevating permissions

While session authentication is helpful for handling Frontend API requests that only require user permission, some requests may require both user and developer permission before they can be processed.

As an example, consider an API that allows a developer to collect credit card payments. Before the payment can be accepted, two pieces of information are required:

1. The user's credit card details
2. The payment amount

While the user's credit card details can safely be collected from a Frontend API, the payment amount must be verified from the developer's backend. This coordination with the developer's backend is a process we call **elevating permissions**.

There are two ways to elevate permissions before processing a request: **backend finalization** and **backend initialization**.

## Backend finalization <a href="backend-finalization" id="backend-finalization"></a>

Backend finalization means that permissions are elevated **after** a request begins on your Frontend API.

Extending on our credit card payments example, we can see an example of backend finalization in Braintree's Credit Card API. The process for collecting a credit card payment with Braintree is:

1. ​[Collect the user's credit card details](https://developers.braintreepayments.com/guides/credit-cards/client-side/javascript/v3) (frontend)
2. ​[Create a transaction](https://developers.braintreepayments.com/guides/credit-cards/server-side/node) with the charge amount (backend)

This is an example of backend finalization because developer permission is gained **after** the user interacts with Braintree from the frontend.

Backend finalization is best in situations where the developer may want to take additional steps after the action is finalized from their backend. Since the developer is already triggering the finalization on their backend, it's easy to take additional steps immediately after it completes.

The challenge you might face with backend finalization, though, is that it's hard to provide an end-to-end experience for the end user (meaning, the developer's customer). For example, since the process is finalized on the developer's backend, you won't be able to present errors associated with finalization directly to the end user.

## Backend initialization <a href="backend-initialization" id="backend-initialization"></a>

Server initialization means that permissions are elevated **before** a request arrives to your Frontend API.

Extending again on our credit card payments example, we can see an example of server initialization in Stripe's Payment Intent API. The process for collecting a credit card payment with Stripe is:

1. ​[Create a Payment Intent](https://stripe.com/docs/payments/accept-a-payment#web-create-payment-intent) object with the charge amount (backend)
2. ​[Collect the user's credit card details](https://stripe.com/docs/payments/accept-a-payment#web-collect-card-details) (frontend)
3. ​[Submit the payment](https://stripe.com/docs/payments/accept-a-payment#web-submit-payment) (frontend)

This is an example of backend initialization because developer permission is gained **before** the user interacts with Stripe from the frontend.

Backend initialization is best in situations where you want to provide an end-to-end experience for the end user, since you can build a Frontend UI to handle and present any finalization errors from your Frontend API. Stripe takes advantage of this capability with [Stripe Checkout](https://stripe.com/payments/checkout).

The challenge you might face with backend initialization, though, is that it's harder for the developer to take additional actions after finalization from the client. When an action is finalized on the frontend, it's not reliable for the developer to trigger additional steps directly from the frontend. Instead, you must build a build asynchronous webhooks to notify their backend when the action is finalized. Stripe [has built webhooks](https://stripe.com/docs/payments/accept-a-payment#web-fulfillment) for this use case.\
