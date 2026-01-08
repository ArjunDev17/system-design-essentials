package com.creational.design.singleton;

public class EagerInitilizationSingleton {
    private static  final EagerInitilizationSingleton INSTANCE=new EagerInitilizationSingleton();
    private EagerInitilizationSingleton(){

    }
    public static EagerInitilizationSingleton getInstance1(){
        return INSTANCE;
    }
}
