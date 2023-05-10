import {createSlice} from "@reduxjs/toolkit";
import {IUser} from "../../models/IUser";

interface UserState {
    token: string
    user: IUser
}


const initialState: UserState = {
    token: "",
    user: <IUser>{},
}

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {

    }
})

export default userSlice.reducer