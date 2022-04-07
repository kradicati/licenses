import {writable} from "svelte/store"
import type auth from "firebase/auth"

const authStore = writable<{
    isLoggedIn: boolean,
    user?: auth.UserInfo
    firebaseControlled: boolean
}>({
    isLoggedIn: false,
    firebaseControlled: false
})

export default {
    subscribe: authStore.subscribe,
    set: authStore.set
}