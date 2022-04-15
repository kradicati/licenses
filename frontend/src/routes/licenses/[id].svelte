<script lang="ts">
    import axios from "axios";
    import {page} from '$app/stores';
    import {Container, ListGroup, ListGroupItem, Spinner} from "sveltestrap";
    import {Col, Row} from "sveltestrap/src";
    import authStore from "../../stores/authStore";
    import Line from "svelte-chartjs/src/Line.svelte"
    import Doughnut from "svelte-chartjs/src/Doughnut.svelte"
    import moment from "moment";
    import {groupWeek, Pair} from "../../lib/utils";

    $: loadData = async () => {
        const id = $page.params.id
        return await axios.get(`http://localhost:8000/api/v1/licenses/${id}`)
    }

    function getLabels() {
        let days = []

        for (let i = 0; i < 7; i++) {
            days.push(moment().subtract(i, "days"))
        }

        return days.map(day => day.format("D MMM")).reverse()
    }

    let grouped: Pair<number[], number[]> = null

    function dataLine(log) {
        grouped = grouped === null ? groupWeek(log) : grouped

        return {
            labels: getLabels(),
            datasets: [
                {
                    label: "Success",
                    fill: true,
                    lineTension: 0.3,
                    backgroundColor: "rgba(159,220,169,0.3)",
                    borderColor: "rgb(158, 220, 169)",
                    borderCapStyle: "round",
                    borderDash: [],
                    borderDashOffset: 0.0,
                    borderJoinStyle: "miter",
                    pointBorderColor: "rgb(128,215,148)",
                    pointBackgroundColor: "rgb(117,215,139)",
                    pointBorderWidth: 5,
                    pointHoverRadius: 2,
                    pointHoverBackgroundColor: "rgb(117,215,139)",
                    pointHoverBorderColor: "rgba(220, 220, 220,1)",
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    data: grouped.x
                },
                {
                    label: "Failure",
                    fill: true,
                    lineTension: 0.3,
                    backgroundColor: "rgba(220,159,159,0.3)",
                    borderColor: "rgb(220,158,158)",
                    borderCapStyle: "round",
                    borderDash: [],
                    borderDashOffset: 0.0,
                    borderJoinStyle: "miter",
                    pointBorderColor: "rgb(215,128,128)",
                    pointBackgroundColor: "rgb(215,117,117)",
                    pointBorderWidth: 5,
                    pointHoverRadius: 2,
                    pointHoverBackgroundColor: "rgb(215,117,117)",
                    pointHoverBorderColor: "rgba(220, 220, 220,1)",
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    data: grouped.y
                },
            ]
        }
    }

    function dataDoughnut(log) {
        grouped = grouped === null ? groupWeek(log) : grouped

        return {
            labels: ["Success", "Failure"],
            datasets: [
                {
                    data: [grouped.x.reduce((a, b) => a + b, 0), grouped.y.reduce((a, b) => a + b, 0)],
                    backgroundColor: ["rgb(103,225,125)", "rgb(225,103,103)"],
                }
            ]
        }
    }
</script>

<Container class="p-3 m-4">
    {#if $authStore.isLoggedIn}
        {#await loadData()}
            <Spinner class="text-primary my-auto mx-auto"/>
        {:then license}
            <div class="text-light">
                <h1 class="text-light mb-4 me-auto">{license.data.id}</h1>
                <Row class="ms-1 mb-4">
                    <Col class="glass col-md-7 me-4">
                        <h4 class="p-3">Activity</h4>
                        <Line data={dataLine(license.data.ip_log)} options={{ responsive: true }}/>
                    </Col>
                    <Col class="glass col-md-3">
                        <h4 class="p-3">Responses</h4>
                        <Doughnut data={dataDoughnut(license.data.ip_log)} options={{ responsive: true }}/>
                    </Col>
                </Row>
                <Row class="ms-1">
                    <Col class="glass me-4 col-md-4">
                        <h4 class="p-3">Information</h4>
                    </Col>
                    <Col class="glass">
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
        content: "";
        position: absolute;
        left: 25%;
        bottom: 0;
        height: 1px;
        width: 50%; /* or 100px */
        border-bottom: 1px solid rgba(255, 255, 255, 0.12);
    }
</style>