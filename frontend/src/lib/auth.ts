import { readable } from 'svelte/store'
import { browser } from '$app/env'
import type { Auth, User } from "firebase/auth"
import firebase from "firebase/compat";
import EmailAuthProvider = firebase.auth.EmailAuthProvider;

const createAuth = () => {
    let auth: Auth

    const { subscribe } = readable<User>(undefined, set => {
        let unsubscribe = () => {}

        async function listen() {
            if (browser) {
                const { app } = await import('./app')
                const { getAuth, onAuthStateChanged } = await import('firebase/auth')

                auth = getAuth(app)

                unsubscribe = onAuthStateChanged(auth, set)
            } else {
                set(null)
            }
        }

        listen()

        return () => unsubscribe()
    })

    async function signIn() {
        const { signInWithRedirect } = await import('firebase/auth')
        await signInWithRedirect(auth, new EmailAuthProvider())
    }

    async function signOut() {
        const { signOut } = await import('firebase/auth')
        await signOut(auth)
    }

    return {
        subscribe,
        signIn,
        signOut,
    }
}

export const auth = createAuth()