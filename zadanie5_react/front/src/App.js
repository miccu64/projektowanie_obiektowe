import './App.css';
import Payments from "./components/Payments";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Cart from "./components/Cart";

function App() {
    return (
        <div style={{textAlign: "center"}}>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Payments/>}>
                        <Route index element={<Payments/>}/>
                    </Route>
                    <Route path="/cart" element={<Cart/>}/>
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;
