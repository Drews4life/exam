import React from 'react'
import LayoutItems from '../../components/LayoutItems'
import useGetResource from '../../hooks/useGetResource'

const BooksPage = () => {
    const [data, fetchAgain] = useGetResource('books')
    return (
        <div style={{marginTop: '25px'}}>
            {data.length && data.map((item, i) => (
                <LayoutItems page="books" fetchAgain={fetchAgain} item={item} key={i}/>
            ))}
        </div>
    )
}

export default BooksPage