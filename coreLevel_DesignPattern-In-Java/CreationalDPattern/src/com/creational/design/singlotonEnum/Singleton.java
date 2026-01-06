package com.creational.design.singlotonEnum;

public enum Singleton {
    INSTANCE;

    public void doSomething() {
        System.out.println("Singleton using enum");
    }
}

