import React, { useState, useEffect } from 'react'
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import useGetResource from '../../hooks/useGetResource'
import usePostResource from '../../hooks/usePostResource';

const AuthorsUpdatePage = ({ match: { params }}) => {
    const [data] = useGetResource(`authors/${params.id}`)  
    const [name, setName] = useState("")
    const [biography, setBiography] = useState("")
    const [dateOfBirth, setDateOfBirth] = useState("")

    useEffect(() => {
        if(Object.keys(data).length > 0) {
            setName(data.name)
            setBiography(data.biography)
            setDateOfBirth(data.dateOfBirth)
        }
    }, [data])

    const onSubmitClick = () => {
        usePostResource(`authors`, 'PUT', {
            name,
            biography,
            dateOfBirth,
            id: Number(params.id),
        })
    }

    if(Object.keys(data).length === 0) { return (<div>Loading</div>)}

    return (
        <div style={styles}>
            <TextField 
                label="Name" 
                value={name} 
                onChange={e => setName(e.target.value)}
            />

            <TextField 
                style={{marginTop: '20px'}}
                label="Biography" 
                value={biography} 
                onChange={e => setBiography(e.target.value)}
            />


            <TextField 
                style={{marginTop: '20px'}}
                label="Date of Birth"
                value={dateOfBirth} 
                onChange={e => setDateOfBirth(e.target.value)}
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

export default AuthorsUpdatePage