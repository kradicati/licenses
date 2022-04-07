<script lang="ts">
    import {Col, Row} from "sveltestrap/src";
    import authStore from "../stores/authStore";
</script>

<Row>
    <Col>
        <h1>Home page
            {#if $authStore.isLoggedIn}
                {$authStore.user.email}
            {/if}
        </h1>
    </Col>
</Row>