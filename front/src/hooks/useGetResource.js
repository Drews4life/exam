import { useEffect, useState } from 'react'
import axios from 'axios'


const useGetResource = (url) => {
    const [response, setResponse] = useState(null)

    const fetchResult = async () => {
        try {
            const res = await axios.get('/api/' + url)
            res.data && setResponse(res.data)
        } catch(e) {
            console.log(`useResource failed for ${url} with error: `, e.response)
        }
    }

    useEffect(() => {
        fetchResult()
    }, [])

    return [response || [], fetchResult]
}

export default useGetResource