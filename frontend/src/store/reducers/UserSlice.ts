import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {IUser} from "../../models/IUser";
import axios, {AxiosError} from "axios";
import {IError} from "../../models/IError";

interface UserState {
    token: string
    user: IUser
    loading: boolean
    error: string | unknown
}

const initialState: UserState = {
    token: "",
    user: <IUser>{},
    loading: false,
    error: ""
}

interface UserRequest {
    token: string
    user: IUser
}

export const fetchUser = createAsyncThunk<UserRequest,
    {login: string, password:string, email:string, isLogin: boolean}, {rejectValue:string}>(
    'user/fetchUser',
    async function ({login, password,email, isLogin}, {rejectWithValue}) {
        try {
            const urlPath = isLogin ? "login" : "registration"
            const url = ["http://localhost:8000/auth/", urlPath].join("")
            const response = await axios.post<UserRequest>(url, {
                username: login,
                password: password,
                email: email,
            })
            return response.data
        } catch (e: unknown) {
            const error = e as AxiosError<IError>;
            if (!error.response) {
                console.log(e)
                throw e;
            }
            console.log(error)
            if (error.response && error.response.status >= 500) {
                console.log(error)
                throw new Error(error.response.data.error)
            }
            console.log(error)
            return rejectWithValue(error.response.data.error)
        }

    }
)

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(fetchUser.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(fetchUser.fulfilled, (state, action) => {
                state.user = action.payload.user
                state.token = action.payload.token
                state.loading = false

            })
            .addCase(fetchUser.rejected, (state, action) => {
                state.error = action.payload
                state.loading = false
            })
    }
})

export default userSlice.reducer
