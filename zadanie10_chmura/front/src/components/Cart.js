import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Cart() {
    const navigate = useNavigate();

    const [cart, setCart] = useState(null);

    React.useEffect(() => {
        const fetchData = async () => {
            const data = await fetch('/api/cart');
            setCart(await data.json());
        }
        fetchData().catch(console.error);
    }, []);

    const handleSubmit = async (event) => {
        event.preventDefault();
        const requestOptions = {
            method: "POST"
        };
        await fetch("/api/cart", requestOptions)
            .then(() => navigate('/'));
    };

    return <div>
        {cart?.map(product => {
            return <div key={product.id} style={{ marginBottom: "10px" }}>
                <h5>Product name: {product.name}</h5>
                <h5>Quantity: {product.qty}</h5>
                <h5>Price: {product.price * product.qty}</h5>
            </div>
        })}
        <form onSubmit={handleSubmit} style={{ marginBottom: "10px" }}>
            <button type={"submit"}>Pay for all</button>
        </form>
        <button onClick={() => navigate('/')}>Go back</button>
    </div>;
}

export default Cart;