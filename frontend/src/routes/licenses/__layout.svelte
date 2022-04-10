<script>
    import {Col, Row} from "sveltestrap/src";
    import Sidebar from "../../lib/Sidebar.svelte";
    import axios from "axios";
    import {getAuth} from "firebase/auth";

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

<Row class="h-25">
    <Col class="col-lg-2 bg-dark shadow-lg p-3">
        <Sidebar/>
    </Col>
    <Col class="bg-black">
        <slot></slot>
    </Col>
</Row>

<!--Sidebar/>
<Col>
    <slot/>
</Col-->

