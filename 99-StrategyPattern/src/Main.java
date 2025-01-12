import com.context.ShippingCostCalculator;
import com.strategy.ExpressShipping;
import com.strategy.FreeShipping;
import com.strategy.StandardShipping;

public class Main {
    public static void main(String[] args) {
        ShippingCostCalculator calculator = new ShippingCostCalculator();

        // Example 1: Standard Shipping
        calculator.setStrategy(new StandardShipping());
        System.out.println("Standard Shipping Cost: " + calculator.calculate(10, 50)); // 10kg, 50km

        // Example 2: Express Shipping
        calculator.setStrategy(new ExpressShipping());
        System.out.println("Express Shipping Cost: " + calculator.calculate(10, 50));

        // Example 3: Free Shipping
        calculator.setStrategy(new FreeShipping());
        System.out.println("Free Shipping Cost: " + calculator.calculate(10, 50));
    }
}
