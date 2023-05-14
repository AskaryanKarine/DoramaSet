import axios from "axios";


export let instance = axios.create({
    withCredentials: true,
    baseURL: "http://localhost:8000"
})