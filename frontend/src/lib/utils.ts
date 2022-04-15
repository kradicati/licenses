import moment from "moment";

export class Pair<X, Y> {
    x: X
    y: Y

    constructor(x: X, y: Y) {
        this.x = x
        this.y = y
    }
}

export function formatDate(millis: number): string {
    return millis === undefined ? "undefined" : moment.unix(millis).format("D MMM, YYYY (hh:mm)")
}

export function groupTimeSeries(timeSeries: object[],
                                periodStart: number,
                                periodEnd: number,
                                delta: number,
                                predicate: (object) => unknown): Map<string, number> {
    let toReturn = new Map<string, number>()

    for (let i = periodStart; i < periodEnd; i += delta) {
        let n = timeSeries
            .filter(function (obj) {
                const time = obj["time"] as number
                return time >= periodStart && time <= periodEnd && predicate(obj)
            })
            .length

        toReturn.set(formatDate(i), n)
    }

    return null
}

export function groupWeek(timeSeries: object[]): Pair<number[], number[]> {
    const day = 1000 * 60 * 60 * 24
    const time = moment().endOf("day").unix()
    const start = time - 7 * day

    let success = []
    let failure = []
    for (let t = start; t < time; t += day) {
        let n = timeSeries
            .filter(function (obj) {
                const time = obj["time"] as number
                return time >= t && time <= t + day
            })

        success.push(n.filter(obj => obj["status"] === true).length)
        failure.push(n.filter(obj => obj["status"] === false).length)
    }

    return new Pair<number[], number[]>(success, failure)
}