import {combineReducers, configureStore} from "@reduxjs/toolkit";
import userReducer from './reducers/UserSlice'

const rootReducer = combineReducers({
    userReducer,
})

export const setupStore = () => {
    return configureStore({
        reducer: rootReducer
    })
}

export const store = setupStore()


export type RootState = ReturnType<typeof rootReducer>
export type AppStore = ReturnType<typeof setupStore>
export type AppDispatch = typeof store.dispatch
