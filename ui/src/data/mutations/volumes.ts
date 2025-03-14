import { CreateResourceInput } from "../../types/resource";

export interface Response {
    ok: boolean;
}

export const createVolume = async (
    input: CreateResourceInput,
): Promise<Response> => {
    return fetch("http://10.100.0.186:8080/v1/resources", {
        method: "POST",
        body: JSON.stringify(input),
        headers: {
            "Content-Type": "application/json",
        },
    }).then((res) => res.json());
};
