export enum Site {
    AO3 = "AO3",
    FFN = "FFN"
}
export interface Story {
    id?: string
    site: Site
}