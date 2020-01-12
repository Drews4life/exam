import React from 'react'
import { useHistory } from 'react-router-dom';
import {
    Grid,
    Button,
} from '@material-ui/core';

const MainPage = () => {
    let history = useHistory();
    return (
        <Grid
            container
            direction="column"
            justify="center"
            alignItems="center"
            style={{ marginTop: "25px" }}
        >
            <div>
                <Button
                    variant="contained"
                    color="secondary"
                    onClick={() => history.push('/authors')}
                >
                    To Authors
                </Button>
                <Button
                    variant="contained"
                    color="secondary"
                    onClick={() => history.push('/add/authors')}
                >
                    +
                </Button>
            </div>
            <div>
                <Button
                    size={'large'}
                    variant="contained"
                    color="primary"
                    onClick={() => history.push('/books')}
                >
                    To Books
                </Button>
                <Button
                    variant="contained"
                    color="primary"
                    onClick={() => history.push('/add/books')}
                >
                    +
                </Button>
            </div>
            <Button
                style={{marginTop: '20px'}}
                size={'large'}
                variant="contained"
                color="primary"
                onClick={() => history.push('/search')}
            >
                To Search
            </Button>
        </Grid>
    )
}

export default MainPage