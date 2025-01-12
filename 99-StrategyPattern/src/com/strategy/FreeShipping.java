package com.strategy;

public class FreeShipping implements ShippingStrategy {
    @Override
    public double calculateShippingCost(double weight, double distance) {
        return 0.0; // Free shipping for premium users
    }
}
