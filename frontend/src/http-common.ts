import axios from "axios";
import {useAppSelector} from "./hooks/redux";

const {token} = useAppSelector(state => state.userReducer)

export default axios.create({
    baseURL: "http://localhost:8000",
    headers: {
        "Content-type": "application/json",
        "Authorization": token,
    }
});