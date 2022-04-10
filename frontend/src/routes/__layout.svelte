<script lang="ts">
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {getAuth, onAuthStateChanged} from "firebase/auth"
    import {onMount} from "svelte";
    import {initializeApp} from "firebase/app";
    import authStore from "../stores/authStore";
    import axios from "axios";
    import {firebaseConfig} from "../lib/app";

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

<slot></slot>