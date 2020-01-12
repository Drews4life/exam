import { useEffect, useState } from 'react'
import axios from 'axios'

const usePostResource = async (url, method, data) => {
    try {
        const res = await axios({
            url: "/api" + url,
            method,
            data,
        })
        return res.data
    } catch(e) {
        console.log('could not perform this action with err: ', e)
        return null
    }
}

export default usePostResource