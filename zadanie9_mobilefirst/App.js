import {createNativeStackNavigator} from "@react-navigation/native-stack";
import {NavigationContainer} from "@react-navigation/native";
import CartScreen from "./components/CartScreen";
import ProductsScreen from "./components/ProductsScreen";

const Stack = createNativeStackNavigator();

export default function App() {
    return (
        <NavigationContainer initialRouteName="ProductsScreen">
            <Stack.Navigator>
                <Stack.Screen name="Products" component={ProductsScreen}/>
                <Stack.Screen name="Cart" component={CartScreen}/>
            </Stack.Navigator>
        </NavigationContainer>
    );
}
