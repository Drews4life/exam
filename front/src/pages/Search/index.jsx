import React, { useState } from 'react'
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import FormControl from '@material-ui/core/FormControl';
import TextField from '@material-ui/core/TextField';
import useDynamicSearch from '../../hooks/useDynamicSearch'
import LayoutItems from '../../components/LayoutItems';
import { useHistory } from 'react-router-dom';

const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
  }));

const SearchPage = () => {
    const [data, setInput] = useDynamicSearch()
    const [type, setType] = useState('books')
    const [query, setQuery] = useState('')

    const history = useHistory()

    const classes = useStyles();

    const handleTypeChange = e => {
        setType(e.target.value)
        setInput(`/${e.target.value}/search/${query}`)
    }

    const handleQueryChange = e => {
        setQuery(e.target.value)
        setInput(`/${type}/search/${e.target.value}`)
    }

    return (
        <div>
            <div style={styles}>
                <FormControl className={classes.formControl}>
                    <InputLabel id="demo-simple-select-label">Age</InputLabel>
                    <Select
                        labelId="demo-simple-select-label"
                        id="demo-simple-select"
                        value={type}
                        onChange={handleTypeChange}
                    >
                        <MenuItem value={'books'}>books</MenuItem>
                        <MenuItem value={'authors'}>authors</MenuItem>
                    </Select>
                    <TextField 
                        style={{marginTop: '20px'}}
                        id="outlined-basic" 
                        label="Search" 
                        variant="outlined" 
                        value={query}
                        onChange={handleQueryChange}
                    />
                </FormControl>
            </div>

            <div style={{marginTop: '20px'}}>
                {data.length > 0 && data.map((item, i) => (
                    <LayoutItems page={type} fetchAgain={() => history.push(`/${type}`)} item={item} key={i}/>
                ))}
            </div>
        </div>
    )
}

const styles = {
    marginTop: '20px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center'
}

export default SearchPage