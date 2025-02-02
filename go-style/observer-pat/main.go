package main

import (
	"fmt"
	"observer-pat/observable"
	"observer-pat/observer"
)

func main() {
	// Create a new observable (IphoneObservableImpl)
	iphoneObservable := observable.NewIphoneObservableImpl()

	// Create observers (EmailAlertObserverImpl)
	emailObserver1 := observer.NewEmailAlertObserverImpl("kabby@example.com")
	emailObserver2 := observer.NewEmailAlertObserverImpl("bob@example.com")

	// Add observers to the observable
	iphoneObservable.Add(emailObserver1)
	iphoneObservable.Add(emailObserver2)

	// Set stock count, which will notify all observers
	iphoneObservable.SetStockCount(50) // Both Alice and Bob will be notified

	// Get the current stock count as a string and print it
	fmt.Println(iphoneObservable.GetStockCount("stockCount"))
}
