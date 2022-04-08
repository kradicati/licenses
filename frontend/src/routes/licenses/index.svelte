<script lang="ts">

    import authStore from "../../stores/authStore";
    import {Col, Row} from "sveltestrap/src";
    import axios from "axios";

    $: loadData = async () => {
        return await axios.get('http://localhost:8000/api/v1/licenses')
    }
</script>

<Row>
    <Col>
        {#if $authStore.isLoggedIn}
            {#await loadData()}
                Waiting...
            {:then list}
                {JSON.stringify(list.data, null, 1)}
            {:catch error}
                <p style="color: red">{error.message}</p>
            {/await}
        {/if}
    </Col>
</Row>