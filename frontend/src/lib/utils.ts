import moment from "moment";

export function formatDate(millis: number): string {
    return millis === undefined ? "undefined" : moment.unix(millis).format("D MMM, YYYY (hh:mm)")
}