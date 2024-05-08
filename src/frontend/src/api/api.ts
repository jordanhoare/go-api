// src/frontend/src/api/api.ts

import { DataResponse } from './models';

const BACKEND_URI = "http://localhost:8080/api";

export async function fetchData(): Promise<DataResponse> {
    try {
        const response = await fetch(`${BACKEND_URI}/hello`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        throw new Error(`There was an error fetching the data: ${error}`);
    }
}
