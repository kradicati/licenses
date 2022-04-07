<script lang="ts">
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {Container} from "sveltestrap/";
    import {getAuth, onAuthStateChanged} from "firebase/auth"
    import {onMount} from "svelte";
    import {initializeApp} from "firebase/app";
    import authStore from "../stores/authStore";

    onMount(() => {
        const firebaseConfig = {
            apiKey: "AIzaSyBrH3kOtyktbNXDS2ZsDER9t8Q7SpPCth4",
            authDomain: "licenses-506d2.firebaseapp.com",
            projectId: "licenses-506d2",
            storageBucket: "licenses-506d2.appspot.com",
            messagingSenderId: "893763060692",
            appId: "1:893763060692:web:69357f5dae5c2feded0c41"
        };

        initializeApp(firebaseConfig);

        const auth = getAuth()
        onAuthStateChanged(auth, (user) => {
            authStore.set({
                isLoggedIn: user !== null,
                user,
                firebaseControlled: true
            })
        })
    })
</script>

<Container>
    <slot/>
</Container>