import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {IUser} from "../../models/IUser";
import axios, {AxiosError} from "axios";
import {IError} from "../../models/IError";
import {AuthState} from "../../models/AuthState";

interface UserState {
    token: string
    user: IUser
    loading: boolean
    error: string | unknown
}

const initialState: UserState = {
    token: "",
    user: {} as IUser,
    loading: false,
    error: ""
}

interface UserRequest {
    token: string
    user: IUser
}

export const fetchUser = createAsyncThunk<UserRequest,
    {login: string, password:string, email:string, authState: AuthState}, {rejectValue:string}>(
    'user/fetchUser',
    async function ({login, password,email, authState}, {rejectWithValue}) {
        try {
            let urlPath: string
            switch (authState) {
                case AuthState.SIGN_IN:
                    urlPath= "login"
                    break
                case AuthState.REGISTRATION:
                    urlPath = "registration"
                    break
            }
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
            if (error.response && error.response.status >= 500) {
                console.log(error)
                throw new Error(error.response.data.error)
            }
            console.log(error)
            console.log(error.response.data.error)
            return rejectWithValue(error.response.data.error)
        }

    }
)

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        // resetState: () => initialState,
        // resetError (state) {
        //     state.error = ""
        // }
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchUser.pending, (state) => {
                console.log(state.error)
                state.loading = true
                state.error = ""
                console.log("load")
            })
            .addCase(fetchUser.fulfilled, (state, action) => {
                state.user = action.payload.user
                state.token = action.payload.token
                state.loading = false
                console.log("fulfilled")
            })
            .addCase(fetchUser.rejected, (state, action) => {
                state.error = action.payload
                state.loading = false
                console.log("err")
                console.log(state.error)
            })
    }
})

// export const {resetState, resetError} = userSlice.actions

export default userSlice.reducer
