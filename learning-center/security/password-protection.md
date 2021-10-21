# Password protection and rules

## Password rules

Clerk refers to the National Institute of Standards and Technology (NIST) guidelines to determine the character rules for passwords:

> Verifiers SHALL require subscriber-chosen memorized secrets to be at least 8 characters in length. Verifiers SHOULD permit subscriber-chosen memorized secrets at least 64 characters in length. All printing ASCII \[RFC 20] characters as well as the space character SHOULD be acceptable in memorized secrets. Unicode \[ISO/ISC 10646] characters SHOULD be accepted as well. To make allowances for likely mistyping, verifiers MAY replace multiple consecutive space characters with a single space character prior to verification, provided that the result is at least 8 characters in length. Truncation of the secret SHALL NOT be performed. For purposes of the above length requirements, each Unicode code point SHALL be counted as a single character.
>
> [NIST Special Publication 800-63B](https://pages.nist.gov/800-63-3/sp800-63b.html#sec5)

While these rules might seem lax independently, NIST's additional [leaked password protection](broken-reference) guidelines do more to prevent the use of unsafe passwords.

Also, please bear in mind, that passwords are not a requirement for using Clerk. Applications can be configured to use a passwordless strategy that relies on your users being sent one-time passwords instead.

## Leaked password protection

Clerk refers to the National Institute of Standards and Technology (NIST) guidelines to determine its handling of leaked passwords:

> When processing requests to establish and change memorized secrets, verifiers SHALL compare the prospective secrets against a list that contains values known to be commonly-used, expected, or compromised. For example, the list MAY include, but is not limited to:
>
> * Passwords obtained from previous breach corpuses.
>
> [NIST Special Publication 800-63B](https://pages.nist.gov/800-63-3/sp800-63b.html#sec5)

Specifically, Clerk contracts with [have i been pwned](https://haveibeenpwned.com) to compare prospective passwords against its corpus of over 10 billion leaked credentials.
