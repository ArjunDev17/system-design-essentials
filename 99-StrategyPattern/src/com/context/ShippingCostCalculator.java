package com.context;

import com.strategy.ShippingStrategy;

public class ShippingCostCalculator {
    private ShippingStrategy strategy;

    // Set the strategy dynamically
    public void setStrategy(ShippingStrategy strategy) {
        this.strategy = strategy;
    }

    public double calculate(double weight, double distance) {
        if (strategy == null) {
            throw new IllegalStateException("Shipping strategy not set");
        }
        return strategy.calculateShippingCost(weight, distance);
    }
}

