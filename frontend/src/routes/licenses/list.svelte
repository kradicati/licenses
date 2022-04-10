<script lang="ts">
    import authStore from "../../stores/authStore";
    import {Col, Row} from "sveltestrap/src";
    import axios from "axios";
    import {Button, Container, Modal, ModalBody, ModalFooter, ModalHeader, Table} from "sveltestrap";
    import {formatDate} from "../../lib/utils";

    let toggled = false

    $: loadData = async () => {
        return await axios.get('http://localhost:8000/api/v1/licenses/')
    }

    $: toggleModal = () => {
        toggled = !toggled
    }
</script>

<Container class="p-3 m-4">
    <Row>
        <h1 class="text-light mb-4">Licenses</h1>
        <Col>
            <Button class="btn-primary" on:click={toggleModal()}>Create</Button>
        </Col>
    </Row>
    {#if $authStore.isLoggedIn}
        <Col>
            <Table hover striped class="bg-dark rounded text-light">
                <thead>
                <tr>
                    <th class="text-light p-3">Key</th>
                    <th class="text-light p-3">Created</th>
                    <th class="text-light p-3">Expires</th>
                    <th class="text-light p-3">Product</th>
                    <th class="text-light"></th>
                </tr>
                </thead>
                <tbody>
                {#await loadData()}
                    Loading...
                {:then list}
                    {#each list.data as item}
                        <tr>
                            <td class="text-light p-3">{item.id}</td>
                            <td class="text-light p-3">{formatDate(item.created)}</td>
                            <td class="text-light p-3">{formatDate(item.expires)}</td>
                            <td class="text-light p-3">{item.product}</td>
                            <td>
                                <button class="btn bg-transparent btn-transparent shadow-none">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="white"
                                         class="bi bi-gear" viewBox="0 0 16 16">
                                        <path d="M8 4.754a3.246 3.246 0 1 0 0 6.492 3.246 3.246 0 0 0 0-6.492zM5.754 8a2.246 2.246 0 1 1 4.492 0 2.246 2.246 0 0 1-4.492 0z"/>
                                        <path d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 0 1-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 0 1-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 0 1 .52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 0 1 1.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 0 1 1.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 0 1 .52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 0 1-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 0 1-1.255-.52l-.094-.319zm-2.633.283c.246-.835 1.428-.835 1.674 0l.094.319a1.873 1.873 0 0 0 2.693 1.115l.291-.16c.764-.415 1.6.42 1.184 1.185l-.159.292a1.873 1.873 0 0 0 1.116 2.692l.318.094c.835.246.835 1.428 0 1.674l-.319.094a1.873 1.873 0 0 0-1.115 2.693l.16.291c.415.764-.42 1.6-1.185 1.184l-.291-.159a1.873 1.873 0 0 0-2.693 1.116l-.094.318c-.246.835-1.428.835-1.674 0l-.094-.319a1.873 1.873 0 0 0-2.692-1.115l-.292.16c-.764.415-1.6-.42-1.184-1.185l.159-.291A1.873 1.873 0 0 0 1.945 8.93l-.319-.094c-.835-.246-.835-1.428 0-1.674l.319-.094A1.873 1.873 0 0 0 3.06 4.377l-.16-.292c-.415-.764.42-1.6 1.185-1.184l.292.159a1.873 1.873 0 0 0 2.692-1.115l.094-.319z"/>
                                    </svg>
                                </button>
                            </td>
                        </tr>
                    {/each}

                {:catch error}
                    <p style="color: red">{error.message}</p>
                {/await}
                </tbody>
            </Table>
        </Col>
    {/if}
</Container>

<Modal isOpen={toggled} xl class="bg-dark">
    <ModalHeader>Create license</ModalHeader>
    <ModalBody>
        Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua.
    </ModalBody>
    <ModalFooter>
        <Button color="primary" on:click={toggleModal()}>Do Something</Button>
        <Button color="secondary" on:click={toggleModal()}>Cancel</Button>
    </ModalFooter>
</Modal>