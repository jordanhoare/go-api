// src/frontend/src/pages/data/Data.tsx

import { useState, useEffect } from 'react';
import { fetchData } from '../../api';
import { DataResponse } from '../../api/models'; 

function HelloPage() {
    const [data, setData] = useState<DataResponse | null>(null);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        fetchData()
            .then(response => setData(response))
            .catch(err => setError(err.message));
    }, []);

    return (
        <div>
            <h1>Data Page</h1>
            {data && <p>{data.message}</p>}
            {error && <p>Error: {error}</p>}
        </div>
    );
}

export default HelloPage;
