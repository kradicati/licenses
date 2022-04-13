<script lang="ts">
    import {onMount} from "svelte";
    import {initializeApp} from "firebase/app";
    import {firebaseConfig} from "../lib/app";
    import {getAuth, onAuthStateChanged} from "firebase/auth";
    import authStore from "../stores/authStore";

    onMount(async () => {
        await initializeApp(firebaseConfig);

        const auth = await getAuth()
        onAuthStateChanged(auth, (user) => {
            authStore.set({
                isLoggedIn: user !== null,
                user,
            })
        })
    })
</script>

<slot/>