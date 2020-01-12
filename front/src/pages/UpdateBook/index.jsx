import React, { useState, useEffect } from 'react'
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import useGetResource from '../../hooks/useGetResource'
import usePostResource from '../../hooks/usePostResource';

const BooksUpdatePage = ({ match: { params }}) => {
    const [data] = useGetResource(`books/${params.id}`)  
    const [title, setTitle] = useState("")
    const [author, setAuthor] = useState("")
    const [year, setYear] = useState("")

    useEffect(() => {
        if(Object.keys(data).length > 0) {
            setTitle(data.title)
            setAuthor(data.author)
            setYear(data.year)
        }
    }, [data])

    const onSubmitClick = () => {
        usePostResource(`books`, 'PUT', {
            title,
            author,
            year,
            id: Number(params.id),
        })
    }

    if(Object.keys(data).length === 0) { return (<div>Loading</div>)}

    return (
        <div style={styles}>
            <TextField 
                label="Title" 
                value={title} 
                onChange={e => setTitle(e.target.value)}
            />

            <TextField 
                style={{marginTop: '20px'}}
                label="Author" 
                value={author} 
                onChange={e => setAuthor(e.target.value)}
            />


            <TextField 
                style={{marginTop: '20px'}}
                label="Year"
                value={year} 
                onChange={e => setYear(e.target.value)}
            />

            <Button style={{marginTop: '25px'}} variant="contained" color="primary" onClick={onSubmitClick}>
                Submit
            </Button>
        </div>
    )
}

const styles = {
    padding: '100px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center'
}

export default BooksUpdatePage