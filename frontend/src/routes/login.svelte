<script lang="ts">
    import {Card, Container} from "sveltestrap";
    import {getAuth, signInWithEmailAndPassword} from "firebase/auth";
    import { goto } from '$app/navigation';

    let email = "", password = ""

    $: submit = async () => {
        const auth = getAuth();

        signInWithEmailAndPassword(auth, email, password)
            .then(() => {
                goto("/")
            })
            .catch((error) => {
                const errorCode = error.code;
                const errorMessage = error.message;

                console.log(errorCode + " " + errorMessage)
            });
    }
</script>

<Container>
    <Card class="p-xl-3 w-25 mx-auto my-auto">
        <form on:submit|preventDefault={submit}>
            <h1 class="h3 mb-3 fw-normal">Log in</h1>

            <div class="form-floating mb-3">
                <input bind:value={email} type="email" class="form-control" placeholder="Email" style="">
                <label>Email</label>
            </div>

            <div class="form-floating mb-3">
                <input bind:value={password} type="password" class="form-control" placeholder="Password">
                <label>Password</label>
            </div>

            <button class="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    </Card>
</Container>