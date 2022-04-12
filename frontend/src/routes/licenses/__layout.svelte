<script>
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {Col, Row} from "sveltestrap/src";
    import Sidebar from "../../lib/Sidebar.svelte";
    import axios from "axios";
    import {getAuth, onAuthStateChanged} from "firebase/auth";
    import {onMount} from "svelte";
    import {initializeApp} from "firebase/app";
    import authStore from "../../stores/authStore";
    import {firebaseConfig} from "../../lib/app";

    let title = "Licenses"

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

    axios.interceptors.request.use(
        async function (request) {
            const user = getAuth().currentUser;

            if (user) {
                try {
                    const token = await user.getIdToken(false); // Force refresh is false
                    request.headers["Authorization"] = "Bearer " + token;
                } catch (error) {
                    console.log("Error obtaining auth token in interceptor, ", error);
                }
            }

            return request;
        })
</script>

<svelte:head>
    <title>{title}</title>
</svelte:head>

<Row>
    <Col class="col-lg-2 bg-dark shadow-lg p-3">
        <Sidebar/>
    </Col>
    <Col class="bg-black">
        <slot/>
    </Col>
</Row>

<!--Sidebar/>
<Col>
    <slot/>
</Col-->

