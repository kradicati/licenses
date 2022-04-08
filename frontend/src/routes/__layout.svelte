<script lang="ts">
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {Container} from "sveltestrap/";
    import {getAuth, onAuthStateChanged} from "firebase/auth"
    import {onMount} from "svelte";
    import {initializeApp} from "firebase/app";
    import authStore from "../stores/authStore";
    import axios from "axios";
    import {firebaseConfig} from "../lib/app";

    axios.interceptors.request.use(
        async function (request) {
            const user = getAuth().currentUser;

            if (user) {
                try {
                    const token = await user.getIdToken(false); // Force refresh is false
                    request.headers["Authorization"] = "Bearer " + token;
                    console.log(token)
                } catch (error) {
                    console.log("Error obtaining auth token in interceptor, ", error);
                }
            }

            return request;
        })

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

<Container>
    <slot/>
</Container>