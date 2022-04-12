<script lang="ts">
    import axios from "axios";
    import { page } from '$app/stores';
    import {Container, Spinner} from "sveltestrap";
    import {Col, Row} from "sveltestrap/src";
    import authStore from "../../stores/authStore";

    $: loadData = async () => {
        const id = $page.params.id
        return await axios.get(`http://localhost:8000/api/v1/licenses/${id}`)
    }
</script>

<Container class="p-3 m-4">
    {#if $authStore.isLoggedIn}
        {#await loadData()}
            <Spinner class="text-primary"/>
        {:then license}
            <Row class="justify-content-between">
                <Col>
                    <h1 class="text-light mb-4 me-auto">{license.data.id}</h1>
                </Col>
            </Row>
        {:catch error}
            <p style="color: red">{error.message}</p>
        {/await}
    {/if}
</Container>