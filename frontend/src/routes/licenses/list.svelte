<script lang="ts">
    import authStore from "../../stores/authStore";
    import {Col, Row} from "sveltestrap/src";
    import axios from "axios";
    import {
        Button,
        Container,
        Form,
        FormGroup,
        Input,
        Modal,
        ModalBody,
        ModalFooter,
        ModalHeader,
        Spinner,
        Table
    } from "sveltestrap";
    import {onMount} from "svelte";
    import moment from "moment";
    import {goto} from '$app/navigation';
    import {formatDate} from "../../lib/utils";

    const regex = RegExp("[0-9A-Z]{4}-[0-9A-Z]{5}-[0-9A-Z]{4}")

    let toggled = false

    $: loadData = async () => {
        return await axios.get('http://localhost:8000/api/v1/licenses/')
    }

    $: toggleModal = () => {
        toggled = !toggled
    }

    let innerDesc: HTMLTextAreaElement;
    let license = ""
    let expiry = ""
    let product = ""
    let description = ""

    const resize = (element: HTMLTextAreaElement) => {
        element.style.height = 'auto';
        element.style.height = 4 + element.scrollHeight + 'px';
    };

    function generateLicense() {
        let result = "";
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
        for (let i = 0; i < 15; i++) {
            if (i == 4 || i == 10) {
                result += "-"
                continue
            }
            result += characters.charAt(Math.floor(Math.random() *
                characters.length));
        }
        return result;
    }

    $: validateLicense = () => {
        return regex.test(license)
    }

    //2022-4-10
    $: create = async () => {
        let response = await axios.post('http://localhost:8000/api/v1/licenses/', {
            "id": license,
            "expires": new moment(expiry).unix(),
            "product": product,
            "additional_info": description
        })

        if (response.status == 201) {
            goto("/licenses/" + license)
        } else {
            toggleModal()
            alert("There was an error while trying to create the license: " + response.data)
        }
    }

    onMount(() => {
        license = generateLicense()
    })
</script>

<Container class="p-3 m-4">
    <Row class="justify-content-between">
        <Col>
            <h1 class="text-light mb-4 me-auto">Licenses</h1>
        </Col>
        <Col sm={{ size: 'auto', offset: 1 }}>
            <Button color="primary" on:click={toggleModal()}>Create</Button>
        </Col>
    </Row>
    {#if $authStore.isLoggedIn}
        <div class="round">
            <Table dark hover striped>
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
                    <Spinner class="text-primary"/>
                {:then list}
                    {#each list.data as item}
                        <tr>
                            <td class="text-light p-3">{item.id}</td>
                            <td class="text-light p-3">{formatDate(item.created)}</td>
                            <td class="text-light p-3">{formatDate(item.expires)}</td>
                            <td class="text-light p-3">{item.product}</td>
                            <td>
                                <button on:click={goto("/licenses/" + item.id)} class="btn bg-transparent btn-transparent shadow-none">
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

                <!--Spinner class="m-auto text-primary"/-->
                </tbody>
            </Table>
        </div>
    {/if}
</Container>

<Modal isOpen={toggled} xl>
    <ModalHeader>Create license</ModalHeader>
    <ModalBody>
        <Form>
            <!--FormGroup floating label="License">
                <Input valid={validLicense} bind:license on:input={validateLicense()} placeholder="Enter a value">asd{license}</Input>
            </FormGroup-->
            <!-- TODO validation -->
            <FormGroup>
                <Input bind:value={license} on:input={validateLicense()} placeholder={license}/>
            </FormGroup>

            <FormGroup floating label="Expiry date">
                <Input bind:value={expiry} placeholder="Expiry date"/>
            </FormGroup>
            <FormGroup floating label="Product">
                <Input bind:value={product} placeholder="Product"/>
            </FormGroup>
            <FormGroup floating label="Additional info">
                <Input rows={1} type="textarea" bind:innerDesc bind:value={description} on:input={resize}
                       placeholder="Additional info"/>
            </FormGroup>
        </Form>
    </ModalBody>
    <ModalFooter>
        <Button color="secondary" on:click={toggleModal()}>Cancel</Button>
        <Button color="primary" on:click={create()}>Create</Button>
    </ModalFooter>
</Modal>

<style>
    .round {
        border-radius: 6px;
        overflow: hidden;
    }
</style>