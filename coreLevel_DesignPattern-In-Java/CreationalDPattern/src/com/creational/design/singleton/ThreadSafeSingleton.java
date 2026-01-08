package com.creational.design.singleton;

public class ThreadSafeSingleton {
    private static ThreadSafeSingleton threadSafeSingleton;
    private ThreadSafeSingleton(){

    }

    public static synchronized ThreadSafeSingleton getLazyInitializationInstance(){
        if(threadSafeSingleton==null){
            ThreadSafeSingleton threadSafeSingleton = new ThreadSafeSingleton();
            return threadSafeSingleton;
        }else {
            return threadSafeSingleton;
        }
    }
}
