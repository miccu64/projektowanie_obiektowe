import Products from "./Products";
import {useNavigate} from "react-router-dom";
import React from 'react';

function Payments() {
    const navigate = useNavigate();
    const handleSubmit = async (event) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        const data = {};
        for (const [key, value] of formData.entries()) {
            data[key] = value;
        }

        const requestOptions = {
            method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify(data)
        };
        await fetch("/payments", requestOptions)
            .then(() => window.location.reload());
    };

    return <div>
        <form onSubmit={handleSubmit} style={{marginBottom: "10px"}}>
            <Products/>
            <button type={"submit"}>Add to cart</button>
        </form>
        <button onClick={() => navigate('/cart')}>Go to cart</button>
    </div>;
}

export default Payments;
