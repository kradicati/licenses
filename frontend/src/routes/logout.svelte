<script lang="ts">
    import authStore from "../stores/authStore";
    import { goto } from '$app/navigation';
    import {getAuth, signOut} from "firebase/auth";

    authStore.subscribe(async info => {
        if (info.isLoggedIn && info.firebaseControlled) {
            const auth = getAuth();

            await signOut(auth)
            await goto("/login")
        }
    })
</script>