package com.creational.design.singleton;

import java.util.Objects;

public class LazyInitialization {
    private LazyInitialization lazyInitializationInstance;
    private LazyInitialization(){

    }

    public LazyInitialization getLazyInitializationInstance(){
        return Objects.requireNonNullElseGet(lazyInitializationInstance, LazyInitialization::new);
    }
}
