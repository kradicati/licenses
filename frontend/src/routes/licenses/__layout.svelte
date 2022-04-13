<script>
    import 'bootstrap/dist/css/bootstrap.min.css';
    import {Col, Row} from "sveltestrap/src";
    import Sidebar from "../../lib/Sidebar.svelte";
    import axios from "axios";
    import {getAuth} from "firebase/auth";
    import '../../licenses.css'

    let title = "Licenses"

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

<main>
    <Row class="h-100 gradient">
        <Col class="sidebar col-lg-2 p-3" style="width: 100px; max-width: 100px; min-width: 100px; border-right: 1px solid rgba(255, 255, 255, 0.12);
        background-color: rgba(3,2,12,0.5);">
            <Sidebar/>
        </Col>
        <Col class="bg-transparent pe-4">
            <slot/>
        </Col>
    </Row>
</main>

<!--Sidebar/>
<Col>
    <slot/>
</Col-->

<style>
    main {
        background: rgb(4, 3, 18);
        background: linear-gradient(to bottom right, rgba(4, 3, 18, 1) 0%, rgba(7, 6, 32, 1) 40%, rgba(4, 3, 18, 1) 100%);
    }
</style>
