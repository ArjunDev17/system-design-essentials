package com.strategy;


public class StandardShipping implements ShippingStrategy {
    @Override
    public double calculateShippingCost(double weight, double distance) {
        return weight * distance * 0.5; // Standard cost formula
    }
}