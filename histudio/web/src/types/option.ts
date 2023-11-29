// START GENERAL
export interface NestedSubOption {
    id: number;
    name: string;
    value: number;
    expanded: boolean;
    children: NestedSubOption[] | null; // for hierarchy view
    label?: string;
    valueString?: string;
    [key: string]: any;
}

export interface FlatOption {
    id: number; //if value === id, can don't pass value
    name: string;
    value?: number;
    valueString?: string;
}
// END GENERAL
