import com.creational.design.singlotonEnum.Singleton;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
void main() {
    //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
    // to see how IntelliJ IDEA suggests fixing it.
    IO.println(String.format("Hello and welcome!"));

    for (int i = 1; i <= 5; i++) {
        //TIP Press <shortcut actionId="Debug"/> to start debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
        // for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.
        IO.println("i = " + i);



    }

    // Accessing Singleton instance FIRST time
    Singleton singleton1 = Singleton.INSTANCE;
    singleton1.doSomething();

    // Accessing Singleton instance SECOND time
    Singleton singleton2 = Singleton.INSTANCE;
    singleton2.doSomething();

    // Check if both references point to same object
    System.out.println(singleton1 == singleton2);
}
