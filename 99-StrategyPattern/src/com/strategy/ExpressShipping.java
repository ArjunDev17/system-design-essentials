package com.strategy;

public class ExpressShipping implements ShippingStrategy {
    @Override
    public double calculateShippingCost(double weight, double distance) {
        return weight * distance * 1.0; // Higher cost for faster delivery
    }
}
