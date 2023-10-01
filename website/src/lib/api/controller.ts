import type { StdError } from "./error";

type CallParams = {
    path: string;
    method: "GET" | "POST" | "PUT" | "DELETE";
    body?: object;
}

class RestController {
    private url: string;
    
    constructor(url: string) {
        this.url = url;
    }

    async call<T>(params: CallParams): Promise<T> {
        const res = await fetch(`${this.url}${params.path}`, {
            method: params.method,
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            body: JSON.stringify(params.body)
        })
    
        const json = await res.json()
    
        if (!res.ok) {
            return Promise.reject(json as StdError)
        }
    
        return Promise.resolve(json as T);
    }
    
    async authCall<T>(params: CallParams): Promise<T> {

        const access_token = localStorage.getItem('access_token')
        const refresh_token = localStorage.getItem('refresh_token')

        if (access_token === null || refresh_token === null) {
            return Promise.reject("unauthorized")
        }

        const res = await fetch(`${this.url}${params.path}`, {
            method: params.method,
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": `Bearer ${access_token}`
            }),
            body: JSON.stringify(params.body)
        })
    
        const json = await res.json()
    
        if (!res.ok) {
            return Promise.reject(json as StdError)
        }
    
        return Promise.resolve(json as T);
    }
}

export const restController = new RestController(import.meta.env.VITE_API_URL);