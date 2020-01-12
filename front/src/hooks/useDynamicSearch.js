import { useEffect, useState } from 'react'
import axios from 'axios'

const useDynamicSearch = () => {
    const [input, setInput] = useState("")
    const [output, setOutput] = useState([])

    const fetchResult = async () => {
        try {
            const res = await axios.get('/api/' + input)
            res.data && setOutput(res.data)
        } catch(e) {
            console.log(`useResource failed for '${input}' with error: `, e.response)
        }
    }

    useEffect(() => {
        input.trim() && fetchResult()
    }, [input])

    return [output, setInput]
}

export default useDynamicSearch