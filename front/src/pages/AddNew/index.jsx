import React, { useState } from 'react'
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import usePostResource from '../../hooks/usePostResource'
import { useHistory } from 'react-router-dom';

const AddNew = ({ match: { params }}) => {
    const type = params.type
    const fieldsSet = type === 'authors' ? ["Name", "Biography", "DateOfBirth"] : ["Title", "Author", "Year"]

    const [firstField, setFirstField] = useState('')
    const [secondField, setSecondField] = useState('')
    const [thirdField, setThirdField] = useState('')
    const history = useHistory();

    const onSubmit = async () => {
        const data = populateParams()
        await usePostResource(type, 'POST', data)
        history.push(`/${type}`)
    }

    const populateParams = () => {
        const values = [firstField, secondField, thirdField]
        const response = {}

        fieldsSet.forEach((field, i) => {
            response[field.toLowerCase()] = values[i]
        })

        return response
    }

    return (
        <div style={styles}>
            <TextField 
                label={fieldsSet[0]}
                value={firstField} 
                onChange={e => setFirstField(e.target.value)}
            />

            <TextField 
                style={{marginTop: '20px'}}
                label={fieldsSet[1]}
                value={secondField} 
                onChange={e => setSecondField(e.target.value)}
            />


            <TextField 
                style={{marginTop: '20px'}}
                label={fieldsSet[2]}
                value={thirdField} 
                onChange={e => setThirdField(e.target.value)}
            />

            <Button style={{marginTop: '25px'}} variant="contained" color="primary" onClick={onSubmit}>
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

export default AddNew