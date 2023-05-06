import React, {useState} from "react";
import {useNavigate} from "react-router-dom";

function Cart() {
    const navigate = useNavigate();

    const [cart, setCart] = useState(null);

    React.useEffect(() => {
        fetch('/cart')
            .then(response => response.json())
            .then(data => setCart(data));
    }, []);

    const handleSubmit = event => {
        event.preventDefault();
        const requestOptions = {
            method: "POST"
        };
        fetch("/cart", requestOptions)
            .then(() => navigate('/'));
    };

    return <div>
        {cart && cart.map(product => {
            return <div style={{marginBottom: "10px"}}>
                <h5>Product name: {product.name}</h5>
                <h5>Quantity: {product.qty}</h5>
                <h5>Price: {product.price * product.qty}</h5>
            </div>
        })}
        <form onSubmit={handleSubmit} style={{marginBottom: "10px"}}>
            <button type={"submit"}>Pay for all</button>
        </form>
        <button onClick={() => navigate('/')}>Go back</button>
    </div>;
}

export default Cart;