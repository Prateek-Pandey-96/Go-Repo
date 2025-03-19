# Payment Gateway Implementation

## Overview
Payments is an integral component of any e-commerce or fintech. With the advent of digital India, we have seen different types of payment ecosystems emerging, including Payment Gateways, UPI, Rupay Network, etc. These solutions are easy to integrate, enabling organizations to start accepting payments in a matter of days.

Organizations often integrate multiple payment gateways and dynamically switch between them based on use cases. Customers can choose between different payment modes such as:

- **Net Banking** (Username & Password)
- **Credit/Debit Card** (Card Number, CVV)
- **UPI** (Upi Id)

## Real-World Example
Flipkart has integrated multiple PGs such as Razorpay, Citrus, PaySafe, and CCAvenue. They use these PGs based on different transaction types. For instance:
- Credit card transactions go to Razorpay.
- Net banking transactions go to CCAvenue.

Each PG integrates directly with multiple banks to facilitate transactions.

## Objective
The goal is to build a **Payment Gateway (PG)** similar to Razorpay or CCAvenue that enables clients (e.g., Flipkart) to onboard and process transactions through different banks.

### Features Required
- Support multiple client onboarding.
- Integrate multiple banks (HDFC, ICICI, SBI, etc.).
- Support different payment methods (UPI, Credit/Debit Card, Net Banking, etc.).
- Implement a **router** to direct transactions to specific banks based on:
  - Payment method (e.g., all credit card transactions go to HDFC).
  - Traffic distribution (e.g., 30% transactions go to Bank1, 70% to Bank2).
- Allow clients to configure supported payment methods (e.g., only UPI, exclude Net Banking, etc.).

## Assumptions
- Banks can randomly return success/failure (mock this behavior).
- Payments should be processed only if the correct parameters are provided.
- Banks require OTP verification (for simplicity, transactions will go through without OTP).

## Code Expectations
- **In-memory storage** for simplicity.
- Use any programming language.
- Basic function implementations without RESTful APIs.
- **Focus on code quality**, OOP principles, and separation of concerns.
- **Testability** should be considered.
- Bonus points for **Test Cases**.

## Required Functions
- `addClient()` - Add a client to PG.
- `removeClient()` - Remove a client from PG.
- `hasClient()` - Check if a client exists in PG.
- `listSupportedPaymodes()` - Show supported payment modes.
- `addSupportForPaymode()` - Add payment mode support.
- `removePaymode()` - Remove a payment mode.
- `makePayment()` - Process a payment.

**Execution Time Limit:** 3 seconds (Java)