package com.creational.design.Factory;

// Payment interface
interface Payment {
    void pay();
}

// UPI payment implementation
class UpiPayment implements Payment {
    @Override
    public void pay() {
        System.out.println("Payment done using UPI");
    }
}

// Card payment implementation
class CardPayment implements Payment {
    @Override
    public void pay() {
        System.out.println("Payment done using Card");
    }
}

// Main class
public class WithoutFactory {

    public static void main(String[] args) {

        String paymentType = "UPI"; // Change to CARD and run again

        Payment payment;

        // ❌ Object creation logic inside client
        if (paymentType.equalsIgnoreCase("UPI")) {
            payment = new UpiPayment();
        } else if (paymentType.equalsIgnoreCase("CARD")) {
            payment = new CardPayment();
        } else {
            throw new IllegalArgumentException("Invalid payment type");
        }

        // Business logic
        payment.pay();
    }
}

