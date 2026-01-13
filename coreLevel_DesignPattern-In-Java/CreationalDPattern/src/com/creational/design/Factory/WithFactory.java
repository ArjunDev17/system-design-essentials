package com.creational.design.Factory;



// Factory class
class PaymentFactory {

    // Factory Method
    public static Payment getPayment(String paymentType) {

        if (paymentType.equalsIgnoreCase("UPI")) {
            return new UpiPayment();
        } else if (paymentType.equalsIgnoreCase("CARD")) {
            return new CardPayment();
        } else {
            throw new IllegalArgumentException("Invalid payment type");
        }
    }
}

// Main class
public class WithFactory {

    public static void main(String[] args) {

        String paymentType = "UPI"; // Change to CARD and run again

        // ✅ Client delegates creation to factory
        Payment payment = PaymentFactory.getPayment(paymentType);

        // Business logic
        payment.pay();
    }
}

