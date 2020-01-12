import React from 'react'
import LayoutItems from '../../components/LayoutItems'
import useGetResource from '../../hooks/useGetResource'

const AuthorsPage = () => {
    const [data, fetchAgain] = useGetResource('authors')
    return (
        <div style={{marginTop: '25px'}}>
            {data.length && data.map((item, i) => (
                <LayoutItems page="authors" fetchAgain={fetchAgain} item={item} key={i}/>
            ))}
        </div>
    )
}

export default AuthorsPage