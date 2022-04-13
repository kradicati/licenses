<script lang="ts">
    import axios from "axios";
    import {page} from '$app/stores';
    import {Container, ListGroup, ListGroupItem, Spinner} from "sveltestrap";
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
            <div class="text-light">
                <h1 class="text-light mb-4 me-auto">{license.data.id}</h1>
                <Row>
                    <Col class="glass">
                        <h4 class="p-3">Information</h4>
                    </Col>
                    <Col class="glass col-md-4">
                        <h4 class="p-3">IP Log</h4>
                        <ListGroup>
                            {#each license.data.ip_log as ip}
                                <ListGroupItem class="bg-transparent text-light lgi">{ip}</ListGroupItem>
                            {/each}
                        </ListGroup>
                    </Col>
                </Row>
            </div>
        {:catch error}
            <p style="color: red">{error.message}</p>
        {/await}
    {/if}
</Container>

<style>

    :global(.lgi) {
        border: 0;
        #border-bottom: 1px solid;
        #border-image: linear-gradient(to right, transparent 0%, rgba(255, 255, 255, 0.12) 50%, transparent 100%) 100% 1;
    }

    :global(.lgi::after) {
        content : "";
        position: absolute;
        left    : 25%;
        bottom  : 0;
        height  : 1px;
        width   : 50%;  /* or 100px */
        border-bottom: 1px solid rgba(255, 255, 255, 0.12);
    }
</style>