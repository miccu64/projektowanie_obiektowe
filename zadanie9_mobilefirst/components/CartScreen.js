import {Button, StyleSheet, Text, View} from "react-native";
import {cart, categories, products} from "../services/DatabaseService";
import {useState} from "react";

const CartScreen = ({navigation}) => {
    const [cartCopy, setCartCopy] = useState({...cart});

    const deleteFromCart = (productId) => {
        delete cart[productId];
        const cartCopy2 = {...cart};
        setCartCopy(cartCopy2);
    }

    const payDeleteAllCart = () => {
        for (const prop in cart) {
            delete cart[prop];
        }
        setCartCopy({});
    }

    return (
        <View style={styles.centerAll}>
            <Text style={styles.header}>Cart details</Text>
            {Object.keys(cartCopy).map(productIdString => {
                const productId = Number(productIdString);
                const product = products.find(p => p.id === productId);
                const categoryName = categories.find(c => c.id === product.categoryId).name;
                const quantity = cartCopy[productId];

                return (<View key={product.id} style={styles.addMargin}>
                    <Text>Name: {product.name}</Text>
                    <Text>Category: {categoryName}</Text>
                    <Text>Quantity: {quantity}</Text>
                    <Text>Price: {product.price * quantity}</Text>
                    <Button title={"Delete from cart"} onPress={() => deleteFromCart(product.id)}></Button>
                </View>)
            })}
            <View style={styles.addMargin}>
                <Button title={"Go back"} onPress={() => navigation.navigate('Products')}></Button>
            </View>
            <View style={styles.addMargin}>
                <Button title={"Pay for all"} onPress={payDeleteAllCart}></Button>
                <Button title={"Delete all from cart"} onPress={payDeleteAllCart}></Button>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    addMargin: {
        margin: 15
    },
    centerAll: {
        flex: 1, alignItems: 'center', justifyContent: 'center'
    },
    header: {
        fontSize: 20, margin: 15
    },
});

export default CartScreen;
