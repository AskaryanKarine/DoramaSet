import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {IUser} from "../../models/IUser";
import {AxiosError} from "axios";
import {IError} from "../../models/IError";
import {AuthState} from "../../models/AuthState";
import {errorHandler} from "../../hooks/errorHandler";
import {instance} from "../../http-common";
import {ISubscribe} from "../../models/ISubscribe";

interface UserState {
    isAuth: boolean
    user: IUser
    loading: boolean
    error: string | unknown
}

const initialState: UserState = {
    isAuth: false,
    user: {} as IUser,
    loading: false,
    error: ""
}

interface UserRequest {
    user: IUser
}

interface Response {
    Data: ISubscribe
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
            const url = ["/auth/", urlPath].join("")
            const response = await instance.post<UserRequest>(url, {
                username: login,
                password: password,
                email: email,
            })
            return response.data
        } catch (e: unknown) {
            return rejectWithValue(errorHandler(e))
        }

    }
)

export const getUser = createAsyncThunk<UserRequest, undefined, {rejectValue:string}>(
    'user/getUser',
    async function (_, {rejectWithValue}) {
        try {
            const response = await instance.get<UserRequest>("/auth/")
            return response.data
        } catch (e: unknown) {
            return rejectWithValue(errorHandler(e))
        }
    }
)

export const logout = createAsyncThunk<void, undefined>(
    'user/logout',
    async function (_) {
        try {
            await instance.get("/auth/logout")
        } catch (e: unknown) {
            const error = e as AxiosError<IError>;
            if (!error.response) {
                console.log(e)
                throw e;
            }
        }
    }
)


export const subscribe = createAsyncThunk<ISubscribe, {idSub: number}, {rejectValue:string}>(
    "user/subscribe",
    async function ({idSub}, {rejectWithValue}){
        const url = ["/subscription/", idSub.toString()].join("")
        try {
            const response = await instance.post<Response>(url)
            return response.data.Data
        } catch (e:unknown) {
            return rejectWithValue(errorHandler(e))
        }
    }
)

export const unsubscribe = createAsyncThunk<void, undefined>(
    'user/unsubscribe',
    async function () {
        try {
            await instance.post("/subscription/")
        } catch (e: unknown) {
            const error = e as AxiosError<IError>;
            if (!error.response) {
                console.log(e)
                throw e;
            }
        }
    }
)

export const earnPoint = createAsyncThunk<number, {points: string}, {rejectValue:string}>(
    "user/earn/point",
    async function ({points}, {rejectWithValue}) {
        try {
            await instance.post("/user/earn/", {
                points: points})
            return parseInt(points)
        } catch (e:unknown) {
            return rejectWithValue(errorHandler(e))
        }
    }
)

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        resetState: () => initialState,
        resetError (state) {
            state.error = ""
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchUser.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(fetchUser.fulfilled, (state, action) => {
                state.user = action.payload.user
                state.isAuth = true
                state.loading = false
            })
            .addCase(fetchUser.rejected, (state, action) => {
                state.error = action.payload
                state.loading = false
            })

            .addCase(getUser.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(getUser.fulfilled, (state, action) => {
                state.user = action.payload.user
                state.isAuth = true
                state.loading = false
            })
            .addCase(getUser.rejected, (state, action) => {
                state.loading = false
                state.error = action.payload
            })

            .addCase(logout.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(logout.fulfilled, (state) => {
                state.user = {} as IUser
                state.isAuth = false
                state.error = ""
                state.loading = false
            })

            .addCase(subscribe.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(subscribe.fulfilled, (state, action) => {
                state.loading = false
                state.error = ""
                state.user.sub = action.payload
                state.user.points -= action.payload.cost
            })
            .addCase(subscribe.rejected, (state, action) => {
                state.loading = false
                state.error = action.payload
            })

            .addCase(unsubscribe.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(unsubscribe.fulfilled, (state) => {
                state.loading = false
                state.error = ""
                getUser()
            })

            .addCase(earnPoint.pending, (state) => {
                state.loading = true
                state.error = ""
            })
            .addCase(earnPoint.fulfilled, (state, action) => {
                state.loading = false
                state.error = ""
                state.user.points += action.payload
            })
            .addCase(earnPoint.rejected, (state, action) => {
                state.loading = false
                state.error = action.payload
            })

    }
})

export const {resetState, resetError} = userSlice.actions

export default userSlice.reducer
