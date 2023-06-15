import {Button, StyleSheet, Text, View} from "react-native";
import {cart, categories, products} from "../services/DatabaseService";

const ProductsScreen = ({navigation}) => {
    const addToCart = (productId) => {
        if (productId in cart) {
            cart[productId] += 1;
        } else {
            cart[productId] = 1;
        }
        alert('Added product to cart');
    }

    return (
        <View style={styles.centerAll}>
            <Text style={styles.header}>List of products</Text>
            {products.map(p => {
                const categoryName = categories.find(c => c.id === p.categoryId).name;
                return (<View key={p.id} style={styles.addMargin}>
                    <Text>Name: {p.name}</Text>
                    <Text>Price: {p.price}</Text>
                    <Text>Category: {categoryName}</Text>
                    <Button title={"Add to cart"} onPress={() => addToCart(p.id)}></Button>
                </View>)
            })}
            <View style={styles.addMargin}>
                <Button title={"Go to cart"} onPress={() => navigation.navigate('Cart')}></Button>
            </View>
        </View>
    );
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

export default ProductsScreen;
