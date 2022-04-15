<script lang="ts">
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {getAuth} from "firebase/auth";
    import '../../../licenses.css'
    import {onMount} from "svelte";
    import {goto} from '$app/navigation';

    onMount(() =>
        getAuth().onAuthStateChanged((auth) => {
            auth.getIdTokenResult().then(token => {
                if ((token.claims.roles as string[]).indexOf("admin") <= -1) {
                    goto("/licenses")
                }
            })
        }))
</script>

<slot/>